CONTAINER_USER := docker
SSH_KEY := ssh/id_rsa
SSH_PORT ?= 2222
SERVER_PORT ?= 8000
CONTAINER_NAME ?= tacc-keys


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
		-i $(SSH_KEY) -p $(SSH_PORT) $(CONTAINER_USER)@localhost

ssh-norm:
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
        -i ssh/id_rsa -p $(SSH_PORT) $(CONTAINER_USER)@localhost

ssh-cmd:
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
		-i $(SSH_KEY) -p $(SSH_PORT) $(CONTAINER_USER)@localhost -- whoami

ssh-norm-cmd:
	ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null \
        -i ssh/id_rsa -p $(SSH_PORT) $(CONTAINER_USER)@localhost -- whoami

clean:
	rm -rf ssh/
