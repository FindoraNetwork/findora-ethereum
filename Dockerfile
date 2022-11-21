# Compile golang
FROM ubuntu:20.04 as golang-builder

RUN mkdir -p /app \
  && chown -R nobody:nogroup /app
WORKDIR /app

RUN apt-get update && apt-get install -y curl make gcc g++ git
ENV GOLANG_VERSION 1.16.8
ENV GOLANG_DOWNLOAD_SHA256 f32501aeb8b7b723bc7215f6c373abb6981bbc7e1c7b44e9f07317e1a300dce2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
  && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
  && tar -C /usr/local -xzf golang.tar.gz \
  && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Compile and Run findora-rosetta
FROM golang-builder as findora-rosetta-builder

RUN apt-get update \
  && apt-get install -y ca-certificates \
  && update-ca-certificates

RUN git clone https://github.com/FindoraNetwork/findora-rosetta.git \
  && cd findora-rosetta \
  && go get \
  && go build \
  && cp findora-rosetta /root/

ENV PORT=8080
ENV RPCURL=http://127.0.0.1
ENV MODE=ONLINE
ENV NETWORK=PRINET


CMD ["/root/findora-rosetta", "run"]
