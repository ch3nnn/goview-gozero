# Protocol 文档
<a name="top"></a>

## 目录

- [pb/screen.proto](#pb_screen-proto)
  - Services
    - [Screen](#Screen)
  
  - Messages
    - [AddScreenDataReq](#addscreendatareq)
    - [AddScreenDataResp](#addscreendataresp)
    - [AddScreenProjectReq](#addscreenprojectreq)
    - [AddScreenProjectResp](#addscreenprojectresp)
    - [DelScreenProjectReq](#delscreenprojectreq)
    - [DelScreenProjectResp](#delscreenprojectresp)
    - [ScreenData](#screendata)
    - [ScreenProject](#screenproject)
    - [ScreenProjectFilter](#screenprojectfilter)
    - [SelectScreenDataByIdReq](#selectscreendatabyidreq)
    - [SelectScreenDataByIdResp](#selectscreendatabyidresp)
    - [SelectScreenProjectByIdReq](#selectscreenprojectbyidreq)
    - [SelectScreenProjectByIdResp](#selectscreenprojectbyidresp)
    - [SelectScreenProjectListReq](#selectscreenprojectlistreq)
    - [SelectScreenProjectListResp](#selectscreenprojectlistresp)
    - [UpdateScreenProjectReq](#updatescreenprojectreq)
    - [UpdateScreenProjectResp](#updatescreenprojectresp)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pb_screen-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pb/screen.proto



<a name="screen-AddScreenDataReq"></a>

### AddScreenDataReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [int64](#int64) |  | project ID |
| user_id | [int64](#int64) |  | 用户 ID |
| content | [string](#string) |  | 内容数据 |






<a name="screen-AddScreenDataResp"></a>

### AddScreenDataResp







<a name="screen-AddScreenProjectReq"></a>

### AddScreenProjectReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 大屏名称 |
| state | [int64](#int64) |  | 发布状态(-1 未发布 1 已发布) |
| user_id | [int64](#int64) |  | 创建用户ID |
| index_img | [string](#string) |  | 缩略图 |
| remark | [string](#string) |  | 备注 |
| is_del | [bool](#bool) |  | 是否删除(0 未删除 1 已删除) |






<a name="screen-AddScreenProjectResp"></a>

### AddScreenProjectResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project | [ScreenProject](#screen-ScreenProject) |  | screen_project |






<a name="screen-DelScreenProjectReq"></a>

### DelScreenProjectReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |






<a name="screen-DelScreenProjectResp"></a>

### DelScreenProjectResp







<a name="screen-ScreenData"></a>

### ScreenData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| project_id | [int64](#int64) |  | project ID |
| user_id | [int64](#int64) |  | 用户 ID |
| content | [string](#string) |  | 内容数据 |
| create_at | [int64](#int64) |  | 创建时间 |






<a name="screen-ScreenProject"></a>

### ScreenProject



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| name | [string](#string) |  | 大屏名称 |
| state | [int64](#int64) |  | 发布状态(-1 未发布 1 已发布) |
| user_id | [int64](#int64) |  | 创建用户ID |
| index_img | [string](#string) |  | 缩略图 |
| remark | [string](#string) |  | 备注 |
| is_del | [bool](#bool) |  | 是否删除(0 未删除 1 已删除) |
| create_at | [int64](#int64) |  | 创建时间 |






<a name="screen-ScreenProjectFilter"></a>

### ScreenProjectFilter



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | optional | ID |
| name | [string](#string) | optional | 大屏名称 |
| state | [int64](#int64) | optional | 发布状态(-1 未发布 1 已发布) |
| user_id | [int64](#int64) | optional | 创建用户ID |
| index_img | [string](#string) | optional | 缩略图 |
| remark | [string](#string) | optional | 备注 |
| is_del | [bool](#bool) | optional | 是否删除(0 未删除 1 已删除) |
| create_at | [int64](#int64) | optional | 创建时间 |






<a name="screen-SelectScreenDataByIdReq"></a>

### SelectScreenDataByIdReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| project_id | [int64](#int64) |  | project ID |






<a name="screen-SelectScreenDataByIdResp"></a>

### SelectScreenDataByIdResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| screen_data | [ScreenData](#screen-ScreenData) |  | screen_data |






<a name="screen-SelectScreenProjectByIdReq"></a>

### SelectScreenProjectByIdReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |






<a name="screen-SelectScreenProjectByIdResp"></a>

### SelectScreenProjectByIdResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| screen_project | [ScreenProject](#screen-ScreenProject) |  | screen_project |






<a name="screen-SelectScreenProjectListReq"></a>

### SelectScreenProjectListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| page_size | [int64](#int64) |  | 每页数量 |
| filter | [ScreenProjectFilter](#screen-ScreenProjectFilter) |  | ScreenProjectFilter |






<a name="screen-SelectScreenProjectListResp"></a>

### SelectScreenProjectListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| count | [int64](#int64) |  | 总数 |
| page_count | [int64](#int64) |  | 页码总数 |
| results | [ScreenProject](#screen-ScreenProject) | repeated | screen_project |






<a name="screen-UpdateScreenProjectReq"></a>

### UpdateScreenProjectReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| name | [string](#string) | optional | 大屏名称 |
| state | [int64](#int64) | optional | 发布状态(-1 未发布 1 已发布) |
| user_id | [int64](#int64) | optional | 创建用户ID |
| index_img | [string](#string) | optional | 缩略图 |
| remark | [string](#string) | optional | 备注 |
| is_del | [bool](#bool) | optional | 是否删除(0 未删除 1 已删除) |






<a name="screen-UpdateScreenProjectResp"></a>

### UpdateScreenProjectResp






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="screen-Screen"></a>

### Screen


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| InsertScreenProject | [AddScreenProjectReq](#addscreenprojectreq) | [AddScreenProjectResp](#addscreenprojectresp) | 创建大屏信息 |
| UpdateScreenProject | [UpdateScreenProjectReq](#updatescreenprojectreq) | [UpdateScreenProjectResp](#updatescreenprojectresp) | 更新大屏信息 |
| DeleteScreenProject | [DelScreenProjectReq](#delscreenprojectreq) | [DelScreenProjectResp](#delscreenprojectresp) | 根据大屏信息ID删除 |
| SelectScreenProjectById | [SelectScreenProjectByIdReq](#selectscreenprojectbyidreq) | [SelectScreenProjectByIdResp](#selectscreenprojectbyidresp) | 根据大屏信息ID获取详情 |
| SelectScreenProjectList | [SelectScreenProjectListReq](#selectscreenprojectlistreq) | [SelectScreenProjectListResp](#selectscreenprojectlistresp) | 大屏信息列表 |
| InsertScreenData | [AddScreenDataReq](#addscreendatareq) | [AddScreenDataResp](#addscreendataresp) | 创建大屏数据 |
| SelectScreenDataById | [SelectScreenDataByIdReq](#selectscreendatabyidreq) | [SelectScreenDataByIdResp](#selectscreendatabyidresp) | 根据大屏数据ID获取详情 |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

