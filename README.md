# goview-gozero

## 一、介绍

GoView Go 后台服务。基于 [gozero](https://github.com/zeromicro/go-zero) 框架, 实现后端接口。

### GoView

![img.png](.github/img/img.png)

[GoView](https://gitee.com/dromara/go-view/tree/master-fetch/) 是一个高效的拖拽式低代码数据可视化开发平台，将图表或页面元素封装为基础组件，无需编写代码即可制作数据大屏，减少心智负担。当然低代码也不是
“银弹”，希望所有人员都能理智看待此技术。

- 文档地址: https://www.mtruning.club/
- 演示地址:  Demo 地址：https://vue.mtruning.club/

## 二、各服务信息

### 一、接口文档

**1. Restful 接口文档**
- [swagger 接口文档](restful%2Fapi%2Frest.swagger.json)
- [接口文档.md](restful%2Fapi%2Frest.swagger.md)

**2. RPC 接口文档**
- [user-rpc.md](service%2Fuser%2Fpb%2Fdoc.md)
- [screen-rpc.md](service%2Fscreen%2Fpb%2Fdoc.md)

### 二、API 服务

| 服务名称    | 端口   | 备注 |
|---------|------|----|
| Restful | 8888 |    |

### 三、RPC 服务

| 服务名称   | 端口   | 备注       |
|--------|------|----------|
| User   | 8080 | 用户服务 RPC |
| Screen | 8090 | 大屏服务 RPC |
