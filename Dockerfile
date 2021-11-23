FROM golang:1.16-alpine

COPY . /app

WORKDIR /app

RUN go build .

CMD /app/family_directory