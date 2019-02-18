# gengou
元号を　取得するやつ　つくりたい

## FOR LIBRARY
### INSTALL

```
go get github.com/deadcheat/gengou
```

### HOW TO USE

Give slashed-style year/month/day formatted string to `gengou.Find` func
```
package main

import (
	"fmt"

	"github.com/deadcheat/gengou"
)

func main() {
	s := "2019/2/5"
	g, err := gengou.Find(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(g)
}

```

## FOR STANDALONE CLI APP

### INSTALL

```
go get github.com/deadcheat/gengou/...
```

### USAGE
```
$ gengou -h 
NAME:
   gengou - A new cli application

USAGE:
   gengou [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     standalone, s  give me year/month/day formatted string, gengou answers gengou to you!!
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

i.e,
```
gengou s 1345/4/17
```
will return
```
興国（南朝）
康永（北朝）
```
```
gengou s 2019/4/17
```
will return
```
平成
```
