# 将数据发送飞书表格

> 本程序所有的数据都存储在内存中。 
> 
> 本程序通过启动http服务器，用户通过请求http的方式向飞书表格发送数据。

## 免责声明

> 该项目是开源项目，如果您在使用的过程中出现问题我是概不负责。

## 文档

[帮助文档](https://docs.notificationfilter.com/docs/category/%E5%88%86%E4%BA%AB%E5%88%B0%E9%A3%9E%E4%B9%A6%E8%A1%A8%E6%A0%BC---%E6%8F%92%E4%BB%B6)

## 使用

在发送数据之前请设置飞书配置。

1. 设置飞书配置：

- 访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”。

2. 获取表格工作表Id：

- 访问“http://你的IP或域名/sheets/密码”。

3. 发送数据：

- 访问“http://你的IP或域名/send/密码”。

## 进阶:

1. 达到设置存储的数据条数后再同步的飞书表格：

- 当内存中存储了10条数据就会将这10条数据一次性发送到飞书表格。如果没有就继续等待满足了后再发送。

- 访问“http://你的IP或域名/sends/密码”。

## QA:

1. 密码有什么用？

- 密码用于防止其他用户恶意插入数据或者造成其他损失。

2. 如何获取密码？

- 访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”地址会返回128位随机字符串。

3. 忘记密码该怎么办？

- 重新启动该服务。再次访问“http://你的IP或域名/set/表格Id/应用Id/应用Secret”。

4. 我想修改飞书配置该如何做？

- 访问“http://你的IP或域名/update/表格Id/应用Id/应用Secret/密码”。

## API说明

### 设置飞书配置
```http request
GET http://IP或域名:8000/set/表格Id/应用Id/应用Secret
```

### 更新飞书配置
```http request
GET http://IP或域名:8000/set/表格Id/应用Id/应用Secret/密码
```

### 查询表格信息
```http request
GET http://IP或域名:8000/sheets/密码
```

### 发送单条通知
```http request
POST http://IP或域名:8000/send/密码
Content-Type: application/json

{
  "valueRange": {
    "range": "工作表Id",
    "values": [
      [1, "发送单条通知", ...]
    ]
  }
}
```

### 发送多条通知
```http request
POST http://localhost:8000/sends/工作表Id/密码
Content-Type: application/json

[1, "发送多条通知", ...]
```