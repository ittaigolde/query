package query

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(obj interface{}) url.Values {
	t := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	ret := make(url.Values)
	for index := 0; index < t.NumField(); index++ {
		f := t.FieldByIndex([]int{index})
		if strings.ToLower(f.Name[0:1]) == f.Name[0:1] {
			//ignore private members
			continue
		}
		n := f.Tag.Get("query")
		if n == "" {
			n = f.Name
		}
		v := val.FieldByIndex([]int{index})
		ret.Add(n, fmt.Sprintf("%v", v))
	}
	return ret
}

func Unmarshal(values url.Values, obj interface{}) error {
	if obj == nil {
		return fmt.Errorf("Object is nil")
	}
	tt := reflect.TypeOf(obj)
	if tt.Kind() != reflect.Ptr {
		return fmt.Errorf("Object type should be ptr not %s", tt.Kind().String())
	}
	t := tt.Elem()
	for k, va := range values {
		for index := 0; index < t.NumField(); index++ {
			nn := t.Field(index)
			if strings.ToLower(nn.Name[0:1]) == nn.Name[0:1] {
				continue
			}
			if strings.ToLower(k) != strings.ToLower(nn.Name) &&
				strings.ToLower(k) != strings.ToLower(nn.Tag.Get("query")) {
				continue
			}
			if nn.Type.Kind() == reflect.String {
				ps := reflect.ValueOf(obj).Elem()
				v := ps.FieldByName(nn.Name)
				v.SetString(va[0])
			} else if nn.Type.Kind() == reflect.Int {
				i, err := strconv.Atoi(va[0])
				if err != nil {
					return err
				}
				ps := reflect.ValueOf(obj).Elem()
				v := ps.FieldByName(nn.Name)
				v.SetInt(int64(i))
			} else {
				return fmt.Errorf("Unsupported field %s type %s", nn.Name, nn.Type.Name())
			}

		}
	}
	return nil
}
