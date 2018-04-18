FROM golang:alpine as builder
RUN apk update && apk add nodejs
COPY . /go/src/github.com/zm-dev/yzm-client
WORKDIR /go/src/github.com/zm-dev/yzm-client
RUN go build -v -o main /go/src/github.com/zm-dev/yzm-client/src/backend/main.go
RUN npm install --registry=https://registry.npm.taobao.org && \
    npm run build

FROM alpine:latest
RUN apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /go/src/github.com/zm-dev/yzm-client/main /app/main
COPY --from=builder /go/src/github.com/zm-dev/yzm-client/static/ /app/static/
WORKDIR /app
RUN chmod +x /app/main
CMD ["./main"]