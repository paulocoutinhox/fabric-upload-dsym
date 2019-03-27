PROJECT=fabric-upload-dsym
LOG_FILE=/var/log/${PROJECT}.log
GOFMT=gofmt -w
GODEPS=go get -u
PACKAGE=github.com/prsolucoes/fabric-upload-dsym

.DEFAULT_GOAL := help

# general
help:
	@echo "Type: make [rule]. Available options are:"
	@echo ""
	@echo "- help"
	@echo "- build"
	@echo "- install"
	@echo "- format"
	@echo "- deps"
	@echo "- stop"
	@echo "- update"
	@echo "- build-all"
	@echo ""

build:
	go build -o ${PROJECT}

install:
	go install

format:
	${GOFMT} main.go

deps:
	${GODEPS} github.com/PuerkitoBio/goquery

stop:
	pkill -f ${PROJECT}

update:
	git pull origin master
	make deps
	make install

build-all:
	rm -rf dist

	mkdir -p dist/linux32
	env GOOS=linux GOARCH=386 go build -o dist/linux32/${PROJECT} -v ${PACKAGE}

	mkdir -p dist/linux64
	env GOOS=linux GOARCH=amd64 go build -o dist/linux64/${PROJECT} -v ${PACKAGE}

	mkdir -p dist/darwin64
	env GOOS=darwin GOARCH=amd64 go build -o dist/darwin64/${PROJECT} -v ${PACKAGE}

	mkdir -p dist/windows32
	env GOOS=windows GOARCH=386 go build -o dist/windows32/${PROJECT}.exe -v ${PACKAGE}

	mkdir -p dist/windows64
	env GOOS=windows GOARCH=amd64 go build -o dist/windows64/${PROJECT}.exe -v ${PACKAGE}