package base

import (
	"fmt"
	"github.com/edunx/lua"
)


type SliceTypeValue int

const (
	ST_STRING SliceTypeValue = iota + 1
	ST_INT
	ST_KEYVAL
)

var sliceTypeNames = [3]string{"string" , "int" , "keyval"}

func (s *SliceTypeValue) String() string {
	return sliceTypeNames[int(*s)]
}

type SliceData struct {
	lua.Super
	t   SliceTypeValue
	v   []interface{}
	cap int
}

func(sd *SliceData) debug(L *lua.LState , args *lua.Args) lua.LValue {
	return lua.LString(fmt.Sprintf("%v" , sd.v))
}

func (sd *SliceData) add(L *lua.LState , args *lua.Args) lua.LValue {
	n := args.Len()
	if n <= 0 {
		return nil
	}

	v := make([]interface{} , n )
	for i:=0;i<n;i++ {
		switch sd.t {
		case ST_STRING:
			v[i] = args.CheckString(L , i + 1)
		case ST_INT:
			v[i] = args.CheckInt(L , i + 1)
		case ST_KEYVAL:
			v[i] = CheckKeyValByArgs(L , args , i + 1)
		default:
			L.RaiseError("#%d invalid type , must be keyval , got " , i + 1 , args.Get(i + 1).Type().String() )
			return nil
		}
	}

	sd.v = append(sd.v ,v...)
	sd.cap = sd.cap + n
	return nil
}

func (sd *SliceData) Index(L *lua.LState , key string) lua.LValue {
	if key == "debug" { return lua.NewGFunction( sd.debug ) }
	if key == "add" { return lua.NewGFunction( sd.add ) }
	return lua.LNil
}

func createSliceData(L *lua.LState , args *lua.Args , st SliceTypeValue ) lua.LValue {
	n := args.Len()
	if n <= 0 {
		return L.NewLightUserData( &SliceData{t: st , cap: 0} )
	}

	obj := &SliceData{
		t: st,
		v: make([]interface{} , n ),
		cap: n,
	}

	for i:=1 ; i<=n ; i++ {
		switch st {
		case ST_STRING:
			obj.v[i - 1] = args.CheckString(L , i)
		case ST_INT:
			obj.v[i - 1] = args.CheckInt(L , i)
		case ST_KEYVAL:
			obj.v[i - 1] = CheckKeyValByArgs(L , args , i)
		default:
			L.RaiseError("#%d invalid type , muset be string or int or keyval" , i)
			return nil
		}
	}

	return L.NewLightUserData( obj )
}

func createStrSlice(L *lua.LState , args *lua.Args) lua.LValue {
	return createSliceData(L , args , ST_STRING )
}

func createIntSlice(L *lua.LState , args *lua.Args) lua.LValue {
	return createSliceData(L , args , ST_INT )
}

func createKeyValSlice(L*lua.LState , args *lua.Args) lua.LValue {
	return createSliceData(L , args , ST_KEYVAL )
}