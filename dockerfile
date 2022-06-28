# Build Stage
# First pull Golang image
FROM golang:1.17-alpine as build-env
 
# Set environment variable
ENV APP_NAME nmapdojo-report
ENV CMD_PATH cmd/report/main.go
 
# Copy application data into image
COPY . $GOPATH/src/nmapdojo/$APP_NAME
WORKDIR $GOPATH/src/nmapdojo/$APP_NAME
 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/nmapdojo/$APP_NAME/$CMD_PATH

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip

# Run Stage
FROM alpine:3.14
 
# Set environment variable
ENV APP_NAME nmapdojo-report
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
 
# Expose application port
EXPOSE 8081
 
# Start app
CMD ./$APP_NAME

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip