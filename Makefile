.PHONY: build

build:
	sam build --no-cached

start:
	sam local start-api --docker-network network

migrate-create:
	migrate create -ext sql -dir ${PWD}/pkg/infrastructure/database/migrations -seq ${class}
# examle: make migrate-create class=users

migrate-run:
	docker run -v ${PWD}/pkg/infrastructure/database/migrations:/migrations --network=network migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(go-mysql)/example" up

migrate-rollback:
	docker run -itv ${PWD}/pkg/infrastructure/database/migrations:/migrations --network=network migrate/migrate -path=/migrations/ -database "mysql://root:password@tcp(go-mysql)/example" down
