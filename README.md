# halodoc

A TLV Solver with Golang for Halodoc recruitment assignment

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/halodoc-tlv < input.txt
```

Or
```console
$ ./bin/halodoc-tlv <input:eg UPPRCS-0005-abcde>
```

To build inside docker
```console
$ make package
```

To run inside docker
```console
$ docker pull prasetyowira/halodoc-tlv:latest
$ docker run --rm ariwira/halodoc bash -c "halodoc-tlv < /opt/halodoc-tlv/input.txt"
or
$ docker run --rm ariwira/halodoc bash -c "halodoc-tlv UPPRCS-0005-abcde"
```

### Testing

``make test``