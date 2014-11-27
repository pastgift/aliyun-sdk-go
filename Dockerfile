FROM docker.cn/docker/golang:1.3.3

RUN mkdir -p /go
WORKDIR /go

COPY . /go

RUN go build /go/src/main.go

CMD ["go", "run", "/go/src/main.go"]
