# Use the offical Golang image to create a build artifact.
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /benchmark

EXPOSE 8080

CMD ["/benchmark"]