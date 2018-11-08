FROM golang:alpine AS builder
ENV WORKPATH=/go/src/github.com/jordyv/wait-for
COPY . $WORKPATH
RUN apk add --update --no-cache curl make git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN cd $WORKPATH && dep ensure -v
RUN cd $WORKPATH && make build

FROM alpine
COPY --from=builder /go/src/github.com/jordyv/wait-for/dist/wait-for /usr/local/bin/wait-for
ENTRYPOINT [ "/usr/local/bin/wait-for" ]
CMD [ "help" ]
