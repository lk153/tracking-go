## GO GUIDE

#### INSTALLATION

1. Compile & build execute binary file
```
make default
```
2. Up server
```
make run
```

#### TIPS

1. Downgrade Go modules
```
go get -u github.com/apache/thrift@v0.13.0
go mod tidy
go clean -modcache
go mod vendor
```


