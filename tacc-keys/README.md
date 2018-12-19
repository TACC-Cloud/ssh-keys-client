# Oficial client for TACC's keys service

Hello, if you have found yourself here then you are probably in need of a
client for TACC's public ssh keys service.

If this is so, then you have found the perfect tool!
This client will allow you to interact with the keys service by listing keys,
creating public and private key pairs, posting public keys, deleting them,
and managing the oauth client you need to interact with TACC.

To install this client you can go and download the appropiate binary from the
[releases page](https://github.com/TACC-Cloud/ssh-keys-client/releases) or if
you already have [Go](https://golang.org/) installed in your system then do:
```
$ go get -u https://github.com/TACC-Cloud/ssh-keys-client/tacc-keys
```

To see how you can get started, 
[read the short documentation for tacc-keys](./docs/tacc-keys.md).
