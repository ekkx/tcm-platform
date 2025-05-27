#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"
.PHONY: help

#? build: アプリケーションのセットアップ
build:
	docker compose build --no-cache
	[ -f ./server/.env ] || cp ./server/.env.development ./server/.env
.PHONY: build

#? up: アプリケーションの起動
up:
	docker compose up -d
.PHONY: up

#? down: アプリケーションの停止
down:
	docker compose down
.PHONY: down

#? shell-http: HTTP サーバーのシェルを起動
shell:
	docker compose exec -it api bash
.PHONY: shell

#? protoc: gRPC のプロトコルバッファをコンパイル
protoc:
	docker compose run --rm api bash -c "\
	  find proto -name '*.proto' | xargs protoc \
	  	--experimental_allow_proto3_optional \
	    --proto_path=proto \
	    --go_out=server/pkg/api \
	    --go-grpc_out=server/pkg/api \
	    --go_opt=paths=source_relative \
	    --go-grpc_opt=paths=source_relative"
.PHONY: protoc

#? api: OpenAPI からコードを生成
api:
	docker compose run --rm api bash -c "oapi-codegen -package api /api/api/openapi.json > /api/server/adapter/api/interface.go"
	npx openapi-typescript ./api/openapi.json -o ./web/app/api/client.ts
.PHONY: api

#? wire: HTTP サーバーの依存関係を自動生成
wire:
	docker compose run --rm api bash -c "cd /api/server/cmd/tcmrsv/wire && wire gen && mv ./wire_gen.go /api/server/cmd/tcmrsv/commands/wire.go"
.PHONY: wire

#? migrate: データベースの構造をマイグレート
migrate:
	docker compose run --rm api bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db:5432/db?sslmode=disable up"
.PHONY: migrate

#? migrate-down: データベースの構造を初期化
migrate-down:
	docker compose run --rm api bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db:5432/db?sslmode=disable down -all"
.PHONY: migrate-down

#? sql-gen: SQL クエリから Go コードを生成
sql-gen:
	docker compose run --rm api bash -c "cd server && sqlc generate"
.PHONY: sql-gen

#? test: テストを実行
test:
	docker compose -f compose.test.yaml -p tcmrsv-test up -d api-test db-test

	@echo "⏳ テスト用 PostgreSQL コンテナを待機中..."
	@until docker compose -p tcmrsv-test exec -T db-test pg_isready -U user -p 5433; do \
		sleep 1; \
	done

	@echo "🧹 マイグレーションを初期化中..."
	docker compose -p tcmrsv-test exec -T api-test bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db-test:5433/db?sslmode=disable down -all"

	@echo "🚀 マイグレーションを開始"
	docker compose -p tcmrsv-test exec -T api-test bash -c "cd server && migrate -source file://migrations -database postgres://user:password@db-test:5433/db?sslmode=disable up"

	@echo "🧪 テスト実行中..."
	docker compose -p tcmrsv-test exec -T api-test bash -c "cd server && go test -v ./internal/..."
.PHONY: test

#? test-down: テスト用コンテナを停止
test-down:
	docker compose -f compose.test.yaml -p tcmrsv-test down
.PHONY: test-down
