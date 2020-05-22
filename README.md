## golang-mvc


### iris 

[iris 11.1.1](https://github.com/kataras/iris)
[iris中文网](https://studyiris.com/example/iris.html)

获取：

        github.com/kataras/iris
        github.com/kataras/golog
        github.com/kataras/pio   

### mysql/go-xorm

        _ "github.com/go-sql-driver/mysql"
        "github.com/go-xorm/core"
        "github.com/go-xorm/xorm"
              
### Casbin

[casbin](https://github.com/casbin/casbin)

[Iris + Casbin 权限控制实战](http://zxc0328.github.io/2018/05/14/casbin-iris/?utm_source=tuicool&utm_medium=referral)

### jwt-go

[jwt-go](https://github.com/dgrijalva/jwt-go)

[token失效解决](https://segmentfault.com/q/1010000010043871)

### go-yaml

    "github.com/go-yaml/yaml"  

### redis

[redis](https://github.com/go-redis/redis.git)


### sms

[alibaba-cloud-sdk](https://github.com/aliyun/alibaba-cloud-sdk-go.git)

### Golang 定时任务

[从99.9%CPU浅谈Golang的定时器实现原理](https://www.jianshu.com/p/c9e8aaa13415)

[golang 实现定时任务](https://www.cnblogs.com/jssyjam/p/11910851.html)


## 项目版本依赖说明

1、casbin 更新对casbin中间件的更新 

## 问题

1、go-xorm group by Vs FindAndCount()
    
    解决办法： 不用group by ,采用 sql 其他过滤方法解决

2、go mod 时出现dial tcp 216.239.37.1:443: connectex: A connection attempt failed这种错误

		最后是因为set GOSUMDB=sum.golang.org
		先go env 看 GOSUMDB 的配置

		go env -w GOSUMDB=off
		把他关掉就好了

