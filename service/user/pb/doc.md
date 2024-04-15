# Protocol 文档
<a name="top"></a>

## 目录

- [pb/user.proto](#pb_user-proto)
  - Services
    - [User](#User)
  
  - Messages
    - [AddSysUserReq](#addsysuserreq)
    - [AddSysUserResp](#addsysuserresp)
    - [SysUser](#sysuser)
    - [Token](#token)
    - [TokenTag](#tokentag)
    - [UserInfo](#userinfo)
    - [UserLoginReq](#userloginreq)
    - [UserLoginResp](#userloginresp)
  
- [Scalar Value Types](#scalar-value-types)



<a name="pb_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## pb/user.proto



<a name="user-AddSysUserReq"></a>

### AddSysUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| nickname | [string](#string) |  | 昵称 |






<a name="user-AddSysUserResp"></a>

### AddSysUserResp







<a name="user-SysUser"></a>

### SysUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | ID |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| nickname | [string](#string) |  | 昵称 |
| dep_id | [int64](#int64) |  | dep_id |
| pos_id | [string](#string) |  | pos_id |






<a name="user-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_name | [string](#string) |  |  |
| token_value | [string](#string) |  |  |
| is_login | [bool](#bool) |  |  |
| login_id | [int64](#int64) |  |  |
| login_type | [string](#string) |  |  |
| token_timeout | [int64](#int64) |  |  |
| session_timeout | [int64](#int64) |  |  |
| token_session_timeout | [int64](#int64) |  |  |
| token_activity_timeout | [int64](#int64) |  |  |
| login_device | [string](#string) |  |  |
| tag | [TokenTag](#user-TokenTag) |  |  |






<a name="user-TokenTag"></a>

### TokenTag







<a name="user-UserInfo"></a>

### UserInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| nickname | [string](#string) |  |  |
| dep_id | [int64](#int64) |  |  |
| pos_id | [string](#string) |  |  |
| dep_name | [string](#string) |  |  |
| pos_name | [string](#string) |  |  |






<a name="user-UserLoginReq"></a>

### UserLoginReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |






<a name="user-UserLoginResp"></a>

### UserLoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [Token](#user-Token) |  |  |
| user_info | [UserInfo](#user-UserInfo) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="user-User"></a>

### User


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UserLogin | [UserLoginReq](#userloginreq) | [UserLoginResp](#userloginresp) | 用户登录 |
| InsertSysUser | [AddSysUserReq](#addsysuserreq) | [AddSysUserResp](#addsysuserresp) | 创建用户表 |

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

