FROM golang

WORKDIR /app

COPY main.go .
COPY go.mod .
COPY README.md .

RUN go build -o main

CMD ["./main"]
