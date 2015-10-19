# strm-append

A driver for the **go-strm** Go programming language library, that provides the **APPEND** command.

**APPEND** returns all rows, with a new row appended to the end of the data table stream.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/strm-append

[![GoDoc](https://godoc.org/github.com/reiver/strm-append?status.svg)](https://godoc.org/github.com/reiver/strm-append)

## Example
```
package main

import (
	. "github.com/reiver/strm-csv"
	. "github.com/reiver/strm-append"
	. "github.com/reiver/strm-stdout"
)

func main() {
	Begin(CSV, "table.csv").
		Strm(APPEND, "blue", 5).
	End(STDOUT, "tsv")
}
```

(Note that in that example dot imports were used.)

## See Also

For more information about **go-strm** and for a list of other drivers, see:
https://github.com/reiver/go-strm
