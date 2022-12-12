FROM golang:1.17-alpine

WORKDIR /home/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o api-umachay

EXPOSE 8000

CMD ["./api-umachay"]
