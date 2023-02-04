FROM golang:1.19-alpine

ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GO111MODULE=on

COPY . .

RUN apk upgrade --update && \
    apk --no-cache add git && \
    apk add --no-cache gcc musl-dev

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]