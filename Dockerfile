FROM golang:alpine
RUN mkdir /app

COPY . /app
COPY /cmd /app

WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]