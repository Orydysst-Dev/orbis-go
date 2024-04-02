docker:
	KO_DOCKER_REPO=ko.local ko build -B --platform=linux/$$(go env GOARCH) ./cmd/orbisd

.PHONY: build
build:
	go build -o build/orbisd ./cmd/orbisd 

run:
	docker-compose -f demo/acp/compose.yaml down -v
	docker-compose -f demo/acp/compose.yaml up


rund:
	docker-compose -f demo/compose.yaml down -v
	docker-compose -f demo/compose.yaml up -d

stop:
	docker-compose -f demo/compose.yaml down -v

## Don't spin up sourcehub, use the one that's already running on the host.
run-no-sourcehub:
	docker-compose -f demo/compose.yaml down -v
	docker-compose -f demo/compose.yaml up
