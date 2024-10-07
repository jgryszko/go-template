# Use the latest Ubuntu base image
FROM ubuntu:24.04

# Set environment variables for non-interactive installation
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    git \
    vim \
    unzip \
    software-properties-common \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install the latest Go version
RUN curl -OL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz \
    && rm go1.22.0.linux-amd64.tar.gz

# Set Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH="/go"

# Install useful Go tools and packages
RUN go install golang.org/x/tools/gopls@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest

# Set up workspace
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

# Expose debugging port
EXPOSE 2345

# Default command to start the shell
CMD ["/bin/bash"]
