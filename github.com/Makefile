ifdef ssh_key_path
	SSH_KEY_PATH="$$(cat $(ssh_key_path))"
else
	SSH_KEY_PATH="$$(cat ~/.ssh/id_rsa)"
endif

docker-down:
	docker-compose down

docker-build:
	docker-compose build --no-cache --build-arg SSH_KEY=$(SSH_KEY_PATH)

docker-up: docker-down docker-build
	docker-compose up


build:
	go build

dev: build
	./private_golang
