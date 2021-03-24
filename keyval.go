package base

import (
	"github.com/edunx/lua"
)

type KeyVal struct {
	lua.Super
	Key    string
	Val    string
}

func (kv *KeyVal) ToLightUserData(L *lua.LState) *lua.LightUserData {
	return L.NewLightUserData(kv)
}

func (kv *KeyVal) LCheck(obj interface{} , set lua.LCallBack) bool {
	v , ok := obj.(*KeyVal)
	if ok {
		if set == nil {
			*v = *kv
			return true
		}
		set(kv)
		return true
	}

	return false
}

func (kv *KeyVal) debug(L *lua.LState , args *lua.Args) lua.LValue{
	return lua.LString("key:"+kv.Key + " , val:" + kv.Val)
}

func (kv *KeyVal) Index(L *lua.LState , key string) lua.LValue {

	if key == "key"   { return lua.LString( kv.Key )}
	if key == "val"   { return lua.LString( kv.Val )}
	if key == "debug" { return lua.NewGFunction( kv.debug )}

	L.RaiseError("invalid key %s in keyval" , key)
	return nil
}

func (kv *KeyVal) NewIndex(L *lua.LState , key string , val lua.LValue) {

	if val.Type() != lua.LTString {
		L.RaiseError("invalid val, must be string , got %s" ,
			val.Type().String())
		return
	}

	if key == "key" { kv.Key = val.String() }
	if key == "val" { kv.Val = val.String() }
}

func createKeyValUserData(L *lua.LState , args *lua.Args) lua.LValue {
	key := args.CheckString(L , 1)
	val := args.CheckString(L , 2)

	return L.NewLightUserData( &KeyVal{ Key: key, Val: val} )
}

func injectKeyVal(L *lua.LState) *lua.GFunction {
	return lua.NewGFunction( createKeyValUserData )
}