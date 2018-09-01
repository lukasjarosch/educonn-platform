FROM golang:1.10 as builder

WORKDIR /go/src/github.com/lukasjarosch/educonn-platform

COPY user .

RUN ls -la

RUN cd cmd/user && go get -d -v
RUN cd cmd/user && CGO_ENABLED=0 go build -ldflags "-X github.com/lukasjarosch/educonn-platform/user/internal/platform/config.Version=${VERSION}" -a -installsuffix cgo -o userd main.go


FROM alpine:latest

ARG ServiceName
ARG VERSION

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/lukasjarosch/educonn-platform/user/cmd/userd/userd .
ENTRYPOINT ["./userd"]
CMD ["./userd"]