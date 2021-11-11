FROM golang:1.17-alpine as builder
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN go build -o /go/bin/opa-keto ./cmd/opa-keto

FROM alpine:3.14
COPY --from=builder /go/bin/opa-keto /
ENTRYPOINT [ "/opa-keto" ]
CMD ["run"]
