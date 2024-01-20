# Start from a base image with Go installed
FROM golang:alpine as builder

ENV PATH="${PATH}:/go/bin"

# Install curl and unzip (needed to download and extract Tailwind CSS CLI)
RUN apk add --no-cache curl unzip

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files (if using Go modules)
COPY go.mod go.sum ./

# Download Go dependencies (including templ)
RUN go mod download

# Copy the source code into the container
COPY . .

# Download Tailwind CSS Standalone CLI
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Run Tailwind CLI to compile CSS
# Replace 'input.css' and 'output.css' with your actual CSS file paths
RUN tailwindcss -i ./web/static/styles/tailwind-input.css -o ./web/static/styles/tailwind-output.css --minify

# Generate templates using templ
# Replace this with the actual command for your templ usage
RUN go install github.com/a-h/templ/cmd/templ@latest && templ generate

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Use a small base image for the final container
FROM alpine:latest  

# Add ca-certificates in case your application makes external HTTPS calls
RUN apk --no-cache add ca-certificates

# Set the working directory inside the final container
WORKDIR /app

# Copy the binary and other necessary files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/web ./web

# Adjust the file paths according to your project structure
COPY data /app/data

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["./main"]
