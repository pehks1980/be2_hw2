
#PROJECT := $(notdir $(shell pwd)) #local path of makefile
PROJECT := k8s-go-app#$(shell pwd) #to check exact path use ***)
GIT_COMMIT := $(shell git rev-parse HEAD) #commit head
VERSION := latest
APP_NAME := k8s-go-app
BUILD := 1.0
USERNAME := pehks1980
PORT := 8080


all: run

run:
	go install -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' -X '$(PROJECT)/version.Commit=$(GIT_COMMIT)' \
-X '$(PROJECT)/version.Build=$(BUILD)'" && $(APP_NAME)
#tab must be before action go
#add ldflags to go app

# ***) here
#go build
#user@MacBook k8s-go-app % go tool nm ./k8s-go-app | grep version - check variables of package 'version'
#100265310 T k8s_go_app/k8s-go-app/server.Server.versionHandler-fm
#10050c150 D k8s_go_app/k8s-go-app/version.Build<<<<<<<<<<<<<<<<<<<<
#10050c160 D k8s_go_app/k8s-go-app/version.Commit
#10050c170 D k8s_go_app/k8s-go-app/version.Version

build_container:
	docker build --build-arg=GIT_COMMIT=$(GIT_COMMIT) --build-arg=VERSION=$(VERSION)  --build-arg=PROJECT=$(PROJECT)\
  -t docker.io/$(USERNAME)/$(APP_NAME):$(VERSION) .

run_container: ## Run container on port configured in `local.env`
	docker run -i -p=$(PORT):$(PORT) --name="$(APP_NAME)" $(USERNAME)/$(APP_NAME)

push_container:
	docker push  docker.io/$(USERNAME)/$(APP_NAME):$(VERSION)

#get name of result image
#docker images | grep pehks
#pehks1980/k8s-go-app          latest...
#load docker image in local kube 'registry'
#user@MacBook k8s-go-app % minikube image load pehks1980/k8s-go-app
#check kube images
#minikube image ls | grep pehks

#check minikube overall operations:
#minikube dashboard

#check events on cubectl:
#cubectl get events


#delete deployment & service:
 #kubectl delete deployment k8s-go-app
 #kubectl delete service k8s-go-app-srv