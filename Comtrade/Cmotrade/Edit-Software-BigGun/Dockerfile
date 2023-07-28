# Use the official Golang base image
FROM golang:1.20.6 AS development

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application as a Linux executable
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Expose the port your application listens on
EXPOSE 8080

# Command to run your Go application
CMD ["go", "run", "main.go", "--start-service"]
