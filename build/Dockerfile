FROM medco/unlynx:v1.2a-alpha as build

# install gopherjs
RUN go get -u github.com/gopherjs/gopherjs && \
    cd "$GOPATH"/src/github.com/gopherjs/gopherjs && \
    go get -u ./... && \
    go install ./... && \
    mkdir -p /medco-unlynx-js-src "$GOPATH"/src/github.com/lca1 && \
    ln -s /medco-unlynx-js-src "$GOPATH"/src/github.com/lca1/medco-unlynx-js

# set up build script
COPY build/build.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/build.sh

VOLUME /medco-unlynx-js-src
ENTRYPOINT build.sh
