package diff

import (
	"bytes"
	"errors"
	"reflect"
	"text/template"
)

var (
	ErrStruct     = errors.New("args must be struct kind")
	ErrDiffStruct = errors.New("can not compare two different structs")
)

type (
	any = interface{}
	st  struct {
		fields []map[string]interface{}
	}
	Diff struct {
		tmpl *template.Template
	}
	options struct {
		tmpl string
	}

	option func(*options)
)

func New(opts ...option) (*Diff, error) {
	d := Diff{}
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}

	tmplSrc := defTmpl
	if len(o.tmpl) > 0 {
		tmplSrc = o.tmpl
	}
	tmpl, err := template.New("diff").Parse(tmplSrc)
	if err != nil {
		return nil, err
	}
	d.tmpl = tmpl
	return &d, nil
}

func WithTmpl(src string) option {
	return func(o *options) {
		o.tmpl = src
	}
}

func (d *Diff) Exec(x, y any, outerKv ...map[string]interface{}) (string, error) {
	st, err := compare(x, y)
	if err != nil {
		return "", err
	}
	args := map[string]interface{}{}
	if len(outerKv) > 0 {
		if outerKv[0] != nil {
			for k, v := range outerKv[0] {
				args[k] = v
			}
		}
	}
	args[fieldsKey] = st.fields
	buff := bytes.Buffer{}
	if err := d.tmpl.Execute(&buff, args); err != nil {
		return "", err
	}
	return buff.String(), nil
}

func compare(x, y any) (*st, error) {
	s := new(st)
	s.fields = make([]map[string]interface{}, 0)
	if x == nil || y == nil {
		return nil, ErrDiffStruct
	}
	v1 := reflect.ValueOf(x)
	v2 := reflect.ValueOf(y)
	if v1.Kind() == reflect.Pointer {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Pointer {
		v2 = v2.Elem()
	}
	if v1.Kind() != reflect.Struct {
		return nil, ErrStruct
	}
	if v1.Type() != v2.Type() {
		return nil, ErrDiffStruct
	}
	s.compare(v1, v2)
	return s, nil
}

func (s *st) compare(v1, v2 reflect.Value) {
	n := v1.NumField()
	for i := 0; i < n; i++ {
		f1, f2 := v1.Field(i), v2.Field(i)
		if f1.Kind() == reflect.Pointer {
			f1, f2 = f1.Elem(), f2.Elem()
		}
		switch f1.Kind() {
		case reflect.Struct:
			s.compare(f1, f2)
		default:
			if equal(f1, f2) {
				continue
			}
			field := v1.Type().Field(i)
			fn := field.Name
			valMp := map[string]interface{}{
				"name":   fn,
				"before": f1,
				"after":  f2,
			}
			kvs := parseKvs(field.Tag)
			for _, kv := range kvs {
				valMp[kv.Key] = kv.Val
			}
			s.fields = append(s.fields, valMp)
		}
	}
}

func equal(x, y reflect.Value) bool {
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.Float32, reflect.Float64:
		return x.Float() == y.Float()
	}
	return true
}
