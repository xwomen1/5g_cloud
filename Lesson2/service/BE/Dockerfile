FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY . ./
COPY data/ data/
COPY models/ models
COPY *.go ./

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
