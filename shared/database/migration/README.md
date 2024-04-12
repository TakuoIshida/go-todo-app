# migration toolのinstall
```
brew install golang-migrate
migrate -version
```

# migrationファイル作成

```
migrate create -ext sql -dir ./migration/common -seq create_user_table
migrate create -ext sql -dir ./migration/tenant -seq create_todo_table
```

- -ext:	マイグレーションファイルの拡張子（今回は SQL としました）
- -dir:	マイグレーションファイルを作成する場所（指定したディレクトリが存在しなければ新規作成されます）
- -seq:	マイグレーションファイルの名前

# migrationファイルの記述
SQLを直接書く

# migrationの実行

```
migrate --path ./migration/tenant --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose up
```

```
migrate --path ./migration/common --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose up
```

# rollback
```
migrate --path ./migration/tenant --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose down
```

```
migrate --path ./migration/common --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose down
```

# 参考
https://zenn.dev/farstep/books/f74e6b76ea7456/viewer/4cd440