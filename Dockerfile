FROM ubuntu:20.04

ENV GO_INSTALL_DIR=/usr/local/go

RUN set -ex; \
    apt-get update && apt-get install -y curl iproute2 iputils-ping vim wget; \
    apt-get clean; \
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

# Copy snippetbox resources to the container
WORKDIR /app
COPY cmd/web/*.go cmd/web/
COPY ui/html/*.tmpl ui/html/
COPY ui/static/css/main.css ui/static/css/main.css
COPY ui/static/img/* ui/static/img/
COPY ui/static/js/main.js ui/static/js/main.js

RUN set -ex; \
    /usr/local/go/bin/go build -o snippetbox cmd/web/*.go

COPY docker_entrypoint.sh /docker_entrypoint.sh
ENTRYPOINT ["/docker_entrypoint.sh"]
    
EXPOSE 4000

CMD ["./snippetbox"]