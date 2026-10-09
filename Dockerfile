FROM golang:1.26
LABEL maintainer="Victor Castell <0x@vcastellm.xyz>"

EXPOSE 8080 8946

RUN mkdir -p /app
WORKDIR /app

ENV GOCACHE=/root/.cache/go-build
ENV GOMODCACHE=/root/.cache/go-build
ENV GO111MODULE=on
ENV GOMAXPROCS=2

# Leverage build cache by copying go.mod and go.sum first
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build go mod download
RUN go mod verify

# Copy the rest of the source code
COPY . .

RUN go install -p 2 ./builtin/...
RUN --mount=type=cache,target=/root/.cache/go-build go build -p 2 -o /go/bin/dkron main.go

CMD ["dkron"]
