put this package go source directory:
only goinct part, no github.com/kitech:
    go/src/cmd/goinct

import "cmd/goinct"

package means go internal cmd tool

all/any go/src/internal or go/src/cmd/internal funcs
can export by this fake go internal package

now exported funcs:

- NMget(file string) []Syminfo
