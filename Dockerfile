FROM golang:latest as build

ENV GOPROXY https://goproxy.cn
ENV GO111MODULE on

WORKDIR /go/cache_daily

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o daily main.go

FROM alpine

COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/daily /daily

RUN chmod +x /daily

EXPOSE 8080

ENTRYPOINT /daily
