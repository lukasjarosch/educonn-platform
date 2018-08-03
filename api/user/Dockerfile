FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/user-apid/user-apid .
ENTRYPOINT ["./user-apid"]
CMD ["./user-apid"]