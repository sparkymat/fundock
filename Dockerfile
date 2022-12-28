FROM golang:1.19-alpine AS builder

RUN apk update && apk add make

COPY . /app/

WORKDIR /app
RUN go generate ./...
RUN make fundock

FROM alpine:3

COPY --from=builder /app/fundock /bin/fundock

WORKDIR /app
COPY public /app/public
COPY migrations /app/migrations

CMD ["/bin/fundock"]
