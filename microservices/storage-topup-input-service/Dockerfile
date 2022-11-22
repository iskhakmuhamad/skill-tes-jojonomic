FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o storage-topup-input-server ./cmd

FROM alpine

WORKDIR /app

COPY --from=builder /app/storage-topup-input-server /usr/bin/

## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY .env .

EXPOSE 7008

CMD /wait && storage-topup-input-server