FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static-debian11 AS runner
COPY --from=builder /app/app /app/app
ENTRYPOINT ["/app/app"]
