REPO=aputrabay
IMAGE=lesson-082
VERSION=v0.1.3

all: build-docker push-docker

build-docker:
	docker build -t ${REPO}/${IMAGE}:${VERSION} -f app/Dockerfile app/
push-docker:
	docker push ${REPO}/${IMAGE}:${VERSION}
