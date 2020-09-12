FROM golang:alpine as builder

WORKDIR /go/src/deploybot/webhook/
COPY . .

RUN go mod download && go build .

FROM alpine:latest as runner

WORKDIR /opt/bin/
COPY --from=builder /go/src/deploybot/webhook/ .


EXPOSE 4949/tcp

ENTRYPOINT ["./deploybot"]
