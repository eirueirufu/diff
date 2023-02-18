package diff_test

import (
	"testing"
	"time"

	"github.com/eirueirufu/diff"
)

type (
	stru struct {
		b    bool
		i    int
		i8   int8
		i16  int16
		i32  int32
		i64  int64
		ui   uint
		ui8  uint8
		ui16 uint16
		ui32 uint32
		ui64 uint64
		f32  float32
		f64  float64
		arr  [3]int64
		c    chan int64
		mp   map[int64]int64
		s    []int64
		str  string
		is   struct {
			ini   int64
			inf   float64
			instr string
		}

		pb    *bool
		pi    *int
		pi8   *int8
		pi16  *int16
		pi32  *int32
		pi64  *int64
		pui   *uint
		pui8  *uint8
		pui16 *uint16
		pui32 *uint32
		pui64 *uint64
		pf32  *float32
		pf34  *float64
		parr  *[3]int64
		pc    *chan int64
		pmp   *map[int64]int64
		ps    *[]int64
		pstr  *string
		pis   *struct {
			inpi64 *int64
			inpf64 *float64
			inpstr *string
		}
	}

	divTagStru struct {
		i64 int64   `diff:"name=divi64"`
		f64 float64 `diff:"name=divf64"`
		str string  `diff:"name=divstr"`
		in  struct {
			ini64 int64   `diff:"name=divini64"`
			inf64 float64 `diff:"name=divinf64"`
			instr string  `diff:"name=divinstr"`
		} `diff:"name=divin"`

		pi64 *int64   `diff:"name=divpi64"`
		pf64 *float64 `diff:"name=divpf64"`
		pstr *string  `diff:"name=divpstr"`
		pin  *struct {
			inpi64 *int64   `diff:"name=divinpi64"`
			inpf64 *float64 `diff:"name=divinpf64"`
			inpstr *string  `diff:"name=divinpstr"`
		} `diff:"name=divpin"`
	}

	divTmplStru struct {
		str string `diff:"k1=v1,k2=v2"`
	}
)

func Test_diff_Exec(t *testing.T) {
	diffOutput := `b changed: before:false after:true
	i changed: before:0 after:1
	i8 changed: before:0 after:1
	i16 changed: before:0 after:1
	i32 changed: before:0 after:1
	i64 changed: before:0 after:1
	ui changed: before:0 after:1
	ui8 changed: before:0 after:1
	ui16 changed: before:0 after:1
	ui32 changed: before:0 after:1
	ui64 changed: before:0 after:1
	f32 changed: before:0 after:1
	f64 changed: before:0 after:1
	str changed: before: after:s
	ini changed: before:0 after:1
	inf changed: before:0 after:1
	instr changed: before: after:s
	`
	diff, err := diff.New()
	if err != nil {
		t.Fatal(err)
	}
	x := stru{}
	y := stru{
		b:    true,
		i:    1,
		i8:   1,
		i16:  1,
		i32:  1,
		i64:  1,
		ui:   1,
		ui8:  1,
		ui16: 1,
		ui32: 1,
		ui64: 1,
		f32:  1.,
		f64:  1.,
		arr:  [3]int64{1, 2, 3},
		c:    make(chan int64),
		mp: map[int64]int64{
			1: 1,
		},
		s:   []int64{1},
		str: "s",
		is: struct {
			ini   int64
			inf   float64
			instr string
		}{
			ini:   1,
			inf:   1.,
			instr: "s",
		},
		pb:    new(bool),
		pi:    new(int),
		pi8:   new(int8),
		pi16:  new(int16),
		pi32:  new(int32),
		pi64:  new(int64),
		pui:   new(uint),
		pui8:  new(uint8),
		pui16: new(uint16),
		pui32: new(uint32),
		pui64: new(uint64),
		pf32:  new(float32),
		pf34:  new(float64),
		parr:  &[3]int64{1, 2, 3},
		pc:    new(chan int64),
		pmp: &map[int64]int64{
			1: 1,
		},
		ps:   &[]int64{1},
		pstr: new(string),
		pis: &struct {
			inpi64 *int64
			inpf64 *float64
			inpstr *string
		}{
			inpi64: new(int64),
			inpf64: new(float64),
			inpstr: new(string),
		},
	}
	type args struct {
		x any
		y any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "error kind",
			args: args{
				x: x,
				y: new(int64),
			},
			wantErr: true,
		},
		{
			name: "nil",
			args: args{
				x: x,
				y: nil,
			},
			wantErr: true,
		},
		{
			name: "diff type",
			args: args{
				x: x,
				y: divTagStru{},
			},
			wantErr: true,
		},
		{
			name: "same vals1",
			args: args{
				x: x,
				y: x,
			},
			want: "",
		},
		{
			name: "same vals2",
			args: args{
				x: y,
				y: y,
			},
			want: "",
		},
		{
			name: "diff vals1",
			args: args{
				x: x,
				y: y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals2",
			args: args{
				x: &x,
				y: y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals3",
			args: args{
				x: x,
				y: &y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals4",
			args: args{
				x: &x,
				y: &y,
			},
			want: diffOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := diff.Exec(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("diff.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("diff.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diff_Exec_DivTag(t *testing.T) {
	diffOutput := `divi64 changed: before:0 after:1
	divf64 changed: before:0 after:1
	divstr changed: before: after:1
	divini64 changed: before:0 after:1
	divinf64 changed: before:0 after:1
	divinstr changed: before: after:1
	`
	diff, err := diff.New()
	if err != nil {
		t.Fatal(err)
	}
	x := divTagStru{}
	y := divTagStru{
		i64: 1,
		f64: 1.,
		str: "1",
		in: struct {
			ini64 int64   "diff:\"name=divini64\""
			inf64 float64 "diff:\"name=divinf64\""
			instr string  "diff:\"name=divinstr\""
		}{
			ini64: 1,
			inf64: 1.,
			instr: "1",
		},
		pi64: new(int64),
		pf64: new(float64),
		pstr: new(string),
		pin: &struct {
			inpi64 *int64   "diff:\"name=divinpi64\""
			inpf64 *float64 "diff:\"name=divinpf64\""
			inpstr *string  "diff:\"name=divinpstr\""
		}{
			inpi64: new(int64),
			inpf64: new(float64),
			inpstr: new(string),
		},
	}
	type args struct {
		x any
		y any
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "same vals1",
			args: args{
				x: x,
				y: x,
			},
			want: "",
		},
		{
			name: "same vals2",
			args: args{
				x: y,
				y: y,
			},
			want: "",
		},
		{
			name: "diff vals1",
			args: args{
				x: x,
				y: y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals2",
			args: args{
				x: &x,
				y: y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals3",
			args: args{
				x: x,
				y: &y,
			},
			want: diffOutput,
		},
		{
			name: "diff vals4",
			args: args{
				x: &x,
				y: &y,
			},
			want: diffOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := diff.Exec(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("diff.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("diff.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diff_Exec_DivTmpl(t *testing.T) {
	divTmpl := `{{- .name}} changed this record:
	{{- range .Fields}}
		{{.name}}: before:{{.before}} after:{{.after}}
	{{- end}}
		record time: {{.time}}`
	diff, err := diff.New(diff.WithTmpl(divTmpl))
	if err != nil {
		t.Fatal(err)
	}
	now := time.UnixDate
	type args struct {
		x       any
		y       any
		outerKv map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no change and not outerKv",
			args: args{
				x: divTmplStru{},
				y: divTmplStru{},
			},
			want: `<no value> changed this record:
		record time: <no value>`,
		},
		{
			name: "no outerKv",
			args: args{
				x: divTmplStru{
					str: "str1",
				},
				y: divTmplStru{
					str: "str2",
				},
			},
			want: `<no value> changed this record:
		str: before:str1 after:str2
		record time: <no value>`,
		},
		{
			name: "with outerKv",
			args: args{
				x: divTmplStru{
					str: "str1",
				},
				y: divTmplStru{
					str: "str2",
				},
				outerKv: map[string]interface{}{
					"name": "admin",
					"time": now,
				},
			},
			want: `admin changed this record:
		str: before:str1 after:str2
		record time: Mon Jan _2 15:04:05 MST 2006`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := diff.Exec(tt.args.x, tt.args.y, tt.args.outerKv)
			if (err != nil) != tt.wantErr {
				t.Errorf("diff.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("diff.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
