FROM golang:1.12 as build

# set up build script
COPY build/build.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/build.sh

VOLUME /src
ENTRYPOINT build.sh
