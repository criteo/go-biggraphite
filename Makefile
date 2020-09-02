include Makefile.common

.PHONY: test
test: common-test

.PHONY: build
build: common-build

.PHONY: coverage
coverage: promu
	GO111MODULE=$(GO111MODULE) $(GOTEST) $(test-flags) -v -covermode=count -coverprofile=coverage.out $(pkgs)
	$(GOPATH)/bin/goveralls -coverprofile=coverage.out -service=travis-pro
