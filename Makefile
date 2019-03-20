.PHONY: install build clean image deploy minikube

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
	kubectl delete -f k8s/deploy.yaml | true
	kubectl delete svc metrix | true
	docker rmi metrix-image:v1 | true
	docker rmi metrix-image | true

image: build
	docker build -t metrix-image .
	docker image tag metrix-image:latest metrix-image:v1

deploy: clean image
	kubectl create -f k8s/deploy.yaml

