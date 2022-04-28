# ginServer框架范例

    ginServer是一个基于 gin 开发的服务端应用，支持mysql,mongodb,rabbitmq,redis等主流数据库，适用于小型项目服务端应用。
   
## 目录结构
+ controller
  api接口控制层，负责从http reguest传入的map参数转换成service层调用的参数，调用service层函数实现api接口逻辑
+ dao
  对各类型数据表操作DAO层
+ global
  定义全局变量，包含全局系统配置及日志等
+ initialize
  项目运行时的初始化，包含创建各类型数据库的连接等
+ middleware
  系统需要的中间件，包含请求日志记录等
+ model 
  数据模型，存放各种对象结构体，如表对象、接口入参对象、出参对象等
+ router
  按模块配置独立的路由信息
+ service
  业务实现层，从controller层传入参数，根据业务需求调用dao等数据库层进行业务逻辑实现
+ Router.go
  http 路由配置入口，可扩展添加router目录内各业务模块配置的路由信息

## 创建新工程的方式

+ 首先采用Git命令行下载这个范例项目，假设自己的新工程名称为 test
    
```shell script
    git clone https://github.com/feiyangderizi/ginServer.git test
```
    
+ 修改 application.yml 文件名为 test.yml
    
```shell script
    cd test
    mv application.yml test.yml
```
    
+ 修改 main.go 中指定的配置文件名
    
```go
    const config_file = "test.yml"
```
    
+ 修改 test.yml 配置中的相应的内容

## 感谢

   本项目源于项目 <a href="https://github.com/maczh/mgin">mgin</a>，将微服务类型应用调整成适用于小型项目的服务端应用。
