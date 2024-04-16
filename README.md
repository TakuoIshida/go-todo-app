# go-todo-app

# mockの生成
```
# moq -pkg パッケージ名 -out 出力先ファイルパス 対象のインターフェース名
moq -pkg todo_test -out ./test/todo_mock.go . ITodoUsecase ITodoService ITodoRepository
```