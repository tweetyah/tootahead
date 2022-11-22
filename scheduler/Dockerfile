FROM golang:1.19-buster

WORKDIR /app
COPY go.* ./
RUN go mod download

COPY . ./
ENV GOOS="linux"
ENV GOARCH="amd64"
RUN go build -v -o scheduler

CMD ["/app/scheduler"]