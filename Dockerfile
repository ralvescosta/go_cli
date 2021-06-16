FROM golang:1.16.0-stretch

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

RUN go install github.com/spf13/cobra/cobra@latest

EXPOSE 3333

CMD ["tail", "-f", "/dev/null"]