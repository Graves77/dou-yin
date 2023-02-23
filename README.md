# dou-yin

# 基本的apps框架和目录
- 选择采用Go-zero的微服务框架
- 采用集中管理的方式，把所有的服务放到一个大仓库中
- dou-yin为工程名，下面有apps和pkg两个目录，其中apps存放的是我们所有的微服务，比如user为用户相关的微服务，pkg目录为所有服务共同依赖的包的存放路径，比如所有的服务都需要依赖鉴权就可以放到pkg目录下。

具体的apps目录：
- video- 视频流微服务

- commaction - 用户评论微服务

- favorite - 用户点赞微服务

- relationfollow - 用户关注微服务

- user - 用户管理微服务


# 抖音极速版日志：
### 2/7 基本的框架搭建
### 2/8 user模块的优化
### 2/11 返回处理
### 2/12 user进一步优化
### 2/13 user的rpc部分完成
### 2/15 user基本完成，video开始搭建
### 2/16 video的api和rpc的基本搭建
### 2/17 video基本功能实现
### 2/18 commaction的搭建
### 2/19 commaction的api搭建
### 2/20 commaction的rpc搭建
### 2/21 关注和关注列表的搭建
### 2/22 关注和关注列表的完善
### 2/23 代码优化
