FROM golang:alpine AS builder
RUN apk --no-cache add build-base git glide
WORKDIR /go/src/github.com/dalmarcogd/twitter-reporter/twitter-reporter-api
ADD glide* ./
RUN glide install
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/dalmarcogd/twitter-reporter/twitter-reporter-api/app /app/
EXPOSE 8000
ENTRYPOINT ./app