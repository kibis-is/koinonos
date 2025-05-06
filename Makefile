scripts_dir := ./scripts

all: init install

dev:
	docker compose -f ./deployment/compose.development.yml up

install:
	go mod download
	go mod verify

init:
	$(scripts_dir)/init.sh
