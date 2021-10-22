# Argument of the this project
ARG GO_VERSION=1.17.1
ARG SERVER_PORT=8080
ARG GO_REPOSITORY=go-twirp-starter-kit

# Start from golang {GO_VERSION} base image
FROM golang:${GO_VERSION} as build

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/pepeunlimited/$GO_REPOSITORY

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Allow the access to the private libs
ARG SSH_PRIVATE_KEY
RUN mkdir -p ~/.ssh && umask 0077 && echo "${SSH_PRIVATE_KEY}" > ~/.ssh/id_rsa \
	&& git config --global url."git@bitbucket.org:".insteadOf https://bitbucket.org/ \
	&& git config --global url."git@github.com:".insteadOf https://github.com/ \
	&& git config --global url."git@gitlab.com:".insteadOf https://gitlab.com/ \
	&& ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts \
	&& ssh-keyscan github.com >> ~/.ssh/known_hosts \
    && ssh-keyscan gitlab.com >> ~/.ssh/known_hosts

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/server ./cmd/server/

###############################################
######## Start a new stage from scratch #######
###############################################

FROM gcr.io/distroless/static:latest as final

LABEL description="GoTwirpStarterKit for the grpc/twirp server"
# This is intended to be run from the hooks/build script, which sets this
# appropriately based on git tags.
ARG version
LABEL com.pepeunlimited.twirpkit.version=$version

# Add Maintainer Info
LABEL maintainer="Pepe Unlimited <info@pepeunlimited.com>"
#RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /go/bin/server .

# This container exposes $SERVER_PORT to the outside world
EXPOSE $SERVER_PORT

# Run the executable
CMD ["./server"]
