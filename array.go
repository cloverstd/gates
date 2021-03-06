package gates

import (
	"math"
	"reflect"
	"strings"
)

var sharedRuntime Runtime

type _Array struct {
	values []Value
}

func Array(values []Value) Value {
	return _Array{
		values: values,
	}
}

func (_Array) IsString() bool   { return false }
func (_Array) IsInt() bool      { return false }
func (_Array) IsFloat() bool    { return false }
func (_Array) IsBool() bool     { return false }
func (_Array) IsFunction() bool { return false }

func (a _Array) ToString() string {
	stringSl := make([]string, 0, len(a.values))
	for _, v := range a.values {
		stringSl = append(stringSl, sharedRuntime.ToValue(v).ToString())
	}
	return strings.Join(stringSl, ",")
}

func (_Array) ToInt() int64         { return 0 }
func (_Array) ToFloat() float64     { return math.NaN() }
func (a _Array) ToNumber() Number   { return Float(a.ToFloat()) }
func (_Array) ToBool() bool         { return true }
func (_Array) ToFunction() Function { return _EmptyFunction }

func (a _Array) ToNative() interface{} {
	result := make([]interface{}, 0, len(a.values))
	for _, value := range a.values {
		result = append(result, value.ToNative())
	}
	return result
}

func (a _Array) Equals(other Value) bool {
	o, ok := other.(_Array)
	if !ok {
		return false
	}
	return reflect.DeepEqual(a.values, o.values)
}

func (a _Array) SameAs(other Value) bool { return false }

func (a _Array) Get(r *Runtime, key Value) Value {
	i := key.ToNumber()
	if i.IsInt() {
		ii := i.ToInt()
		if ii < 0 || ii >= int64(len(a.values)) {
			return Null
		}
		return a.values[ii]
	}

	switch key.ToString() {
	case "length":
		return Int(len(a.values))
	}

	return Null
}

func (a _Array) Set(r *Runtime, key, value Value) {
	if !key.IsInt() {
		return
	}
	i := key.ToInt()
	if i < 0 || i >= int64(len(a.values)) {
		return
	}
	a.values[i] = value
}
