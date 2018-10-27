## tacc-keys completion

Generate bash completion scripts and documentation

### Synopsis

completion will generate bash completion scripts and write them to 
standard out. To load completion run
    
    . <(tacc-keys completion)

To configure your bash shell to load completions for each session add to your bashrc

    # ~/.bashrc or ~/.profile
    . <(tacc-keys completion)
    

```
tacc-keys completion [flags]
```

### Options

```
  -h, --help   help for completion
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.tacc-keys.yaml)
```

### SEE ALSO

* [tacc-keys](README.md)	 - Client for TACC's public ssh keys service
