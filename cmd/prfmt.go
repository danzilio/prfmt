package main

import "os"
import "fmt"
import "strings"

import "github.com/danzilio/prfmt"

func main() {
  args := os.Args[1:]

  if len(args) < 3 {
    fmt.Println("Usage: prfmt <type> <title> attribute=value ...")
    os.Exit(1)
  }

  r := prfmt.PuppetResource{Type: args[0], Title: args[1]}
  r.Attributes = argsToAttributes(args[2:])
  r.PrintResource()
}

func argsToAttributes(a []string) map[string]string {
  attribs := make(map[string]string)
  for _, e := range(a) {
    s := strings.Split(e, "=")
    if len(s) < 2 {
      fmt.Printf("ERROR: %s is not a key=value pair!\n", e)
      os.Exit(1)
    }
    attribs[s[0]] = s[1]
  }
  return attribs
}
