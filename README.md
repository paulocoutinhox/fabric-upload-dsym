# Fabric Upload DSYM Tool

This is a simple command line tool made with Go that upload any DSYM to *Fabric.io*

# Usage

```
fabric-upload-dsym --bundleid=[YOUR-APP-BUNDLE] --fabricapikey=[YOUR-FABRIC-API-KEY] --file=[ZIPPED-DSYM-FILE]
```

# Get, build and install

```
go install github.com/prsolucoes/fabric-upload-dsym
```

# Help

```
fabric-upload-dsym -help
```

# Dependencies to build

```
go get github.com/PuerkitoBio/goquery
```