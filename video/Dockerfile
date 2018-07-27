FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/videod/videod .
ENTRYPOINT ["./videod"]
CMD ["./videod"]