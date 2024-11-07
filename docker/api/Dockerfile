# builder
FROM golang:1.23.2-bullseye

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "cmd/main.go"]

# runtime
# FROM golang:1.23.2-bullseye

# WORKDIR /

# COPY --from=builder /api-server .

# EXPOSE 8080

# CMD ["/api-server"]

