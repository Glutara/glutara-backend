# Build stage
FROM golang:1.20 as build

# Set working directory inside the container
WORKDIR /go/src/app

# Copy everything from the current directory into the container's working directory
COPY . .

# Disable CGO to ensure Go binaries are statically linked
ENV CGO_ENABLED=0

# Fetch dependencies
RUN go get -d -v ./...

# Install the dependencies
RUN go install -v ./...

# Build the Go application binary
RUN go build -v -o go-app

# Run stage
FROM alpine:3.11

# Set working directory inside the container
WORKDIR /app

# Copy the built binary from the build stage to the final stage
COPY --from=build /go/src/app/go-app .

# Copy the .env file from the build context into the container
COPY .env .

# Set the PORT environment variable to 8080
ENV PORT=8080

# Define the default command to run the application when the container starts
CMD ["./go-app"]
