FROM golang:alpine as builder

MAINTAINER quangdangfit<quangdangfit@gmail.com>

WORKDIR /go/src/getjob
COPY . /go/src/getjob
RUN go build -o ./dist/getjob

FROM alpine:3.11.3
RUN apk add --update ca-certificates
RUN apk add --no-cache tzdata && \
  cp -f /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime && \
  apk del tzdata

COPY ./config/config.yaml .
COPY --from=builder /go/src/getjob/dist/getjob .
EXPOSE 8888
ENTRYPOINT ["./getjob"]
