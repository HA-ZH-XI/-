package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/structs"
	"github.com/mitchellh/mapstructure"
	"reflect"
	"strings"
)

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc map转struct
 */
func MapToStruct(m map[string]any, t any) any {
	err := mapstructure.Decode(m, &t)
	if err != nil {
		return nil
	}
	return t
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc map转结构体
 */
func MapToStruct2(m map[string]any, t any) any {
	arr, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(arr, &t)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return t
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc struct转map
 */
func StructsToMap(content any) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				name[k] = v
			}
		}
	}
	return name
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc  struct转map
 */
func ReflectMethod(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc struct转map
 */
func StructsToMap2(obj interface{}) map[string]interface{} {
	return structs.ToMap(obj)
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc json转struct
 */
func JsonToStruct(jsonStr string, t any) any {
	err := json.Unmarshal([]byte(jsonStr), &t)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return t
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc json转struct
 */
func StructToJson(t any) string {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonBytes)
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc json转struct
 */
func JsonToMap(jsonStr string) map[string]any {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return mapResult
}

/**
 * @author feige
 * @date 2023-10-17
 * @version 1.0
 * @desc json转struct
 */
func MapToJson(t map[string]any) string {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonBytes)
}
