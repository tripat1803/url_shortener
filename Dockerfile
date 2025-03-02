FROM golang:latest

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./.env ./
COPY ./main.go ./
COPY ./config ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

CMD ["/build"]