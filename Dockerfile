FROM golang:latest

EXPOSE 8000

WORKDIR /go/src/github.com/RadicalPixels/server

RUN set -x && wget https://dist.ipfs.io/go-ipfs/v0.4.14/go-ipfs_v0.4.14_linux-amd64.tar.gz -O /tmp/go-ipfs.tar.gz && tar xvfz /tmp/go-ipfs.tar.gz && cp go-ipfs/ipfs /usr/local/bin

RUN ipfs init
RUN ipfs key gen default --type=rsa --size=2048

# temp
RUN go get github.com/ethereum/go-ethereum
RUN go get github.com/patrickmn/go-cache
RUN rm -rf /go/src/github.com/RadicalPixels/server/vendor
COPY . /go/src/github.com/RadicalPixels/server

ENTRYPOINT make start
