package base

import (
	"github.com/edunx/lua"
)

func injectBase(L *lua.LState) *lua.UserKV {
	base := &lua.UserKV{}
	base.Set("keyval" , injectKeyVal(L)     )
	base.Set("timer"  , injectTimer(L)      )
	return base
}


func injectSlice(L *lua.LState) *lua.UserKV {
	obj := &lua.UserKV{}
	obj.Set("keyval" , lua.NewGFunction( createKeyValSlice ))
	obj.Set("str"    , lua.NewGFunction( createStrSlice ))
	obj.Set("int"    , lua.NewGFunction( createIntSlice ))
	return obj
}



func LuaInjectApi(L *lua.LState , parent *lua.LTable) {

	L.SetField(parent , "system" ,  injectSystem(L)      )
	L.SetField(parent , "base" ,    injectBase( L )      )
	L.SetField(parent , "slice",    injectSlice( L )     )
	L.SetField(parent , "request" , injectHttpRequest(L) )
	L.SetField(parent , "stdout" ,  injectStdout(L)      )
}