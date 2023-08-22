FROM golang:latest as builder

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main


FROM alpine:latest
RUN apk --no-cache add ca-certificates openssl

WORKDIR /root
COPY --from=builder /app/main ./
COPY ./data.json ./

EXPOSE 80

CMD ["./main"]
