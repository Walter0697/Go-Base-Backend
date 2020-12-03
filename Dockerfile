FROM golang:1.15

ENV GIN_MODE=release
ENV PORT=5000

WORKDIR /go/src/go-docker-dev.to

COPY . .
RUN go get
RUN go build

EXPOSE ${PORT}

ENTRYPOINT ["./GinBackend"]