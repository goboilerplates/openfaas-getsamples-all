# Boilerplate for GetSamplesAll OpenFaas Function in Go
[![Build Status](https://travis-ci.org/goboilerplates/openfaas-getsamples-all.svg?branch=master)](https://travis-ci.org/goboilerplates/openfaas-getsamples-all)
[![codecov](https://codecov.io/gh/goboilerplates/openfaas-getsamples-all/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/openfaas-getsamples-all)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/openfaas-getsamples-all)](https://goreportcard.com/report/github.com/goboilerplates/openfaas-getsamples-all)
[![GoDoc](https://godoc.org/github.com/goboilerplates/openfaas-getsamples-all/function?status.svg)](https://godoc.org/github.com/goboilerplates/openfaas-getsamples-all/function)
[![](https://images.microbadger.com/badges/image/goboilerplates/openfaas-getsamples-all.svg)](https://microbadger.com/images/goboilerplates/openfaas-getsamples-all)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/openfaas-getsamples-all/blob/master/LICENSE)

## Features
- OpenFaas Function
- CI with Travis
- Docker Build
- Docker Compose

## Installation

Get the openfaas-getsamples-all repository

```
go get github.com/goboilerplates/openfaas-getsamples-all

cd echo $GOPATH/src/github.com/goboilerplates/openfaas-getsamples-all
```
## Running the tests

Run all tests

```
go test ./test/...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Docker support 

Deploy OpenFaas with Docker Swarm

```
docker swarm init

bash script/deploy_stack.sh
```

Build function docker image

```
bash script/Dockerbuild.sh
```

Deploy function to OpenFaas

```
bash script/deploy.sh
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/openfaas-getsamples-all/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details