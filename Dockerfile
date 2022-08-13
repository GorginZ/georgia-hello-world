FROM golang:1.18-alpine as ci
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2

FROM golang:1.18-alpine as builder
WORKDIR /go/src/github.com/GorginZ/basic-app
COPY . /go/src/github.com/GorginZ/basic-app

RUN go mod download
RUN go mod tidy

ENV CGO_ENABLED=0
RUN go test ./... -v

RUN CGO_ENABLED=0 go build -o /_build/basic-app

FROM scratch
COPY --from=builder /_build/basic-app /app
ENTRYPOINT ["/app"]
