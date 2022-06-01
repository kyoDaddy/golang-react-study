# AWS ECR 에서의 이미지 크기 감소는 비용 감소로 이뤄지니까 최적화 수정 적용
FROM  golang:1.18-buster as builder
WORKDIR /tmp/tiny-golang-image
COPY . .

RUN go mod tidy \
    && go get -u -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o main cmd/api/*.go

FROM scratch
COPY --from=builder /tmp/tiny-golang-image /

ENV GO_PROFILE=dev
CMD ["/main"]