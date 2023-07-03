# OpenSergo APIs and Common Configuration Definitions

This repository defines component-level APIs and common configuration formats for OpenSergo.
These definitions are specified using the [protobuf](https://github.com/google/protobuf) syntax.

## Generate files

Generate CRD files and .pb.go:

```shell
./gen.sh
```

To generate go-client code, you'll need to put the project under `$GOPATH` directory, then:

```shell
# The project should be located at $GOPATH/opensergo.io/api
cd client && make
```