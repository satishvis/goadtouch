package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var usage string = `
	Program: Advance Touch in Golang
		 ╔════════════════╗
		 ║  Advance Touch ║
		 ║     Usage      ║
		 ╚════════════════╝
 ╭────────────────────────────────────────────────────────────────────╮
 │ $> goadtouch topdir/{subdir/first.go,fileinsidetopdir.go} \        │
 │ hello/anothersubdir/anotherfile.go independentdir/ independentfile │
 ╰────────────────────────────────────────────────────────────────────╯
 Will give us
.
├── hello
│   └── anothersubdir
│       └── anotherfile.go
├── independentdir
├── independentfile
└── topdir
    ├── fileinsidetopdir.go
    └── subdir
        └── first.go
# You can copy this file to your bin directory.
# You can also rename it to something easier like "ad"
`

func main() {
	a := os.Args[1:]

	if len(a) <= 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	for _, s := range a {
		dir, file := filepath.Split(s)
		if len(dir) > 0 {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				fmt.Println(err, "Exiting...")
			}
		}
		if len(file) > 0 && len(file) < 255 {
			_, err := os.Create(s)
			if err != nil {
				fmt.Println(err, "Exiting...")
			}
		}
	}
	fmt.Println("Done...")
}
