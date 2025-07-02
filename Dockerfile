FROM ubuntu:22.04

# Hindari interaksi saat build
ENV DEBIAN_FRONTEND=noninteractive

# Install system dependencies & Golang
RUN apt-get update && apt-get install -y \
    iputils-ping \
    curl \
    git \
    unzip \
    zip \
    nano \
    ca-certificates \
 && curl -OL https://go.dev/dl/go1.22.4.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz \
 && ln -s /usr/local/go/bin/go /usr/local/bin/go \
 && ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt \
 && rm go1.22.4.linux-amd64.tar.gz \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

ENV PATH="/usr/local/go/bin:$PATH"

WORKDIR /app

# Salin semua source ke container
COPY . .

# Build aplikasi sekali saja saat build image
RUN go mod tidy && go build -o server

# Expose port yang sesuai dengan yang dipakai di aplikasi (8010)
EXPOSE 8010

# Jalankan server
CMD ["./server"]