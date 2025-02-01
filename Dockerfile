FROM golang:1.23 AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o yagokv .

FROM scratch

WORKDIR /root/
COPY --from=builder /app/yagokv .

EXPOSE 8080
CMD ["./yagokv"]