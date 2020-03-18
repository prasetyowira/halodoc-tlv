# halodoc

A TLV Solver with Golang for Halodoc recruitment assignment

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/halodoc < input.txt
```

Or
```console
$ ./bin/halodoc <input:eg UPPRCS-0005-abcde>
```

To build inside docker
```console
$ make package
```

To run inside docker
```console
$ docker pull ariwira/halodoc:latest
$ docker run --rm ariwira/halodoc bash -c "halodoc < /opt/halodoc-tlv/input.txt"
or
$ docker run --rm ariwira/halodoc bash -c "halodoc UPPRCS-0005-abcde"
```

### Testing

``make test``