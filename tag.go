package diff

import (
	"reflect"
	"strings"
)

const tagName = "diff"

type (
	kv struct {
		Key string
		Val string
	}
)

var trimSpace = strings.TrimSpace

func parseKvs(tag reflect.StructTag) []kv {
	kvs := make([]kv, 0)
	tval := tag.Get(tagName)
	vals := strings.Split(tval, ",")
	for _, val := range vals {
		strs := strings.SplitN(val, "=", 2)
		key := trimSpace(strs[0])
		if len(key) == 0 {
			continue
		}
		v := kv{
			Key: key,
		}
		if len(strs) > 1 {
			v.Val = trimSpace(strs[1])
		}
		kvs = append(kvs, v)
	}
	return kvs
}
