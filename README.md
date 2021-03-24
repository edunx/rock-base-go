# rock-base-go

## 说明
- 基础数据结构和方法

## 主要功能 

### rock.base.keyval
- 说明: 基础key , val数据结构
- 语法: rock.base.keyval(string , string)
```lua
    local kv = rock.base.keyval
    local v = kv("key1" , "value1")
    v.key = "name"
    v.val = "edunx"

    print(v.debug())
```
### rock.base.timer.*
- 函数: rock.base.timer.now
- 说明: 获取当前时间函数 
- 语法: timer.now() , timer.now("2006-01-02 15:04:05")
```lua
    local timer = rock.base.timer
    local now = timer.now()
    local format = timer.now("2006-01-02 15:04:05")
    print(now)
    print(format)
```
- 函数: rock.base.timer.sleep
- 说明: 延时函数 , 单位毫秒
- 语法: timer.sleep(1000)
```lua
    local timer = rock.base.timer
    timer.sleep(1000) --毫秒
```
   
### rock.request.*
`发起http请求的客户端`
- 函数： rock.request.GET,POST,PUT ...
- 说明： 发起对应动作的http请求
- 语法： rock.request.GET("www.baidu.com")
```lua
    local resp = rock.request.GET("www.baidu.com")
    print(resp.code)
    print(resp.body)
```
- 函数： rock.request.output
- 说明： 确定返回包的保存位置
- 语法： r = rock.request.output(string)
```lua
    local r = rock.request
    local resp = r.output("1.html").GET("www.baidu.com")
    print(resp.code)
    print(resp.body)
```
- 返回对象包含
```lua
    local response = r.GET("www.baidu.com") 
    print(response.code)
    print(response.body)
    print(response.err)
```

## rock.slice.*
`常用的分片结构的映射固定数据累心`
- 函数: rock.slice.str(string , string , string)
- 函数: rock.slice.int(int, int, int)
- 函数: rock.slice.keyval(kv , kv , kv)

- 返回数据对象包含:
```lua
    local slice = rock.slice
    local obj = slice.str("edunx" , "ok")
    obj.add("goo" , "d") --追加 
    obj.debug() --调试
    --其他对象类似
```

## rock.system.*   
`常见系统信号监听`
- 函数: rock.system.notify
- 说明: 监听syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT
- 语法: rock.system.notify
```lua
    local system = rock.system
    system.notify()
```

- 函数: rock.system.reg
- 说明: 服务注册功能， 当notify 监听到信号是会关闭注册后的服务， 注册的对象必须要有close方法
- 语法: rock.system.reg( obj , obj)
```lua
    local system = rock.system
    system.reg(ud1 , ud2)
    system.notify()

```

# 安装使用
```go
    import base "github.com/edunx/rock-base-go"

    //注入 API 函数 
    base.LuaInjecApi(L , rock)
```
