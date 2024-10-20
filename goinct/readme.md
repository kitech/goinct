put this package go source directory:
only goinct part, no github.com/kitech:
    $GOROOT/src/.git

```
import "goinct"         // << $GOROOT/internal/*
import "cmd/cmdin"      // << $GOROOT/cmd/internal/*
import "runtime/rtin"   // << $GOROOT/runtime/internal/*
```

package means go internal cmd tool

all/any go/src/internal or go/src/cmd/internal funcs
can export by this fake go internal package

now exported funcs:

- NMget(file string) []Syminfo
