FROM alpine:3.21
RUN apk --no-cache add ca-certificates
COPY gitchangelog /gitchangelog
EXPOSE 8081
ENTRYPOINT ["/gitchangelog"]
