package main

import "os"
import "fmt"
import "strings"

type PuppetResource struct {
  Type       string
  Title      string
  Attributes map[string]string
}

func main() {
  args := os.Args[1:]

  if len(args) < 3 {
    fmt.Println("Usage: prfmt <type> <title> attribute=value ...")
    os.Exit(1)
  }

  r := PuppetResource{Type: args[0], Title: args[1]}
  r.Attributes = argsToAttributes(args[2:])
  r.printResource()
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

func (p *PuppetResource) printResource() {
  fmt.Printf("%s { '%s':\n", p.Type, p.Title)
  printAttributes(1, p.Attributes)
  fmt.Println("}")
}

func printAttributes(indentLevel int, attributes map[string]string) {
  longestKeyLength := longestKeyLength(attributes)
  for k, v := range attributes {
    keyLength := len(k)
    numberOfSpaces := longestKeyLength - keyLength + 1
    spaces := strings.Repeat(" ", numberOfSpaces)
    indent := strings.Repeat(" ", indentLevel * 2)
    fmt.Printf("%s%s%s=> '%s',\n", indent, k, spaces, v)
  }
}

func longestKeyLength(attributes map[string]string) int {
  var length int
  for k, _ := range attributes {
    l := len(k)
    if l > length {
      length = l
    }
  }
  return length
}
