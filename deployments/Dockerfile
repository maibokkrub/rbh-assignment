FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./


# todo: create multi stage build

# Build
RUN go build -o /main
EXPOSE 8080

# Run
CMD ["/main"]