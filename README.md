### README

# Quantified Self - Rails

### Table Of Contents
- [Versions/Prerequisites](#versions-prerequisites)
- [Setup](#setup)
- [The Test Suite](#the-test-suite)
- [Endpoints](#endpoints)
- [Authors](#authors)

### Versions/Prerequisites

The prerequisites for this application are:
- Postgresql 10+
- Golang 1.10+
- [Go Dep](https://github.com/golang/dep)

### Setup
- Make sure you have go language installed and your GOPATH set
- From your go path get the go library and change directory into the project directory:
```
go get github.com/anon0mys/qs_golang
cd qs_golang
```
- Run go build and install to install the gem dependencies:
```
go build
go install
```
- Run the go app locally
```
qs_golang
```

### The Test Suite
- The test suite is written in Ginkgo. To run the test suite, from the root project folder run:
```
ginkgo -r -v
```

### Endpoints
- Documentation for all endpoints is here:
[Endpoint Explanation](https://github.com/anon0mys/qs_golang/blob/master/endpoint.md)

### Contributions
Quantified self is open source and welcomes contributions. If you would like to contribute please follow this workflow:
- Ensure you have a working Ruby environment with the appropriate [Versions/Prerequisites](#versions-prerequisites)
- Fork, then clone the repository
- Follow the [Setup](#setup) instructions
- Make your desired changes and accompanying tests
- Open a PR to the anon0mys/qs_golang repository
- An app administrator will conduct code review and contact you once the fix is accepted or rejected

### Authors
- [Evan Wheeler](https://github.com/anon0mys)
