# Use Golang base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o todo ./cmd/todo

# Expose service port
EXPOSE 8080

# Run the binary
CMD ["./todo"]
