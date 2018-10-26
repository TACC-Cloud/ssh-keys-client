This program can list all the keys for a given user and delete keys given their
id (the key id can be obtained by listing them first).


Compile it.
```
go build -o keys
```

## Create ssh keys
```
$ ./keys --create --user username --keys "id_rsa"
```

## Delete a key                                                                 
```
$ ./keys-client --delete --user docker --id 60
key 60 was successfully deleted
``` 

## List all the keys for a given user
```
$ ./keys-client --list --user docker
56 keyservice-test
ssh-rsa AAAAB3NzaC1yc2EA1OFclnupLBSwrs8UUsgYIdiCBFOScKFiBtJEOkhBn+HK/LnreuyQ==
57 keyservice-test
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDO1UFxpdd//1p9Kl8U1J6R5r1pT7cf/CLlXL1Fz7
```
