# nsid

A tool for generating IDs based on Unix Nanoseconds since Epoch.

BaseAlpha is designed to be alpha sortable.

## Building

```
$ git clone https://github.com/Lavos/nsid
$ cd commands
$ go build -o nsid main.go
```

## Usage

```
$ nsid -format base64
FZfjHbAjM1E
```

```
Usage of ./nsid:
  -format string
        Encoding format for the generated id. Formats: all unixnano phrase base16 base62 base64 basealpha base85 (default "base64")
  -n    If passed, suppress newline after outputting generated id.
```

Here's an example output of format `all` that shows all representations of a single timestamp:

```
UnixNano:	1556669143458873642
Phrase:		atop feed dial cure fled gab dust both
Base16:		159a6656a0ad812a
Base64:		FZpmVqCtgSo
Base62:		1QZyNdN9lCW
BaseAlpha:	KnybzEsCZMq
Base85:		6Zns}W7l+h
```
