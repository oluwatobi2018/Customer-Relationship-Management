FROM golang:latest as builder
LABEL maintainer="qwertmax@gmail.com"

WORKDIR /app
COPY . .
RUN ["./setup.sh"]
RUN go build -o app .

FROM alpine:latest
LABEL maintainer="qwertmax@gmail.com"
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app
COPY --from=builder /app/app .

ARG VERSION
ENV VERSION=$VERSION
CMD ["/app/app"]

EXPOSE 80
