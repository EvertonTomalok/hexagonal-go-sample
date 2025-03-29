#!/bin/bash

# Set version of migrate
VERSION="v4.16.0"  # Replace with latest version

# Detect OS
ios="$(uname | tr '[:upper:]' '[:lower:]')"
if [[ "$ios" == "linux" ]]; then
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - && \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list && \
    apt-get update && \
    apt-get install -y migrate
elif [[ "$ios" == "darwin" ]]; then
    brew install golang-migrate
else
    echo "Unsupported OS: $ios"
    exit 1
fi