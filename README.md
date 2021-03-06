# tmaic saas 多租户（Multi-Tenant ）框架

#### 介绍
tmaic saas多租户（Multi-Tenant ），即一个Tenant，一个Database“的数据存储方式。隔离级别最高、最安全，说到底是要解决数据存储的问题，golang语言saas框架唯数不多的框架之一。
#### 软件架构

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

### 校验密码是否正确
user.Password：用户密码（已加密）

requestData.Password：要校验的密码（明文）
```golang
if !crypt.BcryptCheck(user.Password, requestData.Password) error{
    
    return "返回提示信息"
}
```



### 注意事项：如果要使用 NSQ 队列的，不用将包注释

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

### NSQ安装请参考：https://nsq.io/overview/quick_start.html
##### nsq启动

1.nsqlookupd

2.nsqd --lookupd-tcp-address=127.0.0.1:4160

3.nsqadmin --lookupd-http-address=127.0.0.1:4161

执行以上命令后：http://127.0.0.1:4171 可查看监控台

### 队列demo (test)

*** 详情见工程项目下

### 入列（我们常说的写入 topic,生产者），以下写入topic例子

```golang 

	test := events.Test{}

	testparam := &pbs.Test{
		Id: uint32(userId),
	}
	test.SetParam(testparam)

	if errs := hub.Emit(&test); errs != nil {
		log.Info("user test", tmaic.Output{"event": test, "errors": errs})
	}

```



### 出列(启动一个观察者)


```golang 
 
 
func init() {
	hub.Register(&Test{})
}

type Test struct {
	user models.User
	hub.Listen
}

func (user *Test) Name() hub.ListenerName {
	return "add-test" //监听器名称 后面我们会用到
}

func (user *Test) Subscribe() (eventPtrList []hub.Eventer) {
	log.Debug("Subscribe-test")
	return []hub.Eventer{
		&events.Test{},
	}

}

func (user *Test) Construct(paramPtr proto.Message) error {
	/** 第一执行这里
	  业务代码
	*/
	log.Debug("Construct-test")
	return nil
}
func (user *Test) Handle() error {
	/**第二执行这里
	 最终实现业务逻辑 ：比如发邮件、消息推送、短信通知等 这些业务通常封装在 service 层，这里只是建议
	*/
	
	// 更新
	t := users.UserService.Get(cast.ToInt64(user.ID))
	t.Email = user.Email //test 其实没有这个字段 根据需要自行组织
	t.Name = user.Name   //test 其实没有这个字段 根据需要自行组织
	_ = users.UserService.Update(t)
	
	log.Debug("Handle-test")
	return nil
}

```
### 实现消费,在根目录与main.go同级，

go run artisan.go queue:listen test-add //test-add 是Name() 返回的
配合supervisor进程守护


## 类型转化助手
### ToInt64
```golang
cast.ToInt64()
```
### ToInt32
```golang
cast.ToInt32()
```
### ToFloat32
```golang
cast.ToFloat32()
```
### 更多方法 查看cast包	

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



# 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


请赞助本项目

如你觉有收获，请给我打赏
