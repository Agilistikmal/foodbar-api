FROM golang:alpine

RUN apk update

WORKDIR /app

COPY . .

RUN go mod tidy 

RUN go build -o binary ./cmd/foodbar

ENTRYPOINT ["/app/binary"]