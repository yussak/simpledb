FROM golang:1.24

WORKDIR /app

# 依存はまだ無いが、増えてきたらここでキャッシュを効かせる
COPY go.mod ./
RUN go mod download

COPY . .

CMD ["go", "run", "."]
