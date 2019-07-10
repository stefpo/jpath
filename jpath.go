package jpath

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"github.com/stefpo/econv"
)

func Parse(js string) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	if e := json.Unmarshal([]byte(js), &obj); e == nil {
		return obj, e
	} else {
		return nil, e
	}
}

func Stringify(obj interface{}) (string, error) {
	if s, e := json.MarshalIndent(obj, "", "  "); e == nil {
		return string(s), nil
	} else {
		return "", e
	}
}

func Get(obj interface{}, path string) (interface{}, error) {

	for path[0] == '/' && len(path) > 1 {
		path = path[1:]
	}
	if path[0] == '/' {
		return nil, fmt.Errorf("Invalid path")
	}

	p := strings.IndexByte(path, '/')

	if p >= 0 {
		key := path[0:p]
		pathLeft := path[p:]
		//fmt.Println("Branch", key)
		switch obj.(type) {
		case map[string]interface{}:
			//fmt.Println("Map")
			o, ok := obj.(map[string]interface{})[key]
			if ok {
				return Get(o, pathLeft)
			} else {
				return nil, fmt.Errorf("Value not found %v", path)
			}
		default:
			return nil, fmt.Errorf("Value not found %v", path)
		}
	} else {
		//fmt.Println("Leaf", path)
		switch obj.(type) {
		case map[string]interface{}:
			//fmt.Println("Map")
			return (obj.(map[string]interface{}))[path], nil
		default:
			//fmt.Println("Value")
			return nil, fmt.Errorf("Value not found %v", path)
		}

	}
}

func Set(obj interface{}, path string, val interface{}) error {

	for path[0] == '/' && len(path) > 1 {
		path = path[1:]
	}
	if path[0] == '/' {
		return fmt.Errorf("Invalid path")
	}

	p := strings.IndexByte(path, '/')

	if p >= 0 {
		key := path[0:p]
		pathLeft := path[p:]
		//fmt.Println("Branch", key)
		switch obj.(type) {
		case map[string]interface{}:
			//fmt.Println("Map")
			om := obj.(map[string]interface{})
			o, ok := om[key]
			if !ok {
				//fmt.Println("New")
				o = make(map[string]interface{})
				om[key] = o
			}
			return Set(o, pathLeft, val)
		default:
			return fmt.Errorf("Path conflict")
		}
	} else {
		//fmt.Println("Leaf", path)
		switch obj.(type) {
		case map[string]interface{}:
			//fmt.Println("Map")
			om := obj.(map[string]interface{})
			om[path] = val
			return nil
		default:
			//fmt.Println("Value")
			return fmt.Errorf("Path conflict")
		}

	}
}



// Map is abreviated type name
type Map = map[string]interface{}

// Slice is abreviated type name
type Slice = []interface{}

// GetMap is Get that returns a Map instead of interface{}
func GetMap(o interface{}, path string) Map {
	if v, e := Get(o, path); e != nil {
		switch v.(type) {
		case Map:
			return v.(Map)
		default:
			return nil
		}
	} else {
		return nil
	}
}

// GetSlice is Get that returns a Slice instead of interface{}
func GetSlice(o interface{}, path string) Slice {
	if v, e := Get(o, path); e != nil {
		switch v.(type) {
		case Slice:
			return v.(Slice)
		default:
			return nil
		}
	} else {
		return nil
	}
}

// GetString is Get that returns a string instead of interface{}
func GetString(o interface{}, path string) string {
	return GetString2(o, path, "")
}

// GetFloat is Get that returns a float64 instead of interface{}
func GetFloat(o interface{}, path string) float64 {
	return GetFloat2(o, path, 0)
}

// GetInt is Get that returns a int64 instead of interface{}
func GetInt(o interface{}, path string) int64 {
	return GetInt2(o, path, 0)
}

// GetBool is Get that returns a bool instead of interface{}
func GetBool(o interface{}, path string) bool {
	return GetBool2(o, path, false)
}

// GetUInt is Get that returns a uint instead of interface{}
func GetUInt(o interface{}, path string) uint64 {
	return GetUInt2(o, path, 0)
}

// GetTime is Get that returns a time.Time instead of interface{}
func GetTime(o interface{}, path string) time.Time {
	return GetTime2(o, path, time.Now())
}

// GetString2 is GetString with custom default value
func GetString2(o interface{}, path string, def string) string {
	if v, e := Get(o, path); e != nil {
		return econv.ToString(v)
	} else {
		return def
	}
}

// GetFloat2 is GetFloat with custom default value
func GetFloat2(o interface{}, path string, def float64) float64 {
	if v, e := Get(o, path); e != nil {
		return econv.ToFloat64(v)
	} else {
		return def
	}
}

// GetInt2 is GetInt with custom default value
func GetInt2(o interface{}, path string, def int64) int64 {
	if v, e := Get(o, path); e != nil {
		return econv.ToInt64(v)
	} else {
		return def
	}
}

// GetBool2 is GetBool with custom default value
func GetBool2(o interface{}, path string, def bool) bool {
	if v, e := Get(o, path); e != nil {
		return econv.ToBool(v)
	} else {
		return def
	}
}

// GetUInt2 is GetUInt with custom default value
func GetUInt2(o interface{}, path string, def uint64) uint64 {
	if v, e := Get(o, path); e != nil {
		return econv.ToUint64(v)
	} else {
		return def
	}
}

// GetTime2 is GetTime with custom default value
func GetTime2(o interface{}, path string, def time.Time) time.Time {
	if v, e := Get(o, path); e != nil {
		return econv.ToTime(v)
	} else {
		return def
	}
}

// FillStruct fills a structure from a jpath.Map
func FillStruct(o Map, dest interface{}) error {
	s,_ := Stringify(o)
	return json.Unmarshal([]byte(s), &dest)
}

// FromStruct create a jpath.Map from a structure
func FromStruct(o interface{}) (Map, error) {
	ret := Map{}
	s,_ := Stringify(o)
	e := json.Unmarshal([]byte(s), &ret)
	return ret, e
}
