package tpl

func EchoServerDockerFileTemplate() []byte {
	return []byte(`FROM golang:1.24.5-bookworm AS builder
# Copy source code
COPY ./ /app/
WORKDIR /app
# Download dependencies and build
RUN CGO_ENABLED='0' go build -o .{{ .AppName }}

# --- Runtime stage ---
FROM debian:bookworm-slim

# Install only runtime dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*
# Copy binary from builder
COPY --from=builder /app/.{{ .AppName }} /usr/local/bin/.{{ .AppName }}
# Create empty config file
RUN touch /root/..{{ .AppName }}.toml
EXPOSE 8090
CMD [".{{ .AppName }}", "server"]
`)
}
