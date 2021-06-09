# Go Mocking & Interfaces Demo

There are two major libs for mocking:
- [golang/mock](https://github.com/golang/mock) (we'll focus on this one)
- [testify/mock](https://github.com/stretchr/testify)

## Setup `mockgen`
```bash
# Go version < 1.16
GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0
```
or
```bash
# Go 1.16+
go install github.com/golang/mock/mockgen@v1.5.0
```

## Generating Mocks

```bash
mockgen -source basic/basic.go -destination basic/mock/basic.go 
mockgen -source ethclient/ethclient.go -destination ethclient/mock/ethclient.go 
```

## Basic Example

This is a simple example of an interface being used for a database. We can generate mocks for this interface and then write tests using it. See `TestStoreMulti`.

Note that we explicitly specify the expected interactions with the interface, thus we can verify exactly which calls are being made. In `TestStoreMultiFailing` we have specified an expected call which is different from the one that is made given the inputs, thus the test fails.

## Ethclient Example

Here we have a struct which uses `ethclient.Client`. 

We have replaced the usage of `ethclient.Client` with an interface - very easy, much wow! 

We can generate mocks for this interface and easily test the functionality of the `ChainReader`.

Once you start using more interfaces, you can refactor your code to be much more modular, like we have with `CheckBlocks` which removes the need for the ChainReader struct.

## Notes

### Interfaces

- By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc. ([source](https://golang.org/doc/effective_go#interface-names))
- Smaller interfaces are better
- Be cautious which types you use as this may restrict the usefulness of an interface