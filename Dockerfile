
FROM golang:latest AS builder
ADD . /app
WORKDIR /app

RUN  go build -a -o /main .


FROM alpine
RUN apk add --no-cache ca-certificates \
	&& update-ca-certificates \
    # cleanup
    && rm -rf /var/cache/apk/*


COPY --from=builder /main .
RUN chmod +x main
EXPOSE 8082
ENTRYPOINT ["/main"]


