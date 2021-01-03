FROM ubuntu:20.04

ENV GO_INSTALL_DIR=/usr/local/go

RUN set -ex; \
    apt-get update && apt-get install -y vim curl wget; \
    mkdir -p "$HOME/go/src" "$HOME"/go/pkg "$HOME"/go/bin; \
    mkdir -p /app /app/cmd/web /app/ui/html /app/ui/static/css /app/ui/static/img /app/ui/static/js

# Install Go
WORKDIR /usr/local
ENV GO_TAR=https://golang.org/dl/go1.15.6.linux-amd64.tar.gz
RUN set -ex; \
    wget -c "$GO_TAR" -O - | tar -xz; \
    echo "GOROOT=$GO_INSTALL_DIR" >> /root/.bashrc; \
    echo "GOPATH=$HOME/go" >> /root/.bashrc; \
    echo "GOBIN=$GO_INSTALL_DIR/bin" >> /root/.bashrc; \
    echo "PATH=$PATH:$GO_INSTALL_DIR/bin" >> /root/.bashrc

WORKDIR /app
COPY cmd/web/*.go cmd/web/
COPY ui/html/*.tmpl ui/html/
COPY ui/static/css/main.css ui/static/css/main.css
COPY ui/static/img/* ui/static/img/
COPY ui/static/js/main.js ui/static/js/main.js

RUN set -ex; \
    /usr/local/go/bin/go build -o snippet cmd/web/*.go 
    
EXPOSE 4000

CMD ["./snippet"]