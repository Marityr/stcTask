build:
	rm -rf build && mkdir build && go build -o build/app_bonus -v cmd/main.go
  
run:
	go run server/cmd/main.go /server

.PHONY: sw
sw:
	swag init -g server/cmd/main.go

test:
	go test

# .PHONY: migratecreate
# migratecreate:
# 	migrate create -ext sql -dir ./schemes -seq init 

# .PHONY: migrateup
# migrateup:
# 	migrate -path ./schemes -database 'postgres://arch:123456@localhost:5432/newbonus?sslmode=disable' up 

# .PHONY: migratedown
# migratedown:
# 	migrate -path ./schemes -database 'postgres://arch:123456@localhost:5432/newbonus?sslmode=disable' down
