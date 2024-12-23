package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("Count:", s.Count("test", "t"))                  // Count
	p("HasPrefix:", s.HasPrefix("test", "te"))         // HasPrefix
	p("HasSuffix:", s.HasSuffix("test", "st"))         // HasSuffix
	p("Index:", s.Index("test", "e"))                  // Index
	p("Join:", s.Join([]string{"test", "exams"}, "-")) // Join
	p("Repeat:", s.Repeat("a", 5))                     // Repeat
	p("Replace:", s.Replace("foo", "o", "0", -1))      // Replace
	p("Replace:", s.Replace("foo", "o", "0", 1))       // Replace
	p("Split:", s.Split("a-b-c-d-e-f", "-"))           // Split
	p("ToLower:", s.ToLower("TEST"))                   // ToLower
	p("ToUpper:", s.ToUpper("test"))                   // ToUpper
}
