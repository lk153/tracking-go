## TRACKING SYSTEM
#### INSTALLATION

1. Compile & build execute binary file
```
make default
```
2. Up server
```
make run
```
3. Install migrate command
```
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.darwin-amd64.tar.gz | tar xvz
mv migrate.darwin-amd64.tar.gz migrate
```
#### TIPS

Downgrade Go modules
```
go get -u github.com/apache/thrift@v0.13.0
go mod tidy
go clean -modcache
go mod vendor
```

#### Kafka Setup
1. Edit configuration of kafka advertise listener to public interface (public IP of VM)
```
KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://broker:29092,PLAINTEXT_HOST://[public IP]:9092
```
[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-2.0-4baaaa.svg)](CODE_OF_CONDUCT.md)
