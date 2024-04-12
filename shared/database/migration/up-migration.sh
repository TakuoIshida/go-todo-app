#!bin/bash
migrate --path . --database 'postgresql://postgres:postgrespw@localhost:25432/local?sslmode=disable' -verbose up