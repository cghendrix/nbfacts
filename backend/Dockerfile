FROM golang:1.19.3-alpine3.17

ARG DB_USER
ENV DB_USER=$DB_USER
ARG DB_PASS
ENV DB_PASS=$DB_PASS

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o server .
CMD ["/app/server"]