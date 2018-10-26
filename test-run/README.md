# Key Service


To build and run the container, which has a sshd service running along with the
`tacc-keys` executable do
```
make run
```

If you want to ssh inside the container you can do so by running the following
command:
```
make ssh
```


For some simple benchmarks comparing the two ways to login:
```
perf stat -r 10 -d make ssh-norm-cmd
```

and 
```
perf stat -r 10 -d make ssh-cmd
```


Note, to kill the container running the ssh server you can try (if it is the
only container running)
```
docker rm -f $(ps -q)
```

## Contents
* [ssh-server](ssh-server) builds an ssh server inside a container.
