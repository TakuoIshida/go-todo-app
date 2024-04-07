# ベースイメージを指定します。
FROM golang:1.21

# アプリケーションのソースコードをコピーします。
COPY . /app

# 作業ディレクトリを設定します。
WORKDIR /app

# アプリケーションをビルドします。
RUN go build -o myapp

# アプリケーションを実行します。
CMD ["./myapp"]
