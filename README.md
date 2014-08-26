# ShortId

[![Build Status](https://travis-ci.org/SKAhack/go-shortid.svg?branch=master)](https://travis-ci.org/SKAhack/go-shortid)

Port of [dylang/shortid](https://github.com/dylang/shortid) to Go

## Usage

```go
package main

import(
  "github.com/SKAhack/go-shortid"
)

func main() {
  g := shortid.Generator()
  g.Generate() // => 9uK7FCrIm
}
```

## Example

```shell
$ ./example
HJpRDHKIP
JJBjvDJ5IQ
HHLZUC9hVP
JuQpRCJhVg
u91ZUCJhIm
u9np_ru5SQ
9HV6UCH5Vg
99_Z_D9hVP
9HsZ_rJKVg
JJOpvDH5Vg
```

## API

See: [dylang/shortid](https://github.com/dylang/shortid)

### Initialize

```go
  g := shortid.Generator()
```

### Generate()

generate a new ShortId

```go
  g.Generate() // => JJOpvDH5Vg
```

### SetSeed(float64)

```go
  g.SetSeed(1)
```

## License

MIT

