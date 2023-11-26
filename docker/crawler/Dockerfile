# Stage 1: Build Golang application
FROM golang:latest AS builder

WORKDIR /app

COPY src/crawler .


RUN ls -la

# Add any additional dependencies or build steps for your Golang application here
# For example:
RUN go get -u crawler

# Build the Golang application
RUN go build -o crawler

# Stage 2: Final image with ChromeDriver and Golang application
FROM zenika/alpine-chrome:with-chromedriver

# # Install necessary dependencies
# RUN apt-get update && \
#     apt-get install -y \
#     wget \
#     unzip \
#     libnss3 \
#     libglib2.0-0 \
#     libfontconfig1 \
#     libxrender1 \
#     libxkbcommon-x11-0 \
#     libx11-xcb1 \
#     libxkbcommon0 \
#     libxss1 \
#     libasound2 \
#     && rm -rf /var/lib/apt/lists/*

# # Download and install ChromeDriver
# RUN wget -O /tmp/chromedriver.zip https://chromedriver.storage.googleapis.com/$(wget -q -O - https://chromedriver.storage.googleapis.com/LATEST_RELEASE)/chromedriver_linux64.zip && \
#     unzip /tmp/chromedriver.zip -d /usr/local/bin/ && \
#     rm /tmp/chromedriver.zip && \
#     chmod +x /usr/local/bin/chromedriver

WORKDIR /app

# # Copy the built Golang application from the previous stage
COPY --from=builder /app/crawler .

# # Expose any necessary ports
# # EXPOSE 8080

# # Set the entry point for the application
CMD ["./crawler"]
