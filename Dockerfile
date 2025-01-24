# Use an official Go runtime as a parent image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Download dependencies using go mod download
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
