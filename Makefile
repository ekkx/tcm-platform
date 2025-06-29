#? help: ヘルプコマンド
help: Makefile
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n "s/^#?//p" $< | column -t -s ":"
.PHONY: help

#? dev-up: 開発環境用サーバーを起動
dev-up:
	cd server && air
.PHONY: dev-up

#? sqlc: SQL クエリを Go コードに変換
sqlc:
	cd server && sqlc generate
.PHONY: sqlc

#? migrate-up: データベースの構造をマイグレート
migrate-up:
	cd server && migrate -source file://migrations -database postgres://tcmrsv:tcmrsv@tcmrsv-db:5432/tcmrsv_db?sslmode=disable up
.PHONY: migrate-up

#? migrate-down: データベースの構造を初期化
migrate-down:
	cd server && migrate -source file://migrations -database postgres://tcmrsv:tcmrsv@tcmrsv-db:5432/tcmrsv_db?sslmode=disable down -all
.PHONY: migrate-down
