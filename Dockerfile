FROM golang:1.16-alpine

ADD . /go/src/event-api
WORKDIR /go/src/event-api

RUN go build -o /event-api

EXPOSE 8080

CMD [ "/event-api" ]