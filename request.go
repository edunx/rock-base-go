package base

import (
	"github.com/edunx/lua"
	"github.com/go-resty/resty/v2"
)

const (
	MethodGet     string = "GET"
	MethodPatch   string = "PATCH"
	MethodTrace   string = "TRACE"
	MethodOptions string = "OPTIONS"
	MethodDelete  string = "DELETE"
	MethodHead    string = "HEAD"
	MethodPut     string = "PUT"
	MethodPost    string = "POST"
)

type httpResponse struct {
	lua.Super
	rc *resty.Response
	err error
}

func (hr *httpResponse) Index( L *lua.LState , key string ) lua.LValue {
	if key == "code" { return lua.LNumber(hr.rc.StatusCode()) }
	if key == "body" { return lua.LString(hr.rc.Body())       }
	if key == "err"  { return lua.LString(hr.err.Error())     }

	return nil
}


type httpRequest struct {
	lua.Super

	client *resty.Client
	r  *resty.Request
}

func (r *httpRequest) ToLightUserData(L *lua.LState) *lua.LightUserData {
	return L.NewLightUserData(r)
}

func (r *httpRequest) output(L *lua.LState , args *lua.Args) lua.LValue {
	r.r.SetOutput(args.CheckString(L , 1))
	return L.NewLightUserData(r)
}

func (r *httpRequest) GET( L *lua.LState , args *lua.Args) lua.LValue  {
	return r.Execute(MethodGet , L , args)
}

func (r *httpRequest) POST( L *lua.LState , args *lua.Args) lua.LValue {
	return r.Execute(MethodPost , L , args)
}

func (r *httpRequest) PUT( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodPut ,L , args)
}

func (r *httpRequest) OPTIONS( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodOptions,L , args)
}
func (r *httpRequest) TRACE( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodTrace,L , args)
}

func (r *httpRequest) HEAD( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodHead,L , args)
}

func (r *httpRequest) DELETE( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodDelete,L , args)
}

func (r *httpRequest) PATCH( L *lua.LState , args *lua.Args) lua.LValue {
	return	r.Execute(MethodPatch ,L , args)
}
func (r *httpRequest) Execute(method string , L *lua.LState, args *lua.Args) lua.LValue {
	rc , err := r.r.Execute( method , args.CheckString(L , 1))
	resp := &httpResponse{ rc: rc , err: err}
	return L.NewLightUserData( resp )
}

func (r *httpRequest) Index( L *lua.LState , key string ) lua.LValue{ //lua代码获取对象的方法

	if key == "OPTIONS"     { return lua.NewGFunction( r.OPTIONS )  }
	if key == "DELETE"      { return lua.NewGFunction( r.DELETE  )  }
	if key == "PATCH"       { return lua.NewGFunction( r.PATCH   )  }
	if key == "TRACE"       { return lua.NewGFunction( r.TRACE   )  }
	if key == "POST"        { return lua.NewGFunction( r.POST    )  }
	if key == "HEAD"        { return lua.NewGFunction( r.HEAD    )  }
	if key == "GET"         { return lua.NewGFunction( r.GET     )  }
	if key == "PUT"         { return lua.NewGFunction( r.PUT     )  }

	if key == "output"      { return lua.NewGFunction( r.output  )  }

	return lua.LNil
}

func injectHttpRequest(L *lua.LState) *lua.LightUserData {
	client := resty.New()

	r := &httpRequest{
		client: client,
		r: client.R(),
	}

	return r.ToLightUserData( L )
}