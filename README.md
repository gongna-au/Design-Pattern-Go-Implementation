# Design-Pattern-Go-Implementation

##  设计模式的go语言实现

#### 目标：

打算实现了一个简单的分布式应用系统（单机版），该系统主要由以下几个模块成：

- 网络 Network，网络功能模块，模拟实现了报文转发、socket通信、http通信等功能。

- 数据库 Db，数据库功能模块，模拟实现了表、事务、dsl等功能。

- 消息队列 Mq，消息队列模块，模拟实现了基于topic的生产者/消费者的消息队列。

- 监控系统 Monitor，监控系统模块，模拟实现了服务日志的收集、分析、存储等功能。

- 边车 Sidecar，边车模块，模拟对网络报文进行拦截，实现access log上报、消息流控等功能。

- 服务 Service，运行服务，当前模拟实现了服务注册中心、在线商城服务集群、服务消息中介等服务。

 设计模式包括23种设计模式：

-  工厂模式

-  原型模式

-  单例模式

-  装饰器模式...

| SOLID原则  | 单一职责原则（SRP）                   | registry.Registry         |
| ---------- | ------------------------------------- | ------------------------- |
|            | 开闭原则（OCP）                       | pipeline.Plugin           |
|            | 里氏替换原则（LSP）                   | pipeline.NewPlugin        |
|            | 接口隔离原则（ISP）                   | mq.Mq                     |
|            | 依赖倒置原则（DIP）                   | db.Db                     |
| 创建型模式 | 单例模式（Singleton）                 | network.network           |
|            | 建造者模式（Builder）                 | model.serviceProfileBuild |
|            | 工厂方法模式（Factory Method）        | sidecar.Factory           |
|            | 抽象工厂模式（Abstract Factory）      | config.Factory            |
|            | 原型模式（Prototype）                 | http.Request              |
| 结构型模式 | 适配器模式（Adapter）                 | db.TableRender            |
|            | 桥接模式（Bridge）                    | pipeline.pipelineTemplate |
|            | 组合模式（Composite）                 | pipeline.pipelineTemplate |
|            | 装饰者模式（Decorator）               | sidecar.FlowCtrlSidecar   |
|            | 外观模式（Facade）                    | shopping.Center           |
|            | 享元模式（Flyweight）                 | model.Region              |
|            | 代理模式（Proxy）                     | db.CacheProxy             |
| 行为型模式 | 责任链模式（Chain Of Responsibility） | filter.Chain              |
|            | 命令模式（Command）                   | db.Command                |
|            | 迭代器模式（Iterator）                | db.TableIterator          |
|            | 中介者模式（Mediator）                | mediator.Mediator         |
|            | 备忘录模式（Memento）                 | db.CmdHistory             |
|            | 观察者模式（Observer）                | network.socketImpl        |
|            | 状态模式（State）                     | flowctrl.state            |
|            | 策略模式（Strategy）                  | input.Plugin              |
|            | 模板方法模式（Template Method）       | flowctrl.stateTemplate    |
|            | 访问者模式（Visitor）                 | db.TableVisitor           |

主要目录结构如下：

```
├── db # 数据库模块，定义Db、Table、TableVisitor等抽象接口和实现
├── monitor # 监控系统模块，采用插件式的架构风格，当前实现access log日志etl功能
│   ├── config  # 监控系统插件配置模块
│   ├── filter # 过滤插件的实现定义
│   ├── input # 输入插件的实现定义
│   ├── output # 输出插件的实现定义
│   ├── pipeline # Pipeline插件的实现定义，一个pipeline表示一个ETL处理流程
│   ├── plugin # 插件抽象接口的定义，比如Plugin、Config等
│   └── model # 监控系统模型对象定义
├── mq # 消息队列模块
├── network  # 网络模块，模拟网络通信，定义了socket、packet等通用类型/接口 
│   └── http # 模拟实现了http通信等服务端、客户端能力
├── service # 服务模块，定义了服务的基本接口
│   ├── mediator # 服务消息中介，作为服务通信的中转方，实现了服务发现，消息转发的能力
│   ├── registry # 服务注册中心，提供服务注册、去注册、更新、 发现、订阅、去订阅、通知等功能
│   │   └── model # 服务注册/发现相关的模型定义
│   └── shopping # 模拟在线商城服务群的定义，包含订单服务、库存服务、支付服务、发货服务
└── sidecar # 边车模块，对socket进行拦截，提供http access log、流控功能
    └── flowctrl # 流控模块，基于消息速率进行随机流控
```