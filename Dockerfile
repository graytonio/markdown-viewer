# Init builder container
FROM golang:alpine AS build
# Install dependencies
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
# Copy source to contianer
COPY . .
RUN go mod tidy
# Build with optimize flags
RUN go build -ldflags="-s -w" -o ./bin/server .
# ================================================================
# Init runner container
FROM nginx:alpine
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
# Copy built binary to runner container
COPY --from=build /go/src/app/bin /go/bin
# Replace default nginx.conf with custom one
COPY nginx.conf /etc/nginx/nginx.conf 
# Declare a volume mount where we expect the markdown files to be
VOLUME ["/markdown"]

ENV GIN_MODE=release
ENV PORT=9090

ENTRYPOINT /go/bin/server