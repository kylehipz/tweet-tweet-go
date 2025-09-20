FROM golang:1.25-alpine3.22 as builder

ARG APP_SERVICE=users
ENV SERVICE=${APP_SERVICE}

WORKDIR /app

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o ./bin/${SERVICE}-service ./cmd/services/${SERVICE}

# Final image
FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/bin/${SERVICE}-service .

CMD ./${SERVICE}-service
