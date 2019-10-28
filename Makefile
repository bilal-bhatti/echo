cnf ?= make.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
.DEFAULT_GOAL := help

github-login: ## Login to GitHub Package Registry
	docker login docker.pkg.github.com --username ${GIT_USER}

publish: build-image push-image ## Build and publish Docker image to GitHub Package Registry

build-image: ## Build linux-amd64 arch Docker image
	env GOOS=linux GOARCH=amd64 go build -o ./dist/echo-linux-amd64
	docker build -t ${GIT_USER}/echo/linux-amd64 .

push-image: ## Push image to GitHub Package Registry
	docker tag ${GIT_USER}/echo/linux-amd64 docker.pkg.github.com/${GIT_USER}/echo/linux-amd64:latest
	docker push docker.pkg.github.com/${GIT_USER}/echo/linux-amd64:latest
	docker tag ${GIT_USER}/echo/linux-amd64 docker.pkg.github.com/${GIT_USER}/echo/linux-amd64:${VERSION}
	docker push docker.pkg.github.com/${GIT_USER}/echo/linux-amd64:${VERSION}
