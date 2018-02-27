# go-skeleton
a go project skeleton

## directory
under GOPATH/src/:
/Users/hma/go/src/github.com/interma/go-skeleton

ln to your workspace for convenience
```
cd /Users/hma/workspace/
ln -s ~/go/src/github.com/interma/go-skeleton
```

## build, run-test, and execution
```
$ make depend
...

$ make unit
gofmt -w -s main.go pkg
goimports -w main.go pkg
ginkgo -r -cover -coverprofile=coverage.out pkg/add
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
	add_test.go:6: in TestAdd
PASS
coverage: 100.0% of statements

Ginkgo ran 1 suite in 269.124388ms
Test Suite Passed

$ make coverage
gofmt -w -s main.go pkg
goimports -w main.go pkg
ginkgo -r -cover -coverprofile=coverage.out pkg/add
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
	add_test.go:6: in TestAdd
PASS
coverage: 100.0% of statements

Ginkgo ran 1 suite in 257.989526ms
Test Suite Passed
go tool cover -func=pkg/add/coverage.out
github.com/interma/go-skeleton/pkg/add/add.go:3:	Add		100.0%
total:							(statements)	100.0%

$ make build_mac
gofmt -w -s main.go pkg
goimports -w main.go pkg
env GOOS=darwin GOARCH=amd64 go build -tags 'main'  -o main_mac -ldflags ""

$ ./main_mac
0.1
3

$ make clean
rm -f main_mac main_linux /Users/hma/go/bin/main pkg/add/coverage.out

```



