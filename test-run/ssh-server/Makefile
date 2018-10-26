SSH_PORT ?= 2222
SERVER_PORT ?= 8000
CONTAINER_NAME ?= ssh-server


.PHONY: build generate-keys run ssh clean 

build: generate-keys
	docker build --force-rm --no-cache -t $(CONTAINER_NAME) .

generate-keys:
	mkdir -p ssh
	ssh-keygen -t rsa -b 4096 -C "docker" -f ssh/id_rsa -N ""
	cat ssh/id_rsa.pub >> ssh/authorized_keys

run: build
	docker run --rm \
		-p $(SERVER_PORT):8000 -p $(SSH_PORT):22 \
		$(CONTAINER_NAME)

ssh:
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
		-i ssh/id_rsa -p $(SSH_PORT) docker@localhost

clean:
	rm -rf ssh/
