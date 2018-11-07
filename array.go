package gates

import (
	"math"
	"strings"
)

var sharedRuntime Runtime

type Array []interface{}

func (Array) IsString() bool { return false }
func (Array) IsInt() bool    { return false }
func (Array) IsFloat() bool  { return false }
func (Array) IsBool() bool   { return false }

func (a Array) ToString() string {
	stringSl := make([]string, 0, len(a))
	for _, v := range a {
		stringSl = append(stringSl, sharedRuntime.ToValue(v).ToString())
	}
	return strings.Join(stringSl, ",")
}

func (Array) ToInt() int64       { return 0 }
func (Array) ToFloat() float64   { return math.NaN() }
func (a Array) ToNumber() Number { return Float(a.ToFloat()) }
func (Array) ToBool() bool       { return true }

func (a Array) Equals(other Value) bool { return Value(a) == other }
func (a Array) SameAs(other Value) bool { return a.Equals(other) }

func (a Array) Get(r *Runtime, key Value) Value {
	if a == nil {
		return Null
	}
	i := key.ToNumber()
	if !i.IsInt() {
		return Null
	}
	ii := i.ToInt()
	if ii < 0 || ii >= int64(len(a)) {
		return Null
	}
	return r.ToValue(a[ii])
}