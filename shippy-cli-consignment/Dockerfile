FROM alpine:latest

RUN mkdir -p /app
WORKDIR /app

RUN apk update && apk upgrade && apk add --no-cache git go

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-cli-consignment

CMD ["./shippy-cli-consignment"]
