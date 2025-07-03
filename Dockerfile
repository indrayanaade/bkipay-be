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

# Buat direktori kerja
WORKDIR /app

COPY .env .env
# Salin file dependency lebih awal untuk cache build
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

# Salin semua source code setelah download dependency
COPY . .

# Build aplikasi sekali saja saat build image
# Ubah path build sesuai lokasi main.go
# Build aplikasi (otomatis lengkapi go.sum juga)
RUN go mod tidy && go build -v -o server .

# Expose port aplikasi
EXPOSE 8010

# Jalankan server
CMD ["/app/server"]