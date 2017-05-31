default: check

deps:
	go get github.com/fatih/color

host:
	go run tests/auth.go tests/ibxhost.go

cname:
	go run tests/auth.go tests/ibxcname.go

text:
	go run tests/auth.go tests/ibxtextrec.go

zone:
	go run tests/auth.go tests/ibxzone.go

network:
	go run tests/auth.go tests/ibxnetwork.go

srv:
	go run tests/auth.go tests/ibxsrv.go

check:
	go fmt skyinfoblox.go
	go vet -v skyinfoblox.go

all: check host cname text zone network srv
