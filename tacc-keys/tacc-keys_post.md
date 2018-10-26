## tacc-keys post

Post a public ssh key to TACC's keys service

### Synopsis

Post a public ssh key to TACC's keys service.
You can also create a public and provate key pair and post the public key to 
the service.

```
tacc-keys post [username] [flags]
```

### Options

```
  -c, --create            Create a pair of public and private ssh keys
  -h, --help              help for post
  -k, --key-name string   Name of key (no extension) (default "id_rsa")
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.tacc-keys.yaml)
```

### SEE ALSO

* [tacc-keys](README.md)	 - Client for TACC's public ssh keys service