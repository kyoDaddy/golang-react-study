FROM golang:1.18-alpine

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o gomovies ./cmd/api

# 다음 프로세스에는 전달안됨..
#RUN ["chmod", "+x", "/app/setup.sh"]
#RUN source /app/setup.sh

CMD ['./gomovies -dsn="host=172.17.0.3 sslmode=disable user=kyo password=1q2w3e4r dbname=postgres"']