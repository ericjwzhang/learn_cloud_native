FROM golang:1.17 AS build
WORKDIR /httpsvr
COPY . .
ENV CGO_ENABLED=0 \
    GO111MODULE=auto \
	GOOS=linux \
	GOARCH=amd64 \
	GOPROXY=https://goproxy.cn,direct
RUN go build -o httpsvr main.go

FROM busybox
COPY --from=build /httpsvr/httpsvr /
EXPOSE 8080
ENTRYPOINT ["/httpsvr"]