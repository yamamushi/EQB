# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/yamamushi/EQB

# Build the EQB command inside the container.
RUN cd /go/src/github.com/yamamushi/EQB && go get -v ./... && go build -v ./... && go install

# Run the EQB command by default when the container starts.
ENTRYPOINT /go/bin/EQB

# Set the working directory to /go/src/github.com/yamamushi/EQB
WORKDIR /go/src/github.com/yamamushi/EQB
