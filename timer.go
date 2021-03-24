package base

import (
	"github.com/edunx/lua"
	"time"
)

func timerSleep( L *lua.LState , args *lua.Args) lua.LValue {
	sleep := args.CheckInt(L , 1)

	time.Sleep(time.Duration(sleep * 1000 ) * time.Microsecond )
	return nil
}

func timerNow( L *lua.LState , args *lua.Args ) lua.LValue {
	if args == nil || args.Len() == 0  {
		return lua.LNumber( time.Now().Unix() )
	}

	format := args.CheckString(L , 1)
	return lua.LString(time.Now().Format( format ))
}

func injectTimer( L *lua.LState ) *lua.UserKV {
	obj := &lua.UserKV{}
	obj.Set("now" ,   lua.NewGFunction( timerNow ))
	obj.Set("sleep" , lua.NewGFunction( timerSleep ))
	return obj
}