# `prfmt`

Formats a Puppet resource based on your CLI arguments.

## Usage

`prfmt <type> <title> [<parameter>=<value>...]`

## Examples

```
$ prfmt file /tmp/test ensure=file content='foo bar baz'
file { '/tmp/test':
  ensure  => 'file',
  content => 'foo bar baz',
}
```

## Limitations

Right now this only formats resources with string parameters so you're SOL if you want to format a resource with a Hash or Array as a parameter value. It's also not very smart about quoting.
