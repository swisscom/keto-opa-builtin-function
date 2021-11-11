FROM golang:1.17-alpine-3.14 as builder
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN go build -o /go/bin/opa

FROM alpine:3.14
COPY --from=builder /go/bin/opa /
ENTRYPOINT [ "/opa" ]
CMD ["run"]
