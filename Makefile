.PHONY: install build clean test image deploy

CWD =  $(shell pwd)
GOPATH = "${CWD}/build"

install:
	[ -d ${GOPATH}] || mkdir -p ${GOPATH}
	cd metrix; GOPATH=${GOPATH} go get -d -v ./...
	cd mx; GOPATH=${GOPATH} go get -d -v ./...
 
build: install test
	cd ${GOPATH}/src/github.com/zilard/metrix; GOPATH=${GOPATH} go install -v ./...

clean:
	rm -fr ${GOPATH}
	kubectl delete -f k8s/deploy.yaml | true
	kubectl delete svc metrix | true
	docker rmi metrix-image:v1 | true
	docker rmi metrix-image | true

test: install
	cd metrix; GOPATH=${GOPATH} go test -v ./...

image: build
	docker build -t metrix-image .
	docker image tag metrix-image:latest metrix-image:v1

deploy: clean image
	kubectl create -f k8s/deploy.yaml

