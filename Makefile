.PHONY: install build clean test image deploy kubeclean kubedeploy

CWD =  $(shell pwd)
GOPATH = "${CWD}/build"

install:
	[ -d ${GOPATH}] || mkdir -p ${GOPATH}
	cd metrix; GOPATH=${GOPATH} go get -d -v ./...
	cd mx; GOPATH=${GOPATH} go get -d -v ./...
 
build: install
	cd metrix; go build .
	cd mx; go build .


clean:
	rm -fr ${GOPATH}
	kubectl delete -f k8s/deploy.yaml | true
	kubectl delete svc metrix | true
	@IMG=$(docker image list | grep "metrix-image:v1 " | awk '{ print $3 }')
	@docker ps -a | awk '{ print $1 " " $2 }' | grep "${IMG}" |  awk '{print $1 }' | xargs -I {} docker rm {} --force
	docker rmi metrix-image:v1 --force | true
	@IMG=$(docker image list | grep "metrix-image " | awk '{ print $3 }')
	@docker ps -a | awk '{ print $1 " " $2 }' | grep "${IMG}" |  awk '{print $1 }' | xargs -I {} docker rm {} --force
	docker rmi metrix-image --force | true

test: install
	cd metrix; GOPATH=${GOPATH} go test -v ./...

kubeclean:
	kubectl delete -f k8s/deploy.yaml | true
	kubectl delete svc metrix | true

kubedeploy: kubeclean
	kubectl create -f k8s/deploy.yaml

image: build
	docker build -t metrix-image .
	docker image tag metrix-image:latest metrix-image:v1

deploy: clean image
	kubectl create -f k8s/deploy.yaml

