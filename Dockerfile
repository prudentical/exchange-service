FROM golang:1.22.2-alpine3.19 as build
WORKDIR /app

ENV GOPROXY https://goproxy.io

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o exchange-service ./cmd/application/


FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app

COPY --from=build /app/internal/database/migrations ./internal/database/migrations
COPY --from=build /app/config.yml .
COPY --from=build /app/exchange-service .

EXPOSE 8001

CMD ["./exchange-service"]