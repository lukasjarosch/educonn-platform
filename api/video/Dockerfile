FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app
ADD ./cmd/video-apid/video-apid .
ENTRYPOINT ["./video-apid"]
CMD ["./video-apid"]