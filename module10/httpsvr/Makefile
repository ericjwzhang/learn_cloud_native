

export tag=v1.0-metrics
root:
	export ROOT=github.com/ericjwzhang/learn_cloud_native/module10/httpsvr

build:
	echo "building httpsvr binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=auto go build -o httpsvr .

release:
	echo "building httpsvr container"
	docker build -t ericjwzhang/httpsvr:${tag} .

push: release
	echo "pushing ericjwzhang/httpsvr"
	docker push ericjwzhang/httpsvr:${tag}