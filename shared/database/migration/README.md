# migration toolのinstall
```
brew install golang-migrate
migrate -version
```

# migrationファイル作成
pathはshared/database/migration直下
```
migrate create -ext sql -dir . -seq create_user_table
```

- -ext:	マイグレーションファイルの拡張子（今回は SQL としました）
- -dir:	マイグレーションファイルを作成する場所（指定したディレクトリが存在しなければ新規作成されます）
- -seq:	マイグレーションファイルの名前

# migrationファイルの記述
SQLを直接書く

# migrationの実行

```
sh up-migration.sh
```

# rollback
```
sh down-migration.sh
migrate --path ./migration --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose down
```

# 参考
https://zenn.dev/farstep/books/f74e6b76ea7456/viewer/4cd440