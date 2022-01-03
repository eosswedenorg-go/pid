# Pid

Helper functions related to process id.

### Install package

``` bash
go get -u github.com/eosswedenorg-go/pid@latest
```

### Functions

```c
func Get() (int)
```

Returns the process id of the caller.

```c
func Save(path string) (error)
```

Saves the process id of the caller to `path` on disk.

### Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)
