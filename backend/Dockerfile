# Use official Golang image matching the module version
FROM golang:1.23.4

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application with a unique binary name to avoid collisions with local folders
RUN go build -o recipe-api .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./recipe-api"]
