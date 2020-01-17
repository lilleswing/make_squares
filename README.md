# Learning Go

## GOPATH vs GOROOT
`$GOPATH` is important
`$GOPATH/src` is where raw go source code lives

`go get github.com/lilleswing/make_squares/` will bring this repo into `$GOPATH`

`$GOROOT` points to the golang binary

## Organizing your repo

```
/pkg/{name}/*.go
# This is where bread and butter code goes
# People should import your libraries from here

/cmd/{name}/main.go
# This is where command line programs live
# E.G. User level entry point to your program
```

## Writing Tests
Create a file {name}_test.go.  This file will have access to all functions/structs from {name}.go







