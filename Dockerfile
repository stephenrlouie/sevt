# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine

ENV ZEUS_TOKEN=8dcccc00
ENV USER=stelouielocal

# Copy the local package files to the container's workspace.
ADD src/github.com/stephenrlouie/web /go/src/github.com/stephenrlouie/web

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

RUN apk add --no-cache git \
    && go get github.com/CiscoZeus/go-zeusclient \
    && apk del git

RUN go install github.com/stephenrlouie/web

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/web

# Document that the service listens on port 8080.
EXPOSE 8080
