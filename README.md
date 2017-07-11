# go-hello-alien

Part of a comparision of [Go HTTP routers](https://github.com/avelino/awesome-go/blob/master/README.md#routers).

1. https://github.com/gernest/alien

## Usage

### Invocation

#### Server

```console
go-hello-alien
```

#### Client

In a web browser:

1. Visit http://localhost:8090/ for "Hello world"

`curl` commands

```console
curl -v -X GET http://localhost:8090/streaming
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR=${GOPATH}/src/github.com/docktermj
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-alien.git
```

#### Download dependencies

```console
cd ${PROJECT_DIR}/go-hello-alien
make dependencies
```

### Build

#### Local build

```console
cd ${PROJECT_DIR}/go-hello-alien
make
```

The results will be in the `${GOPATH}/bin` directory.
