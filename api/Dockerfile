# Build the binary
FROM golang:1.14 as builder

WORKDIR /app/

# Copy the go source
COPY . .

# Build
RUN make build

FROM alpine:3.12

WORKDIR /app/

COPY --from=builder /app/bin/order-viewer .

EXPOSE 8000

ENTRYPOINT ["/app/order-viewer"]
