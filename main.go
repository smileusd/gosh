package main

/*
   Copyright (C) 2014 Kouhei Maeda <mkouhei@palmtb.net>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"flag"
	"fmt"
)

var version string
var show_version = flag.Bool("version", false, "show_version")

var license = `Gosh %s
Copyright (C) 2014 Kouhei Maeda
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>.
This is free software, and you are welcome to redistribute it.
There is NO WARRANTY, to the extent permitted by law.
`

func main() {
	d := flag.Bool("d", false, "debug mode")
	flag.Parse()
	if *show_version {
		fmt.Printf("version: %s\n", version)
		return
	}
	cleanDirs()
	e := newEnv(*d)
	fmt.Println(goVersion())
	fmt.Printf(license, version)
	e.shell(nil)
}
