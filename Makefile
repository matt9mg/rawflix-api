.PHONY: build
build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/rawflix

.PHONY: up
up: build
	@docker compose up -d
	$(MAKE) logs

.PHONY: restart
restart: build
	@docker compose stop app
	@docker compose up app -d
	$(MAKE) logs

.PHONY: stop
stop:
	@docker compose stop

.PHONY: logs
logs:
	@docker logs -f rawflix_app

.PHONY: enter
enter:
	@docker exec -it rawflix_app bash