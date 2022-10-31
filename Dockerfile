##
## Build
##
FROM golang:1.19-alpine as dev-env

# Copy application data into image
COPY . /go/src/redis-lock
WORKDIR /go/src/redis-lock

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy only `.go` files, if you want all files to be copied then replace `with `COPY . .` for the code below.
COPY *.go .

# Build our application.
# RUN go build -o /go/src/redis-lock/bin/mullberry-backend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags "all=-N -l" -o /main

##
## Deploy
##
FROM alpine:latest
RUN mkdir /data

COPY --from=dev-env /main ./
EXPOSE 8080
ENTRYPOINT ["./main"]