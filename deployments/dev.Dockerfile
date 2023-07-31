# syntax=docker/dockerfile:1

#
# using official dockerfile
FROM golang:1.18

# Expected all files to be mounted at /app
WORKDIR /app

# Download Go modules
RUN go mod download
EXPOSE 8080

# Run
CMD ["go run main.go"]