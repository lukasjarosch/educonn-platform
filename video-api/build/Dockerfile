FROM golang:1.10 as builder

MAINTAINER lukas.jarosch@mail.com

# dep
RUN wget https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -O /usr/bin/dep && chmod +x /usr/bin/dep

# gopath
RUN mkdir -p /go/src/github.com/lukasjarosch/educonn-platform/video-api/
WORKDIR /go/src/github.com/lukasjarosch/educonn-platform/video-api/

# vendoring
COPY video-api/Gopkg.lock user-api/Gopkg.toml ./
RUN dep ensure -vendor-only -v

# build
COPY video-api/ ./
RUN scripts/build.sh

# ----------

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/lukasjarosch/educonn-platform/video-api/cmd/video-api/video-api .
ENTRYPOINT ["./video-api"]
CMD ["./video-api"]
