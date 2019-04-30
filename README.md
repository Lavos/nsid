# nsid
A tool for generating IDs based on Unix Nanoseconds since Epoch

## usage

```
$ nsid -format <format>
FZfjHbAjM1E
```

```
Usage of ./nsid:
  -format string
        Encoding format for the generated id. Formats: all unixnano phrase base16 base62 base64 basealpha base85 (default "base64")
  -n    If passed, suppress newline after outputting generated id.
```
