package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i< val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i:= 0; i<val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	}
}
func getValue(x interface{}) reflect.Value{
	val:=reflect.ValueOf(x)
	//指针类型的value不能使用NumField方法，需要调用Elem()提取底层值
	if val.Kind()==reflect.Ptr {
		val=val.Elem()
	}
	return val
}