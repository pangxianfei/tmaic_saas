# tmaic

框架思想借鉴：https://gitee.com/daitougege/GinSkeleton

golang iris版本：

https://gitee.com/pangxianfei/tmaic-iris-mvc

#### 介绍
tmaic 是一套简洁、优雅的Golang Web开发框架(GoLang Web Framework)。支持mysql,mssql等多类型数据库，它可以让你从面条一样杂乱的代码中解脱出来；它可以帮你构建一个完美的网络应用，而且每行代码都可以简洁、富于表达力。

#### 软件架构


# 问题反馈

1.提交问题请在项目顶栏的issue直接添加问题，基本上都是每天处理当天上报的问题。

2.本项目优先关注 https://gitee.com/pangxianfei/tmaic 仓库的所有问题, github 太卡严重影响效率。

3.QQ群：50403087 欢迎喜欢gin框架go开发者一期参与讨论.

# 招募共同开发者

参与方式：简单的东西直接提交PR,如果想法比较多，需要改动大段代码，你也可以直接加我 qq：421339244 ，直接添加至开发组，共同商讨开发的功能，约定规范，提交代码。

# 数据模型  
### 以下是例子
```golang

package models

import (
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/helpers/ptr"
	"github.com/pangxianfei/framework/helpers/zone"
	"github.com/pangxianfei/framework/model"
)
type User struct {
    model.BaseModel
    ID        *uint      `gorm:"column:user_id;primary_key;auto_increment"`
    Name      *string    `gorm:"column:user_name;type:varchar(100)"`
    Email     *string    `gorm:"column:user_email;type:varchar(100);unique_index;not null"`
    Password  *string    `gorm:"column:user_password;type:varchar(100);not null"`
    CreatedAt *zone.Time `gorm:"column:user_created_at"`
    UpdatedAt zone.Time  `gorm:"column:user_updated_at"`
    DeletedAt *zone.Time `gorm:"column:user_deleted_at"`
}

func (user *User) TableName() string {
    return user.SetTableName("user")
}

func (user *User) SetNameAttribute(value interface{}) {
    user.Name = user.Email
}

func (user *User) GetUpdatedAtAttribute(value interface{}) interface{} {
    return user.UpdatedAt //查询后这里可以其他处理，如格式化处理，如果是时间time 可以格式成 2020-10-01 18：00：20
}
```

# 生成公钥文件

tmaic.CreateRsaKey()

# 系统函数
### 加密解密函数
### 公钥加密

s,_:= tmaic.Encryption("Golang使用RSA进行公钥加密私钥解密,私钥加密公钥解密的实现")

### 私钥解密
dd,_ := tmaic.Decrypt(s)

### 私钥加密
s,_:= tmaic.PrivateEncryption("Golang使用RSA进行公钥加密私钥解密,私钥加密公钥解密的实现")

### 公钥解密
dd,_ := tmaic.PublicDecrypt(s)

### 注意事项：如果要使用 NSQ 队列的，不用将包注释


### 校验密码是否正确
user.Password：用户密码（已加密）

requestData.Password：要校验的密码（明文）
```golang
if !crypt.BcryptCheck(user.Password, requestData.Password) error{
    
    return "返回提示信息"
}
```

导入包：

"github.com/pangxianfei/framework/queue"

"tmaic/app/events"

"tmaic/app/jobs"

"tmaic/app/listeners"

以下去掉注释即可：

//queue.Initialize()

//jobs.Initialize()

//events.Initialize()

//listeners.Initialize()
### 打印

tmaic.Dump(mugs)
```golang
    mugs := map[string]interface{} {
        "password2" : map[string]string {
        "password3" : "The name cannot be empty",
        },
    }
    
    //效果
    {
        "password2": {
            "password3": "The name cannot be empty"
        }
    }
    
```


### NSQ安装请参考：https://nsq.io/overview/quick_start.html


# 性能测试

### 虚拟机环境下：
![avatar](/docs/img/centos.png)

### 压力测试
ab -c 1000 -n 5000  -k http://127.0.0.1/all

![avatar](/docs/img/cpu1.jpg)




# 安装教程

1.    git clone git@gitee.com:pangxianfei/tmaic.git
2.    go mod init // 初始化go mod 包管理
3.    go mod tidy // 加载依赖包
4.    go mod vendor // 将依赖包拷贝到项目目录中去
5.    go run main.go

# 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

# 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


