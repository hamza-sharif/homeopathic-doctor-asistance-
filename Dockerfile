
FROM golang:1.21-alpine

RUN apk add build-base

RUN mkdir /app

WORKDIR /app

ADD . /app

RUN GO111MODULE=on

COPY go.mod .

COPY go.sum .

RUN go build  cmd/main.go

CMD ["./main"]


#RUN apt-get update

#RUN apt-get install ca-certificates -y

#COPY bin/homeopathic-doctor-assistant-server /usr/local/bin/

#CMD ["/usr/local/bin/homeopathic-doctor-assistant-server"]
