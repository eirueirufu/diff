# diff

Compare the fields in go struct and customize the output.

<p align="left">
<a href="https://github.com/eirueirufu/diff/actions"><img src="https://github.com/eirueirufu/diff/workflows/go/badge.svg?branch=main" alt="Build Status"></a>
<a href="https://codecov.io/github/eirueirufu/diff"><img src="https://codecov.io/github/eirueirufu/diff/branch/main/graph/badge.svg" alt="codeCov"></a>
<a href="https://pkg.go.dev/github.com/eirueirufu/diff"><img src="https://pkg.go.dev/badge/github.com/eirueirufu/diff" alt="Go Reference"></a>
</p>

## Quick start

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/eirueirufu/diff"
)

func main() {
	type Info struct {
		Name string `diff:"alias=username"`
		Age  int
	}
	bf := Info{
		Name: "foo",
	}
	af := bf
	af.Name = "bar"
	d, err := diff.New(diff.WithTmpl(`{{- .name}} changed this record:
{{- range .Fields}}
	{{.alias}}: before:{{.before}} after:{{.after}}
{{- end}}
	record time: {{.time}}`))
	if err != nil {
		log.Fatal(err)
	}
	out, err := d.Exec(bf, af, map[string]interface{}{
		"name": "admin",
		"time": time.UnixDate,
	})
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write([]byte(out))
	// Output:admin changed this record:
	// 	username: before:foo after:bar
	// 	record time: Mon Jan _2 15:04:05 MST 2006
}
```

## template

Here is an example template, it uses [template](https://pkg.go.dev/text/template) to render.

```text/plain
{{- .name}} changed this record:
{{- range .Fields}}
	{{.alias}}: before:{{.before}} after:{{.after}}
{{- end}}
	record time: {{.time}}
```

- `Fields`: fields in struct to range.
  - `name`: the field name.
  - `before`: the value before change.
  - `after`: the value after change.
  - others in field: define in field tag, for example `diff:"alias=username,anyKey=anyVal"`
- others: like the `time` in example
    ```go
    out, err := d.Exec(bf, af, map[string]interface{}{
        "time": time.UnixDate,
    })
    ```

Use `diff.New(diff.WithTmpl(myTmpl))` to define your template.

## Field type

Basic data types will be compared, such as: Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, String, Bool, Float32, Float64, Complex64, Complex128.
Other data types will be ignored, such as: slice, map, chan, etc.

