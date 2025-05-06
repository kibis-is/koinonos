scripts_dir := ./scripts

all: setup install

dev:
	docker compose -f ./deployment/compose.development.yml up

install:
	go mod download
	go mod verify

setup:
	@go install github.com/conventionalcommit/commitlint@latest
	$(scripts_dir)/setup.sh
