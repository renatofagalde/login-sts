#build stage
FROM golang:1.22.4-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build .

## RUN STAGE
#FROM scratch
FROM golang:1.22.4-alpine3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env ./app.env
COPY start.sh ./start.sh
COPY wait-for.sh ./wait-for.sh

EXPOSE 8080
CMD ["/app/main"]