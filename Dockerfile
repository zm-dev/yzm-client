FROM golang:alpine as builder
RUN apk update && apk add nodejs
COPY . /go/src/github.com/zm-dev/yzm-client
WORKDIR /go/src/github.com/zm-dev/yzm-client
RUN go build -v -o main /go/src/github.com/zm-dev/yzm-client/main.go
RUN npm install --registry=https://registry.npm.taobao.org && \
    npm run build

FROM alpine:latest
COPY --from=builder /go/src/github.com/zm-dev/yzm-client/main /app/main
COPY --from=builder /go/src/github.com/zm-dev/yzm-client/static/ /app/static/
COPY --from=builder /go/src/github.com/zm-dev/yzm-client/mappings/ /app/mappings/
WORKDIR /app
RUN chmod +x /app/main
CMD ["./main"]