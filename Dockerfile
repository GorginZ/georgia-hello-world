FROM golang:1.18-alpine as ci
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.46.2

FROM golang:1.18-alpine as builder
WORKDIR /go/src/github.com/GorginZ/basic-app
COPY . /go/src/github.com/GorginZ/basic-app

RUN go mod download
RUN go mod tidy
ENV CGO_ENABLED=0

ARG version
ARG sha
ARG description

RUN CGO_ENABLED=0 go build -v -o /_build/basic-app -ldflags "-X 'github.com/gorginz/georgia-hello-world/app-metadata/app-metadata.Version=${version}' -X 'github.com/gorginz/georgia-hello-world/app-metadata/app-metadata.Sha=${sha}' -X 'github.com/gorginz/georgia-hello-world/app-metadata/app-metadata.Description=${description}'"

FROM scratch
COPY --from=builder /_build/basic-app /app

ENTRYPOINT ["/app"]

