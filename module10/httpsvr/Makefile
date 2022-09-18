

export tag=v1.0-metrics
root:
	export ROOT=github.com/ericjwzhang/learn_cloud_native/module10/httpsvr

build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=auto go build -o httpsvr .

release: build
	echo "building httpserver container"
	docker build -t ericjwzhang/httpsvr:${tag} .

push: release
	echo "pushing ericjwzhang/httpserver"
	docker push ericjwzhang/httpsvr:${tag}