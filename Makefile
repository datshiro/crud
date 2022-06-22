include .env
export

.PHONY: help
help: 
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

docker/up:
	docker-compose -f docker/docker-compose.yaml up  -d

docker/down:
	docker-compose -f docker/docker-compose.yaml down

app/run:
	go run cmd/server/main.go

setup:
	@scripts/setup.sh

# Basic commands: up/down/drop/force\ <version>
db/%:
	@scripts/migrate.sh $*

db/gen:
	sqlboiler psql --wipe --config ./db/sqlboiler.toml --add-soft-deletes 
