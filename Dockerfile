FROM golang:1.23-alpine as builder



WORKDIR /app

COPY go.mod . 

RUN go mod download

COPY . .


RUN go build cmd/app/main.go


FROM alpine:3


COPY --from=builder /app/main .
COPY --from=builder /app/config.yml .

EXPOSE 50051
EXPOSE 8080
EXPOSE 9100

ENTRYPOINT ["./main"]