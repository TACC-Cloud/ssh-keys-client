# Welcome to the client for TACC's public ssh keys server

This repo contains:
* [authorized keys client](./authorizedkeycommand) A minimal version of the
client for the keys service that will list all public keys for a given user in
a format that is accepted for sshd (see `AuthorizedKeysCommand` in `sshd_config(5)`).

* [oficial TACC keys client](./tacc-keys) A full client for the keys service. 
Once you have a registered oauth agave client, you can use this client to 
create public and private ssh keys for you, post the public key, delete keys, 
list keys, etc.

* [test run](./test-run) Provides a way to locally test the keys server
functionality and the authorized keys command by setting up an ssh server in a
container and trying to access it via a saved public ssh key and by using the
keys server.

## Notes
In case you want to cross compile any of the clients:
```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo
```
