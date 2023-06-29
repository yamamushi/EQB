# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/yamamushi/EQB

# Build the EQB command inside the container.
RUN go install github.com/yamamushi/EQB@latest

# Run the EQB command by default when the container starts.
ENTRYPOINT /go/bin/EQB

# Set the working directory to /go/src/github.com/yamamushi/EQB
WORKDIR /go/src/github.com/yamamushi/EQB
