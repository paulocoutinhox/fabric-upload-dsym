PROJECT=fabric-upload-dsym
LOG_FILE=/var/log/${PROJECT}.log
GOFMT=gofmt -w
GODEPS=go get -u
PACKAGE=github.com/prsolucoes/gohc

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
	rm -rf build

	mkdir -p build/linux32
	env GOOS=linux GOARCH=386 go build -o build/linux32/${PROJECT} -v ${PACKAGE}

	mkdir -p build/linux64
	env GOOS=linux GOARCH=amd64 go build -o build/linux64/${PROJECT} -v ${PACKAGE}

	mkdir -p build/darwin64
	env GOOS=darwin GOARCH=amd64 go build -o build/darwin64/${PROJECT} -v ${PACKAGE}

	mkdir -p build/windows32
	env GOOS=windows GOARCH=386 go build -o build/windows32/${PROJECT}.exe -v ${PACKAGE}

	mkdir -p build/windows64
	env GOOS=windows GOARCH=amd64 go build -o build/windows64/${PROJECT}.exe -v ${PACKAGE}