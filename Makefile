.PHONY: help build up down shell

#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"

#? build: アプリケーションのセットアップ
build:
	docker compose build --no-cache
	[ -f ./server/.env ] || cp ./server/.env.development ./server/.env

#? up: アプリケーションの起動
up:
	docker compose up -d

#? down: アプリケーションの停止
down:
	docker compose down

#? shell-http: HTTP サーバーのシェルを起動
shell:
	docker compose exec -it api bash

.PHONY: api wire

#? api: OpenAPI からコードを生成
api:
	docker compose run --rm api bash -c "oapi-codegen -package api /api/api/openapi.json > /api/server/adapter/api/interface.go"

#? wire: HTTP サーバーの依存関係を自動生成
wire:
	docker compose run --rm api bash -c "cd /api/server/cmd/server/wire && wire gen && mv ./wire_gen.go /api/server/cmd/server/commands/wire.go"

.PHONY: migrate migrate-down sql-gen

#? migrate: データベースの構造をマイグレート
migrate:
	docker compose run --rm api bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db:5432/db?sslmode=disable up"

#? migrate-down: データベースの構造を初期化
migrate-down:
	docker compose run --rm api bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db:5432/db?sslmode=disable down -all"

#? sql-gen: SQL クエリから Go コードを生成
sql-gen:
	docker compose run --rm api bash -c "cd server && sqlc generate"
