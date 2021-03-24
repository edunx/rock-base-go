package base

import "github.com/edunx/lua"

func CheckKeyValByArgs(L *lua.LState , args *lua.Args , idx int ) *KeyVal {
	ud := args.CheckLightUserData( L , idx )
	kv , ok := ud.Value.(*KeyVal)
	if ok {
		return kv
	}
	L.RaiseError("invalid type , #%id must be keyval" , idx)
	return nil
}
