# goのイメージをDockerHubから流用する(Alpine Linux)
FROM golang:latest
# ログのタイムゾーンを指定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
# コンテナ内の作業ディレクトリを指定
WORKDIR /app
# ソースコードをコンテナ内にコピー
COPY . ./
# /app/go.modに記載された依存関係の解決＋必要なパッケージのダウンロードを実行
RUN go mod tidy
# Airのバイナリをインストール
RUN go install github.com/cosmtrek/air@latest
# コンテナの公開するポートを指定
EXPOSE 5050
# 起動時のコマンド(airを使用するため)
CMD ["air", "-c", ".air.toml"]