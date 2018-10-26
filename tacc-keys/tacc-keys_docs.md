## tacc-keys docs

Generate bash completion scripts and documentation

### Synopsis

docs will generate bash completion scripts, write them to standard out
and generate markdown documentation to /tmp/tacc-keys.

To load completion run
    
    . <(tacc-keys completion)

To configure your bash shell to load completions for each session add to your bashrc

    # ~/.bashrc or ~/.profile
    . <(tacc-keys completion)
    

```
tacc-keys docs [flags]
```

### Options

```
  -h, --help   help for docs
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.tacc-keys.yaml)
```

### SEE ALSO

* [tacc-keys](README.md)	 - Client for TACC's public ssh keys service
