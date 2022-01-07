HOSTNAME=github.com
NAME?=avengers
BINARY?=terraform-example-${NAME}
VERSION=$(shell cat version/version.go | grep 'Version = ' | sed "s/Version = //g" | sed 's/"//g' | awk '{$$1=$$1};1')
OS_ARCH?=amd64
GIT_COMMIT?=dev
APP_IMAGE=${BINARY}:${VERSION}
DOCKER?=docker

BUILD_DATE?= $(shell date -u +"%Y-%m-%dT%TZ")

default: install

build: clean
	go build -o ./build/_output/bin/${BINARY}

install: clean build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/${VERSION}/${OS_ARCH}
	mv ./build/_output/bin/${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAME}/${VERSION}/${OS_ARCH}

.PHONY: ruclean
clean:
	rm -rf build/_output

# Build an operator image for deployment
.PHONY: image
image: build
	@${DOCKER} build . \
	-f ./Dockerfile \
	-t ${APP_IMAGE} \
	--pull \
	--no-cache \
	--build-arg RELEASE=${GIT_COMMIT} \
	--build-arg BUILD_DATE=${BUILD_DATE} \
	--build-arg VERSION=${VERSION}