# Change Log

A distributed middleware to promote redis concurrent efficiency

# [version:0.0.1]
### Features
初始框架

# [version:0.0.2]
### Features
实现基于zookeeper的分布式读写锁

# [version:0.0.3]
### Features
引入gin框架，实现部分rest接口框架

# [version:0.0.4]
### Features
实现redis读写接口

# [version:0.0.5]
### Features
引入grpc框架

# [version:0.0.6]
### Features
实现基本Set/Get功能

# [version:0.0.7]
### Features
修复grpc错误

# [version:0.0.8]
### Features
优化zookeeper连接，采用长连接

# [version:0.0.9]
### Features
分布式读写方案修改为 分布式写锁+本地读写锁

# [version:1.0.0]
### Features
完成性能测试，提供初始可用版本

# [version:1.0.1]
### Features
Dockerfile

# [version:1.0.2]
### Features
更新vendor

# [version:1.0.3]
### Features
优化zookeeper获取锁流程，获取锁失败则删除当前节点

# [version:1.0.4]
### Features
初始化时，删除zookeeper锁目录下的所有子节点

# [version:1.0.5]
### Features
增加定时清理超时锁和待处理事务问题