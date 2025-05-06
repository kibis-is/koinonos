scripts_dir := ./scripts

all: install setup

dev:
	docker compose -f ./deployment/compose.development.yml up --build

install:
	@go install github.com/conventionalcommit/commitlint@latest
	go mod download
	go mod verify

setup:
	$(scripts_dir)/setup.sh
