#Postgres-Docker環境構築
postgres: 
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

#ポスグレDB作成
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bankapp

#ポスグレDB削除
dropdb:
	docker exec -it postgres12 dropdb bankapp

#マイグレ

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose down

#SQLC CRUD用Goコード自動作成
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc