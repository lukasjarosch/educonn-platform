FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/lessond/lessond .
ENTRYPOINT ["./lessond"]
CMD ["./lessond"]