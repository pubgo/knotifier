## 任务通知系统设计

## 对外调用方式
```
对外提供http API
对外能够接入mq等
```

## 数据库设计
### 通知信息管理
```
- 通知ID
- 通知名称
- 通知权限
- 通知执行方式
- 通知类型标签分类
- 通知相关链接和参数
- 最大通知次数
- 通知超时时间
```

### 回调任务管理系统
```
- 通知内容输入
- 通知状态(通知结果成功失败)
- 通知结果
- 本次任务ID
- 通知ID
- 创建时间
- 完成时间
- 本次任务执行标签
```