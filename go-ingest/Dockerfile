FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main .

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/templates /app/templates

EXPOSE 8080
CMD ["./main"]
