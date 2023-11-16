# goweb-docker-cg


## 快速开始
### go version
GO >= 1.13


### gomodule环境变量设置

```
env -w GO111MODULE=on
env -w GOPROXY=https://goproxy.cn,direct
```


### 服务构建和启动

* 使用`gowatch`实现热编译(推荐)：
```
# 提前安装
go get github.com/silenceper/gowatch
# 启动项目
gowatch
```

* 直接启动
```
go run main.go
```

### 生成你的项目 ✨✨✨✨

#### 命令行方式

在终端执行以下命令，你将在 `goweb-docker-cg` 同级目录得到名为 `projectName` 的项目。
```bash
sh cg.sh projectName
```

> 命令行方式可以基于 cg.sh (code generator) 快速复制出一个命名为 `projectName` 的项目，
> 项目包含各种使用示例，业务生成后一般会根据自己的需求进行裁剪--移除一些不需要使用资源配置、代码示例等。


## 框架规范

  强烈建议按照以下目录规范来规范你的项目：

 ```
|-goweb-docker-cg
    |-api 访问第三方请求目录
    |-components 组件，可被其他所有目录依赖
    |-conf 配置文件目录
        |-app 用户自定义配置，跟随代码发布的业务配置
        |-mount 用来放置环境相关的配置，可通过配置中心发布的配置
    |-controllers 控制器目录
        |-http http控制器目录
        |-command 任务控制器入口，包括cycle任务、crontab任务、一次性任务
        |-mq 消息队列回调入口
    |-data 数据层。当项目比较复杂时，可以增加data层用于组装数据，包括不限于数据库查询到的数据、api调用后查询到的数据
    |-helpers 公共类目录，可以用来初始化一些全局变量
    |-models 数据模型访问目录。数据库相关调用。
    |-middleware 业务中间件
    |-router 路由目录，一般对应controllers目录结构
        |-http http路由
        |-command 人物类路由
        |-mq 消息队列路由
    |-service 业务逻辑聚合目录。主要强调业务逻辑，能够看出一个功能的核心处理流程。
    |-sql mysql建库建表相关语句
    |-go.mod go module使用，记录项目的依赖
    |-go.sum go mod tidy 后生成，记录依赖的详细依赖
    |-main.go 程序执行入口
  ```