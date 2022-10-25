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
FROM alpine:3.16
RUN apk --no-cache add ca-certificates
WORKDIR /app
# Copy built binary to runner container
COPY --from=build /go/src/app/bin /app
# Declare a volume mount where we expect the markdown files to be
VOLUME ["/markdown"]

# Copy static files to static directory
COPY ./static /app/static
COPY ./templates /app/templates

ENV GIN_MODE=release

EXPOSE 9090

CMD [ "/app/server" ]