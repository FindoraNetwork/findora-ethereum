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

# Compile findora platform
FROM golang-builder as findora-platform-builder
ENV PATH=/root/.cargo/bin:$PATH
RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install libc-dev make git curl wget jq ssh python3-pip clang libclang-dev llvm-dev libleveldb-dev musl-tools pkg-config libssl-dev build-essential librocksdb-dev vim ca-certificates -y
RUN pip3 install toml-cli
RUN pip3 install toml
RUN pip3 install web3
RUN mkdir /root/.findora
RUN echo "export OPENSSL_LIB_DIR=/usr/lib/x86_64-linux-gnu" >>/etc/profile
RUN echo "export OPENSSL_INCLUDE_DIR=/usr/include/openssl" >>/etc/profile
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs >/root/rust_install.sh
RUN chmod +x /root/rust_install.sh
RUN /root/rust_install.sh --profile complete -y
RUN /bin/bash -c "source /root/.profile"
RUN /bin/bash -c "source /root/.bashrc"
RUN /bin/bash -c "source /root/.cargo/env"
RUN /bin/bash /root/.cargo/env
RUN curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
RUN cargo install crm
RUN crm best
RUN rustup target add x86_64-unknown-linux-musl
RUN rustup install nightly
RUN rustup target add wasm32-unknown-unknown --toolchain nightly
RUN cargo install cargo-tarpaulin
RUN git clone http://github.com/FindoraNetwork/platform.git -b v0.3.27-release \
  && mv platform /root/ \
  && cd /root/platform  \
  && make build_release_debug

# Compile and Run findora-rosetta
FROM findora-platform-builder as findora-rosetta-builder
RUN apt-get update \
  && apt-get install -y ca-certificates \
  && update-ca-certificates

RUN git clone https://github.com/FindoraNetwork/findora-rosetta.git \
  && cd findora-rosetta \
  && go get \
  && go build \
  && cp findora-rosetta /root/

ENV PORT=8080
ENV RPCURL=http://127.0.0.1:8545
ENV NETWORK=PRINET

COPY ./docker-run.sh /root/
RUN chmod a+x /root/docker-run.sh
 
ENTRYPOINT ["/bin/bash", "/root/docker-run.sh"]
