# Use official Golang image
FROM golang:1.22

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 4000

# Run the application
CMD ["./main"]
