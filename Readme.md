### Requirement
- Golang version > 1.11 (using with module)

### Function
- config env
- middleware group api
- mysql
- unitest
- intergration test
- validate input
- redis
- api document

### Install dependency library
```
$ go mod init <name-project>
Ex: $ go mod init start
```

### To run 
```
$ ./start.sh
```

### Build binary
```
$ go build
```

### Run
development without integration test:
```
$ GIN_MODE=release ./start -ENV=development
```

development with integration test: (need config sqlite3: $GOPATH/bin/caro.db)
```
$ GIN_MODE=release ./start -ENV=development -TEST=true
```

production:
```
$ GIN_MODE=release ./start -ENV=production
```

auto refresh: (only use with development ENV)
```
$ gin --a 3001 main.go 
```

### Unit test
```
$ go test -v -count=1 ./...
```

