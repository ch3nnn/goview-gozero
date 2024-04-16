# 接口文档

## 接口概览（总计11个）

### restful/screen

| **路径**                                                   | **功能**                   | **请求方式** | **是否需要鉴权** |
|------------------------------------------------------------|----------------------------|--------------|------------------|
| [/api/goview/project/create](#创建大屏信息)                | 创建大屏信息               | POST         | true             |
| [/api/goview/project/delete](#根据大屏信息ID删除)          | 根据大屏信息ID删除         | DELETE       | true             |
| [/api/goview/project/edit](#更新大屏信息)                  | 更新大屏信息               | POST         | true             |
| [/api/goview/project/getData](#根据大屏信息ID获取大屏数据) | 根据大屏信息ID获取大屏数据 | GET          | false            |
| [/api/goview/project/list](#大屏信息列表)                  | 大屏信息列表               | GET          | true             |
| [/api/goview/project/publish](#更新大屏信息发布状态)       | 更新大屏信息发布状态       | PUT          | true             |
| [/api/goview/project/save/data](#保存大屏数据)             | 保存大屏数据               | POST         | false            |
| [/api/goview/project/upload](#文件上传)                    | 文件上传                   | POST         | true             |

### restful/sys

| **路径**                                   | **功能**    | **请求方式** | **是否需要鉴权** |
|--------------------------------------------|-------------|--------------|------------------|
| [/api/goview/sys/getOssInfo](#获取oss信息) | 获取oss信息 | GET          | false            |
| [/api/goview/sys/login](#用户登录)         | 用户登录    | POST         | false            |
| [/api/goview/sys/logout](#用户登出)        | 用户登出    | GET          | false            |

## 接口详情

### restful/screen

### 创建大屏信息

[返回概览](#restfulscreen)

---

POST /api/goview/project/create  
Content-Type: application/json

请求参数：

| **来源** | **参数**    | **描述** | **类型** | **约束** | **说明** |
|----------|-------------|----------|----------|----------|----------|
| body     | indexImage  | 缩略图   | string   | 非必填   |          |
| body     | projectName | 大屏名称 | string   | 必填     |          |
| body     | remarks     | 备注     | string   | 非必填   |          |

请求示例：

```json
{
  "indexImage": "indexImage",
  "projectName": "projectName",
  "remarks": "remarks"
}
```

---

Content-Type: application/json

响应参数：

| **参数**     | **描述**                     | **类型** | **说明** |
|--------------|------------------------------|----------|----------|
| createTime   | 创建时间                     | string   |          |
| createUserId | 创建用户ID                   | integer  |          |
| id           | ID                           | integer  |          |
| indexImage   | 缩略图                       | string   |          |
| isDelete     | 是否删除(0 未删除 1 已删除)  | boolean  |          |
| projectName  | 大屏名称                     | string   |          |
| remarks      | 备注                         | string   |          |
| state        | 发布状态(-1 未发布 1 已发布) | integer  |          |

响应示例：

```json
{
  "createTime": "createTime",
  "createUserId": 1,
  "id": 1,
  "indexImage": "indexImage",
  "isDelete": false,
  "projectName": "projectName",
  "remarks": "remarks",
  "state": 1
}
```

### 根据大屏信息ID删除

[返回概览](#restfulscreen)

---

DELETE /api/goview/project/delete  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| query    | ids      | ID       | integer  | 必填     |          |

请求示例：

```
Query:
/api/goview/project/delete?ids=1
```

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|

### 更新大屏信息

[返回概览](#restfulscreen)

---

POST /api/goview/project/edit  
Content-Type: application/json

请求参数：

| **来源** | **参数**    | **描述** | **类型** | **约束** | **说明** |
|----------|-------------|----------|----------|----------|----------|
| body     | id          | ID       | string   | 必填     |          |
| body     | indexImage  | 缩略图   | string   | 非必填   |          |
| body     | projectName | 大屏名称 | string   | 非必填   |          |
| body     | remarks     | 备注     | string   | 非必填   |          |

请求示例：

```json
{
  "id": "id",
  "indexImage": "indexImage",
  "projectName": "projectName",
  "remarks": "remarks"
}
```

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|

### 根据大屏信息ID获取大屏数据

[返回概览](#restfulscreen)

---

GET /api/goview/project/getData

请求参数：

| **来源** | **参数**  | **描述**   | **类型** | **约束** | **说明** |
|----------|-----------|------------|----------|----------|----------|
| query    | projectId | 大屏信息ID | integer  | 必填     |          |

请求示例：

```
Query:
/api/goview/project/getData?projectId=1
```

---

Content-Type: application/json

响应参数：

| **参数**     | **描述** | **类型** | **说明** |
|--------------|----------|----------|----------|
| content      |          | string   |          |
| createTime   |          | string   |          |
| createUserId |          | integer  |          |
| id           |          | string   |          |
| indexImage   |          | string   |          |
| isDelete     |          | integer  |          |
| projectName  |          | string   |          |
| remarks      |          | string   |          |
| state        |          | integer  |          |

响应示例：

```json
{
  "content": "content",
  "createTime": "createTime",
  "createUserId": 1,
  "id": "id",
  "indexImage": "indexImage",
  "isDelete": 1,
  "projectName": "projectName",
  "remarks": "remarks",
  "state": 1
}
```

### 大屏信息列表

[返回概览](#restfulscreen)

---

GET /api/goview/project/list

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| query    | page     | 页码     | integer  | 必填     |          |
| query    | limit    | 每页数量 | integer  | 必填     |          |

请求示例：

```
Query:
/api/goview/project/list?limit=1&page=1
```

---

Content-Type: application/json

响应参数：

| **参数**            | **描述**                     | **类型**     | **说明** |
|---------------------|------------------------------|--------------|----------|
| count               | 总数                         | integer      |          |
| results             | screen_project               | object array |          |
| &emsp; createTime   | 创建时间                     | string       |          |
| &emsp; createUserId | 创建用户ID                   | integer      |          |
| &emsp; id           | ID                           | integer      |          |
| &emsp; indexImage   | 缩略图                       | string       |          |
| &emsp; isDelete     | 是否删除(0 未删除 1 已删除)  | boolean      |          |
| &emsp; projectName  | 大屏名称                     | string       |          |
| &emsp; remarks      | 备注                         | string       |          |
| &emsp; state        | 发布状态(-1 未发布 1 已发布) | integer      |          |

响应示例：

```json
{
  "count": 1,
  "results": [
    {
      "createTime": "createTime",
      "createUserId": 1,
      "id": 1,
      "indexImage": "indexImage",
      "isDelete": false,
      "projectName": "projectName",
      "remarks": "remarks",
      "state": 1
    }
  ]
}
```

### 更新大屏信息发布状态

[返回概览](#restfulscreen)

---

PUT /api/goview/project/publish  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述**                     | **类型** | **约束** | **说明** |
|----------|----------|------------------------------|----------|----------|----------|
| body     | id       | ID                           | integer  | 必填     |          |
| body     | state    | 发布状态(-1 未发布 1 已发布) | integer  | 必填     |          |

请求示例：

```json
{
  "id": 1,
  "state": 1
}
```

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|

### 保存大屏数据

[返回概览](#restfulscreen)

---

POST /api/goview/project/save/data  
Content-Type: multipart/form-data

请求参数：

| **来源** | **参数**  | **描述** | **类型** | **约束** | **说明** |
|----------|-----------|----------|----------|----------|----------|
| formData | projectId |          | integer  | 必填     |          |
| formData | content   |          | string   | 必填     |          |

请求示例：

```
Form Data:
projectId: 1
content: content
```

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|

### 文件上传

[返回概览](#restfulscreen)

---

POST /api/goview/project/upload  
Content-Type: multipart/form-data

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| formData | filename | 文件名   | string   | 非必填   |          |
| formData | object   | 文件     | string   | 非必填   |          |

请求示例：

```
Form Data:
filename: filename
object: object
```

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|
| fileName | 文件名   | string   |          |

响应示例：

```json
{
  "fileName": "fileName"
}
```

### restful/sys

### 获取oss信息

[返回概览](#restfulsys)

---

GET /api/goview/sys/getOssInfo

---

Content-Type: application/json

响应参数：

| **参数**  | **描述** | **类型** | **说明** |
|-----------|----------|----------|----------|
| bucketURL |          | string   |          |

响应示例：

```json
{
  "bucketURL": "bucketURL"
}
```

### 用户登录

[返回概览](#restfulsys)

---

POST /api/goview/sys/login  
Content-Type: application/json

请求参数：

| **来源** | **参数** | **描述** | **类型** | **约束** | **说明** |
|----------|----------|----------|----------|----------|----------|
| body     | password | 密码     | string   | 必填     |          |
| body     | username | 用户名   | string   | 必填     |          |

请求示例：

```json
{
  "password": "password",
  "username": "username"
}
```

---

Content-Type: application/json

响应参数：

| **参数**                    | **描述**          | **类型** | **说明** |
|-----------------------------|-------------------|----------|----------|
| token                       |                   | object   |          |
| &emsp; isLogin              |                   | boolean  |          |
| &emsp; loginDevice          |                   | string   |          |
| &emsp; loginId              |                   | string   |          |
| &emsp; loginType            |                   | string   |          |
| &emsp; sessionTimeout       |                   | integer  |          |
| &emsp; tag                  | map[string]string | object   |          |
| &emsp; tokenActivityTimeout |                   | integer  |          |
| &emsp; tokenName            |                   | string   |          |
| &emsp; tokenSessionTimeout  |                   | integer  |          |
| &emsp; tokenTimeout         |                   | integer  |          |
| &emsp; tokenValue           |                   | string   |          |
| userinfo                    |                   | object   |          |
| &emsp; depId                |                   | integer  |          |
| &emsp; depName              |                   | string   |          |
| &emsp; id                   |                   | string   |          |
| &emsp; nickname             |                   | string   |          |
| &emsp; password             |                   | string   |          |
| &emsp; posId                |                   | string   |          |
| &emsp; posName              |                   | string   |          |
| &emsp; username             |                   | string   |          |

响应示例：

```json
{
  "token": {
    "isLogin": false,
    "loginDevice": "loginDevice",
    "loginId": "loginId",
    "loginType": "loginType",
    "sessionTimeout": 1,
    "tag": null,
    "tokenActivityTimeout": 1,
    "tokenName": "tokenName",
    "tokenSessionTimeout": 1,
    "tokenTimeout": 1,
    "tokenValue": "tokenValue"
  },
  "userinfo": {
    "depId": 1,
    "depName": "depName",
    "id": "id",
    "nickname": "nickname",
    "password": "password",
    "posId": "posId",
    "posName": "posName",
    "username": "username"
  }
}
```

### 用户登出

[返回概览](#restfulsys)

---

GET /api/goview/sys/logout

---

Content-Type: application/json

响应参数：

| **参数** | **描述** | **类型** | **说明** |
|----------|----------|----------|----------|
