# Build stage
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application, outputting the binary named 'main'
RUN go build -o main . && ls -la


# Run stage (using a lightweight image)
FROM debian:latest

# Set the working directory inside the runtime container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Set executable permission for the binary
RUN chmod +x ./main

FROM scratch
ADD main /
CMD ["./main"]
