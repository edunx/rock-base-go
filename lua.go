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



func LuaInjectApi(L *lua.LState , parent *lua.UserKV) {

	parent.Set("system" ,  injectSystem(L)      )
	parent.Set("base" ,    injectBase( L )      )
	parent.Set("slice",    injectSlice( L )     )
	parent.Set("request" , injectHttpRequest(L) )
	parent.Set("stdout" ,  injectStdout(L)      )
}