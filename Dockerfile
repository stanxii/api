# Use Golang version 1.11 as a base image
FROM golang:1.11 as builder

# Opt into Go modules available in Golang version 1.11
ENV GO111MODULE=on

# Set the working directory
WORKDIR /app

# Copy Go modules files so dependencies aren't redownloaded each time the image is built
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Copy all files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

# Copy binary to a lightweight container
FROM scratch
COPY --from=builder /app/main /app/

# Expose container port 8000
EXPOSE 8000

# Start the application
ENTRYPOINT [ "/app/main" ]
