FROM golang:1.10 as builder

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/lukasjarosch/educonn-platform
COPY . .

RUN cd user && dep ensure -v
RUN user/scripts/build.sh

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/lukasjarosch/educonn-platform/user/cmd/user/user .
ENTRYPOINT ["./userd"]
CMD ["./userd"]