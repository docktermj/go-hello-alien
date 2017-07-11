# go-hello-alien

Part of a comparision of [Go HTTP routers](https://github.com/avelino/awesome-go/blob/master/README.md#routers).

1. https://github.com/gernest/alien

## Usage

### Invocation

```console
go-hello-alien
```

In a web browser, visit http://localhost:8090/

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

