# Default build target
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

build: build-cmd build-plugin-fiz build-plugin-buz

build-cmd:  ## build the core cli
	CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -v \
		-o "./dist/${GOOS}/${GOARCH}/ankorcli" ./cmd

build-plugin-%:  ## build the plugin
	CGO_ENABLED=1 GOOS=${GOOS} GOARCH=${GOARCH} go build -v \
	-buildmode=plugin \
		-o "./dist/${GOOS}/${GOARCH}/$*.so" ./plugins/$*