.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## デプロイするためのDockerイメージをビルドする
	docker build -t budougumi0617/gotodo:${DOCKER_TAG} \
		--target deploy ./

build-local: ## ローカル開発へのDockerイメージの構築
	docker compose build --no-cache

up: ## ホットリロードでdocker compose upを行う
	docker compose up -d

down: ## docker compose downを行う
	docker compose down

logs: ## docker composeのログを尾行する
	docker compose logs -f

ps: ## コンテナの状態を確認する
	docker compose ps

test: ## テストの実行
	go test -race -shuffle=on ./...

help: ## オプションを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'