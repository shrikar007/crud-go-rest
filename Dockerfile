
FROM golang:latest AS builder
ADD . /app
WORKDIR /app

RUN  go build -a -o /main .


FROM alpine

COPY --from=builder /main /app/main
RUN chmod +x /app/main
ENTRYPOINT ["/app/main"]

EXPOSE 8082
