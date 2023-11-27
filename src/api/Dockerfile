# TODO: Add Multi Stage Strategy to decrease image size, preferably using Alpine

# Stage 1: Build Golang application
FROM golang:latest AS builder

ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY
ARG AWS_DEFAULT_REGION
ARG SQS_QUEUE_URL

ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY
ENV AWS_DEFAULT_REGION=$AWS_DEFAULT_REGION
ENV SQS_QUEUE_URL=$SQS_QUEUE_URL

WORKDIR /app

COPY . .

# Add any additional dependencies or build steps for your Golang application here
# For example:
RUN go get -u api

# Build the Golang application
RUN go build -o api


# # Expose any necessary ports
EXPOSE 8080

# # Set the entry point for the application
CMD ["./api"]
