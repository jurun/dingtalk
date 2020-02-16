## 接口体验测试
> 可以使用钉钉提供的工具进行测试：
> [Api Explorer](https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=/get_jsapi_ticket)


## SDK调用方式
```go
import (
    "github.com/jurun/dingtalk"
    "github.com/jurun/dingtalk/user"
)

dingtalk.Option.AppKey    = "xxx"
dingtalk.Option.AppSecret = ""

userid,isSys,sysLevel,err := user.Authorized("xxxx")
if err != nil {
    //
}

```