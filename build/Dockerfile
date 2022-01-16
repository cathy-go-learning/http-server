FROM golang:1.14.1-stretch as builder

ARG GITHUB_TOKEN
ARG GITHUB_USER

WORKDIR /build/src/github.com/cathy-go-learning/http-server

ENV GOPATH /build
ENV GOBIN $GOPATH/bin
ENV GOPRIVATE github.com/cathy-go-learning/*

COPY ./go.mod ./

## Download dependencies
RUN git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/" && \
    go mod download -x -json && \
    git config --global --remove-section url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/"

COPY . $GOPATH/src/github.com/cathy-go-learning/http-server


RUN cd $GOPATH/src/github.com/cathy-go-learning/http-server/cmd  && \
  GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /build/http-server .


FROM alpine as binary
WORKDIR /app

RUN apk --no-cache add git ca-certificates

COPY --from=builder /build/http-server .

RUN chmod +x http-server

CMD ["./http-server"]
