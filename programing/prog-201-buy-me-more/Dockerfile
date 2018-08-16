FROM golang:1.10.3-alpine3.7 as builder

RUN apk add --no-cache git gcc libc-dev curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/jusmistic/programming/prog-201-buy-me-more/src

COPY ./src/Gopkg.toml ./src/Gopkg.lock ./
RUN dep ensure --vendor-only
COPY ./src .
RUN GOOS=linux go build -v -o /go/bin/buy-me-more .

FROM alpine:3.7

WORKDIR /app

EXPOSE 9000

RUN chown -R 1337:1337 .
USER 1337

COPY --from=builder /go/bin/buy-me-more .
COPY ./src/tmpl tmpl

ENTRYPOINT [ "./buy-me-more" ]
