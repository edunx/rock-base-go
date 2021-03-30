package base

import (
	"github.com/edunx/lua"
	pub "github.com/edunx/rock-public-go"
	"os"
	"os/signal"
	"syscall"
)

var hub = service{}

type Notify interface {
	Close()
}

type service []Notify

func (s *service) add( v Notify ) {
	a := *s
	a = append(a , v)
	*s = a
}

func (s *service) close() {
	a := *s
	for i := 0; i< len(a); i++ {
		a[i].Close()
	}
}

func reg(L *lua.LState , args *lua.Args) lua.LValue {
	if args == nil {
		return nil
	}

	n := args.Len()
	for i := 1 ; i<=n ; i++ {
		ud := args.CheckLightUserData( L , i )
		v ,ok := ud.Value.(Notify)
		if !ok {
			L.RaiseError("invalid type , #%d must be notify , got %s" , i , ud.Type().String())
		}
		hub.add(v)
	}
	return nil
}

func notify(L *lua.LState , args *lua.Args) lua.LValue {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c
	hub.close()
	pub.Out.Err("exit by sigal %v", s)
	return nil
}

func injectSystem( L *lua.LState ) *lua.UserKV {

	obj := &lua.UserKV{}
	obj.Set("notify" , lua.NewGFunction( notify ) )
	obj.Set("reg"    , lua.NewGFunction(reg))
	return obj
}