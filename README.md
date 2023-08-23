# 测试题：基于 FB message api 的 chatbot
## 需求分析
需要实现 2个 关键功能：

1. 收集用户反馈并存储

2. 根据特定事件推送消息

## 需求拆解
 
 1. 通过 api 接收来自公共主页用户的 message 
 2. 将 message，consumer info 写入数据库存储
 3. 对 message 数据进行展现
 4. 监听外部事件（mq or http 接口）
 5. 发送模版消息到用户

## 其他需求
### 部署需求
第一期需求是云服务环境部署，或者至少 docker 部署；第二期仅要求机器部署
### 测试需求
要实现单元测试
### 语言限制
golang
### bonus
语义情感分析

## 方案设计
### 技术调研和难点
1. facebook messenger platform 需要应用暴露 “/webhook” https 接口，供消息推送和身份验证，因此首先需要一个稳定的网络地址，另外需要对应证书。\
2. 云服务（以 aws 为例）提供的 ecs 和 eks 各自有差别，需要针对性学习。\
3. Go 语言框架都比较简易，gin 和 gorm 使用较多，但缺乏良好的项目结构。\
4. ML 不了解，云服务厂商通常有方案，之前 aws 就有鉴别黄色和暴力内容的服务，但时间有限

### 解决方案
1. 公共 ip 即使通过动态域名也太麻烦了，因此购买云服务是最好的方案；另外，自签名证书可能会有问题，Let's encrypt 证书需要域名，我找到了 zerossl 签发了一个 IP ssl 证书。
2. 搭建一套简易的 k8s 环境，降低心智负担，我在 azure 上购买了一台 2u 16g 的虚拟机，用 rancher 快速启动了一套 k3s （精简版 k8s）环境，服务开放 nodeport, 通过 nginx 反向代理到 service。
3. 项目采用单体结构（monolith），但设计上可以拆分多个模块，比如 message-admin模块，Facebook messenger 模块，event listener模块。对于 gin 项目结构最佳实践我比较生疏了，我按照习惯的 spring mvc 的方式拆分项目结构，主要按 constroller，service，dao 分层。
4. 数据库没时间做太多选型，直接用 mysql rds 存储，拆分了 message 表 和 consumer 表，直接通过 gorm 生成了表结构。fb api 对消息的发送评论有限制，预估可能写入数量不高。如果有大量写入的 OLAP 场景，可以再接入列式数据库。
5. 关于事件触发后发送模版消息的需求，理想情况是该应用接入 mq，达到解耦合和流量平滑的效果，但时间有限，我仅暴露一个 "/api/message" 的 post 接口发送文本消息。
6. 业务熔断降级，目前主要的外部接口是 fb api，我看到 Java 体系的 hystrix 在 Go 下也有实现，但时间有限不展开。
7. openapi 接入，简单看了看 go swagger，不如 Java 原生注解方便，未接入。

### 接口定义
1. GET "/consumer": 传入 consumer_id 获取单个用户的 profile 信息，不传该参数则获取所有用户的 profile 信息
2. GET "/message": 传入 consumer_id 获取单个用户的所有消息，不传该参数则获取所有用户的消息
3. POST "/message": 按 form-data 传入 consumer_id，msg_text 可以给 consumer_id 用户发送内容为 msg_text 的文本消息。预留的 event_type 字段用来支持实践类型匹配不同模版（TODO）

### 测试方法
我申请的公共主页 https://facebook.com/erikxsc \
申请的云主机公网 ip 为 52.184.83.251，443 端口运行 rancher dashboard 服务，chatbot 应用运行在 8443 端口\
由于我的测试 app 未获得企业认证，因此仅有 app 绑定的开发者发送的消息才能到我的 app，因此数据目前仅有我自己(9805893816149728)给主页()发送的消息，但可以通过 curl 验证。\
1. curl --location 'http://52.184.83.251:8080/api/consumer'
2. curl --location 'https://52.184.83.251:8443/api/message'
3. curl --location 'https://52.184.83.251:8443/api/message' \
--form 'consumer_id="9805893816149728"' \
--form 'msg_text="hello world"' \
--form 'event_type=""'
