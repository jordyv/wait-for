FROM golang:alpine AS builder
WORKDIR /go/src/github.com/jordyv/wait-for
COPY . .
RUN apk add make git
RUN wget -O - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN make build

FROM alpine
COPY --from=builder /go/src/github.com/jordyv/wait-for/dist/wait-for /usr/local/bin/wait-for
ENTRYPOINT [ "/usr/local/bin/wait-for" ]
CMD [ "help" ]
