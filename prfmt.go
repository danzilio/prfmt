package prfmt

import "fmt"
import "strings"

type PuppetResource struct {
  Type       string
  Title      string
  Attributes map[string]string
}

func (p *PuppetResource) PrintResource() {
  for _, e := range(p.FormatResource()) {
    fmt.Println(e)
  }
}

func (p *PuppetResource) FormatResource() []string {
  var resource []string
  resource = append(resource, fmt.Sprintf("%s { '%s':", p.Type, p.Title))
  resource = append(resource, formatAttributes(1, p.Attributes)...)
  resource = append(resource, "}")
  return resource
}

func formatAttributes(indentLevel int, attributes map[string]string) []string {
  var attr []string
  longestKeyLength := longestKeyLength(attributes)
  for k, v := range attributes {
    keyLength := len(k)
    numberOfSpaces := longestKeyLength - keyLength + 1
    spaces := strings.Repeat(" ", numberOfSpaces)
    indent := strings.Repeat(" ", indentLevel * 2)
    attr = append(attr, fmt.Sprintf("%s%s%s=> '%s',", indent, k, spaces, v))
  }
  return attr
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
