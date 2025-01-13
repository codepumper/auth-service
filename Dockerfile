# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o auth-service .

# Expose the application port
EXPOSE 50051

# Command to run the service
CMD ["./auth-service"]
