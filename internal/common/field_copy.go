package common

import "reflect"

// CopyFields 通过反射将结构体A的字段值赋值给结构体B
func CopyFields(a any, b any) {
	valA := reflect.ValueOf(a)
	for valA.Kind() == reflect.Ptr {
		valA = valA.Elem()
	}

	valB := reflect.ValueOf(b) // 获取指针指向的值
	for valB.Kind() == reflect.Ptr {
		if valB.IsNil() {
			valB.Set(reflect.New(valB.Type().Elem())) // 将新创建的指针赋值给 valB
		}
		valB = valB.Elem()
	}

	if valA.Type().Kind() == reflect.Slice { // 如果a是切片，则遍历切片
		if valB.Len() < valA.Len() { // 如果b的长度小于a的长度，则扩容
			valB.Set(reflect.MakeSlice(valB.Type(), valA.Len(), valA.Len()))
		}
		for i := 0; i < valA.Len(); i++ {
			CopyFields(valA.Index(i).Interface(), valB.Index(i).Addr().Interface())
		}
		return
	}

	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Type().Field(i)          // 获取字段类型信息
		fieldB := valB.FieldByName(fieldA.Name) // 根据字段名获取B中的字段

		if fieldB.IsValid() && fieldB.Type() == fieldA.Type { // 检查字段是否存在且类型相同
			//如果是结构体，递归调用copyFields
			if fieldA.Type.Kind() == reflect.Struct {
				CopyFields(valA.Field(i).Interface(), fieldB.Addr().Interface())
			} else {
				fieldB.Set(valA.Field(i)) // 赋值
			}
		} else if fieldB.Kind() == reflect.Interface { // 字段类型为interface{}，尝试赋值
			fieldB.Set(reflect.ValueOf(valA.Field(i).Interface()))
		} else if fieldB.Kind() == reflect.Slice && fieldA.Type.Kind() == reflect.Slice { //如果是切片继续递归调用copyFields
			CopyFields(valA.Field(i).Interface(), fieldB.Addr().Interface())
		}
	}
}
