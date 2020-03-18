# halodoc
![Docker Image CI](https://github.com/prasetyowira/halodoc-tlv/workflows/Docker%20Image%20CI/badge.svg?branch=master)

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
$ docker login docker.pkg.github.com -u prasetyowira --password 5205317c242488d2b6a1d55ce911acc1ccbd3a58
$ docker pull docker.pkg.github.com/prasetyowira/halodoc-tlv/halodoc-tlv:latest
$ docker run --rm docker.pkg.github.com/prasetyowira/halodoc-tlv/halodoc-tlv bash -c "halodoc-tlv < /opt/halodoc-tlv/input.txt"
or
$ docker run --rm docker.pkg.github.com/prasetyowira/halodoc-tlv/halodoc-tlv bash -c "halodoc-tlv UPPRCS-0005-abcde"
```

### Testing

``make test``