FROM golang:1.13-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/opa

FROM gcr.io/distroless/cc

COPY --from=build /go/bin/opa /

ENTRYPOINT [ "/opa" ]

CMD ["run"]
