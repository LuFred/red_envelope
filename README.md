# 口令红包服务

## 完成：
* 用户创建红包  
* 用户领取红包  
* 用户查看余额  
* 用户查看红包领取历史记录

## 未完成  
* 将过期红包数据插入红包过期清退表
* 处理红包过期清退表数据

## 整体项目重点细节设计  
> 架构：api+redis+rpc红包服务  
> api通过grpc调用红包服务实现红包创建、领取、余额查看、历史记录查询  
> 创建的红包存入redis中  
> 金额实施计算后从从红包余额中扣除[规则：最小=1分，最大=(余额-剩余红包数)／剩余红包数]  
> redis维护每个红包的用户领取历史    
> redis通过加锁防止并发领取同一个红包溢出的情况  

## 项目目录结构
### Service: 服务目录
### util：相关涉及工具类
### Service/api_service: api服务
* config：配置文件目录  
* core：api核心目录，包括自定义context，自定义handle函数  
* handler：路由实际执行函数(即MVC中的C)  
* logic_service:业务逻辑层，供handler调用
  * 红包处理逻辑就在本层
* microservice_client:rpc服务客户端实例创建层
* middleware:中间件层 包括auth、cors
* model：即MVC中的M dto对象定义层
* router：路由定义层
* swagger：api展示层 
  * api-json：api列表json文件  

### Service/red_envelope_service: 红包数据操作服务  
* config：配置文件目录  
* db: 数据库操作层  
  * core:数据库操作对象封装层
  * entity:数据库表实体映射即CRUD基本操作封装
* doc:存放初始化表结构脚本
* handle：rpc函数实现层
* proto：protobuf定义层
