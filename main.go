package main

import (
	"net/http"
	"net/http/cookiejar"
	"log"
	"os"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"mime/multipart"
	"io"
	"encoding/json"
	"path/filepath"
	"flag"
)

type FabricConfigData struct {
	DeveloperToken string `json:"developer_token"`
}

func debug(data interface{}) {
	log.Printf("> %v", data)
}

func debugFatal(data interface{}) {
	log.Printf("> Error: %v", data)
	os.Exit(1)
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)

	if err == nil {
		request.Header.Set("Content-Type", writer.FormDataContentType())
	}

	return request, err
}

func main() {
	appBundleId := flag.String("bundleid", "", "the app Bundle ID")
	fabricApiKey := flag.String("fabricapikey", "", "the fabric Api Key")
	fileName := flag.String("file", "", "the 'dsym.zip' file path")

	flag.Parse()

	if *appBundleId == "" {
		log.Fatal("The App Bundle ID is not defined")
	}

	if *fabricApiKey == "" {
		log.Fatal("The Fabric Api Key is not defined")
	}

	if *fileName == "" {
		log.Fatal("The FileName is not defined")
	}

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	////////////////////////////////////////////////////////////////////////////////
	// GET CSRF TOKEN
	////////////////////////////////////////////////////////////////////////////////
	debug("Getting CSRF token...")

	var response *http.Response
	var responseError error

	response, responseError = client.Get("https://fabric.io/login")

	if (responseError != nil) {
		debugFatal(fmt.Sprintf("Error: %v", responseError))
	}

	doc, err := goquery.NewDocumentFromResponse(response)

	if err != nil {
		debugFatal(err)
	}

	var fabricCsrfToken string = ""

	doc.Find("meta[name=csrf-token]").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("content")

		if (exists) {
			fabricCsrfToken = val
		}
	})

	if fabricCsrfToken == "" {
		debugFatal("CSRF token not found")
	} else {
		debug(fmt.Sprintf("CSRF token: %v", fabricCsrfToken))
	}

	////////////////////////////////////////////////////////////////////////////////
	// GET DEVELOPER TOKEN
	////////////////////////////////////////////////////////////////////////////////
	debug("Getting developer token...")

	reqDT, _ := http.NewRequest("GET", "https://fabric.io/api/v2/client_boot/config_data", nil)
	reqDT.Header.Set("X-CSRF-Token", fabricCsrfToken)

	resp, err := client.Do(reqDT)

	if (err != nil) {
		debugFatal(fmt.Sprintf("Error: %v", responseError))
	}

	decoder := json.NewDecoder(resp.Body)

	configData := &FabricConfigData{}

	err = decoder.Decode(&configData)

	if err != nil {
		debugFatal(fmt.Sprintf("Error: %v", responseError))
	}

	debug(fmt.Sprintf("Developer token: %v", configData.DeveloperToken))

	////////////////////////////////////////////////////////////////////////////////
	// UPLOAD FILE
	////////////////////////////////////////////////////////////////////////////////

	debug("Uploading file...")

	extraParams := map[string]string{
		"project[identifier]": *appBundleId,
		"code_mapping[type]": "dsym",
	}

	requestU, err := newfileUploadRequest("https://cm.crashlytics.com/api/v3/platforms/ios/code_mappings", extraParams, "code_mapping[file]", *fileName)

	if err != nil {
		debugFatal(err)
	}

	requestU.Header.Set("X-CRASHLYTICS-API-KEY", *fabricApiKey)
	requestU.Header.Set("X-CRASHLYTICS-DEVELOPER-TOKEN", configData.DeveloperToken)

	resp, err = client.Do(requestU)

	if err != nil {
		debugFatal(err)
	}

	uploadStatus := resp.StatusCode

	debug(fmt.Sprintf("Upload status: %v", uploadStatus))

	if uploadStatus >= 200 && uploadStatus <= 299 {
		debug("SUCCESS")
	} else {
		debug("ERROR")
	}
}