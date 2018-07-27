FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/maild/maild .
ENTRYPOINT ["./maild"]
CMD ["./maild"]