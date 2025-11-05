FROM golang:1.24.2-alpine AS builder

WORKDIR /app
RUN go mod init github.com/orbit-alliance/orbit-backend
COPY . .
RUN go mod tidy
RUN go build -o app ./cmd
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/app .
COPY .env .env
ENV $(cat .env | xargs)
EXPOSE 8888
CMD ["./app"]
