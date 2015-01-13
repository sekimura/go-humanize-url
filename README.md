# go-humanize-url [![Build Status](https://travis-ci.org/sekimura/go-humanize-url.svg?branch=master)](https://travis-ci.org/sekimura/go-humanize-url)

> Humanize a URL: `http://www.sekimura.org` → `sekimura.org`


## Install

```
$ go get github.com/sekimura/go-humanize-url
```


## Usage

```go
package main

import (
    "github.com/sekimura/go-humanize-url"
)

func main () {
    s, _ := humanizeurl.Humanize("http://www.sekimura.org/")
    println(s)
    //=> sekimura.org
}
```

## License

MIT © [Masayoshi Sekimura](http://sekimura.org)
