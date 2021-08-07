FROM golang:1.16

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . .

RUN go test -v ./...

RUN go build -o hackpoints

ENTRYPOINT [ "./hackpoints" ]
