FROM golang:1.17 AS build
WORKDIR /httpsvr/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpsvr main.go

FROM busybox
COPY --from=build /httpsvr/httpsvr /httpsvr/httpsvr
EXPOSE 8080
ENV ENV local
WORKDIR /httpsvr/
ENTRYPOINT ["./httpsvr"]


