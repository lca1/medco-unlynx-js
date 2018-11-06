FROM golang:1.11

COPY build/build.sh /usr/local/bin/

# install gopherjs
RUN go get -u github.com/gopherjs/gopherjs && \
    cd "$GOPATH"/src/github.com/gopherjs/gopherjs && \
    go get -u ./... && \
    go install ./... && \
    mkdir -p /medco-unlynx-js-src "$GOPATH"/src/github.com/lca1 && \
    ln -s /medco-unlynx-js-src "$GOPATH"/src/github.com/lca1/medco-unlynx-js && \
    chmod +x /usr/local/bin/build.sh

VOLUME /medco-unlynx-js-src
ENTRYPOINT build.sh
