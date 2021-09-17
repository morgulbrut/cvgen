FROM golang:1.17-alpine

RUN apk add texlive
RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main","-P"]

# CMD [ "go","version" ]
# CMD [ "whomai" ]
# CMD [ "ls"]