FROM golang:1.19-buster

WORKDIR /app
COPY dist/scheduler ./
COPY .env ./

CMD ["/app/scheduler"]