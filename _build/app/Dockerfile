FROM golang:1.10.4-stretch

WORKDIR /go/src/ciao-universe
COPY . .

RUN go get -d -v ./cmd/app
RUN go install -v ./cmd/app

CMD ["app"]