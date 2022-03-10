FROM golang:alpine AS development
WORKDIR $GOPATH/src
COPY . .
RUN GOPROXY="https://nexus3.corp.youdao.com/repository/go-all" CGO_ENABLED=0 GOOS=linux go build -o cyandevbak

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/cyandevbak .
COPY --from=development /go/src/configonline.json config.json
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update && apk add curl bash tree tzdata \
    && cp -r -f /usr/share/zoneinfo/Hongkong /etc/localtime \
    && echo -ne "Alpine Linux 3.4 image. (`uname -rsv`)\n" >> /root/.built
EXPOSE 5000
ENTRYPOINT ["./cyandevbak"]