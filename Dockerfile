# syntax=docker/dockerfile:1

FROM golang:1.21.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go mod tidy
RUN go build cmd/main.go

EXPOSE 4001

# Run
CMD ["./main"]