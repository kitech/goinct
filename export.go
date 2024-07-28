package goinct

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"cmd/internal/objfile"
)

func Keep() {
	objfile.Open(""	)
}

 func errorf(format string, args ...any) {
	log.Printf(format, args...)
	//exitCode = 1
}

var (
	sortOrder = flag.String("sort", "name", "")
	printSize = flag.Bool("size", false, "")
	printType = flag.Bool("type", false, "")

	filePrefix = false
)


func NM(file string) {
	f, err := objfile.Open(file)
	if err != nil {
		errorf("%v", err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(os.Stdout)

	entries := f.Entries()

	var found bool

	for _, e := range entries {
		syms, err := e.Symbols()
		if err != nil {
			errorf("reading %s: %v", file, err)
		}
		if len(syms) == 0 {
			continue
		}

		found = true

		switch *sortOrder {
		case "address":
			sort.Slice(syms, func(i, j int) bool { return syms[i].Addr < syms[j].Addr })
		case "name":
			sort.Slice(syms, func(i, j int) bool { return syms[i].Name < syms[j].Name })
		case "size":
			sort.Slice(syms, func(i, j int) bool { return syms[i].Size > syms[j].Size })
		}

		for _, sym := range syms {
			if len(entries) > 1 {
				name := e.Name()
				if name == "" {
					fmt.Fprintf(w, "%s(%s):\t", file, "_go_.o")
				} else {
					fmt.Fprintf(w, "%s(%s):\t", file, name)
				}
			} else if filePrefix {
				fmt.Fprintf(w, "%s:\t", file)
			}
			if sym.Code == 'U' {
				fmt.Fprintf(w, "%8s", "")
			} else {
				fmt.Fprintf(w, "%8x", sym.Addr)
			}
			if *printSize {
				fmt.Fprintf(w, " %10d", sym.Size)
			}
			fmt.Fprintf(w, " %c %s", sym.Code, sym.Name)
			if *printType && sym.Type != "" {
				fmt.Fprintf(w, " %s", sym.Type)
			}
			fmt.Fprintf(w, "\n")
		}
	}

	if !found {
		errorf("reading %s: no symbols", file)
	}

	w.Flush()
}
