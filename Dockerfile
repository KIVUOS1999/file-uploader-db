FROM golang:1.23.2

LABEL maintainer="souvik souviksarkar.ronnie@gmail.com"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY configs/.env ./configs/.env
COPY file-uploader-db .

EXPOSE 8001

CMD ["./file-uploader-db"]