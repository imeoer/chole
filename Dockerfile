FROM golang:latest
RUN go get -u github.com/imeoer/chole
EXPOSE 1025-65535
WORKDIR /go/src/github.com/imeoer/chole
CMD ["chole", "-s"]