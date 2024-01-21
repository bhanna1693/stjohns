FROM alpine:latest  

# Set the working directory inside the container
WORKDIR /app

RUN apk --no-cache add ca-certificates

# Copy the binary file from the build context to the container
COPY ./main .

# Copy static files
COPY ./web ./web
COPY ./data ./data

# Expose the port on which the application is running
EXPOSE 8080

# Run the binary
CMD ["./main"]
