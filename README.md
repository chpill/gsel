Gsel
====

A command line tool to wrap this [json-select go implementation].

**Usage**:

```
# in this directory
gsel .a.b < test_data/simple.json
# => [true]
```

**WARNING**: this is a very alpha project, mostly a pretext to play with
go for my own amusement. Use this code at your own risk.

**Dependencies**: `go get github.com/coddingtonbear/go-jsonselect github.com/coddingtonbear/go-jsonselect`

**Build**: `go build`

**Test**: `go test`

[json-select go implementation]: https://github.com/coddingtonbear/go-jsonselect
