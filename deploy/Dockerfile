FROM golang:1.21 AS builder

COPY ./  /go/src/app/

WORKDIR /go/src/app

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app ./main.go

# Copy the executable into alpine image which results in smaller image
FROM alpine:latest  

RUN mkdir -p /go/src/app

COPY --from=builder /go/bin/app /go/bin/app

WORKDIR /go/src/app

EXPOSE 5000

ENTRYPOINT ["/go/bin/app"]
