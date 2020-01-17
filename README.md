# Learning Go

## GOPATH vs GOROOT
`$GOPATH` is important

`go install` will put things in "main" package in `$GOPATH/bin`

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







