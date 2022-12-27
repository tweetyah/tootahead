FROM golang:1.19-buster

WORKDIR /app
COPY go.mod ./
COPY main.go ./
RUN go get -u
RUN go build -o scheduler .

CMD ["/app/scheduler"]