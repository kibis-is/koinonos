scripts_dir := ./scripts

all: install setup

dev:
	docker compose -f compose.development.yml up

install:
	@go install github.com/conventionalcommit/commitlint@latest
	go mod download
	go mod verify

setup:
	$(scripts_dir)/setup.sh
