#write here path for your project
PROJECT := $(shell pwd)
GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION := latest
APP_NAME := k8s-go-app

all: run

run:
  go install -ldflags="-X '$(PROJECT)/version.Version=$(VERSION)' \
  -X '$(PROJECT)/version.Commit=$(GIT_COMMIT)'" && $(APP_NAME)

