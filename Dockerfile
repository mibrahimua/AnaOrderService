# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

ADD go.mod .

ADD go.sum .

RUN go mod download

# Copy the Go application files to the container
COPY . .

# Build the Go application
RUN go build -o app .

# Run the application when the container starts
CMD ["./app"]