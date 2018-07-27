FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/userd/userd .
ENTRYPOINT ["./userd"]
CMD ["./userd"]