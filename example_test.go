package diff_test

import (
	"log"
	"os"
	"time"

	"github.com/eirueirufu/diff"
)

func ExampleDiff_Exec() {
	type Info struct {
		Name string
	}
	bf := Info{
		Name: "foo",
	}
	af := bf
	af.Name = "bar"
	d, err := diff.New()
	if err != nil {
		log.Fatal(err)
	}
	out, err := d.Exec(bf, af)
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.Write([]byte(out))
	// Output:Name changed: before:foo after:bar
}

func ExampleWithTmpl() {
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
