FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o synapsis-try

EXPOSE 7000

CMD ./synapsis-try