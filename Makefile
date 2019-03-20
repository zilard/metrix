.PHONY: install build clean image

CWD =  $(shell pwd)
GOPATH = "${CWD}/build"

install:
	[ -d ${GOPATH}] || mkdir -p ${GOPATH}
	$(shell cd metrix; GOPATH=${GOPATH} go get -d -v ./...)
	$(shell cd mx; GOPATH=${GOPATH} go get -d -v ./...)
 
build: install
	$(shell cd ${GOPATH}/src/github.com/zilard/metrix; \
          GOPATH=${GOPATH} go install -v ./...)

clean:
	rm -fr ${GOPATH}

image: build
	docker build -t metrix-image . 
