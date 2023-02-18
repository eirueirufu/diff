package diff

import (
	"reflect"
	"testing"
)

func Test_parseKvs(t *testing.T) {
	type args struct {
		tag reflect.StructTag
	}
	tests := []struct {
		name string
		args args
		want []kv
	}{
		{
			name: "empty tag",
			args: args{
				reflect.StructTag(""),
			},
			want: []kv{},
		},
		{
			name: "diff tag1",
			args: args{
				reflect.StructTag("diff:\"k1\""),
			},
			want: []kv{{Key: "k1", Val: ""}},
		},
		{
			name: "diff tag2",
			args: args{
				reflect.StructTag("diff:\"k1=v1\""),
			},
			want: []kv{{Key: "k1", Val: "v1"}},
		},
		{
			name: "diff tag3",
			args: args{
				reflect.StructTag("diff:\" k1 = v1 \""),
			},
			want: []kv{{Key: "k1", Val: "v1"}},
		},
		{
			name: "diff tag4",
			args: args{
				reflect.StructTag("diff:\"k1=v1,k2=v2\""),
			},
			want: []kv{
				{Key: "k1", Val: "v1"},
				{Key: "k2", Val: "v2"},
			},
		},
		{
			name: "diff tag5",
			args: args{
				reflect.StructTag("diff:\" k1 = v1 , k2 = v2 \""),
			},
			want: []kv{
				{Key: "k1", Val: "v1"},
				{Key: "k2", Val: "v2"},
			},
		},
		{
			name: "other tag1",
			args: args{
				reflect.StructTag("key:\"val\""),
			},
			want: []kv{},
		},
		{
			name: "other tag2",
			args: args{
				reflect.StructTag("key1:\"val1\" key2:\"val2\""),
			},
			want: []kv{},
		},
		{
			name: "json and diff tag",
			args: args{
				reflect.StructTag("key1:\"val1\" key2:\"val2\" diff:\" k1 = v1 , k2 = v2 \""),
			},
			want: []kv{
				{Key: "k1", Val: "v1"},
				{Key: "k2", Val: "v2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseKvs(tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKvs() = %v, want %v", got, tt.want)
			}
		})
	}
}
