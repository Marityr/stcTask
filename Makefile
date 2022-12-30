.PHONY: srart
start:
	docker-compose build && docker-compose up
  
run:
	cd server/ && go run cmd/main.go

.PHONY: sw
sw:
	cd server/ && swag init -g cmd/main.go

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
