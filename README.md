# Go API client for openapi

CODING 提供了丰富的 API
接口，注册应用即可使用，无需审核，支持两种认证方式：[OAuth 2.0 协议](#oauth-认证)、[个人访问令牌](#个人令牌认证)、[项目令牌](#项目令牌认证)。

# [OAuth 认证](#oauth)

#### [生态令牌权限](#ecology-scope)

令牌权限说明
> Scope的权限分为只读（ro）和读写(rw) 俩种，用户可通过权限点+权限方式获取完整的权限 Scope 信息。例如，用户信息授权为只读时，用户的令牌权限 Scope 为 user:profile:ro

| 名称            | 描述信息                 | Scope 信息                  | 权限范围  | 示例                           |
|---------------|----------------------|---------------------------|-------|------------------------------|
| 用户信息          | 管理用户的基础信息。           | user:profile              | ro    | user:profile:ro              |
| 用户邮箱          | 管理用户的电子邮件地址。         | user:email                | ro    | user:email:ro                |
| 用户通知          | 管理用户的站内通知。           | user:notification         | ro、rw | user:notification:rw         |
| 用户公钥          | 管理用户配置的个人公钥和部署公钥信息。  | user:public-key           | ro、rw | user:public-key:rw           |
| 团队信息          | 管理团队基本信息。            | team:profile              | ro    | team:profile:ro              |
| 团队成员          | 管理团队成员信息以及团队成员相关操作。  | team:member               | ro、rw | team:member:rw               |
| 项目信息          | 管理项目基本信息。            | project:profile           | ro、rw | project:profile:rw           |
| 项目成员          | 管理项目成员。              | project:member            | ro、rw | project:member:rw            |
| 项目令牌          | 管理项目令牌。              | project:token             | ro、rw | project:token:rw             |
| 项目公告          | 管理项目公告。              | project:notice            | ro、rw | project:notice:rw            |
| 项目标签          | 管理项目标签。              | project:label             | ro、rw | project:label:rw             |
| 项目集信息         | 管理项目集基本信息。           | program:profile           | ro、rw | program:profile:rw           |
| 项目集项目         | 管理项目集下的项目列表。         | program:project           | ro、rw | program:project:rw           |
| 项目集成员         | 管理项目集下的成员列表。         | program:member            | ro、rw | program:member:rw            |
| 关联资源          | 管理团队和项目资源关联关系。       | related-resource:resource | ro、rw | related-resource:resource:rw |
| 凭据信息          | 管理团队凭据。              | credential:profile        | ro、rw | credential:profile:rw        |
| Service Hooks | 管理和配置 Service Hooks。 | service-hook:profile      | ro、rw | service-hook:profile:rw      |
| 权限组           | 管理权限组。               | ram:policy                | ro、rw | ram:policy:rw                |
| 授权            | 配置权限授权。              | ram:grant                 | ro、rw | ram:grant:rw                 |
| 用户组           | 管理权限用户组。             | ram:user-group            | ro    | ram:user-group:ro            |
| 研发度量数据集       | 研发度量数据集              | performance:dataset       | ro    | performance:dataset:ro       |
| 项目协同          | 配置和使用项目协同功能。         | collaboration:issue       | ro、rw | collaboration:issue:rw       |
| 知识管理          | 管理知识空间和撰写知识文档。       | document:knowledge        | ro、rw | document:knowledge:rw        |
| 文件网盘          | 管理上传、分享和下载文件等。       | document:file             | ro、rw | document:file:rw             |
| API 文档        | 发布、授权发布 API 文档。      | document:api-doc          | ro、rw | document:api-doc:rw          |
| 代码仓库          | 管理仓库                 | vcs:repository            | ro、rw | vcs:repository:rw            |
| 合并请求          | 管理代码仓库的合并请求。         | vcs:merge                 | ro、rw | vcs:merge:rw                 |
| 部署公钥          | 管理代码仓库的部署公钥。         | vcs:ssh-key               | ro、rw | vcs:ssh-key:rw               |
| 版本管理          | 管理代码仓库的版本信息。         | vcs:release               | ro、rw | vcs:release:rw               |
| 外部仓库          | 管理关联的外部仓库信息。         | depot:external-repository | ro、rw | depot:external-repository:rw |
| 测试管理          | 管理测试计、测试用例和测试报告等。    | testing:profile           | ro、rw | testing:profile:rw           |
| 持续部署数据统计      | 持续部署发布数据统计。          | cd:statistics             | ro    | cd:statistics:ro             |
| 持续部署主机组       | 管理持续部署主机组。           | cd:host-server            | ro、rw | cd:host-server:rw            |
| 持续部署云账号       | 管理持续部署云账号。           | cd:cloud-account          | ro、rw | cd:cloud-account:rw          |
| 持续部署应用        | 管理和配置持续部署应用。         | cd:application            | ro、rw | cd:application:rw            |
| 持续部署流程        | 管理和配置持续部署流程。         | cd:pipeline               | ro、rw | cd:pipeline:rw               |
| 制品库仓库         | 管理制品库仓库。             | artifact:repository       | ro、rw | artifact:repository:rw       |
| 制品库版本         | 管理制品版本信息。            | artifact:version          | ro、rw | artifact:version:rw          |
| 制品库配置         | 管理制品库配置。             | artifact:properties       | ro、rw | artifact:properties:rw       |
| 制品库包          | 管理制品库包。              | artifact:package          | ro    | artifact:package:ro          |
| 资产列表          | 管理资产列表               | assets:list               | ro、rw | assets:list:rw               |
| 资产属性          | 管理资产属性               | assets:attribute          | ro、rw | assets:attribute:rw          |

#### [创建 CODING 应用](#创建-CODING-应用)

##### 1、新建应用

点击【团队设置】->【生态能力】->【发布应用】->【新建应用】，填写信息。「回调地址」处填写回调服务地址，这里以codesign为例。

![](https://help-assets.codehub.cn/enterprise/202309201515397.png)

![](https://help-assets.codehub.cn/enterprise/202309201519877.png)

##### 2、获取 Client ID 和 Client Secret

点击创建好的应用，点击客户端秘钥，即可看到客户端ID与客户端秘钥

![](https://help-assets.codehub.cn/enterprise/202309201520614.png)

##### 注意事项：
- 更换令牌时，先新建令牌，再删除旧令牌
- 最多同时可创建 5 个令牌
- 令牌只有创建时可查看，创建后，任何人无法查看，请妥善保管。

#### 3、修改令牌权限

根据需求修改令牌权限，这里设置了用户信息只读权限与项目信息读写权限作为示例

![](https://help-assets.codehub.cn/enterprise/202309201520611.png)

#### [用户授权](#oauth-scopes)

浏览器访问以下链接，进入到授权登录页面：

```shell
GET https://{your-team}.coding.net/oauth_authorize.html?client_id={client_id}&redirect_uri={redirect_uri}&response_type=code&state=${state}&scope={scope}

```

参数说明：

- your-team：团队域名前缀；

- client_id：应用详情页的 Client ID；

- redirect_uri：应用详情页填写的回调地址；

- response_type：返回类型，固定为 code；

- state: 业务侧可以传任何值，这个值会原封不动传递回去，用来给业务识别一些场景用的。

- scope：授权范围，多个权限间以逗号分隔，如：user:profile:ro,project:profile:rw

![](https://help-assets.codehub.cn/enterprise/202309201520497.png)

点击授权后，浏览器将带着授权码（code）参数和业务状态码（state）跳转到回调地址，如：

```shell
https://codesign.qq.com/?code={code}&state={state}&team={teamGk}&scope=user%3Aprofile%3Aro,project%3Aprofile%3Arw
```

#### [获取 access_token](#oauth-access-token)

获取授权码（code）后，开发者的后端程序向 CODING 发送请求，获取 access_token。

请求链接：

```shell
POST https://{your-team}.coding.net/api/oauth/access_token
Content-Type: application/x-www-form-urlencoded

client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&code={code}
```

参数说明：

- your-team：团队域名前缀；

- client_id：应用详情页的 Client ID；

- client_secret：应用详情页的 Client Secret；

- code：上一步获取的授权码，须注意每个 code 仅能使用一次，且有效期仅5分钟；

- grant_type：授权类型，固定为 authorization_code；

返回值：

```json
{
  \"access_token\": \"RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975\",
  \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",
  \"scope\": \"user:profile:ro project:profile:rw\",
  \"team\": \"jackwhu-test\",
  \"token_type\": \"Bearer\",
  \"expires_in\": \"7200\"
}
```

返回值：
- access_token： access_token值，可用于调用OpenAPI接口，建议按expires_in保存access_token
- refresh_token： 刷新时使用的token，有效期90天。access_token过期后可用于刷新access_token
- scope：令牌权限范围
- team： 团队gk
- token_type：token类型
- expires_in：过期时间，单位秒


#### [刷新 access_token](#oauth-access-token)

通过上面获取到的 refresh_token，开发者的后端程序可以向 CODING 发送请求，刷新 access_token。  
注意：调用刷新接口后，即使 access_token 未过期，原 access_token 也将失效

请求链接：

```shell
POST https://{your-team}.coding.net/api/oauth/access_token
Content-Type: application/x-www-form-urlencoded

client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&refresh_token={refresh_token}
```

参数说明：

- your-team：团队域名前缀；

- client_id：应用详情页的 Client ID；

- client_secret：应用详情页的 Client Secret；

- refresh_token：获取 access_token接口返回的refresh_token字段

- grant_type：授权类型，固定为 refresh_token；

返回值：

```json
{
  \"access_token\": \"q4qIkUGhJ2qfcdSV3bWx0YfQj-WjLqXG7LSdP9Oo3sOAjmuY-Bb_QJ6ORt-By-bc7WFFT7PyH7RXEvPLBM5lFU9PBOofgzXN9Lh5zp3FWRdyV_4XGno4U_S7zMkixWnv\",
  \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",
  \"scope\": \"user:profile:ro project:profile:rw\",
  \"team\": \"jackwhu-test\",
  \"token_type\": \"Bearer\",
  \"expires_in\": \"7200\"
}
```

#### [获取当前用户信息](#oauth-get-user)

请求链接：

```shell
POST https://{your-team}.coding.net/open-api
Authorization: Bearer RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975
Content-Type: application/json

{
    \"Action\": \"DescribeCodingCurrentUser\"
}
```

Headers说明：

- Authorization：\"Bearer {access_token}\"

参数说明：

- Action：固定为DescribeCodingCurrentUser，令牌须打开user:profile:ro权限

返回值：

```json
{
  \"Response\": {
    \"User\": {
      \"Id\": 183478,
      \"Status\": 1,
      \"Email\": \"test@coding.com\",
      \"GlobalKey\": \"anywhere\",
      \"Avatar\": \"https://e.coding.net/static/fruit_avatar/Fruit-20.png\",
      \"Gravatar\": \"\",
      \"Name\": \"qqq\",
      \"NamePinYin\": \"qqq\",
      \"Phone\": \"18800000000\",
      \"PhoneValidation\": 1,
      \"EmailValidation\": 1,
      \"PhoneRegionCode\": \"+86\",
      \"Region\": \"cn\",
      \"TeamId\": 1,
      \"WeComId\": \"woG7VgCgAAov2F-sAQkD_YPsLNJiP1Vg\"
    },
    \"RequestId\": \"133e152f-8852-4d99-b704-c7ff245a1640\"
  }
}

```

# [个人令牌认证](#personal_token)

#### [获取个人令牌](#personal-token-create)

- 点击左下角的【个人账户设置】>【访问令牌】>【新建访问令牌】，勾选相关权限后会生成「个人访问令牌」。若刷新页面令牌会消失，需输入账号密码后重新生成。
- 令牌权限与[OAuth令牌权限](#生态令牌权限)一样

![](https://help-assets.codehub.cn/enterprise/202309201521630.png)

#### [获取当前用户信息](#personal-token-get-user)

请求链接：

```shell
POST https://{your-team}.coding.net/open-api
Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8
Content-Type: application/json

{
    \"Action\": \"DescribeCodingCurrentUser\"
}
```

header：

```text
Authorization: token {访问令牌}

```

请求示例：

```shell
curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\
--header 'Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8' \\
--header 'Content-Type: application/json' \\
--data-raw '{
    \"Action\": \"DescribeCodingCurrentUser\"
}'
```

# [项目令牌认证](#personal_token)

#### [项目令牌Scope](#project-token-scope)

| 名称      | 描述                        | Scope 信息                  | 权限范围 |
|---------|---------------------------|---------------------------|------|
| 项目协同    | 读取与操作项目协同                 | project_issue_rw          | 读写   |
| 文件      | 新建、查询、删除、编辑文件             | file_rw                   | 读写   |
| WIKI    | 新建、查询、删除、编辑Wiki           | wiki_rw                   | 读写   |
| 项目公告    | 新建、查询、删除、编辑项目公告           | project_tweet_rw          | 读写   |
| API 文档  | 发布 API 文档                 | api_doc_release           | 读写   |
| 关联资源    | 新建、查询、删除关联资源              | resource_reference_rw     | 读写   |
| 项目成员    | 读取与操作项目成员                 | project_member_rw         | 读写   |
| 项目权限    | 读取与操作项目权限                 | project_permission_rw     | 读写   |
| 项目标签    | 新建、查询、删除、编辑项目标签           | project_label_rw          | 读写   |
| 测试协同    | 执行 API 自动化测试              | ifbook_run_task           | 读写   |
| 测试协同    | API 文档导入与导出               | ifbook_import_export      | 读写   |
| 读取代码仓库  | 读取代码仓库                    | depot_read                | 只读   |
| 推送至代码仓库 | 推送至代码仓库                   | depot_write               | 读写   |
| MR      | 新建、查询、删除、编辑合并请求           | merge_request_rw          | 读写   |
| 版本发布    | 新建、查询、删除、编辑版本发布           | release_rw                | 读写   |
| 制品库     | 拉取制品库                     | artifact_r                | 只读   |
| 制品库     | 拉取、推送制品库                  | artifact_rw               | 读写   |
| 制品属性    | 新建、查询、删除、编辑制品属性           | artifact_version_props_rw | 读写   |
| 构建节点    | 允许构建节点接入                  | ci_agent_register         | 读写   |
| API触发   | 构建计划管理/触发构建               | ci_manager                | 读写   |
| 构建计划    | 构建计划管理/触发构建（仅用于 Open API） | open_ci_manager           | 读写   |

#### [获取项目令牌](#project-token-create)

1. 点击【项目设置】>【开发者选项】>【项目令牌】>【新建项目令牌】，勾选相关权限后会生成「项目令牌」。点击查看密码即可获取密码信息

![](https://help-assets.codehub.cn/enterprise/202309201843081.png)

2. Basic认证：将用户名与密码通过”用户名:密码“方式拼接后，使用Base64进行编码，获取凭证

#### [获取当前项目信息](#project-token-get-user)

请求链接：

```shell
POST https://{your-team}.coding.net/open-api
Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE
Content-Type: application/json

{
    \"Action\": \"DescribeOneProject\",
    \"ProjectId\": 11595365
}
```

header：

```text
Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE
```

请求示例：

```shell
curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\
--header 'Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE' \\
--header 'Content-Type: application/json' \\
--data-raw '{
    \"Action\": \"DescribeOneProject\",
    \"ProjectId\": 11595365
}'
```



## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Generator version: 7.9.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://e.coding.net/open-api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ArchiveTestRun**](docs/DefaultApi.md#archivetestrun) | **Post** /?action&#x3D;ArchiveTestRun | 测试计划归档
*DefaultApi* | [**AttachResourceScopeToPolicy**](docs/DefaultApi.md#attachresourcescopetopolicy) | **Post** /?action&#x3D;AttachResourceScopeToPolicy | 权限组添加可用的资源，原有其它资源不受影响，若已存在的资源不再进行追加
*DefaultApi* | [**AttachToResource**](docs/DefaultApi.md#attachtoresource) | **Post** /?action&#x3D;AttachToResource | 授权追加，原有其它授权不受影响，若授权已存在不再进行追加
*DefaultApi* | [**BindCdApplicationToProject**](docs/DefaultApi.md#bindcdapplicationtoproject) | **Post** /?action&#x3D;BindCdApplicationToProject | 绑定 CD 应用到项目
*DefaultApi* | [**BoundExternalDepot**](docs/DefaultApi.md#boundexternaldepot) | **Post** /?action&#x3D;BoundExternalDepot | 外部仓库关联
*DefaultApi* | [**CancelCdPipeline**](docs/DefaultApi.md#cancelcdpipeline) | **Post** /?action&#x3D;CancelCdPipeline | CD 部署流程取消执行
*DefaultApi* | [**ClearCodingCIJobCache**](docs/DefaultApi.md#clearcodingcijobcache) | **Post** /?action&#x3D;ClearCodingCIJobCache | 构建计划缓存清理
*DefaultApi* | [**CreateArtifactCredit**](docs/DefaultApi.md#createartifactcredit) | **Post** /?action&#x3D;CreateArtifactCredit | 制品授信清单创建
*DefaultApi* | [**CreateArtifactProperties**](docs/DefaultApi.md#createartifactproperties) | **Post** /?action&#x3D;CreateArtifactProperties | 制品属性新增（指定版本）
*DefaultApi* | [**CreateArtifactRepository**](docs/DefaultApi.md#createartifactrepository) | **Post** /?action&#x3D;CreateArtifactRepository | 制品仓库创建
*DefaultApi* | [**CreateAttachmentPrepareSignUrl**](docs/DefaultApi.md#createattachmentpreparesignurl) | **Post** /?action&#x3D;CreateAttachmentPrepareSignUrl | 附件预上传信息生成
*DefaultApi* | [**CreateBinaryFiles**](docs/DefaultApi.md#createbinaryfiles) | **Post** /CreateBinaryFiles | Git文件-Git仓库创建二进制文件
*DefaultApi* | [**CreateBranchProtection**](docs/DefaultApi.md#createbranchprotection) | **Post** /CreateBranchProtection | 仓库设置-新增代码保护规则
*DefaultApi* | [**CreateBranchProtectionMember**](docs/DefaultApi.md#createbranchprotectionmember) | **Post** /CreateBranchProtectionMember | 仓库设置-新增保护分支规则管理员
*DefaultApi* | [**CreateCaseResult**](docs/DefaultApi.md#createcaseresult) | **Post** /?action&#x3D;CreateCaseResult | 测试用例添加测试结果
*DefaultApi* | [**CreateCdCloudAccount**](docs/DefaultApi.md#createcdcloudaccount) | **Post** /?action&#x3D;CreateCdCloudAccount | CD 云账号添加
*DefaultApi* | [**CreateCdHostServerGroup**](docs/DefaultApi.md#createcdhostservergroup) | **Post** /?action&#x3D;CreateCdHostServerGroup | CD 主机组添加
*DefaultApi* | [**CreateCdPipeline**](docs/DefaultApi.md#createcdpipeline) | **Post** /?action&#x3D;CreateCdPipeline | CD 部署流程创建
*DefaultApi* | [**CreateCdTask**](docs/DefaultApi.md#createcdtask) | **Post** /?action&#x3D;CreateCdTask | CD 任务执行
*DefaultApi* | [**CreateCodingCIJob**](docs/DefaultApi.md#createcodingcijob) | **Post** /?action&#x3D;CreateCodingCIJob | 构建计划创建
*DefaultApi* | [**CreateCodingCIJobByTeamTemplate**](docs/DefaultApi.md#createcodingcijobbyteamtemplate) | **Post** /?action&#x3D;CreateCodingCIJobByTeamTemplate | 构建计划-根据团队模版创建
*DefaultApi* | [**CreateCodingProject**](docs/DefaultApi.md#createcodingproject) | **Post** /?action&#x3D;CreateCodingProject | 项目创建
*DefaultApi* | [**CreateDepartment**](docs/DefaultApi.md#createdepartment) | **Post** /?action&#x3D;CreateDepartment | 部门创建
*DefaultApi* | [**CreateDepotByTemplate**](docs/DefaultApi.md#createdepotbytemplate) | **Post** /CreateDepotByTemplate | 仓库信息-根据模板创建仓库
*DefaultApi* | [**CreateDepotFilePushRule**](docs/DefaultApi.md#createdepotfilepushrule) | **Post** /CreateDepotFilePushRule | 仓库设置-新增git仓库文件推送规则
*DefaultApi* | [**CreateDepotFilePushRulePrivilege**](docs/DefaultApi.md#createdepotfilepushruleprivilege) | **Post** /CreateDepotFilePushRulePrivilege | 仓库设置-新增git仓库文件推送规则特权者
*DefaultApi* | [**CreateFile**](docs/DefaultApi.md#createfile) | **Post** /?action&#x3D;CreateFile | 文件创建
*DefaultApi* | [**CreateFolder**](docs/DefaultApi.md#createfolder) | **Post** /?action&#x3D;CreateFolder | 文件夹创建
*DefaultApi* | [**CreateGitBranch**](docs/DefaultApi.md#creategitbranch) | **Post** /CreateGitBranch | 仓库分支-用于代码仓库新建分支
*DefaultApi* | [**CreateGitCommit**](docs/DefaultApi.md#creategitcommit) | **Post** /CreateGitCommit | Git提交-创建一次提交
*DefaultApi* | [**CreateGitCommitComment**](docs/DefaultApi.md#creategitcommitcomment) | **Post** /CreateGitCommitComment | Git提交-为某次提交创建一条评论
*DefaultApi* | [**CreateGitCommitNote**](docs/DefaultApi.md#creategitcommitnote) | **Post** /CreateGitCommitNote | Git提交-创建提交注释。注意：对于 NotesRef 参数建议默认为空，因为git会使用默认的ref ：refs/notes/commits，如果填了这个参数，会用这个参数指定的ref来保存您的git note，有可能会覆盖您原有的ref。
*DefaultApi* | [**CreateGitDeployKey**](docs/DefaultApi.md#creategitdeploykey) | **Post** /CreateGitDeployKey | 仓库设置-新建部署公钥
*DefaultApi* | [**CreateGitDepot**](docs/DefaultApi.md#creategitdepot) | **Post** /?action&#x3D;CreateGitDepot | 仓库信息-创建代码仓库
*DefaultApi* | [**CreateGitFiles**](docs/DefaultApi.md#creategitfiles) | **Post** /CreateGitFiles | Git文件-创建仓库文件
*DefaultApi* | [**CreateGitMergeReq**](docs/DefaultApi.md#creategitmergereq) | **Post** /CreateGitMergeReq | 合并请求-创建git合并请求
*DefaultApi* | [**CreateGitMergeRequest**](docs/DefaultApi.md#creategitmergerequest) | **Post** /CreateGitMergeRequest | 合并请求-创建Git合并请求mr
*DefaultApi* | [**CreateGitProtectedTagRule**](docs/DefaultApi.md#creategitprotectedtagrule) | **Post** /CreateGitProtectedTagRule | 仓库设置-创建标签保护规则
*DefaultApi* | [**CreateGitProtectedTagRules**](docs/DefaultApi.md#creategitprotectedtagrules) | **Post** /CreateGitProtectedTagRules | 仓库设置-批量创建标签保护规则
*DefaultApi* | [**CreateGitRelease**](docs/DefaultApi.md#creategitrelease) | **Post** /CreateGitRelease | 版本信息-新建git版本
*DefaultApi* | [**CreateGitTag**](docs/DefaultApi.md#creategittag) | **Post** /CreateGitTag | 标签信息-创建标签
*DefaultApi* | [**CreateIssue**](docs/DefaultApi.md#createissue) | **Post** /CreateIssue | 事项创建
*DefaultApi* | [**CreateIssueBlock**](docs/DefaultApi.md#createissueblock) | **Post** /?action&#x3D;CreateIssueBlock | 前置事项添加
*DefaultApi* | [**CreateIssueComment**](docs/DefaultApi.md#createissuecomment) | **Post** /?action&#x3D;CreateIssueComment | 事项评论创建
*DefaultApi* | [**CreateIssueModule**](docs/DefaultApi.md#createissuemodule) | **Post** /CreateIssueModule | 事项模块创建
*DefaultApi* | [**CreateIssueWorkHours**](docs/DefaultApi.md#createissueworkhours) | **Post** /?action&#x3D;CreateIssueWorkHours | 工时登记
*DefaultApi* | [**CreateIteration**](docs/DefaultApi.md#createiteration) | **Post** /?action&#x3D;CreateIteration | 迭代创建
*DefaultApi* | [**CreateMemberSshKey**](docs/DefaultApi.md#createmembersshkey) | **Post** /CreateMemberSshKey | 仓库设置-导入团队成员SSH公钥
*DefaultApi* | [**CreateMergeRequestNote**](docs/DefaultApi.md#createmergerequestnote) | **Post** /CreateMergeRequestNote | 合并请求-创建合并请求行评论和改动文件diff行评论
*DefaultApi* | [**CreateMergeRequestReviewer**](docs/DefaultApi.md#createmergerequestreviewer) | **Post** /CreateMergeRequestReviewer | 合并请求-新增合并请求评审者
*DefaultApi* | [**CreatePolicy**](docs/DefaultApi.md#createpolicy) | **Post** /?action&#x3D;CreatePolicy | 权限组创建
*DefaultApi* | [**CreateProgram**](docs/DefaultApi.md#createprogram) | **Post** /?action&#x3D;CreateProgram | 项目集创建
*DefaultApi* | [**CreateProgramMemberPolicy**](docs/DefaultApi.md#createprogrammemberpolicy) | **Post** /?action&#x3D;CreateProgramMemberPolicy | 项目集成员权限组添加
*DefaultApi* | [**CreateProgramProjects**](docs/DefaultApi.md#createprogramprojects) | **Post** /?action&#x3D;CreateProgramProjects | 项目集中添加项目
*DefaultApi* | [**CreateProjectAnnouncement**](docs/DefaultApi.md#createprojectannouncement) | **Post** /?action&#x3D;CreateProjectAnnouncement | 项目公告创建
*DefaultApi* | [**CreateProjectLabel**](docs/DefaultApi.md#createprojectlabel) | **Post** /?action&#x3D;CreateProjectLabel | 项目标签创建
*DefaultApi* | [**CreateProjectMemberPrincipal**](docs/DefaultApi.md#createprojectmemberprincipal) | **Post** /?action&#x3D;CreateProjectMemberPrincipal | 项目成员主体新增(包含用户组、部门、成员)
*DefaultApi* | [**CreateProjectWithTemplate**](docs/DefaultApi.md#createprojectwithtemplate) | **Post** /?action&#x3D;CreateProjectWithTemplate | 模版项目创建
*DefaultApi* | [**CreateReadOnlyRef**](docs/DefaultApi.md#createreadonlyref) | **Post** /?action&#x3D;CreateReadOnlyRef | 仓库分支-创建只读分支
*DefaultApi* | [**CreateRelease**](docs/DefaultApi.md#createrelease) | **Post** /CreateRelease | 版本创建
*DefaultApi* | [**CreateReport**](docs/DefaultApi.md#createreport) | **Post** /?action&#x3D;CreateReport | 测试报告创建
*DefaultApi* | [**CreateRequirementDefectRelation**](docs/DefaultApi.md#createrequirementdefectrelation) | **Post** /?action&#x3D;CreateRequirementDefectRelation | 需求关联缺陷
*DefaultApi* | [**CreateSshKey**](docs/DefaultApi.md#createsshkey) | **Post** /?action&#x3D;CreateSshKey | 仓库设置-导入用户SSH公钥
*DefaultApi* | [**CreateTestCase**](docs/DefaultApi.md#createtestcase) | **Post** /?action&#x3D;CreateTestCase | 测试用例创建
*DefaultApi* | [**CreateTestCaseSection**](docs/DefaultApi.md#createtestcasesection) | **Post** /?action&#x3D;CreateTestCaseSection | 测试用例分组创建
*DefaultApi* | [**CreateTestDefect**](docs/DefaultApi.md#createtestdefect) | **Post** /?action&#x3D;CreateTestDefect | 测试任务关联缺陷
*DefaultApi* | [**CreateTestResult**](docs/DefaultApi.md#createtestresult) | **Post** /?action&#x3D;CreateTestResult | 测试任务添加测试结果
*DefaultApi* | [**CreateTestResults**](docs/DefaultApi.md#createtestresults) | **Post** /?action&#x3D;CreateTestResults | 测试任务状态批量更新
*DefaultApi* | [**CreateTestRun**](docs/DefaultApi.md#createtestrun) | **Post** /?action&#x3D;CreateTestRun | 测试计划创建
*DefaultApi* | [**CreateTestStepResult**](docs/DefaultApi.md#createteststepresult) | **Post** /?action&#x3D;CreateTestStepResult | 测试任务添加某步骤的测试结果
*DefaultApi* | [**CreateUserGroup**](docs/DefaultApi.md#createusergroup) | **Post** /?action&#x3D;CreateUserGroup | 用户组创建
*DefaultApi* | [**CreateUserGroupUsers**](docs/DefaultApi.md#createusergroupusers) | **Post** /?action&#x3D;CreateUserGroupUsers | 用户组添加用户
*DefaultApi* | [**DeleteAllUsersOnGroup**](docs/DefaultApi.md#deleteallusersongroup) | **Post** /?action&#x3D;DeleteAllUsersOnGroup | 用户组清理用户
*DefaultApi* | [**DeleteArtifactProperties**](docs/DefaultApi.md#deleteartifactproperties) | **Post** /?action&#x3D;DeleteArtifactProperties | 制品属性删除
*DefaultApi* | [**DeleteBranchProtection**](docs/DefaultApi.md#deletebranchprotection) | **Post** /DeleteBranchProtection | 仓库设置-删除保护分支规则
*DefaultApi* | [**DeleteBranchProtectionMember**](docs/DefaultApi.md#deletebranchprotectionmember) | **Post** /DeleteBranchProtectionMember | 仓库设置-删除保护分支规则管理员
*DefaultApi* | [**DeleteCdCloudAccount**](docs/DefaultApi.md#deletecdcloudaccount) | **Post** /?action&#x3D;DeleteCdCloudAccount | CD 云账号删除
*DefaultApi* | [**DeleteCdHostServerGroup**](docs/DefaultApi.md#deletecdhostservergroup) | **Post** /?action&#x3D;DeleteCdHostServerGroup | CD 主机组删除
*DefaultApi* | [**DeleteCdPipeline**](docs/DefaultApi.md#deletecdpipeline) | **Post** /?action&#x3D;DeleteCdPipeline | CD 部署流程删除
*DefaultApi* | [**DeleteCodingCIBuild**](docs/DefaultApi.md#deletecodingcibuild) | **Post** /?action&#x3D;DeleteCodingCIBuild | 构建删除
*DefaultApi* | [**DeleteCodingCIJob**](docs/DefaultApi.md#deletecodingcijob) | **Post** /?action&#x3D;DeleteCodingCIJob | 构建计划删除
*DefaultApi* | [**DeleteDepartment**](docs/DefaultApi.md#deletedepartment) | **Post** /?action&#x3D;DeleteDepartment | 部门删除
*DefaultApi* | [**DeleteDepotFilePushRule**](docs/DefaultApi.md#deletedepotfilepushrule) | **Post** /DeleteDepotFilePushRule | 仓库设置-删除git仓库文件推送规则
*DefaultApi* | [**DeleteDepotFilePushRuleDenyPrivilege**](docs/DefaultApi.md#deletedepotfilepushruledenyprivilege) | **Post** /DeleteDepotFilePushRuleDenyPrivilege | 仓库设置-删除git仓库特权者文件推送权限
*DefaultApi* | [**DeleteGitBranch**](docs/DefaultApi.md#deletegitbranch) | **Post** /DeleteGitBranch | 仓库分支-删除代码仓库分支
*DefaultApi* | [**DeleteGitDeployKey**](docs/DefaultApi.md#deletegitdeploykey) | **Post** /DeleteGitDeployKey | 仓库设置-删除部署公钥
*DefaultApi* | [**DeleteGitDepot**](docs/DefaultApi.md#deletegitdepot) | **Post** /DeleteGitDepot | 仓库信息-删除git仓库
*DefaultApi* | [**DeleteGitFiles**](docs/DefaultApi.md#deletegitfiles) | **Post** /DeleteGitFiles | Git文件-删除文件并提交
*DefaultApi* | [**DeleteGitMergedBranches**](docs/DefaultApi.md#deletegitmergedbranches) | **Post** /DeleteGitMergedBranches | 仓库分支-删除已合并到默认分支的分支（此操作不会删除受保护的分支)
*DefaultApi* | [**DeleteGitProtectedTagRule**](docs/DefaultApi.md#deletegitprotectedtagrule) | **Post** /DeleteGitProtectedTagRule | 标签信息-删除标签保护规则
*DefaultApi* | [**DeleteGitRelease**](docs/DefaultApi.md#deletegitrelease) | **Post** /DeleteGitRelease | 版本信息-删除仓库版本
*DefaultApi* | [**DeleteGitTag**](docs/DefaultApi.md#deletegittag) | **Post** /DeleteGitTag | 标签信息-代码仓库删除tag
*DefaultApi* | [**DeleteIssue**](docs/DefaultApi.md#deleteissue) | **Post** /?action&#x3D;DeleteIssue | 事项删除
*DefaultApi* | [**DeleteIssueBlock**](docs/DefaultApi.md#deleteissueblock) | **Post** /?action&#x3D;DeleteIssueBlock | 前置事项删除
*DefaultApi* | [**DeleteIssueModule**](docs/DefaultApi.md#deleteissuemodule) | **Post** /DeleteIssueModule | 事项模块删除
*DefaultApi* | [**DeleteIssueWorkHours**](docs/DefaultApi.md#deleteissueworkhours) | **Post** /?action&#x3D;DeleteIssueWorkHours | 工时日志删除
*DefaultApi* | [**DeleteIteration**](docs/DefaultApi.md#deleteiteration) | **Post** /?action&#x3D;DeleteIteration | 迭代删除
*DefaultApi* | [**DeleteMemberSshKey**](docs/DefaultApi.md#deletemembersshkey) | **Post** /DeleteMemberSshKey | 仓库设置-删除团队成员的SSH公钥
*DefaultApi* | [**DeleteMergeRequestNote**](docs/DefaultApi.md#deletemergerequestnote) | **Post** /DeleteMergeRequestNote | 合并请求-删除合并请求行评论和改动文件diff行评论
*DefaultApi* | [**DeleteMergeRequestReviewer**](docs/DefaultApi.md#deletemergerequestreviewer) | **Post** /DeleteMergeRequestReviewer | 合并请求-删除mr评审者
*DefaultApi* | [**DeleteOneProject**](docs/DefaultApi.md#deleteoneproject) | **Post** /?action&#x3D;DeleteOneProject | 单个项目删除
*DefaultApi* | [**DeletePoliciesById**](docs/DefaultApi.md#deletepoliciesbyid) | **Post** /?action&#x3D;DeletePoliciesById | 权限组批量删除
*DefaultApi* | [**DeleteProgramMemberPolicy**](docs/DefaultApi.md#deleteprogrammemberpolicy) | **Post** /?action&#x3D;DeleteProgramMemberPolicy | 项目集成员权限组删除
*DefaultApi* | [**DeleteProjectAnnouncement**](docs/DefaultApi.md#deleteprojectannouncement) | **Post** /?action&#x3D;DeleteProjectAnnouncement | 项目公告删除
*DefaultApi* | [**DeleteProjectLabel**](docs/DefaultApi.md#deleteprojectlabel) | **Post** /?action&#x3D;DeleteProjectLabel | 项目标签删除
*DefaultApi* | [**DeleteProjectMemberPrincipal**](docs/DefaultApi.md#deleteprojectmemberprincipal) | **Post** /?action&#x3D;DeleteProjectMemberPrincipal | 项目成员主体删除(包含用户组、部门、成员)
*DefaultApi* | [**DeleteRelease**](docs/DefaultApi.md#deleterelease) | **Post** /DeleteRelease | 版本删除
*DefaultApi* | [**DeleteReport**](docs/DefaultApi.md#deletereport) | **Post** /?action&#x3D;DeleteReport | 测试报告删除
*DefaultApi* | [**DeleteRequirementDefectRelation**](docs/DefaultApi.md#deleterequirementdefectrelation) | **Post** /?action&#x3D;DeleteRequirementDefectRelation | 需求取消关联缺陷
*DefaultApi* | [**DeleteSshKey**](docs/DefaultApi.md#deletesshkey) | **Post** /?action&#x3D;DeleteSshKey | 仓库设置-删除当前用户的SSH公钥
*DefaultApi* | [**DeleteTeamLevelDepotSpec**](docs/DefaultApi.md#deleteteamleveldepotspec) | **Post** /DeleteTeamLevelDepotSpec | 仓库设置-删除团队级别的分支规范
*DefaultApi* | [**DeleteTeamMember**](docs/DefaultApi.md#deleteteammember) | **Post** /?action&#x3D;DeleteTeamMember | 团队成员删除
*DefaultApi* | [**DeleteTestCase**](docs/DefaultApi.md#deletetestcase) | **Post** /?action&#x3D;DeleteTestCase | 测试用例删除
*DefaultApi* | [**DeleteTestCaseSection**](docs/DefaultApi.md#deletetestcasesection) | **Post** /?action&#x3D;DeleteTestCaseSection | 测试用例分组删除
*DefaultApi* | [**DeleteTestRun**](docs/DefaultApi.md#deletetestrun) | **Post** /?action&#x3D;DeleteTestRun | 测试计划删除
*DefaultApi* | [**DeleteUserGroupByIds**](docs/DefaultApi.md#deleteusergroupbyids) | **Post** /?action&#x3D;DeleteUserGroupByIds | 用户组删除
*DefaultApi* | [**DeleteUserGroupUsers**](docs/DefaultApi.md#deleteusergroupusers) | **Post** /?action&#x3D;DeleteUserGroupUsers | 用户组删除用户
*DefaultApi* | [**DescribeAgentSecret**](docs/DefaultApi.md#describeagentsecret) | **Post** /?action&#x3D;DescribeAgentSecret | 堡垒机安装 Secret
*DefaultApi* | [**DescribeAllMergeRequestNotes**](docs/DefaultApi.md#describeallmergerequestnotes) | **Post** /DescribeAllMergeRequestNotes | 合并请求-获取所有合并请求行评论和改动文件diff行评论
*DefaultApi* | [**DescribeAllProjectLabels**](docs/DefaultApi.md#describeallprojectlabels) | **Post** /?action&#x3D;DescribeAllProjectLabels | 项目标签查询
*DefaultApi* | [**DescribeAllProjectsIssueWorkLogList**](docs/DefaultApi.md#describeallprojectsissueworkloglist) | **Post** /?action&#x3D;DescribeAllProjectsIssueWorkLogList | 工时日志列表查询
*DefaultApi* | [**DescribeArtifactChecksums**](docs/DefaultApi.md#describeartifactchecksums) | **Post** /?action&#x3D;DescribeArtifactChecksums | 制品Checksum列表查询
*DefaultApi* | [**DescribeArtifactCredit**](docs/DefaultApi.md#describeartifactcredit) | **Post** /?action&#x3D;DescribeArtifactCredit | 查询制品授信清单详情
*DefaultApi* | [**DescribeArtifactCreditList**](docs/DefaultApi.md#describeartifactcreditlist) | **Post** /?action&#x3D;DescribeArtifactCreditList | 制品授信清单列表查询
*DefaultApi* | [**DescribeArtifactFileDownloadUrl**](docs/DefaultApi.md#describeartifactfiledownloadurl) | **Post** /?action&#x3D;DescribeArtifactFileDownloadUrl | 制品文件临时下载链接获取
*DefaultApi* | [**DescribeArtifactPackageList**](docs/DefaultApi.md#describeartifactpackagelist) | **Post** /?action&#x3D;DescribeArtifactPackageList | 制品包（镜像）列表查询
*DefaultApi* | [**DescribeArtifactProperties**](docs/DefaultApi.md#describeartifactproperties) | **Post** /?action&#x3D;DescribeArtifactProperties | 制品属性列表查询
*DefaultApi* | [**DescribeArtifactRepositoryFileList**](docs/DefaultApi.md#describeartifactrepositoryfilelist) | **Post** /?action&#x3D;DescribeArtifactRepositoryFileList | 制品仓库下可下载的文件列表获取
*DefaultApi* | [**DescribeArtifactRepositoryList**](docs/DefaultApi.md#describeartifactrepositorylist) | **Post** /?action&#x3D;DescribeArtifactRepositoryList | 制品仓库列表查询
*DefaultApi* | [**DescribeArtifactVersionFileList**](docs/DefaultApi.md#describeartifactversionfilelist) | **Post** /?action&#x3D;DescribeArtifactVersionFileList | 制品版本可下载的文件列表获取
*DefaultApi* | [**DescribeArtifactVersionList**](docs/DefaultApi.md#describeartifactversionlist) | **Post** /?action&#x3D;DescribeArtifactVersionList | 制品版本列表查询
*DefaultApi* | [**DescribeAvailablePoliciesOnResource**](docs/DefaultApi.md#describeavailablepoliciesonresource) | **Post** /?action&#x3D;DescribeAvailablePoliciesOnResource | 权限组列表查询（指定资源）
*DefaultApi* | [**DescribeBlockIssueList**](docs/DefaultApi.md#describeblockissuelist) | **Post** /?action&#x3D;DescribeBlockIssueList | 后置事项查询
*DefaultApi* | [**DescribeBlockedByIssueList**](docs/DefaultApi.md#describeblockedbyissuelist) | **Post** /?action&#x3D;DescribeBlockedByIssueList | 前置事项查询
*DefaultApi* | [**DescribeBranchProtection**](docs/DefaultApi.md#describebranchprotection) | **Post** /DescribeBranchProtection | 仓库设置-查询单个保护分支规则
*DefaultApi* | [**DescribeBranchProtectionMembers**](docs/DefaultApi.md#describebranchprotectionmembers) | **Post** /DescribeBranchProtectionMembers | 仓库设置-查询保护分支规则下所有管理员信息
*DefaultApi* | [**DescribeBranchProtections**](docs/DefaultApi.md#describebranchprotections) | **Post** /DescribeBranchProtections | 仓库设置-查询仓库保护分支规则集合
*DefaultApi* | [**DescribeCanMerge**](docs/DefaultApi.md#describecanmerge) | **Post** /?action&#x3D;DescribeCanMerge | 合并请求-查看两个分支是否可以合并
*DefaultApi* | [**DescribeCdAgentMachines**](docs/DefaultApi.md#describecdagentmachines) | **Post** /?action&#x3D;DescribeCdAgentMachines | CD 堡垒机列表获取
*DefaultApi* | [**DescribeCdApplication**](docs/DefaultApi.md#describecdapplication) | **Post** /?action&#x3D;DescribeCdApplication | CD 应用详情获取
*DefaultApi* | [**DescribeCdApplications**](docs/DefaultApi.md#describecdapplications) | **Post** /?action&#x3D;DescribeCdApplications | CD 应用列表获取
*DefaultApi* | [**DescribeCdApplicationsByProject**](docs/DefaultApi.md#describecdapplicationsbyproject) | **Post** /?action&#x3D;DescribeCdApplicationsByProject | 关联应用列表获取（指定项目名）
*DefaultApi* | [**DescribeCdCloudAccounts**](docs/DefaultApi.md#describecdcloudaccounts) | **Post** /?action&#x3D;DescribeCdCloudAccounts | CD 云账号列表获取
*DefaultApi* | [**DescribeCdDeployCountByApplications**](docs/DefaultApi.md#describecddeploycountbyapplications) | **Post** /?action&#x3D;DescribeCdDeployCountByApplications | 发布次数-根据应用名列表获取
*DefaultApi* | [**DescribeCdDeployCountByProject**](docs/DefaultApi.md#describecddeploycountbyproject) | **Post** /?action&#x3D;DescribeCdDeployCountByProject | 关联应用的发布次数获取（指定项目名）
*DefaultApi* | [**DescribeCdDeployTimeByApplications**](docs/DefaultApi.md#describecddeploytimebyapplications) | **Post** /?action&#x3D;DescribeCdDeployTimeByApplications | 发布时长-根据应用名列表获取
*DefaultApi* | [**DescribeCdDeployTimeByProject**](docs/DefaultApi.md#describecddeploytimebyproject) | **Post** /?action&#x3D;DescribeCdDeployTimeByProject | 关联应用的发布时长-根据项目名获取
*DefaultApi* | [**DescribeCdDeployTrendByApplications**](docs/DefaultApi.md#describecddeploytrendbyapplications) | **Post** /?action&#x3D;DescribeCdDeployTrendByApplications | 发布趋势-根据应用名列表获取
*DefaultApi* | [**DescribeCdDeployTrendByProject**](docs/DefaultApi.md#describecddeploytrendbyproject) | **Post** /?action&#x3D;DescribeCdDeployTrendByProject | 关联应用的发布趋势-根据项目名获取
*DefaultApi* | [**DescribeCdHostServerGroup**](docs/DefaultApi.md#describecdhostservergroup) | **Post** /?action&#x3D;DescribeCdHostServerGroup | CD 主机组获取
*DefaultApi* | [**DescribeCdHostServerGroups**](docs/DefaultApi.md#describecdhostservergroups) | **Post** /?action&#x3D;DescribeCdHostServerGroups | CD 主机组列表获取
*DefaultApi* | [**DescribeCdPipeline**](docs/DefaultApi.md#describecdpipeline) | **Post** /?action&#x3D;DescribeCdPipeline | CD 部署流程执行记录获取
*DefaultApi* | [**DescribeCdPipelineConfig**](docs/DefaultApi.md#describecdpipelineconfig) | **Post** /?action&#x3D;DescribeCdPipelineConfig | CD 部署流程配置-根据名称获取
*DefaultApi* | [**DescribeCdPipelineConfigs**](docs/DefaultApi.md#describecdpipelineconfigs) | **Post** /?action&#x3D;DescribeCdPipelineConfigs | CD 应用下的所有部署流程配置获取
*DefaultApi* | [**DescribeCdTask**](docs/DefaultApi.md#describecdtask) | **Post** /?action&#x3D;DescribeCdTask | CD 任务执行记录获取
*DefaultApi* | [**DescribeCodeSearch**](docs/DefaultApi.md#describecodesearch) | **Post** /DescribeCodeSearch | 仓库信息-查询代码片段详细列表
*DefaultApi* | [**DescribeCodingCIBuild**](docs/DefaultApi.md#describecodingcibuild) | **Post** /?action&#x3D;DescribeCodingCIBuild | 构建记录详情查询
*DefaultApi* | [**DescribeCodingCIBuildArtifacts**](docs/DefaultApi.md#describecodingcibuildartifacts) | **Post** /?action&#x3D;DescribeCodingCIBuildArtifacts | 构建任务制品查询
*DefaultApi* | [**DescribeCodingCIBuildEnvs**](docs/DefaultApi.md#describecodingcibuildenvs) | **Post** /?action&#x3D;DescribeCodingCIBuildEnvs | 构建计划环境变量获取
*DefaultApi* | [**DescribeCodingCIBuildHtmlReports**](docs/DefaultApi.md#describecodingcibuildhtmlreports) | **Post** /?action&#x3D;DescribeCodingCIBuildHtmlReports | 构建任务网页报告查询
*DefaultApi* | [**DescribeCodingCIBuildLog**](docs/DefaultApi.md#describecodingcibuildlog) | **Post** /?action&#x3D;DescribeCodingCIBuildLog | 构建日志获取
*DefaultApi* | [**DescribeCodingCIBuildLogRaw**](docs/DefaultApi.md#describecodingcibuildlograw) | **Post** /?action&#x3D;DescribeCodingCIBuildLogRaw | 构建完整日志查询（原始日志 Raw）
*DefaultApi* | [**DescribeCodingCIBuildMetrics**](docs/DefaultApi.md#describecodingcibuildmetrics) | **Post** /?action&#x3D;DescribeCodingCIBuildMetrics | 构建计划度量查询
*DefaultApi* | [**DescribeCodingCIBuildStage**](docs/DefaultApi.md#describecodingcibuildstage) | **Post** /?action&#x3D;DescribeCodingCIBuildStage | 构建任务阶段获取
*DefaultApi* | [**DescribeCodingCIBuildStatistics**](docs/DefaultApi.md#describecodingcibuildstatistics) | **Post** /?action&#x3D;DescribeCodingCIBuildStatistics | 构建任务统计
*DefaultApi* | [**DescribeCodingCIBuildStep**](docs/DefaultApi.md#describecodingcibuildstep) | **Post** /?action&#x3D;DescribeCodingCIBuildStep | 构建任务指定阶段的步骤获取
*DefaultApi* | [**DescribeCodingCIBuildStepLog**](docs/DefaultApi.md#describecodingcibuildsteplog) | **Post** /?action&#x3D;DescribeCodingCIBuildStepLog | 构建步骤日志获取
*DefaultApi* | [**DescribeCodingCIBuilds**](docs/DefaultApi.md#describecodingcibuilds) | **Post** /?action&#x3D;DescribeCodingCIBuilds | 构建计划的构建列表获取
*DefaultApi* | [**DescribeCodingCIJob**](docs/DefaultApi.md#describecodingcijob) | **Post** /?action&#x3D;DescribeCodingCIJob | 构建计划详情获取
*DefaultApi* | [**DescribeCodingCIJobs**](docs/DefaultApi.md#describecodingcijobs) | **Post** /?action&#x3D;DescribeCodingCIJobs | 构建计划查询（通过项目ID）
*DefaultApi* | [**DescribeCodingCurrentUser**](docs/DefaultApi.md#describecodingcurrentuser) | **Post** /?action&#x3D;DescribeCodingCurrentUser | 当前用户信息查询
*DefaultApi* | [**DescribeCodingProjects**](docs/DefaultApi.md#describecodingprojects) | **Post** /?action&#x3D;DescribeCodingProjects | 项目列表查询
*DefaultApi* | [**DescribeCommitRefs**](docs/DefaultApi.md#describecommitrefs) | **Post** /DescribeCommitRefs | Git提交-查询commit的ref信息
*DefaultApi* | [**DescribeCommitsBetweenCommitAndCommit**](docs/DefaultApi.md#describecommitsbetweencommitandcommit) | **Post** /DescribeCommitsBetweenCommitAndCommit | Git提交-查询两个请求之间的请求列表（source target顺序正常）
*DefaultApi* | [**DescribeConfigTemplateList**](docs/DefaultApi.md#describeconfigtemplatelist) | **Post** /DescribeConfigTemplateList | 配置方案获取
*DefaultApi* | [**DescribeDepartment**](docs/DefaultApi.md#describedepartment) | **Post** /?action&#x3D;DescribeDepartment | 部门详情查询
*DefaultApi* | [**DescribeDepartmentMembers**](docs/DefaultApi.md#describedepartmentmembers) | **Post** /?action&#x3D;DescribeDepartmentMembers | 部门成员列表查询
*DefaultApi* | [**DescribeDepotByNameInfo**](docs/DefaultApi.md#describedepotbynameinfo) | **Post** /DescribeDepotByNameInfo | 仓库信息-查询项目下所有的仓库信息列表
*DefaultApi* | [**DescribeDepotDefaultBranch**](docs/DefaultApi.md#describedepotdefaultbranch) | **Post** /DescribeDepotDefaultBranch | 仓库分支-查询仓库的默认分支
*DefaultApi* | [**DescribeDepotFilePushRules**](docs/DefaultApi.md#describedepotfilepushrules) | **Post** /DescribeDepotFilePushRules | 仓库设置-查询git仓库文件推送规则
*DefaultApi* | [**DescribeDepotMergeRequests**](docs/DefaultApi.md#describedepotmergerequests) | **Post** /DescribeDepotMergeRequests | 合并请求-查询仓库合并请求列表
*DefaultApi* | [**DescribeDepotPushSetting**](docs/DefaultApi.md#describedepotpushsetting) | **Post** /DescribeDepotPushSetting | 仓库设置-查询仓库推送设置
*DefaultApi* | [**DescribeDepotSpecDetail**](docs/DefaultApi.md#describedepotspecdetail) | **Post** /DescribeDepotSpecDetail | 仓库设置-查询仓库规范详情
*DefaultApi* | [**DescribeDepotSpecs**](docs/DefaultApi.md#describedepotspecs) | **Post** /DescribeDepotSpecs | 仓库设置-查询仓库规范列表
*DefaultApi* | [**DescribeDifferentBetween2Commits**](docs/DefaultApi.md#describedifferentbetween2commits) | **Post** /?action&#x3D;DescribeDifferentBetween2Commits | Git提交-两次提交之间的文件差异（source target顺序正常）
*DefaultApi* | [**DescribeDifferentBetweenTwoCommits**](docs/DefaultApi.md#describedifferentbetweentwocommits) | **Post** /DescribeDifferentBetweenTwoCommits | Git提交-获取两次commit之间的文件差异详情(废弃，source target顺序不一致)
*DefaultApi* | [**DescribeGitBlameInfo**](docs/DefaultApi.md#describegitblameinfo) | **Post** /DescribeGitBlameInfo | Git提交-获取指定commit下某文件指定代码行的最后一次提交
*DefaultApi* | [**DescribeGitBlob**](docs/DefaultApi.md#describegitblob) | **Post** /DescribeGitBlob | Git文件-查询GitBlob
*DefaultApi* | [**DescribeGitBlobRaw**](docs/DefaultApi.md#describegitblobraw) | **Post** /DescribeGitBlobRaw | Git文件-查询Git Blob raw信息
*DefaultApi* | [**DescribeGitBranch**](docs/DefaultApi.md#describegitbranch) | **Post** /DescribeGitBranch | 仓库分支-查询代码仓库单个分支
*DefaultApi* | [**DescribeGitBranchList**](docs/DefaultApi.md#describegitbranchlist) | **Post** /DescribeGitBranchList | 仓库分支-查询仓库分支列表
*DefaultApi* | [**DescribeGitBranches**](docs/DefaultApi.md#describegitbranches) | **Post** /DescribeGitBranches | 仓库分支-查询仓库下所有分支列表
*DefaultApi* | [**DescribeGitBranchesBySha**](docs/DefaultApi.md#describegitbranchesbysha) | **Post** /DescribeGitBranchesBySha | 仓库分支-根据sha值查询所在分支
*DefaultApi* | [**DescribeGitCommitComments**](docs/DefaultApi.md#describegitcommitcomments) | **Post** /DescribeGitCommitComments | Git提交-获取commit评论
*DefaultApi* | [**DescribeGitCommitDiff**](docs/DefaultApi.md#describegitcommitdiff) | **Post** /DescribeGitCommitDiff | Git提交-查询某次提交的diff信息
*DefaultApi* | [**DescribeGitCommitFilePathList**](docs/DefaultApi.md#describegitcommitfilepathlist) | **Post** /DescribeGitCommitFilePathList | Git提交-查询仓库某次提交改动的文件路径列表
*DefaultApi* | [**DescribeGitCommitInfo**](docs/DefaultApi.md#describegitcommitinfo) | **Post** /DescribeGitCommitInfo | Git提交-查询单个请求详情信息
*DefaultApi* | [**DescribeGitCommitInfos**](docs/DefaultApi.md#describegitcommitinfos) | **Post** /DescribeGitCommitInfos | Git提交-查询仓库分支下提交列表
*DefaultApi* | [**DescribeGitCommitNote**](docs/DefaultApi.md#describegitcommitnote) | **Post** /DescribeGitCommitNote | Git提交-获取提交注释
*DefaultApi* | [**DescribeGitCommitStatus**](docs/DefaultApi.md#describegitcommitstatus) | **Post** /DescribeGitCommitStatus | Git提交-查询提交对应的流水线状态
*DefaultApi* | [**DescribeGitCommitsInPage**](docs/DefaultApi.md#describegitcommitsinpage) | **Post** /DescribeGitCommitsInPage | Git提交-查询仓库分支下提交列表
*DefaultApi* | [**DescribeGitContributors**](docs/DefaultApi.md#describegitcontributors) | **Post** /DescribeGitContributors | 仓库信息-查询git仓库的贡献者
*DefaultApi* | [**DescribeGitDepot**](docs/DefaultApi.md#describegitdepot) | **Post** /DescribeGitDepot | 仓库信息-根据代码仓库id获取代码仓库信息
*DefaultApi* | [**DescribeGitDepotDeployKeys**](docs/DefaultApi.md#describegitdepotdeploykeys) | **Post** /DescribeGitDepotDeployKeys | 仓库设置-查询某仓库下的部署公钥列表
*DefaultApi* | [**DescribeGitFile**](docs/DefaultApi.md#describegitfile) | **Post** /DescribeGitFile | Git文件-获取文件详情
*DefaultApi* | [**DescribeGitFileContent**](docs/DefaultApi.md#describegitfilecontent) | **Post** /DescribeGitFileContent | Git提交-查询某次提交某文件的内容
*DefaultApi* | [**DescribeGitFileStat**](docs/DefaultApi.md#describegitfilestat) | **Post** /DescribeGitFileStat | Git文件-检查仓库文件是否存在
*DefaultApi* | [**DescribeGitFiles**](docs/DefaultApi.md#describegitfiles) | **Post** /DescribeGitFiles | Git文件-查询仓库目录下文件和文件夹名字
*DefaultApi* | [**DescribeGitMergeBase**](docs/DefaultApi.md#describegitmergebase) | **Post** /DescribeGitMergeBase | 仓库分支-查询两个分支的公共祖先
*DefaultApi* | [**DescribeGitMergeRequestDiffDetail**](docs/DefaultApi.md#describegitmergerequestdiffdetail) | **Post** /DescribeGitMergeRequestDiffDetail | 合并请求-查询合并请求文件的 diff 详情
*DefaultApi* | [**DescribeGitMergeRequestDiffs**](docs/DefaultApi.md#describegitmergerequestdiffs) | **Post** /DescribeGitMergeRequestDiffs | 合并请求-查询合并请求diff信息的文件列表
*DefaultApi* | [**DescribeGitMergeRequestParticipants**](docs/DefaultApi.md#describegitmergerequestparticipants) | **Post** /DescribeGitMergeRequestParticipants | 合并请求-获取合并请求的参与者
*DefaultApi* | [**DescribeGitMergeRequestsBySha**](docs/DefaultApi.md#describegitmergerequestsbysha) | **Post** /DescribeGitMergeRequestsBySha | 合并请求-查询含有某次提交的合并请求
*DefaultApi* | [**DescribeGitProjectDeployKeys**](docs/DefaultApi.md#describegitprojectdeploykeys) | **Post** /DescribeGitProjectDeployKeys | 仓库设置-查询某项目下的部署公钥列表
*DefaultApi* | [**DescribeGitProtectedTags**](docs/DefaultApi.md#describegitprotectedtags) | **Post** /DescribeGitProtectedTags | 标签信息-查询受保护的标签列表
*DefaultApi* | [**DescribeGitProtectedTagsByRule**](docs/DefaultApi.md#describegitprotectedtagsbyrule) | **Post** /DescribeGitProtectedTagsByRule | 标签信息-根据标签保护规则查询受保护的标签列表
*DefaultApi* | [**DescribeGitRef**](docs/DefaultApi.md#describegitref) | **Post** /DescribeGitRef | 仓库分支-获取分支或标签信息
*DefaultApi* | [**DescribeGitRefsBySha**](docs/DefaultApi.md#describegitrefsbysha) | **Post** /DescribeGitRefsBySha | Git提交-查询含有某次提交的标签或分支列表
*DefaultApi* | [**DescribeGitReleaseDetail**](docs/DefaultApi.md#describegitreleasedetail) | **Post** /DescribeGitReleaseDetail | 版本信息-查询仓库的版本详情
*DefaultApi* | [**DescribeGitReleases**](docs/DefaultApi.md#describegitreleases) | **Post** /DescribeGitReleases | 版本信息-查询仓库的版本列表
*DefaultApi* | [**DescribeGitTag**](docs/DefaultApi.md#describegittag) | **Post** /DescribeGitTag | 标签信息-查询单个tag
*DefaultApi* | [**DescribeGitTags**](docs/DefaultApi.md#describegittags) | **Post** /DescribeGitTags | 标签信息-查询当前仓库下所有tags
*DefaultApi* | [**DescribeGitTagsByBranch**](docs/DefaultApi.md#describegittagsbybranch) | **Post** /DescribeGitTagsByBranch | 标签信息-根据分支获取标签列表
*DefaultApi* | [**DescribeGitTagsBySha**](docs/DefaultApi.md#describegittagsbysha) | **Post** /DescribeGitTagsBySha | 标签信息-查询含有某次提交的标签列表
*DefaultApi* | [**DescribeGitTree**](docs/DefaultApi.md#describegittree) | **Post** /DescribeGitTree | 仓库信息-查询git仓库的树
*DefaultApi* | [**DescribeGrantObjectsOnResource**](docs/DefaultApi.md#describegrantobjectsonresource) | **Post** /?action&#x3D;DescribeGrantObjectsOnResource | 授权对象列表分页查询
*DefaultApi* | [**DescribeGrantUsersOnResource**](docs/DefaultApi.md#describegrantusersonresource) | **Post** /?action&#x3D;DescribeGrantUsersOnResource | 授权用户列表分页查询
*DefaultApi* | [**DescribeHostServerInstance**](docs/DefaultApi.md#describehostserverinstance) | **Post** /?action&#x3D;DescribeHostServerInstance | CD 主机组部署详情获取
*DefaultApi* | [**DescribeIssue**](docs/DefaultApi.md#describeissue) | **Post** /DescribeIssue | 事项详情查询
*DefaultApi* | [**DescribeIssueAttachmentPreSignedUrl**](docs/DefaultApi.md#describeissueattachmentpresignedurl) | **Post** /?action&#x3D;DescribeIssueAttachmentPreSignedUrl | 预签名信息获取
*DefaultApi* | [**DescribeIssueCommentList**](docs/DefaultApi.md#describeissuecommentlist) | **Post** /?action&#x3D;DescribeIssueCommentList | 事项评论列表查询
*DefaultApi* | [**DescribeIssueCustomFieldLogList**](docs/DefaultApi.md#describeissuecustomfieldloglist) | **Post** /DescribeIssueCustomFieldLogList | 事项的自定义属性变更日志查询
*DefaultApi* | [**DescribeIssueFileUrl**](docs/DefaultApi.md#describeissuefileurl) | **Post** /?action&#x3D;DescribeIssueFileUrl | 事项附件的下载地址查询
*DefaultApi* | [**DescribeIssueFilterList**](docs/DefaultApi.md#describeissuefilterlist) | **Post** /?action&#x3D;DescribeIssueFilterList | 事项筛选器列表查询
*DefaultApi* | [**DescribeIssueList**](docs/DefaultApi.md#describeissuelist) | **Post** /?action&#x3D;DescribeIssueList | 事项列表查询
*DefaultApi* | [**DescribeIssueListWithPage**](docs/DefaultApi.md#describeissuelistwithpage) | **Post** /DescribeIssueListWithPage | 事项查询,返回分页信息
*DefaultApi* | [**DescribeIssueModuleList**](docs/DefaultApi.md#describeissuemodulelist) | **Post** /DescribeIssueModuleList | 事项模块列表查询
*DefaultApi* | [**DescribeIssueReferenceResources**](docs/DefaultApi.md#describeissuereferenceresources) | **Post** /DescribeIssueReferenceResources | 事项的引用资源列表查询
*DefaultApi* | [**DescribeIssueRelatedWorkItemList**](docs/DefaultApi.md#describeissuerelatedworkitemlist) | **Post** /DescribeIssueRelatedWorkItemList | 事项关联的项目集中的工作项查询
*DefaultApi* | [**DescribeIssueReleaseList**](docs/DefaultApi.md#describeissuereleaselist) | **Post** /DescribeIssueReleaseList | 事项加入的版本查询
*DefaultApi* | [**DescribeIssueStatusChangeLogList**](docs/DefaultApi.md#describeissuestatuschangeloglist) | **Post** /DescribeIssueStatusChangeLogList | 事项的状态变更记录查询
*DefaultApi* | [**DescribeIssueWorkLogList**](docs/DefaultApi.md#describeissueworkloglist) | **Post** /DescribeIssueWorkLogList | 事项的工时日志查询
*DefaultApi* | [**DescribeIteration**](docs/DefaultApi.md#describeiteration) | **Post** /?action&#x3D;DescribeIteration | 迭代详情查询
*DefaultApi* | [**DescribeIterationList**](docs/DefaultApi.md#describeiterationlist) | **Post** /?action&#x3D;DescribeIterationList | 迭代列表获取
*DefaultApi* | [**DescribeMemberSshKey**](docs/DefaultApi.md#describemembersshkey) | **Post** /DescribeMemberSshKey | 仓库设置-获取团队成员的SSH公钥列表
*DefaultApi* | [**DescribeMergeReqCommits**](docs/DefaultApi.md#describemergereqcommits) | **Post** /DescribeMergeReqCommits | 合并请求-查询合并请求列表
*DefaultApi* | [**DescribeMergeReqInfo**](docs/DefaultApi.md#describemergereqinfo) | **Post** /DescribeMergeReqInfo | 合并请求-查询合并请求详情
*DefaultApi* | [**DescribeMergeRequest**](docs/DefaultApi.md#describemergerequest) | **Post** /DescribeMergeRequest | 合并请求-查询合并请求详情信息
*DefaultApi* | [**DescribeMergeRequestFileDiff**](docs/DefaultApi.md#describemergerequestfilediff) | **Post** /DescribeMergeRequestFileDiff | 合并请求-获取合并请求文件修改记录
*DefaultApi* | [**DescribeMergeRequestLog**](docs/DefaultApi.md#describemergerequestlog) | **Post** /DescribeMergeRequestLog | 合并请求-查询合并请求操作记录
*DefaultApi* | [**DescribeMergeRequestReviewers**](docs/DefaultApi.md#describemergerequestreviewers) | **Post** /DescribeMergeRequestReviewers | 合并请求-获取合并请求的评审者
*DefaultApi* | [**DescribeMyDepots**](docs/DefaultApi.md#describemydepots) | **Post** /DescribeMyDepots | 仓库信息-获取当前用户拥有读权限的仓库列表
*DefaultApi* | [**DescribeNotesByCommits**](docs/DefaultApi.md#describenotesbycommits) | **Post** /DescribeNotesByCommits | 仓库信息-获取提交的note信息
*DefaultApi* | [**DescribeOneProject**](docs/DefaultApi.md#describeoneproject) | **Post** /?action&#x3D;DescribeOneProject | 单个项目查询
*DefaultApi* | [**DescribePersonalExternalDepots**](docs/DefaultApi.md#describepersonalexternaldepots) | **Post** /?action&#x3D;DescribePersonalExternalDepots | 个人外部仓库获取
*DefaultApi* | [**DescribePinyin**](docs/DefaultApi.md#describepinyin) | **Post** /?action&#x3D;DescribePinyin | 汉字转拼音
*DefaultApi* | [**DescribePoliciesOnResourceType**](docs/DefaultApi.md#describepoliciesonresourcetype) | **Post** /?action&#x3D;DescribePoliciesOnResourceType | 权限组列表查询（指定资源类型）
*DefaultApi* | [**DescribePolicy**](docs/DefaultApi.md#describepolicy) | **Post** /?action&#x3D;DescribePolicy | 权限组详情获取
*DefaultApi* | [**DescribePreSignUploadUrl**](docs/DefaultApi.md#describepresignuploadurl) | **Post** /?action&#x3D;DescribePreSignUploadUrl | 预签名URL获取
*DefaultApi* | [**DescribePredicatePolicyOnResource**](docs/DefaultApi.md#describepredicatepolicyonresource) | **Post** /?action&#x3D;DescribePredicatePolicyOnResource | 资源权限判定模式获取
*DefaultApi* | [**DescribeProgramProjects**](docs/DefaultApi.md#describeprogramprojects) | **Post** /?action&#x3D;DescribeProgramProjects | 项目集下项目列表查询
*DefaultApi* | [**DescribePrograms**](docs/DefaultApi.md#describeprograms) | **Post** /?action&#x3D;DescribePrograms | 项目集列表查询
*DefaultApi* | [**DescribeProjectAnnouncement**](docs/DefaultApi.md#describeprojectannouncement) | **Post** /?action&#x3D;DescribeProjectAnnouncement | 项目公告查询
*DefaultApi* | [**DescribeProjectAnnouncements**](docs/DefaultApi.md#describeprojectannouncements) | **Post** /?action&#x3D;DescribeProjectAnnouncements | 项目公告列表查询
*DefaultApi* | [**DescribeProjectByName**](docs/DefaultApi.md#describeprojectbyname) | **Post** /?action&#x3D;DescribeProjectByName | 项目查询(通过项目名称)
*DefaultApi* | [**DescribeProjectCredentials**](docs/DefaultApi.md#describeprojectcredentials) | **Post** /?action&#x3D;DescribeProjectCredentials | 项目凭据列表查询
*DefaultApi* | [**DescribeProjectDepotBranches**](docs/DefaultApi.md#describeprojectdepotbranches) | **Post** /?action&#x3D;DescribeProjectDepotBranches | 仓库分支列表获取
*DefaultApi* | [**DescribeProjectDepotCommits**](docs/DefaultApi.md#describeprojectdepotcommits) | **Post** /?action&#x3D;DescribeProjectDepotCommits | 分支下的提交列表获取
*DefaultApi* | [**DescribeProjectDepotInfoList**](docs/DefaultApi.md#describeprojectdepotinfolist) | **Post** /DescribeProjectDepotInfoList | 仓库信息-查询项目下所有的仓库信息列表
*DefaultApi* | [**DescribeProjectDepotTags**](docs/DefaultApi.md#describeprojectdepottags) | **Post** /?action&#x3D;DescribeProjectDepotTags | 仓库的标签列表获取
*DefaultApi* | [**DescribeProjectDepots**](docs/DefaultApi.md#describeprojectdepots) | **Post** /?action&#x3D;DescribeProjectDepots | 项目仓库列表获取
*DefaultApi* | [**DescribeProjectIssueFieldList**](docs/DefaultApi.md#describeprojectissuefieldlist) | **Post** /?action&#x3D;DescribeProjectIssueFieldList | 具体事项类型的属性列表查询
*DefaultApi* | [**DescribeProjectIssueStatusList**](docs/DefaultApi.md#describeprojectissuestatuslist) | **Post** /?action&#x3D;DescribeProjectIssueStatusList | 具体事项类型的状态列表查询
*DefaultApi* | [**DescribeProjectIssueTypeList**](docs/DefaultApi.md#describeprojectissuetypelist) | **Post** /?action&#x3D;DescribeProjectIssueTypeList | 项目事项类型列表查询
*DefaultApi* | [**DescribeProjectLabels**](docs/DefaultApi.md#describeprojectlabels) | **Post** /?action&#x3D;DescribeProjectLabels | 项目列表查询-指定项目标签
*DefaultApi* | [**DescribeProjectMemberPrincipals**](docs/DefaultApi.md#describeprojectmemberprincipals) | **Post** /?action&#x3D;DescribeProjectMemberPrincipals | 项目成员主体查询(包含用户组、部门、成员)
*DefaultApi* | [**DescribeProjectMembers**](docs/DefaultApi.md#describeprojectmembers) | **Post** /?action&#x3D;DescribeProjectMembers | 项目成员列表查询
*DefaultApi* | [**DescribeProjectMergeRequests**](docs/DefaultApi.md#describeprojectmergerequests) | **Post** /DescribeProjectMergeRequests | 合并请求-获取项目下的合并请求列表
*DefaultApi* | [**DescribeProjectRoles**](docs/DefaultApi.md#describeprojectroles) | **Post** /?action&#x3D;DescribeProjectRoles | 项目用户组查询
*DefaultApi* | [**DescribeProjectsByFeature**](docs/DefaultApi.md#describeprojectsbyfeature) | **Post** /?action&#x3D;DescribeProjectsByFeature | 项目查询（通过一级菜单名）
*DefaultApi* | [**DescribeProtectedBranch**](docs/DefaultApi.md#describeprotectedbranch) | **Post** /DescribeProtectedBranch | 仓库设置-查询保护分支详情
*DefaultApi* | [**DescribeProtectedBranchMembers**](docs/DefaultApi.md#describeprotectedbranchmembers) | **Post** /DescribeProtectedBranchMembers | 仓库设置-查询保护分支成员
*DefaultApi* | [**DescribeProtectedBranches**](docs/DefaultApi.md#describeprotectedbranches) | **Post** /DescribeProtectedBranches | 仓库设置-查询保护分支列表
*DefaultApi* | [**DescribeRelatedCaseList**](docs/DefaultApi.md#describerelatedcaselist) | **Post** /DescribeRelatedCaseList | 事项关联的测试用例查询
*DefaultApi* | [**DescribeRelease**](docs/DefaultApi.md#describerelease) | **Post** /DescribeRelease | 版本详情查询
*DefaultApi* | [**DescribeReleaseIssueList**](docs/DefaultApi.md#describereleaseissuelist) | **Post** /DescribeReleaseIssueList | 版本发布范围查询
*DefaultApi* | [**DescribeReleaseList**](docs/DefaultApi.md#describereleaselist) | **Post** /DescribeReleaseList | 版本列表查询
*DefaultApi* | [**DescribeReport**](docs/DefaultApi.md#describereport) | **Post** /?action&#x3D;DescribeReport | 测试报告详情
*DefaultApi* | [**DescribeReportList**](docs/DefaultApi.md#describereportlist) | **Post** /?action&#x3D;DescribeReportList | 测试报告列表
*DefaultApi* | [**DescribeRequirementDefectRelation**](docs/DefaultApi.md#describerequirementdefectrelation) | **Post** /?action&#x3D;DescribeRequirementDefectRelation | 需求关联缺陷列表查询
*DefaultApi* | [**DescribeRequirementTestCaseList**](docs/DefaultApi.md#describerequirementtestcaselist) | **Post** /?action&#x3D;DescribeRequirementTestCaseList | 需求关联的测试用例列表
*DefaultApi* | [**DescribeResourceReferences**](docs/DefaultApi.md#describeresourcereferences) | **Post** /?action&#x3D;DescribeResourceReferences | 资源引用的资源列表，如 开发任务中引用了多个需求，获取任务引用的需求列表
*DefaultApi* | [**DescribeResourceReferencesCited**](docs/DefaultApi.md#describeresourcereferencescited) | **Post** /?action&#x3D;DescribeResourceReferencesCited | 被引用资源列表查询
*DefaultApi* | [**DescribeResourceReferencesCiting**](docs/DefaultApi.md#describeresourcereferencesciting) | **Post** /?action&#x3D;DescribeResourceReferencesCiting | 引用资源列表查询
*DefaultApi* | [**DescribeResourceScopeListOnPolicy**](docs/DefaultApi.md#describeresourcescopelistonpolicy) | **Post** /?action&#x3D;DescribeResourceScopeListOnPolicy | 权限组可用资源范围分页查询
*DefaultApi* | [**DescribeSelfMergeRequests**](docs/DefaultApi.md#describeselfmergerequests) | **Post** /DescribeSelfMergeRequests | 合并请求-获取自己的合并请求列表
*DefaultApi* | [**DescribeSingeMergeRequestNotes**](docs/DefaultApi.md#describesingemergerequestnotes) | **Post** /DescribeSingeMergeRequestNotes | 合并请求-获取单个合并请求行评论和改动文件diff行评论
*DefaultApi* | [**DescribeSshKey**](docs/DefaultApi.md#describesshkey) | **Post** /?action&#x3D;DescribeSshKey | 仓库设置-获取当前用户所有SSH公钥
*DefaultApi* | [**DescribeSubIssueList**](docs/DefaultApi.md#describesubissuelist) | **Post** /?action&#x3D;DescribeSubIssueList | 子事项列表查询
*DefaultApi* | [**DescribeTeam**](docs/DefaultApi.md#describeteam) | **Post** /?action&#x3D;DescribeTeam | 团队信息查询
*DefaultApi* | [**DescribeTeamAdminMembers**](docs/DefaultApi.md#describeteamadminmembers) | **Post** /?action&#x3D;DescribeTeamAdminMembers | 团队管理员查询
*DefaultApi* | [**DescribeTeamArtifacts**](docs/DefaultApi.md#describeteamartifacts) | **Post** /?action&#x3D;DescribeTeamArtifacts | 制品列表查询
*DefaultApi* | [**DescribeTeamDepotInfoList**](docs/DefaultApi.md#describeteamdepotinfolist) | **Post** /DescribeTeamDepotInfoList | 仓库信息-获取团队下仓库列表，仅团队所有者或团队管理员可以调用。
*DefaultApi* | [**DescribeTeamDepots**](docs/DefaultApi.md#describeteamdepots) | **Post** /?action&#x3D;DescribeTeamDepots | 团队下可访问的所有仓库列表获取
*DefaultApi* | [**DescribeTeamIssueTypeList**](docs/DefaultApi.md#describeteamissuetypelist) | **Post** /?action&#x3D;DescribeTeamIssueTypeList | 企业事项类型列表查询
*DefaultApi* | [**DescribeTeamMember**](docs/DefaultApi.md#describeteammember) | **Post** /?action&#x3D;DescribeTeamMember | 团队成员信息查询（通过用户 ID）
*DefaultApi* | [**DescribeTeamMemberByEmail**](docs/DefaultApi.md#describeteammemberbyemail) | **Post** /?action&#x3D;DescribeTeamMemberByEmail | 团队成员信息查询（通过用户 Email）
*DefaultApi* | [**DescribeTeamMembers**](docs/DefaultApi.md#describeteammembers) | **Post** /?action&#x3D;DescribeTeamMembers | 团队成员列表查询
*DefaultApi* | [**DescribeTest**](docs/DefaultApi.md#describetest) | **Post** /?action&#x3D;DescribeTest | 测试任务详情
*DefaultApi* | [**DescribeTestCase**](docs/DefaultApi.md#describetestcase) | **Post** /?action&#x3D;DescribeTestCase | 测试用例详情
*DefaultApi* | [**DescribeTestCaseList**](docs/DefaultApi.md#describetestcaselist) | **Post** /?action&#x3D;DescribeTestCaseList | 测试用例列表
*DefaultApi* | [**DescribeTestCaseSectionList**](docs/DefaultApi.md#describetestcasesectionlist) | **Post** /?action&#x3D;DescribeTestCaseSectionList | 测试用例分组列表
*DefaultApi* | [**DescribeTestDefectList**](docs/DefaultApi.md#describetestdefectlist) | **Post** /?action&#x3D;DescribeTestDefectList | 测试任务关联的缺陷列表
*DefaultApi* | [**DescribeTestList**](docs/DefaultApi.md#describetestlist) | **Post** /?action&#x3D;DescribeTestList | 测试任务列表
*DefaultApi* | [**DescribeTestRun**](docs/DefaultApi.md#describetestrun) | **Post** /?action&#x3D;DescribeTestRun | 测试计划详情
*DefaultApi* | [**DescribeTestRunList**](docs/DefaultApi.md#describetestrunlist) | **Post** /?action&#x3D;DescribeTestRunList | 测试计划列表
*DefaultApi* | [**DescribeUserGroups**](docs/DefaultApi.md#describeusergroups) | **Post** /?action&#x3D;DescribeUserGroups | 用户组列表分页查询
*DefaultApi* | [**DescribeUserProjects**](docs/DefaultApi.md#describeuserprojects) | **Post** /?action&#x3D;DescribeUserProjects | 项目列表查询（已加入的项目）
*DefaultApi* | [**DescribeUsersByGroupId**](docs/DefaultApi.md#describeusersbygroupid) | **Post** /?action&#x3D;DescribeUsersByGroupId | 用户列表查询（根据用户组id分页查询）
*DefaultApi* | [**DescribeUsersOnResourceAndGrantObject**](docs/DefaultApi.md#describeusersonresourceandgrantobject) | **Post** /?action&#x3D;DescribeUsersOnResourceAndGrantObject | 授权用户列表分页查询（指定资源）
*DefaultApi* | [**DescribeWorkItemSalvage**](docs/DefaultApi.md#describeworkitemsalvage) | **Post** /DescribeWorkItemSalvage | 事项分解信息查询
*DefaultApi* | [**DescribeWorkbenchIssueList**](docs/DefaultApi.md#describeworkbenchissuelist) | **Post** /DescribeWorkbenchIssueList | 用户在团队内的所有代办事项查询
*DefaultApi* | [**DetachFromResource**](docs/DefaultApi.md#detachfromresource) | **Post** /?action&#x3D;DetachFromResource | 授权收回，只收回参数指定的授权，已有其它授权不受影响
*DefaultApi* | [**DetachResourceScopeOnPolicy**](docs/DefaultApi.md#detachresourcescopeonpolicy) | **Post** /?action&#x3D;DetachResourceScopeOnPolicy | 权限组可用资源删除，只删除参数指定的资源，已有的其它资源不受影响
*DefaultApi* | [**ForbiddenArtifactVersion**](docs/DefaultApi.md#forbiddenartifactversion) | **Post** /?action&#x3D;ForbiddenArtifactVersion | 制品版本下载 禁止、解禁
*DefaultApi* | [**ModifyArtifactCredit**](docs/DefaultApi.md#modifyartifactcredit) | **Post** /?action&#x3D;ModifyArtifactCredit | 制品授信清单修改
*DefaultApi* | [**ModifyArtifactProperties**](docs/DefaultApi.md#modifyartifactproperties) | **Post** /?action&#x3D;ModifyArtifactProperties | 制品属性修改
*DefaultApi* | [**ModifyBranchProtection**](docs/DefaultApi.md#modifybranchprotection) | **Post** /ModifyBranchProtection | 仓库设置-修改保护分支规则相关信息
*DefaultApi* | [**ModifyBranchProtectionMemberPermission**](docs/DefaultApi.md#modifybranchprotectionmemberpermission) | **Post** /ModifyBranchProtectionMemberPermission | 仓库设置-更改保护分支管理员权限
*DefaultApi* | [**ModifyCdCloudAccount**](docs/DefaultApi.md#modifycdcloudaccount) | **Post** /?action&#x3D;ModifyCdCloudAccount | CD 云账号更新
*DefaultApi* | [**ModifyCdHostServerGroup**](docs/DefaultApi.md#modifycdhostservergroup) | **Post** /?action&#x3D;ModifyCdHostServerGroup | CD 主机组修改
*DefaultApi* | [**ModifyCdPipeline**](docs/DefaultApi.md#modifycdpipeline) | **Post** /?action&#x3D;ModifyCdPipeline | CD 部署流程修改
*DefaultApi* | [**ModifyChooseDepotSpec**](docs/DefaultApi.md#modifychoosedepotspec) | **Post** /ModifyChooseDepotSpec | 仓库设置-使用、取消使用仓库规范
*DefaultApi* | [**ModifyCloseMR**](docs/DefaultApi.md#modifyclosemr) | **Post** /ModifyCloseMR | 合并请求-关闭合并请求
*DefaultApi* | [**ModifyCodingCIAgentEnable**](docs/DefaultApi.md#modifycodingciagentenable) | **Post** /?action&#x3D;ModifyCodingCIAgentEnable | 自定义构建节点启用、禁用
*DefaultApi* | [**ModifyCodingCIJob**](docs/DefaultApi.md#modifycodingcijob) | **Post** /?action&#x3D;ModifyCodingCIJob | 构建计划修改
*DefaultApi* | [**ModifyDefaultBranch**](docs/DefaultApi.md#modifydefaultbranch) | **Post** /ModifyDefaultBranch | 仓库设置-修改仓库默认分支
*DefaultApi* | [**ModifyDefectRelatedRequirement**](docs/DefaultApi.md#modifydefectrelatedrequirement) | **Post** /?action&#x3D;ModifyDefectRelatedRequirement | 缺陷所属的需求修改
*DefaultApi* | [**ModifyDepartment**](docs/DefaultApi.md#modifydepartment) | **Post** /?action&#x3D;ModifyDepartment | 部门信息修改
*DefaultApi* | [**ModifyDepartmentAssignee**](docs/DefaultApi.md#modifydepartmentassignee) | **Post** /?action&#x3D;ModifyDepartmentAssignee | 部门负责人管理
*DefaultApi* | [**ModifyDepartmentMember**](docs/DefaultApi.md#modifydepartmentmember) | **Post** /?action&#x3D;ModifyDepartmentMember | 部门成员管理
*DefaultApi* | [**ModifyDepotDescription**](docs/DefaultApi.md#modifydepotdescription) | **Post** /ModifyDepotDescription | 仓库信息-修改仓库描述
*DefaultApi* | [**ModifyDepotFilePushRule**](docs/DefaultApi.md#modifydepotfilepushrule) | **Post** /ModifyDepotFilePushRule | 仓库设置-修改git仓库文件推送规则
*DefaultApi* | [**ModifyDepotFilePushRuleDenyPrivilege**](docs/DefaultApi.md#modifydepotfilepushruledenyprivilege) | **Post** /ModifyDepotFilePushRuleDenyPrivilege | 仓库设置-修改 git 仓库特权者文件推送权限
*DefaultApi* | [**ModifyDepotLevelDepotSpec**](docs/DefaultApi.md#modifydepotleveldepotspec) | **Post** /ModifyDepotLevelDepotSpec | 仓库设置-修改、新增仓库级别的仓库规范
*DefaultApi* | [**ModifyDepotName**](docs/DefaultApi.md#modifydepotname) | **Post** /ModifyDepotName | 仓库信息-修改仓库名称
*DefaultApi* | [**ModifyDepotPushSetting**](docs/DefaultApi.md#modifydepotpushsetting) | **Post** /ModifyDepotPushSetting | 仓库设置-修改仓库推送设置
*DefaultApi* | [**ModifyDepotQuota**](docs/DefaultApi.md#modifydepotquota) | **Post** /ModifyDepotQuota | 仓库信息-修改仓库容量
*DefaultApi* | [**ModifyDepotSettings**](docs/DefaultApi.md#modifydepotsettings) | **Post** /ModifyDepotSettings | 仓库设置-修改仓库设置
*DefaultApi* | [**ModifyDepotSharedSetting**](docs/DefaultApi.md#modifydepotsharedsetting) | **Post** /ModifyDepotSharedSetting | 仓库设置-修改仓库是否开源设置
*DefaultApi* | [**ModifyGitCherryPick**](docs/DefaultApi.md#modifygitcherrypick) | **Post** /ModifyGitCherryPick | Git提交-将某次提交cherry-pick到指定分支
*DefaultApi* | [**ModifyGitCommitRevert**](docs/DefaultApi.md#modifygitcommitrevert) | **Post** /ModifyGitCommitRevert | Git提交-还原某次提交
*DefaultApi* | [**ModifyGitCommitStatus**](docs/DefaultApi.md#modifygitcommitstatus) | **Post** /ModifyGitCommitStatus | Git提交-修改提交对应的流水线状态
*DefaultApi* | [**ModifyGitDepotArchive**](docs/DefaultApi.md#modifygitdepotarchive) | **Post** /ModifyGitDepotArchive | 仓库设置-仓库归档
*DefaultApi* | [**ModifyGitDepotUnarchive**](docs/DefaultApi.md#modifygitdepotunarchive) | **Post** /ModifyGitDepotUnarchive | 仓库设置-仓库解除归档
*DefaultApi* | [**ModifyGitFiles**](docs/DefaultApi.md#modifygitfiles) | **Post** /ModifyGitFiles | Git提交-修改仓库某文件
*DefaultApi* | [**ModifyGitMergeBranch**](docs/DefaultApi.md#modifygitmergebranch) | **Post** /ModifyGitMergeBranch | 合并请求-将源分支的改动合并到目标分支
*DefaultApi* | [**ModifyGitMergeRequest**](docs/DefaultApi.md#modifygitmergerequest) | **Post** /ModifyGitMergeRequest | 合并请求-修改合并请求的标题和描述信息
*DefaultApi* | [**ModifyGitMergeRequestRebase**](docs/DefaultApi.md#modifygitmergerequestrebase) | **Post** /ModifyGitMergeRequestRebase | 合并请求-合并请求中的源分支进行rebase操作
*DefaultApi* | [**ModifyGitRebase**](docs/DefaultApi.md#modifygitrebase) | **Post** /ModifyGitRebase | 仓库信息-git变基操作
*DefaultApi* | [**ModifyGitRelease**](docs/DefaultApi.md#modifygitrelease) | **Post** /ModifyGitRelease | 版本信息-修改仓库版本信息
*DefaultApi* | [**ModifyGitTransfer**](docs/DefaultApi.md#modifygittransfer) | **Post** /ModifyGitTransfer | 仓库信息-仓库转移至同团队下的其他项目中
*DefaultApi* | [**ModifyIssue**](docs/DefaultApi.md#modifyissue) | **Post** /ModifyIssue | 事项修改
*DefaultApi* | [**ModifyIssueComment**](docs/DefaultApi.md#modifyissuecomment) | **Post** /?action&#x3D;ModifyIssueComment | 事项评论修改
*DefaultApi* | [**ModifyIssueDescription**](docs/DefaultApi.md#modifyissuedescription) | **Post** /?action&#x3D;ModifyIssueDescription | 事项描述修改
*DefaultApi* | [**ModifyIssueParentRequirement**](docs/DefaultApi.md#modifyissueparentrequirement) | **Post** /?action&#x3D;ModifyIssueParentRequirement | 事项父需求修改
*DefaultApi* | [**ModifyIteration**](docs/DefaultApi.md#modifyiteration) | **Post** /?action&#x3D;ModifyIteration | 迭代修改
*DefaultApi* | [**ModifyMergeMR**](docs/DefaultApi.md#modifymergemr) | **Post** /ModifyMergeMR | 合并信息-执行合并
*DefaultApi* | [**ModifyMergeRequestBasicSettings**](docs/DefaultApi.md#modifymergerequestbasicsettings) | **Post** /ModifyMergeRequestBasicSettings | 仓库设置-修改合并请求基础设置
*DefaultApi* | [**ModifyMergeRequestMergeCommitMessageTemplate**](docs/DefaultApi.md#modifymergerequestmergecommitmessagetemplate) | **Post** /ModifyMergeRequestMergeCommitMessageTemplate | 仓库设置-修改合并请求合并提交消息模版
*DefaultApi* | [**ModifyMergeRequestSquashCommitMessageTemplate**](docs/DefaultApi.md#modifymergerequestsquashcommitmessagetemplate) | **Post** /ModifyMergeRequestSquashCommitMessageTemplate | 仓库设置-修改合并请求合并压缩提交消息模版
*DefaultApi* | [**ModifyPolicy**](docs/DefaultApi.md#modifypolicy) | **Post** /?action&#x3D;ModifyPolicy | 权限组修改
*DefaultApi* | [**ModifyProject**](docs/DefaultApi.md#modifyproject) | **Post** /?action&#x3D;ModifyProject | 项目信息修改
*DefaultApi* | [**ModifyProjectAnnouncement**](docs/DefaultApi.md#modifyprojectannouncement) | **Post** /?action&#x3D;ModifyProjectAnnouncement | 项目公告更新
*DefaultApi* | [**ModifyProjectLabel**](docs/DefaultApi.md#modifyprojectlabel) | **Post** /?action&#x3D;ModifyProjectLabel | 项目标签修改
*DefaultApi* | [**ModifyProjectPermission**](docs/DefaultApi.md#modifyprojectpermission) | **Post** /?action&#x3D;ModifyProjectPermission | 项目成员权限配置(老权限)
*DefaultApi* | [**ModifyRelease**](docs/DefaultApi.md#modifyrelease) | **Post** /ModifyRelease | 版本修改
*DefaultApi* | [**ModifyTeamLevelDepotSpec**](docs/DefaultApi.md#modifyteamleveldepotspec) | **Post** /ModifyTeamLevelDepotSpec | 仓库设置-修改、新增团队级别的仓库规范
*DefaultApi* | [**ModifyTeamMemberLocked**](docs/DefaultApi.md#modifyteammemberlocked) | **Post** /?action&#x3D;ModifyTeamMemberLocked | 团队成员锁定
*DefaultApi* | [**ModifyTeamMemberUnlocked**](docs/DefaultApi.md#modifyteammemberunlocked) | **Post** /?action&#x3D;ModifyTeamMemberUnlocked | 团队成员解锁
*DefaultApi* | [**ModifyTestCase**](docs/DefaultApi.md#modifytestcase) | **Post** /?action&#x3D;ModifyTestCase | 测试用例修改
*DefaultApi* | [**ModifyTestCaseSection**](docs/DefaultApi.md#modifytestcasesection) | **Post** /?action&#x3D;ModifyTestCaseSection | 测试用例分组修改
*DefaultApi* | [**ModifyTestRun**](docs/DefaultApi.md#modifytestrun) | **Post** /?action&#x3D;ModifyTestRun | 测试计划修改
*DefaultApi* | [**ModifyWorkItemSplitIssues**](docs/DefaultApi.md#modifyworkitemsplitissues) | **Post** /ModifyWorkItemSplitIssues | 项目集工作项分解&amp;取消分解到项目中的事项
*DefaultApi* | [**PlanIterationIssue**](docs/DefaultApi.md#planiterationissue) | **Post** /?action&#x3D;PlanIterationIssue | 迭代批量规划
*DefaultApi* | [**ReleaseArtifactVersion**](docs/DefaultApi.md#releaseartifactversion) | **Post** /?action&#x3D;ReleaseArtifactVersion | 制品版本发布
*DefaultApi* | [**ReorderCdPipelines**](docs/DefaultApi.md#reordercdpipelines) | **Post** /?action&#x3D;ReorderCdPipelines | 部署流程重排序
*DefaultApi* | [**SetGrantToResource**](docs/DefaultApi.md#setgranttoresource) | **Post** /?action&#x3D;SetGrantToResource | 授权设置，收回授权体在资源中的所有授权，重新设置为参数指定的授权
*DefaultApi* | [**SetPredicatePolicyOnResource**](docs/DefaultApi.md#setpredicatepolicyonresource) | **Post** /?action&#x3D;SetPredicatePolicyOnResource | 资源权限判定策略设置
*DefaultApi* | [**StopCodingCIBuild**](docs/DefaultApi.md#stopcodingcibuild) | **Post** /?action&#x3D;StopCodingCIBuild | 构建停止
*DefaultApi* | [**TriggerCdPipeline**](docs/DefaultApi.md#triggercdpipeline) | **Post** /?action&#x3D;TriggerCdPipeline | 部署流程触发
*DefaultApi* | [**TriggerCodingCIBuild**](docs/DefaultApi.md#triggercodingcibuild) | **Post** /?action&#x3D;TriggerCodingCIBuild | 构建触发
*DefaultApi* | [**TriggerCodingScan**](docs/DefaultApi.md#triggercodingscan) | **Post** /?action&#x3D;TriggerCodingScan | 代码扫描触发
*DefaultApi* | [**UpdateUserGroupById**](docs/DefaultApi.md#updateusergroupbyid) | **Post** /?action&#x3D;UpdateUserGroupById | 用户组更新
*APIAPI* | [**CreateAPIDoc**](docs/APIAPI.md#createapidoc) | **Post** /?action&#x3D;CreateAPIDoc | API 文档创建
*APIAPI* | [**DeleteAPIDoc**](docs/APIAPI.md#deleteapidoc) | **Post** /?action&#x3D;DeleteAPIDoc | API 文档删除
*APIAPI* | [**DescribeAPIDoc**](docs/APIAPI.md#describeapidoc) | **Post** /?action&#x3D;DescribeAPIDoc | API 文档详情获取
*APIAPI* | [**DescribeAPIDocList**](docs/APIAPI.md#describeapidoclist) | **Post** /?action&#x3D;DescribeAPIDocList | API 文档列表查询
*APIAPI* | [**ModifyAPIDocBaseInfo**](docs/APIAPI.md#modifyapidocbaseinfo) | **Post** /?action&#x3D;ModifyAPIDocBaseInfo | API 文档基础信息修改
*APIAPI* | [**ModifyAPIDocContent**](docs/APIAPI.md#modifyapidoccontent) | **Post** /?action&#x3D;ModifyAPIDocContent | API 文档内容发布
*ServiceHookAPI* | [**CreateServiceHook**](docs/ServiceHookAPI.md#createservicehook) | **Post** /?action&#x3D;CreateServiceHook | Service Hook 创建
*ServiceHookAPI* | [**DeleteServiceHook**](docs/ServiceHookAPI.md#deleteservicehook) | **Post** /?action&#x3D;DeleteServiceHook | Service Hook 删除
*ServiceHookAPI* | [**DescribeEvents**](docs/ServiceHookAPI.md#describeevents) | **Post** /?action&#x3D;DescribeEvents | Service Hook 事件列表查询
*ServiceHookAPI* | [**DescribeServiceHook**](docs/ServiceHookAPI.md#describeservicehook) | **Post** /?action&#x3D;DescribeServiceHook | Service Hook 查询单条
*ServiceHookAPI* | [**DescribeServiceHookLogs**](docs/ServiceHookAPI.md#describeservicehooklogs) | **Post** /?action&#x3D;DescribeServiceHookLogs | Service Hook 发送记录分页查询
*ServiceHookAPI* | [**DescribeServiceHooks**](docs/ServiceHookAPI.md#describeservicehooks) | **Post** /?action&#x3D;DescribeServiceHooks | Service Hook 列表分页查询
*ServiceHookAPI* | [**EnabledServiceHook**](docs/ServiceHookAPI.md#enabledservicehook) | **Post** /?action&#x3D;EnabledServiceHook | Service Hook 事件开关
*ServiceHookAPI* | [**ModifyServiceHook**](docs/ServiceHookAPI.md#modifyservicehook) | **Post** /?action&#x3D;ModifyServiceHook | Service Hook 修改
*ServiceHookAPI* | [**PingServiceHook**](docs/ServiceHookAPI.md#pingservicehook) | **Post** /?action&#x3D;PingServiceHook | Service Hook 测试
*ServiceHookAPI* | [**ResendServiceHookLog**](docs/ServiceHookAPI.md#resendservicehooklog) | **Post** /?action&#x3D;ResendServiceHookLog | Service Hook 重发
*WikiAPI* | [**CreateUploadToken**](docs/WikiAPI.md#createuploadtoken) | **Post** /?action&#x3D;CreateUploadToken | 上传文件的Token获取
*WikiAPI* | [**CreateWiki**](docs/WikiAPI.md#createwiki) | **Post** /?action&#x3D;CreateWiki | Wiki创建
*WikiAPI* | [**CreateWikiByZip**](docs/WikiAPI.md#createwikibyzip) | **Post** /?action&#x3D;CreateWikiByZip | Wiki 通过zip包上传
*WikiAPI* | [**DeleteWiki**](docs/WikiAPI.md#deletewiki) | **Post** /?action&#x3D;DeleteWiki | Wiki 移至回收站
*WikiAPI* | [**DescribeImportJobStatus**](docs/WikiAPI.md#describeimportjobstatus) | **Post** /?action&#x3D;DescribeImportJobStatus | zip包创建wiki的任务状态查询
*WikiAPI* | [**DescribeUpdateJobStatus**](docs/WikiAPI.md#describeupdatejobstatus) | **Post** /?action&#x3D;DescribeUpdateJobStatus | zip包更新wiki的任务状态查询
*WikiAPI* | [**DescribeWiki**](docs/WikiAPI.md#describewiki) | **Post** /?action&#x3D;DescribeWiki | Wiki 详情获取
*WikiAPI* | [**DescribeWikiList**](docs/WikiAPI.md#describewikilist) | **Post** /?action&#x3D;DescribeWikiList | Wiki 列表详情获取
*WikiAPI* | [**ModifyWiki**](docs/WikiAPI.md#modifywiki) | **Post** /?action&#x3D;ModifyWiki | Wiki 更新
*WikiAPI* | [**ModifyWikiByZip**](docs/WikiAPI.md#modifywikibyzip) | **Post** /?action&#x3D;ModifyWikiByZip | 通过zip包更新wiki
*WikiAPI* | [**ModifyWikiOrder**](docs/WikiAPI.md#modifywikiorder) | **Post** /?action&#x3D;ModifyWikiOrder | Wiki 父级修改
*WikiAPI* | [**ModifyWikiTitle**](docs/WikiAPI.md#modifywikititle) | **Post** /?action&#x3D;ModifyWikiTitle | Wiki 标题更新


## Documentation For Models

 - [APIDoc](docs/APIDoc.md)
 - [AgentMachine](docs/AgentMachine.md)
 - [AgentMachineSecret](docs/AgentMachineSecret.md)
 - [ApiFileFile](docs/ApiFileFile.md)
 - [ApiFileFolder](docs/ApiFileFolder.md)
 - [ApiFilePreSignUploadUrlData](docs/ApiFilePreSignUploadUrlData.md)
 - [ApiIssueIssueAttachmentPreSign](docs/ApiIssueIssueAttachmentPreSign.md)
 - [ApiUserMemberRef](docs/ApiUserMemberRef.md)
 - [ApiUserRole](docs/ApiUserRole.md)
 - [ApiUserUserData](docs/ApiUserUserData.md)
 - [ApiUserUserDepartmentMember](docs/ApiUserUserDepartmentMember.md)
 - [ArchiveTestRun200Response](docs/ArchiveTestRun200Response.md)
 - [ArchiveTestRunRequest](docs/ArchiveTestRunRequest.md)
 - [ArtifactChecksum](docs/ArtifactChecksum.md)
 - [ArtifactFilterRule](docs/ArtifactFilterRule.md)
 - [ArtifactFilterRuleDetail](docs/ArtifactFilterRuleDetail.md)
 - [ArtifactPackageBean](docs/ArtifactPackageBean.md)
 - [ArtifactPackagePageBean](docs/ArtifactPackagePageBean.md)
 - [ArtifactProperty](docs/ArtifactProperty.md)
 - [ArtifactPropertyBean](docs/ArtifactPropertyBean.md)
 - [ArtifactRepositoryBean](docs/ArtifactRepositoryBean.md)
 - [ArtifactRepositoryFile](docs/ArtifactRepositoryFile.md)
 - [ArtifactRepositoryFileListData](docs/ArtifactRepositoryFileListData.md)
 - [ArtifactRepositoryPageBean](docs/ArtifactRepositoryPageBean.md)
 - [ArtifactVersionBean](docs/ArtifactVersionBean.md)
 - [ArtifactVersionFileBean](docs/ArtifactVersionFileBean.md)
 - [ArtifactVersionPageBean](docs/ArtifactVersionPageBean.md)
 - [ArtifactsOpenApiArtifactCreditsData](docs/ArtifactsOpenApiArtifactCreditsData.md)
 - [ArtifactsOpenApiArtifactCreditsRangeData](docs/ArtifactsOpenApiArtifactCreditsRangeData.md)
 - [ArtifactsOpenApiArtifactCreditsRuleData](docs/ArtifactsOpenApiArtifactCreditsRuleData.md)
 - [ArtifactsOpenApiCreateArtifactCreditsRangeData](docs/ArtifactsOpenApiCreateArtifactCreditsRangeData.md)
 - [ArtifactsOpenApiCreditRemoteTeam](docs/ArtifactsOpenApiCreditRemoteTeam.md)
 - [ArtifactsOpenApiProjectData](docs/ArtifactsOpenApiProjectData.md)
 - [ArtifactsOpenApiRemoteRepoData](docs/ArtifactsOpenApiRemoteRepoData.md)
 - [ArtifactsOpenApiRemoteTeamData](docs/ArtifactsOpenApiRemoteTeamData.md)
 - [ArtifactsOpenApiRepositoryData](docs/ArtifactsOpenApiRepositoryData.md)
 - [AttachResourceScopeToPolicyRequest](docs/AttachResourceScopeToPolicyRequest.md)
 - [Attachment](docs/Attachment.md)
 - [AttachmentPrepare](docs/AttachmentPrepare.md)
 - [AttachmentPrepareData](docs/AttachmentPrepareData.md)
 - [BindCdApplicationToProjectRequest](docs/BindCdApplicationToProjectRequest.md)
 - [BlobDetail](docs/BlobDetail.md)
 - [BoundExternalDepotRequest](docs/BoundExternalDepotRequest.md)
 - [BranchProtection](docs/BranchProtection.md)
 - [BranchProtectionMember](docs/BranchProtectionMember.md)
 - [BuildDescribeCodingCIBuildArtifactsData](docs/BuildDescribeCodingCIBuildArtifactsData.md)
 - [BuildDescribeCodingCIBuildArtifactsResponseDataReports](docs/BuildDescribeCodingCIBuildArtifactsResponseDataReports.md)
 - [BuildDescribeCodingCIBuildHtmlReportsData](docs/BuildDescribeCodingCIBuildHtmlReportsData.md)
 - [BuildDescribeCodingCIBuildHtmlReportsResponseDataReports](docs/BuildDescribeCodingCIBuildHtmlReportsResponseDataReports.md)
 - [CIBuildTestResult](docs/CIBuildTestResult.md)
 - [CIJobCachePath](docs/CIJobCachePath.md)
 - [CIJobEnv](docs/CIJobEnv.md)
 - [CIJobSchedule](docs/CIJobSchedule.md)
 - [CancelCdPipelineRequest](docs/CancelCdPipelineRequest.md)
 - [Case](docs/Case.md)
 - [CaseData](docs/CaseData.md)
 - [CaseDataSpecial](docs/CaseDataSpecial.md)
 - [CaseSpecial](docs/CaseSpecial.md)
 - [CasesData](docs/CasesData.md)
 - [CdApplication](docs/CdApplication.md)
 - [CdDeployCount](docs/CdDeployCount.md)
 - [CdDeployCountDetail](docs/CdDeployCountDetail.md)
 - [CdDeployOpenApiHostServerInstance](docs/CdDeployOpenApiHostServerInstance.md)
 - [CdDeployTime](docs/CdDeployTime.md)
 - [CdDeployTimeDetail](docs/CdDeployTimeDetail.md)
 - [CdDeployTrend](docs/CdDeployTrend.md)
 - [CdDeployTrendDetail](docs/CdDeployTrendDetail.md)
 - [CdDeployTrendTotal](docs/CdDeployTrendTotal.md)
 - [ClearCodingCIJobCacheRequest](docs/ClearCodingCIJobCacheRequest.md)
 - [CloudAccount](docs/CloudAccount.md)
 - [CloudAccountCredential](docs/CloudAccountCredential.md)
 - [CodingCIBuild](docs/CodingCIBuild.md)
 - [CodingCIBuildStageData](docs/CodingCIBuildStageData.md)
 - [CodingCIBuildStepData](docs/CodingCIBuildStepData.md)
 - [CodingCIDepotDetail](docs/CodingCIDepotDetail.md)
 - [CodingCIJob](docs/CodingCIJob.md)
 - [CodingCIJobCachePath](docs/CodingCIJobCachePath.md)
 - [CodingCIJobEnv](docs/CodingCIJobEnv.md)
 - [CodingCIJobSchedule](docs/CodingCIJobSchedule.md)
 - [CodingCIPersonalExternalDepot](docs/CodingCIPersonalExternalDepot.md)
 - [CodingCIPersonalExternalDepotData](docs/CodingCIPersonalExternalDepotData.md)
 - [CodingCIProjectDepot](docs/CodingCIProjectDepot.md)
 - [CodingCITeamDepot](docs/CodingCITeamDepot.md)
 - [CodingUser](docs/CodingUser.md)
 - [Commit](docs/Commit.md)
 - [CommitComment](docs/CommitComment.md)
 - [CommitData](docs/CommitData.md)
 - [CommitFile](docs/CommitFile.md)
 - [CommitInfo](docs/CommitInfo.md)
 - [CommitNote](docs/CommitNote.md)
 - [CommitRef](docs/CommitRef.md)
 - [Committer](docs/Committer.md)
 - [ConfigTemplate](docs/ConfigTemplate.md)
 - [Contributor](docs/Contributor.md)
 - [CreateAPIDocRequest](docs/CreateAPIDocRequest.md)
 - [CreateArtifactCredit200Response](docs/CreateArtifactCredit200Response.md)
 - [CreateArtifactCredit200ResponseResponse](docs/CreateArtifactCredit200ResponseResponse.md)
 - [CreateArtifactCreditRequest](docs/CreateArtifactCreditRequest.md)
 - [CreateArtifactRepository200Response](docs/CreateArtifactRepository200Response.md)
 - [CreateArtifactRepository200ResponseResponse](docs/CreateArtifactRepository200ResponseResponse.md)
 - [CreateArtifactRepositoryRequest](docs/CreateArtifactRepositoryRequest.md)
 - [CreateAttachmentPrepareSignUrl200Response](docs/CreateAttachmentPrepareSignUrl200Response.md)
 - [CreateAttachmentPrepareSignUrl200ResponseResponse](docs/CreateAttachmentPrepareSignUrl200ResponseResponse.md)
 - [CreateAttachmentPrepareSignUrlRequest](docs/CreateAttachmentPrepareSignUrlRequest.md)
 - [CreateBinaryFiles200Response](docs/CreateBinaryFiles200Response.md)
 - [CreateBinaryFilesRequest](docs/CreateBinaryFilesRequest.md)
 - [CreateBranchProtectionMemberRequest](docs/CreateBranchProtectionMemberRequest.md)
 - [CreateBranchProtectionRequest](docs/CreateBranchProtectionRequest.md)
 - [CreateCaseResultRequest](docs/CreateCaseResultRequest.md)
 - [CreateCdCloudAccount200Response](docs/CreateCdCloudAccount200Response.md)
 - [CreateCdCloudAccount200ResponseResponse](docs/CreateCdCloudAccount200ResponseResponse.md)
 - [CreateCdCloudAccountRequest](docs/CreateCdCloudAccountRequest.md)
 - [CreateCdCloudAccountResponseData](docs/CreateCdCloudAccountResponseData.md)
 - [CreateCdHostServerGroup200Response](docs/CreateCdHostServerGroup200Response.md)
 - [CreateCdHostServerGroup200ResponseResponse](docs/CreateCdHostServerGroup200ResponseResponse.md)
 - [CreateCdHostServerGroupRequest](docs/CreateCdHostServerGroupRequest.md)
 - [CreateCdHostServerGroupResponseData](docs/CreateCdHostServerGroupResponseData.md)
 - [CreateCdPipeline200Response](docs/CreateCdPipeline200Response.md)
 - [CreateCdPipeline200ResponseResponse](docs/CreateCdPipeline200ResponseResponse.md)
 - [CreateCdPipelineRequest](docs/CreateCdPipelineRequest.md)
 - [CreateCdPipelineResponseData](docs/CreateCdPipelineResponseData.md)
 - [CreateCdTask200Response](docs/CreateCdTask200Response.md)
 - [CreateCdTask200ResponseResponse](docs/CreateCdTask200ResponseResponse.md)
 - [CreateCdTaskRequest](docs/CreateCdTaskRequest.md)
 - [CreateCdTaskResponseData](docs/CreateCdTaskResponseData.md)
 - [CreateCodingCIJob200Response](docs/CreateCodingCIJob200Response.md)
 - [CreateCodingCIJob200ResponseResponse](docs/CreateCodingCIJob200ResponseResponse.md)
 - [CreateCodingCIJobByTeamTemplate200Response](docs/CreateCodingCIJobByTeamTemplate200Response.md)
 - [CreateCodingCIJobByTeamTemplate200ResponseResponse](docs/CreateCodingCIJobByTeamTemplate200ResponseResponse.md)
 - [CreateCodingCIJobByTeamTemplateRequest](docs/CreateCodingCIJobByTeamTemplateRequest.md)
 - [CreateCodingCIJobData](docs/CreateCodingCIJobData.md)
 - [CreateCodingCIJobRequest](docs/CreateCodingCIJobRequest.md)
 - [CreateCodingProject200Response](docs/CreateCodingProject200Response.md)
 - [CreateCodingProjectRequest](docs/CreateCodingProjectRequest.md)
 - [CreateDepartment200Response](docs/CreateDepartment200Response.md)
 - [CreateDepartmentRequest](docs/CreateDepartmentRequest.md)
 - [CreateDepotByTemplate200Response](docs/CreateDepotByTemplate200Response.md)
 - [CreateDepotByTemplateRequest](docs/CreateDepotByTemplateRequest.md)
 - [CreateDepotFilePushRuleRequest](docs/CreateDepotFilePushRuleRequest.md)
 - [CreateFile200Response](docs/CreateFile200Response.md)
 - [CreateFile200ResponseResponse](docs/CreateFile200ResponseResponse.md)
 - [CreateFileRequest](docs/CreateFileRequest.md)
 - [CreateFolder200Response](docs/CreateFolder200Response.md)
 - [CreateFolder200ResponseResponse](docs/CreateFolder200ResponseResponse.md)
 - [CreateFolderRequest](docs/CreateFolderRequest.md)
 - [CreateGitBranchRequest](docs/CreateGitBranchRequest.md)
 - [CreateGitCommit200Response](docs/CreateGitCommit200Response.md)
 - [CreateGitCommitComment200Response](docs/CreateGitCommitComment200Response.md)
 - [CreateGitCommitComment200ResponseResponse](docs/CreateGitCommitComment200ResponseResponse.md)
 - [CreateGitCommitCommentRequest](docs/CreateGitCommitCommentRequest.md)
 - [CreateGitCommitNoteRequest](docs/CreateGitCommitNoteRequest.md)
 - [CreateGitCommitRequest](docs/CreateGitCommitRequest.md)
 - [CreateGitDeployKeyRequest](docs/CreateGitDeployKeyRequest.md)
 - [CreateGitDepot200Response](docs/CreateGitDepot200Response.md)
 - [CreateGitDepot200ResponseResponse](docs/CreateGitDepot200ResponseResponse.md)
 - [CreateGitDepotRequest](docs/CreateGitDepotRequest.md)
 - [CreateGitFiles200Response](docs/CreateGitFiles200Response.md)
 - [CreateGitFilesRequest](docs/CreateGitFilesRequest.md)
 - [CreateGitMergeReq200Response](docs/CreateGitMergeReq200Response.md)
 - [CreateGitMergeReq200ResponseResponse](docs/CreateGitMergeReq200ResponseResponse.md)
 - [CreateGitMergeReqRequest](docs/CreateGitMergeReqRequest.md)
 - [CreateGitMergeRequest200Response](docs/CreateGitMergeRequest200Response.md)
 - [CreateGitMergeRequest200ResponseResponse](docs/CreateGitMergeRequest200ResponseResponse.md)
 - [CreateGitMergeRequestRequest](docs/CreateGitMergeRequestRequest.md)
 - [CreateGitProtectedTagRuleRequest](docs/CreateGitProtectedTagRuleRequest.md)
 - [CreateGitProtectedTagRulesRequest](docs/CreateGitProtectedTagRulesRequest.md)
 - [CreateGitReleaseRequest](docs/CreateGitReleaseRequest.md)
 - [CreateGitTag200Response](docs/CreateGitTag200Response.md)
 - [CreateGitTag200ResponseResponse](docs/CreateGitTag200ResponseResponse.md)
 - [CreateGitTagRequest](docs/CreateGitTagRequest.md)
 - [CreateIssue200Response](docs/CreateIssue200Response.md)
 - [CreateIssueBlockRequest](docs/CreateIssueBlockRequest.md)
 - [CreateIssueCommentRequest](docs/CreateIssueCommentRequest.md)
 - [CreateIssueModule200Response](docs/CreateIssueModule200Response.md)
 - [CreateIssueModule200ResponseResponse](docs/CreateIssueModule200ResponseResponse.md)
 - [CreateIssueModuleRequest](docs/CreateIssueModuleRequest.md)
 - [CreateIssueRequest](docs/CreateIssueRequest.md)
 - [CreateIssueWorkHoursRequest](docs/CreateIssueWorkHoursRequest.md)
 - [CreateIteration200Response](docs/CreateIteration200Response.md)
 - [CreateIterationRequest](docs/CreateIterationRequest.md)
 - [CreateMemberSshKey200Response](docs/CreateMemberSshKey200Response.md)
 - [CreateMemberSshKey200ResponseResponse](docs/CreateMemberSshKey200ResponseResponse.md)
 - [CreateMemberSshKeyRequest](docs/CreateMemberSshKeyRequest.md)
 - [CreateMergeRequestNote200Response](docs/CreateMergeRequestNote200Response.md)
 - [CreateMergeRequestNote200ResponseResponse](docs/CreateMergeRequestNote200ResponseResponse.md)
 - [CreateMergeRequestNoteRequest](docs/CreateMergeRequestNoteRequest.md)
 - [CreateMergeRequestReviewer200Response](docs/CreateMergeRequestReviewer200Response.md)
 - [CreateMergeRequestReviewer200ResponseResponse](docs/CreateMergeRequestReviewer200ResponseResponse.md)
 - [CreateMergeRequestReviewerRequest](docs/CreateMergeRequestReviewerRequest.md)
 - [CreatePolicyRequest](docs/CreatePolicyRequest.md)
 - [CreateProgram200Response](docs/CreateProgram200Response.md)
 - [CreateProgram200ResponseResponse](docs/CreateProgram200ResponseResponse.md)
 - [CreateProgramMemberPolicy200Response](docs/CreateProgramMemberPolicy200Response.md)
 - [CreateProgramMemberPolicy200ResponseResponse](docs/CreateProgramMemberPolicy200ResponseResponse.md)
 - [CreateProgramMemberPolicyRequest](docs/CreateProgramMemberPolicyRequest.md)
 - [CreateProgramProjects200Response](docs/CreateProgramProjects200Response.md)
 - [CreateProgramProjects200ResponseResponse](docs/CreateProgramProjects200ResponseResponse.md)
 - [CreateProgramProjectsRequest](docs/CreateProgramProjectsRequest.md)
 - [CreateProgramRequest](docs/CreateProgramRequest.md)
 - [CreateProjectAnnouncement200Response](docs/CreateProjectAnnouncement200Response.md)
 - [CreateProjectAnnouncement200ResponseResponse](docs/CreateProjectAnnouncement200ResponseResponse.md)
 - [CreateProjectAnnouncementRequest](docs/CreateProjectAnnouncementRequest.md)
 - [CreateProjectLabelRequest](docs/CreateProjectLabelRequest.md)
 - [CreateProjectMemberPrincipalRequest](docs/CreateProjectMemberPrincipalRequest.md)
 - [CreateProjectWithTemplate200Response](docs/CreateProjectWithTemplate200Response.md)
 - [CreateProjectWithTemplate200ResponseResponse](docs/CreateProjectWithTemplate200ResponseResponse.md)
 - [CreateProjectWithTemplateRequest](docs/CreateProjectWithTemplateRequest.md)
 - [CreateReadOnlyRefRequest](docs/CreateReadOnlyRefRequest.md)
 - [CreateReleaseRequest](docs/CreateReleaseRequest.md)
 - [CreateReport200Response](docs/CreateReport200Response.md)
 - [CreateReport200ResponseResponse](docs/CreateReport200ResponseResponse.md)
 - [CreateReportRequest](docs/CreateReportRequest.md)
 - [CreateRequirementDefectRelationRequest](docs/CreateRequirementDefectRelationRequest.md)
 - [CreateServiceHook200Response](docs/CreateServiceHook200Response.md)
 - [CreateServiceHookRequest](docs/CreateServiceHookRequest.md)
 - [CreateSshKey200Response](docs/CreateSshKey200Response.md)
 - [CreateSshKeyRequest](docs/CreateSshKeyRequest.md)
 - [CreateTestCase200Response](docs/CreateTestCase200Response.md)
 - [CreateTestCase200ResponseResponse](docs/CreateTestCase200ResponseResponse.md)
 - [CreateTestCaseRequest](docs/CreateTestCaseRequest.md)
 - [CreateTestCaseSectionRequest](docs/CreateTestCaseSectionRequest.md)
 - [CreateTestDefectRequest](docs/CreateTestDefectRequest.md)
 - [CreateTestResultRequest](docs/CreateTestResultRequest.md)
 - [CreateTestResultsRequest](docs/CreateTestResultsRequest.md)
 - [CreateTestRunRequest](docs/CreateTestRunRequest.md)
 - [CreateTestStepResultRequest](docs/CreateTestStepResultRequest.md)
 - [CreateUploadToken200Response](docs/CreateUploadToken200Response.md)
 - [CreateUploadToken200ResponseResponse](docs/CreateUploadToken200ResponseResponse.md)
 - [CreateUploadTokenRequest](docs/CreateUploadTokenRequest.md)
 - [CreateUserGroup200Response](docs/CreateUserGroup200Response.md)
 - [CreateUserGroup200ResponseResponse](docs/CreateUserGroup200ResponseResponse.md)
 - [CreateUserGroupRequest](docs/CreateUserGroupRequest.md)
 - [CreateUserGroupUsersRequest](docs/CreateUserGroupUsersRequest.md)
 - [CreateWiki200Response](docs/CreateWiki200Response.md)
 - [CreateWikiByZip200Response](docs/CreateWikiByZip200Response.md)
 - [CreateWikiByZip200ResponseResponse](docs/CreateWikiByZip200ResponseResponse.md)
 - [CreateWikiByZipRequest](docs/CreateWikiByZipRequest.md)
 - [CreateWikiRequest](docs/CreateWikiRequest.md)
 - [Credential](docs/Credential.md)
 - [CurrentUser](docs/CurrentUser.md)
 - [CustomFieldChangeLog](docs/CustomFieldChangeLog.md)
 - [CustomFields](docs/CustomFields.md)
 - [CustomStep](docs/CustomStep.md)
 - [DefectType](docs/DefectType.md)
 - [DeleteAPIDoc200Response](docs/DeleteAPIDoc200Response.md)
 - [DeleteAPIDoc200ResponseResponse](docs/DeleteAPIDoc200ResponseResponse.md)
 - [DeleteAPIDocRequest](docs/DeleteAPIDocRequest.md)
 - [DeleteAllUsersOnGroupRequest](docs/DeleteAllUsersOnGroupRequest.md)
 - [DeleteArtifactPropertiesRequest](docs/DeleteArtifactPropertiesRequest.md)
 - [DeleteBranchProtectionMemberRequest](docs/DeleteBranchProtectionMemberRequest.md)
 - [DeleteBranchProtectionRequest](docs/DeleteBranchProtectionRequest.md)
 - [DeleteCdCloudAccountRequest](docs/DeleteCdCloudAccountRequest.md)
 - [DeleteCdHostServerGroupRequest](docs/DeleteCdHostServerGroupRequest.md)
 - [DeleteCdPipelineRequest](docs/DeleteCdPipelineRequest.md)
 - [DeleteCodingCIJobRequest](docs/DeleteCodingCIJobRequest.md)
 - [DeleteDepartmentRequest](docs/DeleteDepartmentRequest.md)
 - [DeleteDepotFilePushRule200Response](docs/DeleteDepotFilePushRule200Response.md)
 - [DeleteDepotFilePushRuleDenyPrivilege200Response](docs/DeleteDepotFilePushRuleDenyPrivilege200Response.md)
 - [DeleteDepotFilePushRuleDenyPrivilegeRequest](docs/DeleteDepotFilePushRuleDenyPrivilegeRequest.md)
 - [DeleteDepotFilePushRuleRequest](docs/DeleteDepotFilePushRuleRequest.md)
 - [DeleteGitBranchRequest](docs/DeleteGitBranchRequest.md)
 - [DeleteGitDeployKeyRequest](docs/DeleteGitDeployKeyRequest.md)
 - [DeleteGitDepotRequest](docs/DeleteGitDepotRequest.md)
 - [DeleteGitFiles200Response](docs/DeleteGitFiles200Response.md)
 - [DeleteGitFilesRequest](docs/DeleteGitFilesRequest.md)
 - [DeleteGitMergedBranchesRequest](docs/DeleteGitMergedBranchesRequest.md)
 - [DeleteGitProtectedTagRuleRequest](docs/DeleteGitProtectedTagRuleRequest.md)
 - [DeleteGitReleaseRequest](docs/DeleteGitReleaseRequest.md)
 - [DeleteGitTagRequest](docs/DeleteGitTagRequest.md)
 - [DeleteIssueBlockRequest](docs/DeleteIssueBlockRequest.md)
 - [DeleteIssueModuleRequest](docs/DeleteIssueModuleRequest.md)
 - [DeleteIssueRequest](docs/DeleteIssueRequest.md)
 - [DeleteIssueWorkHoursRequest](docs/DeleteIssueWorkHoursRequest.md)
 - [DeleteIterationRequest](docs/DeleteIterationRequest.md)
 - [DeleteMemberSshKeyRequest](docs/DeleteMemberSshKeyRequest.md)
 - [DeleteMergeRequestNote200Response](docs/DeleteMergeRequestNote200Response.md)
 - [DeleteMergeRequestNoteRequest](docs/DeleteMergeRequestNoteRequest.md)
 - [DeleteMergeRequestReviewer200Response](docs/DeleteMergeRequestReviewer200Response.md)
 - [DeleteMergeRequestReviewer200ResponseResponse](docs/DeleteMergeRequestReviewer200ResponseResponse.md)
 - [DeleteMergeRequestReviewerRequest](docs/DeleteMergeRequestReviewerRequest.md)
 - [DeleteOneProjectRequest](docs/DeleteOneProjectRequest.md)
 - [DeletePoliciesById200Response](docs/DeletePoliciesById200Response.md)
 - [DeletePoliciesById200ResponseResponse](docs/DeletePoliciesById200ResponseResponse.md)
 - [DeletePoliciesByIdRequest](docs/DeletePoliciesByIdRequest.md)
 - [DeleteProgramMemberPolicy200Response](docs/DeleteProgramMemberPolicy200Response.md)
 - [DeleteProgramMemberPolicy200ResponseResponse](docs/DeleteProgramMemberPolicy200ResponseResponse.md)
 - [DeleteProgramMemberPolicyRequest](docs/DeleteProgramMemberPolicyRequest.md)
 - [DeleteProjectAnnouncementRequest](docs/DeleteProjectAnnouncementRequest.md)
 - [DeleteProjectLabel200Response](docs/DeleteProjectLabel200Response.md)
 - [DeleteProjectLabelRequest](docs/DeleteProjectLabelRequest.md)
 - [DeleteProjectMemberPrincipalRequest](docs/DeleteProjectMemberPrincipalRequest.md)
 - [DeleteReleaseRequest](docs/DeleteReleaseRequest.md)
 - [DeleteReportRequest](docs/DeleteReportRequest.md)
 - [DeleteServiceHookRequest](docs/DeleteServiceHookRequest.md)
 - [DeleteSshKeyRequest](docs/DeleteSshKeyRequest.md)
 - [DeleteTeamLevelDepotSpecRequest](docs/DeleteTeamLevelDepotSpecRequest.md)
 - [DeleteTeamMemberRequest](docs/DeleteTeamMemberRequest.md)
 - [DeleteTestCaseRequest](docs/DeleteTestCaseRequest.md)
 - [DeleteTestCaseSectionRequest](docs/DeleteTestCaseSectionRequest.md)
 - [DeleteTestRunRequest](docs/DeleteTestRunRequest.md)
 - [DeleteUserGroupByIdsRequest](docs/DeleteUserGroupByIdsRequest.md)
 - [DeleteUserGroupUsersRequest](docs/DeleteUserGroupUsersRequest.md)
 - [DeleteWikiRequest](docs/DeleteWikiRequest.md)
 - [DepartmentDepartmentData](docs/DepartmentDepartmentData.md)
 - [DepartmentDepartmentMemberRef](docs/DepartmentDepartmentMemberRef.md)
 - [DepartmentDepartmentMembersData](docs/DepartmentDepartmentMembersData.md)
 - [DepartmentDepartmentTreeData](docs/DepartmentDepartmentTreeData.md)
 - [DepartmentDescribeDepartmentMemberPageList](docs/DepartmentDescribeDepartmentMemberPageList.md)
 - [DeployKeyInfo](docs/DeployKeyInfo.md)
 - [DepotBranchType](docs/DepotBranchType.md)
 - [DepotBranchTypeParam](docs/DepotBranchTypeParam.md)
 - [DepotData](docs/DepotData.md)
 - [DepotDetailData](docs/DepotDetailData.md)
 - [DepotInfo](docs/DepotInfo.md)
 - [DepotMergeRequestRule](docs/DepotMergeRequestRule.md)
 - [DepotMergeRequestRuleParam](docs/DepotMergeRequestRuleParam.md)
 - [DepotPushSetting](docs/DepotPushSetting.md)
 - [DepotSpec](docs/DepotSpec.md)
 - [DepotSpecDepotLevelParam](docs/DepotSpecDepotLevelParam.md)
 - [DepotSpecDetail](docs/DepotSpecDetail.md)
 - [DepotSpecTeamLevelParam](docs/DepotSpecTeamLevelParam.md)
 - [DepotUser](docs/DepotUser.md)
 - [DescribeAPIDocList200Response](docs/DescribeAPIDocList200Response.md)
 - [DescribeAPIDocList200ResponseResponse](docs/DescribeAPIDocList200ResponseResponse.md)
 - [DescribeAPIDocListRequest](docs/DescribeAPIDocListRequest.md)
 - [DescribeAPIDocRequest](docs/DescribeAPIDocRequest.md)
 - [DescribeAgentSecret200Response](docs/DescribeAgentSecret200Response.md)
 - [DescribeAgentSecret200ResponseResponse](docs/DescribeAgentSecret200ResponseResponse.md)
 - [DescribeAllMergeRequestNotes200Response](docs/DescribeAllMergeRequestNotes200Response.md)
 - [DescribeAllMergeRequestNotes200ResponseResponse](docs/DescribeAllMergeRequestNotes200ResponseResponse.md)
 - [DescribeAllMergeRequestNotesRequest](docs/DescribeAllMergeRequestNotesRequest.md)
 - [DescribeAllProjectLabels200Response](docs/DescribeAllProjectLabels200Response.md)
 - [DescribeAllProjectLabels200ResponseResponse](docs/DescribeAllProjectLabels200ResponseResponse.md)
 - [DescribeAllProjectLabelsRequest](docs/DescribeAllProjectLabelsRequest.md)
 - [DescribeAllProjectsIssueWorkLogList200Response](docs/DescribeAllProjectsIssueWorkLogList200Response.md)
 - [DescribeAllProjectsIssueWorkLogListRequest](docs/DescribeAllProjectsIssueWorkLogListRequest.md)
 - [DescribeArtifactChecksums200Response](docs/DescribeArtifactChecksums200Response.md)
 - [DescribeArtifactChecksums200ResponseResponse](docs/DescribeArtifactChecksums200ResponseResponse.md)
 - [DescribeArtifactChecksumsRequest](docs/DescribeArtifactChecksumsRequest.md)
 - [DescribeArtifactCredit200Response](docs/DescribeArtifactCredit200Response.md)
 - [DescribeArtifactCredit200ResponseResponse](docs/DescribeArtifactCredit200ResponseResponse.md)
 - [DescribeArtifactCreditList200Response](docs/DescribeArtifactCreditList200Response.md)
 - [DescribeArtifactCreditList200ResponseResponse](docs/DescribeArtifactCreditList200ResponseResponse.md)
 - [DescribeArtifactCreditRequest](docs/DescribeArtifactCreditRequest.md)
 - [DescribeArtifactFileDownloadUrl200Response](docs/DescribeArtifactFileDownloadUrl200Response.md)
 - [DescribeArtifactFileDownloadUrl200ResponseResponse](docs/DescribeArtifactFileDownloadUrl200ResponseResponse.md)
 - [DescribeArtifactFileDownloadUrlRequest](docs/DescribeArtifactFileDownloadUrlRequest.md)
 - [DescribeArtifactPackageList200Response](docs/DescribeArtifactPackageList200Response.md)
 - [DescribeArtifactPackageList200ResponseResponse](docs/DescribeArtifactPackageList200ResponseResponse.md)
 - [DescribeArtifactPackageListRequest](docs/DescribeArtifactPackageListRequest.md)
 - [DescribeArtifactProperties200Response](docs/DescribeArtifactProperties200Response.md)
 - [DescribeArtifactProperties200ResponseResponse](docs/DescribeArtifactProperties200ResponseResponse.md)
 - [DescribeArtifactRepositoryFileList200Response](docs/DescribeArtifactRepositoryFileList200Response.md)
 - [DescribeArtifactRepositoryFileList200ResponseResponse](docs/DescribeArtifactRepositoryFileList200ResponseResponse.md)
 - [DescribeArtifactRepositoryFileListRequest](docs/DescribeArtifactRepositoryFileListRequest.md)
 - [DescribeArtifactRepositoryList200Response](docs/DescribeArtifactRepositoryList200Response.md)
 - [DescribeArtifactRepositoryList200ResponseResponse](docs/DescribeArtifactRepositoryList200ResponseResponse.md)
 - [DescribeArtifactRepositoryListRequest](docs/DescribeArtifactRepositoryListRequest.md)
 - [DescribeArtifactVersionFileList200Response](docs/DescribeArtifactVersionFileList200Response.md)
 - [DescribeArtifactVersionFileList200ResponseResponse](docs/DescribeArtifactVersionFileList200ResponseResponse.md)
 - [DescribeArtifactVersionFileListRequest](docs/DescribeArtifactVersionFileListRequest.md)
 - [DescribeArtifactVersionList200Response](docs/DescribeArtifactVersionList200Response.md)
 - [DescribeArtifactVersionList200ResponseResponse](docs/DescribeArtifactVersionList200ResponseResponse.md)
 - [DescribeArtifactVersionListRequest](docs/DescribeArtifactVersionListRequest.md)
 - [DescribeAvailablePoliciesOnResource200Response](docs/DescribeAvailablePoliciesOnResource200Response.md)
 - [DescribeAvailablePoliciesOnResource200ResponseResponse](docs/DescribeAvailablePoliciesOnResource200ResponseResponse.md)
 - [DescribeAvailablePoliciesOnResourceRequest](docs/DescribeAvailablePoliciesOnResourceRequest.md)
 - [DescribeAvailablePoliciesOnResourceRequestFilter](docs/DescribeAvailablePoliciesOnResourceRequestFilter.md)
 - [DescribeAvailablePoliciesOnResponsePageData](docs/DescribeAvailablePoliciesOnResponsePageData.md)
 - [DescribeBlockIssueList200Response](docs/DescribeBlockIssueList200Response.md)
 - [DescribeBlockIssueList200ResponseResponse](docs/DescribeBlockIssueList200ResponseResponse.md)
 - [DescribeBlockIssueListRequest](docs/DescribeBlockIssueListRequest.md)
 - [DescribeBlockedByIssueList200Response](docs/DescribeBlockedByIssueList200Response.md)
 - [DescribeBlockedByIssueList200ResponseResponse](docs/DescribeBlockedByIssueList200ResponseResponse.md)
 - [DescribeBranchProtection200Response](docs/DescribeBranchProtection200Response.md)
 - [DescribeBranchProtectionMembers200Response](docs/DescribeBranchProtectionMembers200Response.md)
 - [DescribeBranchProtectionMembers200ResponseResponse](docs/DescribeBranchProtectionMembers200ResponseResponse.md)
 - [DescribeBranchProtectionMembersRequest](docs/DescribeBranchProtectionMembersRequest.md)
 - [DescribeBranchProtectionRequest](docs/DescribeBranchProtectionRequest.md)
 - [DescribeBranchProtections200Response](docs/DescribeBranchProtections200Response.md)
 - [DescribeBranchProtections200ResponseResponse](docs/DescribeBranchProtections200ResponseResponse.md)
 - [DescribeBranchProtectionsRequest](docs/DescribeBranchProtectionsRequest.md)
 - [DescribeCanMerge200Response](docs/DescribeCanMerge200Response.md)
 - [DescribeCanMerge200ResponseResponse](docs/DescribeCanMerge200ResponseResponse.md)
 - [DescribeCanMergeRequest](docs/DescribeCanMergeRequest.md)
 - [DescribeCdAgentMachines200Response](docs/DescribeCdAgentMachines200Response.md)
 - [DescribeCdAgentMachines200ResponseResponse](docs/DescribeCdAgentMachines200ResponseResponse.md)
 - [DescribeCdAgentMachinesRequest](docs/DescribeCdAgentMachinesRequest.md)
 - [DescribeCdAgentMachinesResponseData](docs/DescribeCdAgentMachinesResponseData.md)
 - [DescribeCdApplication200Response](docs/DescribeCdApplication200Response.md)
 - [DescribeCdApplication200ResponseResponse](docs/DescribeCdApplication200ResponseResponse.md)
 - [DescribeCdApplicationRequest](docs/DescribeCdApplicationRequest.md)
 - [DescribeCdApplicationResponseData](docs/DescribeCdApplicationResponseData.md)
 - [DescribeCdApplications200Response](docs/DescribeCdApplications200Response.md)
 - [DescribeCdApplications200ResponseResponse](docs/DescribeCdApplications200ResponseResponse.md)
 - [DescribeCdApplicationsByProject200Response](docs/DescribeCdApplicationsByProject200Response.md)
 - [DescribeCdApplicationsByProject200ResponseResponse](docs/DescribeCdApplicationsByProject200ResponseResponse.md)
 - [DescribeCdApplicationsByProjectResponseData](docs/DescribeCdApplicationsByProjectResponseData.md)
 - [DescribeCdApplicationsResponseData](docs/DescribeCdApplicationsResponseData.md)
 - [DescribeCdCloudAccounts200Response](docs/DescribeCdCloudAccounts200Response.md)
 - [DescribeCdCloudAccounts200ResponseResponse](docs/DescribeCdCloudAccounts200ResponseResponse.md)
 - [DescribeCdCloudAccountsRequest](docs/DescribeCdCloudAccountsRequest.md)
 - [DescribeCdCloudAccountsResponseData](docs/DescribeCdCloudAccountsResponseData.md)
 - [DescribeCdDeployCountByApplications200Response](docs/DescribeCdDeployCountByApplications200Response.md)
 - [DescribeCdDeployCountByApplications200ResponseResponse](docs/DescribeCdDeployCountByApplications200ResponseResponse.md)
 - [DescribeCdDeployCountByApplicationsRequest](docs/DescribeCdDeployCountByApplicationsRequest.md)
 - [DescribeCdDeployCountByApplicationsResponseData](docs/DescribeCdDeployCountByApplicationsResponseData.md)
 - [DescribeCdDeployCountByProject200Response](docs/DescribeCdDeployCountByProject200Response.md)
 - [DescribeCdDeployCountByProject200ResponseResponse](docs/DescribeCdDeployCountByProject200ResponseResponse.md)
 - [DescribeCdDeployCountByProjectResponseData](docs/DescribeCdDeployCountByProjectResponseData.md)
 - [DescribeCdDeployTimeByApplications200Response](docs/DescribeCdDeployTimeByApplications200Response.md)
 - [DescribeCdDeployTimeByApplications200ResponseResponse](docs/DescribeCdDeployTimeByApplications200ResponseResponse.md)
 - [DescribeCdDeployTimeByApplicationsRequest](docs/DescribeCdDeployTimeByApplicationsRequest.md)
 - [DescribeCdDeployTimeByApplicationsResponseData](docs/DescribeCdDeployTimeByApplicationsResponseData.md)
 - [DescribeCdDeployTimeByProject200Response](docs/DescribeCdDeployTimeByProject200Response.md)
 - [DescribeCdDeployTimeByProject200ResponseResponse](docs/DescribeCdDeployTimeByProject200ResponseResponse.md)
 - [DescribeCdDeployTimeByProjectRequest](docs/DescribeCdDeployTimeByProjectRequest.md)
 - [DescribeCdDeployTimeByProjectResponseData](docs/DescribeCdDeployTimeByProjectResponseData.md)
 - [DescribeCdDeployTrendByApplications200Response](docs/DescribeCdDeployTrendByApplications200Response.md)
 - [DescribeCdDeployTrendByApplications200ResponseResponse](docs/DescribeCdDeployTrendByApplications200ResponseResponse.md)
 - [DescribeCdDeployTrendByApplicationsRequest](docs/DescribeCdDeployTrendByApplicationsRequest.md)
 - [DescribeCdDeployTrendByApplicationsResponseData](docs/DescribeCdDeployTrendByApplicationsResponseData.md)
 - [DescribeCdDeployTrendByProject200Response](docs/DescribeCdDeployTrendByProject200Response.md)
 - [DescribeCdDeployTrendByProject200ResponseResponse](docs/DescribeCdDeployTrendByProject200ResponseResponse.md)
 - [DescribeCdDeployTrendByProjectRequest](docs/DescribeCdDeployTrendByProjectRequest.md)
 - [DescribeCdDeployTrendByProjectResponseData](docs/DescribeCdDeployTrendByProjectResponseData.md)
 - [DescribeCdHostServerGroup200Response](docs/DescribeCdHostServerGroup200Response.md)
 - [DescribeCdHostServerGroup200ResponseResponse](docs/DescribeCdHostServerGroup200ResponseResponse.md)
 - [DescribeCdHostServerGroupRequest](docs/DescribeCdHostServerGroupRequest.md)
 - [DescribeCdHostServerGroupResponseData](docs/DescribeCdHostServerGroupResponseData.md)
 - [DescribeCdHostServerGroups200Response](docs/DescribeCdHostServerGroups200Response.md)
 - [DescribeCdHostServerGroups200ResponseResponse](docs/DescribeCdHostServerGroups200ResponseResponse.md)
 - [DescribeCdHostServerGroupsRequest](docs/DescribeCdHostServerGroupsRequest.md)
 - [DescribeCdHostServerGroupsResponseData](docs/DescribeCdHostServerGroupsResponseData.md)
 - [DescribeCdPipeline200Response](docs/DescribeCdPipeline200Response.md)
 - [DescribeCdPipeline200ResponseResponse](docs/DescribeCdPipeline200ResponseResponse.md)
 - [DescribeCdPipelineConfig200Response](docs/DescribeCdPipelineConfig200Response.md)
 - [DescribeCdPipelineConfig200ResponseResponse](docs/DescribeCdPipelineConfig200ResponseResponse.md)
 - [DescribeCdPipelineConfigRequest](docs/DescribeCdPipelineConfigRequest.md)
 - [DescribeCdPipelineConfigResponseData](docs/DescribeCdPipelineConfigResponseData.md)
 - [DescribeCdPipelineConfigs200Response](docs/DescribeCdPipelineConfigs200Response.md)
 - [DescribeCdPipelineConfigs200ResponseResponse](docs/DescribeCdPipelineConfigs200ResponseResponse.md)
 - [DescribeCdPipelineConfigsRequest](docs/DescribeCdPipelineConfigsRequest.md)
 - [DescribeCdPipelineConfigsResponseData](docs/DescribeCdPipelineConfigsResponseData.md)
 - [DescribeCdPipelineRequest](docs/DescribeCdPipelineRequest.md)
 - [DescribeCdPipelineResponseData](docs/DescribeCdPipelineResponseData.md)
 - [DescribeCdTask200Response](docs/DescribeCdTask200Response.md)
 - [DescribeCdTask200ResponseResponse](docs/DescribeCdTask200ResponseResponse.md)
 - [DescribeCdTaskRequest](docs/DescribeCdTaskRequest.md)
 - [DescribeCdTaskResponseData](docs/DescribeCdTaskResponseData.md)
 - [DescribeCodeSearch200Response](docs/DescribeCodeSearch200Response.md)
 - [DescribeCodeSearch200ResponseResponse](docs/DescribeCodeSearch200ResponseResponse.md)
 - [DescribeCodeSearchRequest](docs/DescribeCodeSearchRequest.md)
 - [DescribeCodingCIBuild200Response](docs/DescribeCodingCIBuild200Response.md)
 - [DescribeCodingCIBuild200ResponseResponse](docs/DescribeCodingCIBuild200ResponseResponse.md)
 - [DescribeCodingCIBuildArtifacts200Response](docs/DescribeCodingCIBuildArtifacts200Response.md)
 - [DescribeCodingCIBuildArtifacts200ResponseResponse](docs/DescribeCodingCIBuildArtifacts200ResponseResponse.md)
 - [DescribeCodingCIBuildArtifactsRequest](docs/DescribeCodingCIBuildArtifactsRequest.md)
 - [DescribeCodingCIBuildEnvs200Response](docs/DescribeCodingCIBuildEnvs200Response.md)
 - [DescribeCodingCIBuildEnvs200ResponseResponse](docs/DescribeCodingCIBuildEnvs200ResponseResponse.md)
 - [DescribeCodingCIBuildEnvsRequest](docs/DescribeCodingCIBuildEnvsRequest.md)
 - [DescribeCodingCIBuildHtmlReports200Response](docs/DescribeCodingCIBuildHtmlReports200Response.md)
 - [DescribeCodingCIBuildHtmlReports200ResponseResponse](docs/DescribeCodingCIBuildHtmlReports200ResponseResponse.md)
 - [DescribeCodingCIBuildLog200Response](docs/DescribeCodingCIBuildLog200Response.md)
 - [DescribeCodingCIBuildLog200ResponseResponse](docs/DescribeCodingCIBuildLog200ResponseResponse.md)
 - [DescribeCodingCIBuildLogData](docs/DescribeCodingCIBuildLogData.md)
 - [DescribeCodingCIBuildLogRaw200Response](docs/DescribeCodingCIBuildLogRaw200Response.md)
 - [DescribeCodingCIBuildLogRaw200ResponseResponse](docs/DescribeCodingCIBuildLogRaw200ResponseResponse.md)
 - [DescribeCodingCIBuildLogRawData](docs/DescribeCodingCIBuildLogRawData.md)
 - [DescribeCodingCIBuildLogRawRequest](docs/DescribeCodingCIBuildLogRawRequest.md)
 - [DescribeCodingCIBuildLogRequest](docs/DescribeCodingCIBuildLogRequest.md)
 - [DescribeCodingCIBuildMetrics](docs/DescribeCodingCIBuildMetrics.md)
 - [DescribeCodingCIBuildMetrics200Response](docs/DescribeCodingCIBuildMetrics200Response.md)
 - [DescribeCodingCIBuildMetrics200ResponseResponse](docs/DescribeCodingCIBuildMetrics200ResponseResponse.md)
 - [DescribeCodingCIBuildMetricsRequest](docs/DescribeCodingCIBuildMetricsRequest.md)
 - [DescribeCodingCIBuildStage200Response](docs/DescribeCodingCIBuildStage200Response.md)
 - [DescribeCodingCIBuildStage200ResponseResponse](docs/DescribeCodingCIBuildStage200ResponseResponse.md)
 - [DescribeCodingCIBuildStageRequest](docs/DescribeCodingCIBuildStageRequest.md)
 - [DescribeCodingCIBuildStatistics200Response](docs/DescribeCodingCIBuildStatistics200Response.md)
 - [DescribeCodingCIBuildStatistics200ResponseResponse](docs/DescribeCodingCIBuildStatistics200ResponseResponse.md)
 - [DescribeCodingCIBuildStatisticsRequest](docs/DescribeCodingCIBuildStatisticsRequest.md)
 - [DescribeCodingCIBuildStatisticsResponseData](docs/DescribeCodingCIBuildStatisticsResponseData.md)
 - [DescribeCodingCIBuildStep200Response](docs/DescribeCodingCIBuildStep200Response.md)
 - [DescribeCodingCIBuildStep200ResponseResponse](docs/DescribeCodingCIBuildStep200ResponseResponse.md)
 - [DescribeCodingCIBuildStepLog200Response](docs/DescribeCodingCIBuildStepLog200Response.md)
 - [DescribeCodingCIBuildStepLog200ResponseResponse](docs/DescribeCodingCIBuildStepLog200ResponseResponse.md)
 - [DescribeCodingCIBuildStepLogData](docs/DescribeCodingCIBuildStepLogData.md)
 - [DescribeCodingCIBuildStepLogRequest](docs/DescribeCodingCIBuildStepLogRequest.md)
 - [DescribeCodingCIBuildStepRequest](docs/DescribeCodingCIBuildStepRequest.md)
 - [DescribeCodingCIBuilds200Response](docs/DescribeCodingCIBuilds200Response.md)
 - [DescribeCodingCIBuilds200ResponseResponse](docs/DescribeCodingCIBuilds200ResponseResponse.md)
 - [DescribeCodingCIBuildsData](docs/DescribeCodingCIBuildsData.md)
 - [DescribeCodingCIBuildsRequest](docs/DescribeCodingCIBuildsRequest.md)
 - [DescribeCodingCIJob200Response](docs/DescribeCodingCIJob200Response.md)
 - [DescribeCodingCIJob200ResponseResponse](docs/DescribeCodingCIJob200ResponseResponse.md)
 - [DescribeCodingCIJobs200Response](docs/DescribeCodingCIJobs200Response.md)
 - [DescribeCodingCIJobs200ResponseResponse](docs/DescribeCodingCIJobs200ResponseResponse.md)
 - [DescribeCodingCIJobsRequest](docs/DescribeCodingCIJobsRequest.md)
 - [DescribeCodingCurrentUser200Response](docs/DescribeCodingCurrentUser200Response.md)
 - [DescribeCodingCurrentUser200ResponseResponse](docs/DescribeCodingCurrentUser200ResponseResponse.md)
 - [DescribeCodingProjects200Response](docs/DescribeCodingProjects200Response.md)
 - [DescribeCodingProjects200ResponseResponse](docs/DescribeCodingProjects200ResponseResponse.md)
 - [DescribeCodingProjectsRequest](docs/DescribeCodingProjectsRequest.md)
 - [DescribeCommitRefs200Response](docs/DescribeCommitRefs200Response.md)
 - [DescribeCommitRefs200ResponseResponse](docs/DescribeCommitRefs200ResponseResponse.md)
 - [DescribeCommitRefsRequest](docs/DescribeCommitRefsRequest.md)
 - [DescribeCommitsBetweenCommitAndCommit200Response](docs/DescribeCommitsBetweenCommitAndCommit200Response.md)
 - [DescribeCommitsBetweenCommitAndCommit200ResponseResponse](docs/DescribeCommitsBetweenCommitAndCommit200ResponseResponse.md)
 - [DescribeCommitsBetweenCommitAndCommitRequest](docs/DescribeCommitsBetweenCommitAndCommitRequest.md)
 - [DescribeConfigTemplateList200Response](docs/DescribeConfigTemplateList200Response.md)
 - [DescribeConfigTemplateList200ResponseResponse](docs/DescribeConfigTemplateList200ResponseResponse.md)
 - [DescribeConfigTemplateListRequest](docs/DescribeConfigTemplateListRequest.md)
 - [DescribeDepartment200Response](docs/DescribeDepartment200Response.md)
 - [DescribeDepartment200ResponseResponse](docs/DescribeDepartment200ResponseResponse.md)
 - [DescribeDepartmentMembers200Response](docs/DescribeDepartmentMembers200Response.md)
 - [DescribeDepartmentMembers200ResponseResponse](docs/DescribeDepartmentMembers200ResponseResponse.md)
 - [DescribeDepartmentMembersRequest](docs/DescribeDepartmentMembersRequest.md)
 - [DescribeDepartmentRequest](docs/DescribeDepartmentRequest.md)
 - [DescribeDepotByNameInfo200Response](docs/DescribeDepotByNameInfo200Response.md)
 - [DescribeDepotByNameInfo200ResponseResponse](docs/DescribeDepotByNameInfo200ResponseResponse.md)
 - [DescribeDepotByNameInfoRequest](docs/DescribeDepotByNameInfoRequest.md)
 - [DescribeDepotDefaultBranch200Response](docs/DescribeDepotDefaultBranch200Response.md)
 - [DescribeDepotDefaultBranch200ResponseResponse](docs/DescribeDepotDefaultBranch200ResponseResponse.md)
 - [DescribeDepotDefaultBranchRequest](docs/DescribeDepotDefaultBranchRequest.md)
 - [DescribeDepotFilePushRules200Response](docs/DescribeDepotFilePushRules200Response.md)
 - [DescribeDepotFilePushRulesRequest](docs/DescribeDepotFilePushRulesRequest.md)
 - [DescribeDepotMergeRequests200Response](docs/DescribeDepotMergeRequests200Response.md)
 - [DescribeDepotMergeRequests200ResponseResponse](docs/DescribeDepotMergeRequests200ResponseResponse.md)
 - [DescribeDepotMergeRequestsRequest](docs/DescribeDepotMergeRequestsRequest.md)
 - [DescribeDepotPushSetting200Response](docs/DescribeDepotPushSetting200Response.md)
 - [DescribeDepotSpecDetailRequest](docs/DescribeDepotSpecDetailRequest.md)
 - [DescribeDepotSpecs200Response](docs/DescribeDepotSpecs200Response.md)
 - [DescribeDepotSpecs200ResponseResponse](docs/DescribeDepotSpecs200ResponseResponse.md)
 - [DescribeDifferentBetween2Commits200Response](docs/DescribeDifferentBetween2Commits200Response.md)
 - [DescribeDifferentBetween2Commits200ResponseResponse](docs/DescribeDifferentBetween2Commits200ResponseResponse.md)
 - [DescribeDifferentBetween2CommitsRequest](docs/DescribeDifferentBetween2CommitsRequest.md)
 - [DescribeDifferentBetweenTwoCommits200Response](docs/DescribeDifferentBetweenTwoCommits200Response.md)
 - [DescribeDifferentBetweenTwoCommits200ResponseResponse](docs/DescribeDifferentBetweenTwoCommits200ResponseResponse.md)
 - [DescribeDifferentBetweenTwoCommitsRequest](docs/DescribeDifferentBetweenTwoCommitsRequest.md)
 - [DescribeEvents200Response](docs/DescribeEvents200Response.md)
 - [DescribeEvents200ResponseResponse](docs/DescribeEvents200ResponseResponse.md)
 - [DescribeGitBlameInfo200Response](docs/DescribeGitBlameInfo200Response.md)
 - [DescribeGitBlameInfo200ResponseResponse](docs/DescribeGitBlameInfo200ResponseResponse.md)
 - [DescribeGitBlameInfoRequest](docs/DescribeGitBlameInfoRequest.md)
 - [DescribeGitBlob200Response](docs/DescribeGitBlob200Response.md)
 - [DescribeGitBlob200ResponseResponse](docs/DescribeGitBlob200ResponseResponse.md)
 - [DescribeGitBlobRaw200Response](docs/DescribeGitBlobRaw200Response.md)
 - [DescribeGitBlobRaw200ResponseResponse](docs/DescribeGitBlobRaw200ResponseResponse.md)
 - [DescribeGitBlobRawRequest](docs/DescribeGitBlobRawRequest.md)
 - [DescribeGitBlobRequest](docs/DescribeGitBlobRequest.md)
 - [DescribeGitBranch200Response](docs/DescribeGitBranch200Response.md)
 - [DescribeGitBranch200ResponseResponse](docs/DescribeGitBranch200ResponseResponse.md)
 - [DescribeGitBranchList200Response](docs/DescribeGitBranchList200Response.md)
 - [DescribeGitBranchList200ResponseResponse](docs/DescribeGitBranchList200ResponseResponse.md)
 - [DescribeGitBranchListRequest](docs/DescribeGitBranchListRequest.md)
 - [DescribeGitBranchRequest](docs/DescribeGitBranchRequest.md)
 - [DescribeGitBranches200Response](docs/DescribeGitBranches200Response.md)
 - [DescribeGitBranches200ResponseResponse](docs/DescribeGitBranches200ResponseResponse.md)
 - [DescribeGitBranchesBySha200Response](docs/DescribeGitBranchesBySha200Response.md)
 - [DescribeGitBranchesBySha200ResponseResponse](docs/DescribeGitBranchesBySha200ResponseResponse.md)
 - [DescribeGitBranchesByShaRequest](docs/DescribeGitBranchesByShaRequest.md)
 - [DescribeGitBranchesRequest](docs/DescribeGitBranchesRequest.md)
 - [DescribeGitCommitComments200Response](docs/DescribeGitCommitComments200Response.md)
 - [DescribeGitCommitComments200ResponseResponse](docs/DescribeGitCommitComments200ResponseResponse.md)
 - [DescribeGitCommitCommentsRequest](docs/DescribeGitCommitCommentsRequest.md)
 - [DescribeGitCommitDiff200Response](docs/DescribeGitCommitDiff200Response.md)
 - [DescribeGitCommitDiff200ResponseResponse](docs/DescribeGitCommitDiff200ResponseResponse.md)
 - [DescribeGitCommitDiffRequest](docs/DescribeGitCommitDiffRequest.md)
 - [DescribeGitCommitFilePathList200Response](docs/DescribeGitCommitFilePathList200Response.md)
 - [DescribeGitCommitFilePathList200ResponseResponse](docs/DescribeGitCommitFilePathList200ResponseResponse.md)
 - [DescribeGitCommitFilePathListRequest](docs/DescribeGitCommitFilePathListRequest.md)
 - [DescribeGitCommitInfo200Response](docs/DescribeGitCommitInfo200Response.md)
 - [DescribeGitCommitInfoRequest](docs/DescribeGitCommitInfoRequest.md)
 - [DescribeGitCommitInfos200Response](docs/DescribeGitCommitInfos200Response.md)
 - [DescribeGitCommitInfos200ResponseResponse](docs/DescribeGitCommitInfos200ResponseResponse.md)
 - [DescribeGitCommitNote200Response](docs/DescribeGitCommitNote200Response.md)
 - [DescribeGitCommitNote200ResponseResponse](docs/DescribeGitCommitNote200ResponseResponse.md)
 - [DescribeGitCommitNoteRequest](docs/DescribeGitCommitNoteRequest.md)
 - [DescribeGitCommitStatus200Response](docs/DescribeGitCommitStatus200Response.md)
 - [DescribeGitCommitStatus200ResponseResponse](docs/DescribeGitCommitStatus200ResponseResponse.md)
 - [DescribeGitCommitStatusRequest](docs/DescribeGitCommitStatusRequest.md)
 - [DescribeGitCommitsInPage200Response](docs/DescribeGitCommitsInPage200Response.md)
 - [DescribeGitCommitsInPage200ResponseResponse](docs/DescribeGitCommitsInPage200ResponseResponse.md)
 - [DescribeGitCommitsInPageRequest](docs/DescribeGitCommitsInPageRequest.md)
 - [DescribeGitContributors200Response](docs/DescribeGitContributors200Response.md)
 - [DescribeGitContributors200ResponseResponse](docs/DescribeGitContributors200ResponseResponse.md)
 - [DescribeGitContributorsRequest](docs/DescribeGitContributorsRequest.md)
 - [DescribeGitDepot200Response](docs/DescribeGitDepot200Response.md)
 - [DescribeGitDepotDeployKeys200Response](docs/DescribeGitDepotDeployKeys200Response.md)
 - [DescribeGitDepotDeployKeys200ResponseResponse](docs/DescribeGitDepotDeployKeys200ResponseResponse.md)
 - [DescribeGitDepotDeployKeysRequest](docs/DescribeGitDepotDeployKeysRequest.md)
 - [DescribeGitDepotRequest](docs/DescribeGitDepotRequest.md)
 - [DescribeGitFile200Response](docs/DescribeGitFile200Response.md)
 - [DescribeGitFile200ResponseResponse](docs/DescribeGitFile200ResponseResponse.md)
 - [DescribeGitFileContent200Response](docs/DescribeGitFileContent200Response.md)
 - [DescribeGitFileContent200ResponseResponse](docs/DescribeGitFileContent200ResponseResponse.md)
 - [DescribeGitFileContentRequest](docs/DescribeGitFileContentRequest.md)
 - [DescribeGitFileRequest](docs/DescribeGitFileRequest.md)
 - [DescribeGitFileStat200Response](docs/DescribeGitFileStat200Response.md)
 - [DescribeGitFileStat200ResponseResponse](docs/DescribeGitFileStat200ResponseResponse.md)
 - [DescribeGitFileStatPayload](docs/DescribeGitFileStatPayload.md)
 - [DescribeGitFileStatRequest](docs/DescribeGitFileStatRequest.md)
 - [DescribeGitFiles200Response](docs/DescribeGitFiles200Response.md)
 - [DescribeGitFiles200ResponseResponse](docs/DescribeGitFiles200ResponseResponse.md)
 - [DescribeGitFilesRequest](docs/DescribeGitFilesRequest.md)
 - [DescribeGitMergeBase200Response](docs/DescribeGitMergeBase200Response.md)
 - [DescribeGitMergeBase200ResponseResponse](docs/DescribeGitMergeBase200ResponseResponse.md)
 - [DescribeGitMergeBaseRequest](docs/DescribeGitMergeBaseRequest.md)
 - [DescribeGitMergeRequestDiffDetail200Response](docs/DescribeGitMergeRequestDiffDetail200Response.md)
 - [DescribeGitMergeRequestDiffDetail200ResponseResponse](docs/DescribeGitMergeRequestDiffDetail200ResponseResponse.md)
 - [DescribeGitMergeRequestDiffDetailRequest](docs/DescribeGitMergeRequestDiffDetailRequest.md)
 - [DescribeGitMergeRequestDiffs200Response](docs/DescribeGitMergeRequestDiffs200Response.md)
 - [DescribeGitMergeRequestDiffs200ResponseResponse](docs/DescribeGitMergeRequestDiffs200ResponseResponse.md)
 - [DescribeGitMergeRequestDiffsRequest](docs/DescribeGitMergeRequestDiffsRequest.md)
 - [DescribeGitMergeRequestParticipants200Response](docs/DescribeGitMergeRequestParticipants200Response.md)
 - [DescribeGitMergeRequestParticipants200ResponseResponse](docs/DescribeGitMergeRequestParticipants200ResponseResponse.md)
 - [DescribeGitMergeRequestParticipantsRequest](docs/DescribeGitMergeRequestParticipantsRequest.md)
 - [DescribeGitMergeRequestsBySha200Response](docs/DescribeGitMergeRequestsBySha200Response.md)
 - [DescribeGitMergeRequestsBySha200ResponseResponse](docs/DescribeGitMergeRequestsBySha200ResponseResponse.md)
 - [DescribeGitMergeRequestsByShaRequest](docs/DescribeGitMergeRequestsByShaRequest.md)
 - [DescribeGitProjectDeployKeysRequest](docs/DescribeGitProjectDeployKeysRequest.md)
 - [DescribeGitProtectedTags200Response](docs/DescribeGitProtectedTags200Response.md)
 - [DescribeGitProtectedTags200ResponseResponse](docs/DescribeGitProtectedTags200ResponseResponse.md)
 - [DescribeGitProtectedTagsByRule200Response](docs/DescribeGitProtectedTagsByRule200Response.md)
 - [DescribeGitProtectedTagsByRuleRequest](docs/DescribeGitProtectedTagsByRuleRequest.md)
 - [DescribeGitProtectedTagsRequest](docs/DescribeGitProtectedTagsRequest.md)
 - [DescribeGitRef200Response](docs/DescribeGitRef200Response.md)
 - [DescribeGitRef200ResponseResponse](docs/DescribeGitRef200ResponseResponse.md)
 - [DescribeGitRefRequest](docs/DescribeGitRefRequest.md)
 - [DescribeGitRefsBySha200Response](docs/DescribeGitRefsBySha200Response.md)
 - [DescribeGitRefsBySha200ResponseResponse](docs/DescribeGitRefsBySha200ResponseResponse.md)
 - [DescribeGitRefsByShaRequest](docs/DescribeGitRefsByShaRequest.md)
 - [DescribeGitReleaseDetail200Response](docs/DescribeGitReleaseDetail200Response.md)
 - [DescribeGitReleaseDetail200ResponseResponse](docs/DescribeGitReleaseDetail200ResponseResponse.md)
 - [DescribeGitReleaseDetailRequest](docs/DescribeGitReleaseDetailRequest.md)
 - [DescribeGitReleases200Response](docs/DescribeGitReleases200Response.md)
 - [DescribeGitReleases200ResponseResponse](docs/DescribeGitReleases200ResponseResponse.md)
 - [DescribeGitReleasesRequest](docs/DescribeGitReleasesRequest.md)
 - [DescribeGitTag200Response](docs/DescribeGitTag200Response.md)
 - [DescribeGitTagRequest](docs/DescribeGitTagRequest.md)
 - [DescribeGitTags200Response](docs/DescribeGitTags200Response.md)
 - [DescribeGitTags200ResponseResponse](docs/DescribeGitTags200ResponseResponse.md)
 - [DescribeGitTagsByBranch200Response](docs/DescribeGitTagsByBranch200Response.md)
 - [DescribeGitTagsByBranch200ResponseResponse](docs/DescribeGitTagsByBranch200ResponseResponse.md)
 - [DescribeGitTagsByBranchRequest](docs/DescribeGitTagsByBranchRequest.md)
 - [DescribeGitTagsBySha200Response](docs/DescribeGitTagsBySha200Response.md)
 - [DescribeGitTagsBySha200ResponseResponse](docs/DescribeGitTagsBySha200ResponseResponse.md)
 - [DescribeGitTagsByShaRequest](docs/DescribeGitTagsByShaRequest.md)
 - [DescribeGitTagsRequest](docs/DescribeGitTagsRequest.md)
 - [DescribeGitTree200Response](docs/DescribeGitTree200Response.md)
 - [DescribeGitTree200ResponseResponse](docs/DescribeGitTree200ResponseResponse.md)
 - [DescribeGitTreeRequest](docs/DescribeGitTreeRequest.md)
 - [DescribeGrantObjectsOnResource200Response](docs/DescribeGrantObjectsOnResource200Response.md)
 - [DescribeGrantObjectsOnResource200ResponseResponse](docs/DescribeGrantObjectsOnResource200ResponseResponse.md)
 - [DescribeGrantObjectsOnResourcePageData](docs/DescribeGrantObjectsOnResourcePageData.md)
 - [DescribeGrantObjectsOnResourceRequest](docs/DescribeGrantObjectsOnResourceRequest.md)
 - [DescribeGrantUsersOnResource200Response](docs/DescribeGrantUsersOnResource200Response.md)
 - [DescribeGrantUsersOnResource200ResponseResponse](docs/DescribeGrantUsersOnResource200ResponseResponse.md)
 - [DescribeGrantUsersOnResourcePageData](docs/DescribeGrantUsersOnResourcePageData.md)
 - [DescribeGrantUsersOnResourceRequest](docs/DescribeGrantUsersOnResourceRequest.md)
 - [DescribeHostServerInstance200Response](docs/DescribeHostServerInstance200Response.md)
 - [DescribeHostServerInstance200ResponseResponse](docs/DescribeHostServerInstance200ResponseResponse.md)
 - [DescribeHostServerInstanceRequest](docs/DescribeHostServerInstanceRequest.md)
 - [DescribeImportJobStatus200Response](docs/DescribeImportJobStatus200Response.md)
 - [DescribeImportJobStatus200ResponseResponse](docs/DescribeImportJobStatus200ResponseResponse.md)
 - [DescribeImportJobStatusRequest](docs/DescribeImportJobStatusRequest.md)
 - [DescribeIssue200Response](docs/DescribeIssue200Response.md)
 - [DescribeIssueAttachmentPreSignedUrl200Response](docs/DescribeIssueAttachmentPreSignedUrl200Response.md)
 - [DescribeIssueAttachmentPreSignedUrl200ResponseResponse](docs/DescribeIssueAttachmentPreSignedUrl200ResponseResponse.md)
 - [DescribeIssueAttachmentPreSignedUrlRequest](docs/DescribeIssueAttachmentPreSignedUrlRequest.md)
 - [DescribeIssueCommentList200Response](docs/DescribeIssueCommentList200Response.md)
 - [DescribeIssueCommentList200ResponseResponse](docs/DescribeIssueCommentList200ResponseResponse.md)
 - [DescribeIssueCommentListRequest](docs/DescribeIssueCommentListRequest.md)
 - [DescribeIssueCustomFieldLogList200Response](docs/DescribeIssueCustomFieldLogList200Response.md)
 - [DescribeIssueCustomFieldLogList200ResponseResponse](docs/DescribeIssueCustomFieldLogList200ResponseResponse.md)
 - [DescribeIssueCustomFieldLogListRequest](docs/DescribeIssueCustomFieldLogListRequest.md)
 - [DescribeIssueFileUrl200Response](docs/DescribeIssueFileUrl200Response.md)
 - [DescribeIssueFileUrl200ResponseResponse](docs/DescribeIssueFileUrl200ResponseResponse.md)
 - [DescribeIssueFileUrlRequest](docs/DescribeIssueFileUrlRequest.md)
 - [DescribeIssueFilterList200Response](docs/DescribeIssueFilterList200Response.md)
 - [DescribeIssueFilterList200ResponseResponse](docs/DescribeIssueFilterList200ResponseResponse.md)
 - [DescribeIssueFilterListRequest](docs/DescribeIssueFilterListRequest.md)
 - [DescribeIssueList200Response](docs/DescribeIssueList200Response.md)
 - [DescribeIssueList200ResponseResponse](docs/DescribeIssueList200ResponseResponse.md)
 - [DescribeIssueListRequest](docs/DescribeIssueListRequest.md)
 - [DescribeIssueListWithPage200Response](docs/DescribeIssueListWithPage200Response.md)
 - [DescribeIssueListWithPage200ResponseResponse](docs/DescribeIssueListWithPage200ResponseResponse.md)
 - [DescribeIssueListWithPageRequest](docs/DescribeIssueListWithPageRequest.md)
 - [DescribeIssueModuleList200Response](docs/DescribeIssueModuleList200Response.md)
 - [DescribeIssueModuleList200ResponseResponse](docs/DescribeIssueModuleList200ResponseResponse.md)
 - [DescribeIssueReferenceResources200Response](docs/DescribeIssueReferenceResources200Response.md)
 - [DescribeIssueReferenceResources200ResponseResponse](docs/DescribeIssueReferenceResources200ResponseResponse.md)
 - [DescribeIssueReferenceResourcesRequest](docs/DescribeIssueReferenceResourcesRequest.md)
 - [DescribeIssueRelatedWorkItemList200Response](docs/DescribeIssueRelatedWorkItemList200Response.md)
 - [DescribeIssueRelatedWorkItemList200ResponseResponse](docs/DescribeIssueRelatedWorkItemList200ResponseResponse.md)
 - [DescribeIssueRelatedWorkItemListRequest](docs/DescribeIssueRelatedWorkItemListRequest.md)
 - [DescribeIssueReleaseList200Response](docs/DescribeIssueReleaseList200Response.md)
 - [DescribeIssueReleaseList200ResponseResponse](docs/DescribeIssueReleaseList200ResponseResponse.md)
 - [DescribeIssueReleaseListRequest](docs/DescribeIssueReleaseListRequest.md)
 - [DescribeIssueRequest](docs/DescribeIssueRequest.md)
 - [DescribeIssueStatusChangeLogList200Response](docs/DescribeIssueStatusChangeLogList200Response.md)
 - [DescribeIssueStatusChangeLogList200ResponseResponse](docs/DescribeIssueStatusChangeLogList200ResponseResponse.md)
 - [DescribeIssueStatusChangeLogListRequest](docs/DescribeIssueStatusChangeLogListRequest.md)
 - [DescribeIssueWorkLogList200Response](docs/DescribeIssueWorkLogList200Response.md)
 - [DescribeIssueWorkLogList200ResponseResponse](docs/DescribeIssueWorkLogList200ResponseResponse.md)
 - [DescribeIssueWorkLogListRequest](docs/DescribeIssueWorkLogListRequest.md)
 - [DescribeIteration200Response](docs/DescribeIteration200Response.md)
 - [DescribeIterationList200Response](docs/DescribeIterationList200Response.md)
 - [DescribeIterationList200ResponseResponse](docs/DescribeIterationList200ResponseResponse.md)
 - [DescribeIterationListData](docs/DescribeIterationListData.md)
 - [DescribeIterationListRequest](docs/DescribeIterationListRequest.md)
 - [DescribeMemberSshKey200Response](docs/DescribeMemberSshKey200Response.md)
 - [DescribeMemberSshKey200ResponseResponse](docs/DescribeMemberSshKey200ResponseResponse.md)
 - [DescribeMemberSshKeyRequest](docs/DescribeMemberSshKeyRequest.md)
 - [DescribeMergeReqCommits200Response](docs/DescribeMergeReqCommits200Response.md)
 - [DescribeMergeReqCommits200ResponseResponse](docs/DescribeMergeReqCommits200ResponseResponse.md)
 - [DescribeMergeReqCommitsRequest](docs/DescribeMergeReqCommitsRequest.md)
 - [DescribeMergeReqInfo200Response](docs/DescribeMergeReqInfo200Response.md)
 - [DescribeMergeReqInfo200ResponseResponse](docs/DescribeMergeReqInfo200ResponseResponse.md)
 - [DescribeMergeReqInfoRequest](docs/DescribeMergeReqInfoRequest.md)
 - [DescribeMergeRequest200Response](docs/DescribeMergeRequest200Response.md)
 - [DescribeMergeRequestFileDiff200Response](docs/DescribeMergeRequestFileDiff200Response.md)
 - [DescribeMergeRequestFileDiff200ResponseResponse](docs/DescribeMergeRequestFileDiff200ResponseResponse.md)
 - [DescribeMergeRequestFileDiffRequest](docs/DescribeMergeRequestFileDiffRequest.md)
 - [DescribeMergeRequestLog200Response](docs/DescribeMergeRequestLog200Response.md)
 - [DescribeMergeRequestLog200ResponseResponse](docs/DescribeMergeRequestLog200ResponseResponse.md)
 - [DescribeMergeRequestRequest](docs/DescribeMergeRequestRequest.md)
 - [DescribeMergeRequestReviewers200Response](docs/DescribeMergeRequestReviewers200Response.md)
 - [DescribeMergeRequestReviewers200ResponseResponse](docs/DescribeMergeRequestReviewers200ResponseResponse.md)
 - [DescribeMergeRequestReviewersRequest](docs/DescribeMergeRequestReviewersRequest.md)
 - [DescribeMergeRequestsData](docs/DescribeMergeRequestsData.md)
 - [DescribeMyDepots200Response](docs/DescribeMyDepots200Response.md)
 - [DescribeMyDepots200ResponseResponse](docs/DescribeMyDepots200ResponseResponse.md)
 - [DescribeMyDepotsRequest](docs/DescribeMyDepotsRequest.md)
 - [DescribeNotesByCommits200Response](docs/DescribeNotesByCommits200Response.md)
 - [DescribeNotesByCommits200ResponseResponse](docs/DescribeNotesByCommits200ResponseResponse.md)
 - [DescribeNotesByCommitsRequest](docs/DescribeNotesByCommitsRequest.md)
 - [DescribeOneProject200Response](docs/DescribeOneProject200Response.md)
 - [DescribeOneProject200ResponseResponse](docs/DescribeOneProject200ResponseResponse.md)
 - [DescribeOneProjectRequest](docs/DescribeOneProjectRequest.md)
 - [DescribePersonalExternalDepots200Response](docs/DescribePersonalExternalDepots200Response.md)
 - [DescribePersonalExternalDepots200ResponseResponse](docs/DescribePersonalExternalDepots200ResponseResponse.md)
 - [DescribePersonalExternalDepotsRequest](docs/DescribePersonalExternalDepotsRequest.md)
 - [DescribePinyin200Response](docs/DescribePinyin200Response.md)
 - [DescribePinyin200ResponseResponse](docs/DescribePinyin200ResponseResponse.md)
 - [DescribePinyinRequest](docs/DescribePinyinRequest.md)
 - [DescribePoliciesOnResourceType200Response](docs/DescribePoliciesOnResourceType200Response.md)
 - [DescribePoliciesOnResourceType200ResponseResponse](docs/DescribePoliciesOnResourceType200ResponseResponse.md)
 - [DescribePoliciesOnResourceTypeRequest](docs/DescribePoliciesOnResourceTypeRequest.md)
 - [DescribePoliciesOnResourceTypeRequestFilter](docs/DescribePoliciesOnResourceTypeRequestFilter.md)
 - [DescribePoliciesOnResourceTypeResponsePageData](docs/DescribePoliciesOnResourceTypeResponsePageData.md)
 - [DescribePolicy200Response](docs/DescribePolicy200Response.md)
 - [DescribePolicy200ResponseResponse](docs/DescribePolicy200ResponseResponse.md)
 - [DescribePolicyRequest](docs/DescribePolicyRequest.md)
 - [DescribePreSignUploadUrl200Response](docs/DescribePreSignUploadUrl200Response.md)
 - [DescribePreSignUploadUrl200ResponseResponse](docs/DescribePreSignUploadUrl200ResponseResponse.md)
 - [DescribePreSignUploadUrlRequest](docs/DescribePreSignUploadUrlRequest.md)
 - [DescribePredicatePolicyOnResource200Response](docs/DescribePredicatePolicyOnResource200Response.md)
 - [DescribePredicatePolicyOnResource200ResponseResponse](docs/DescribePredicatePolicyOnResource200ResponseResponse.md)
 - [DescribePredicatePolicyOnResourceRequest](docs/DescribePredicatePolicyOnResourceRequest.md)
 - [DescribeProgramProjects200Response](docs/DescribeProgramProjects200Response.md)
 - [DescribeProgramProjects200ResponseResponse](docs/DescribeProgramProjects200ResponseResponse.md)
 - [DescribeProgramProjectsRequest](docs/DescribeProgramProjectsRequest.md)
 - [DescribePrograms200Response](docs/DescribePrograms200Response.md)
 - [DescribePrograms200ResponseResponse](docs/DescribePrograms200ResponseResponse.md)
 - [DescribeProgramsRequest](docs/DescribeProgramsRequest.md)
 - [DescribeProjectAnnouncementRequest](docs/DescribeProjectAnnouncementRequest.md)
 - [DescribeProjectAnnouncements200Response](docs/DescribeProjectAnnouncements200Response.md)
 - [DescribeProjectAnnouncements200ResponseResponse](docs/DescribeProjectAnnouncements200ResponseResponse.md)
 - [DescribeProjectAnnouncementsRequest](docs/DescribeProjectAnnouncementsRequest.md)
 - [DescribeProjectByNameRequest](docs/DescribeProjectByNameRequest.md)
 - [DescribeProjectCredentials200Response](docs/DescribeProjectCredentials200Response.md)
 - [DescribeProjectCredentials200ResponseResponse](docs/DescribeProjectCredentials200ResponseResponse.md)
 - [DescribeProjectCredentialsData](docs/DescribeProjectCredentialsData.md)
 - [DescribeProjectCredentialsRequest](docs/DescribeProjectCredentialsRequest.md)
 - [DescribeProjectDepotBranches200Response](docs/DescribeProjectDepotBranches200Response.md)
 - [DescribeProjectDepotBranches200ResponseResponse](docs/DescribeProjectDepotBranches200ResponseResponse.md)
 - [DescribeProjectDepotBranchesRequest](docs/DescribeProjectDepotBranchesRequest.md)
 - [DescribeProjectDepotCommitsRequest](docs/DescribeProjectDepotCommitsRequest.md)
 - [DescribeProjectDepotInfoList200Response](docs/DescribeProjectDepotInfoList200Response.md)
 - [DescribeProjectDepotInfoList200ResponseResponse](docs/DescribeProjectDepotInfoList200ResponseResponse.md)
 - [DescribeProjectDepotInfoListRequest](docs/DescribeProjectDepotInfoListRequest.md)
 - [DescribeProjectDepots200Response](docs/DescribeProjectDepots200Response.md)
 - [DescribeProjectDepots200ResponseResponse](docs/DescribeProjectDepots200ResponseResponse.md)
 - [DescribeProjectDepotsData](docs/DescribeProjectDepotsData.md)
 - [DescribeProjectIssueFieldList200Response](docs/DescribeProjectIssueFieldList200Response.md)
 - [DescribeProjectIssueFieldList200ResponseResponse](docs/DescribeProjectIssueFieldList200ResponseResponse.md)
 - [DescribeProjectIssueFieldListRequest](docs/DescribeProjectIssueFieldListRequest.md)
 - [DescribeProjectIssueStatusList200Response](docs/DescribeProjectIssueStatusList200Response.md)
 - [DescribeProjectIssueStatusList200ResponseResponse](docs/DescribeProjectIssueStatusList200ResponseResponse.md)
 - [DescribeProjectIssueTypeList200Response](docs/DescribeProjectIssueTypeList200Response.md)
 - [DescribeProjectIssueTypeList200ResponseResponse](docs/DescribeProjectIssueTypeList200ResponseResponse.md)
 - [DescribeProjectIssueTypeListRequest](docs/DescribeProjectIssueTypeListRequest.md)
 - [DescribeProjectLabels200Response](docs/DescribeProjectLabels200Response.md)
 - [DescribeProjectLabels200ResponseResponse](docs/DescribeProjectLabels200ResponseResponse.md)
 - [DescribeProjectLabelsRequest](docs/DescribeProjectLabelsRequest.md)
 - [DescribeProjectMemberPrincipals200Response](docs/DescribeProjectMemberPrincipals200Response.md)
 - [DescribeProjectMemberPrincipals200ResponseResponse](docs/DescribeProjectMemberPrincipals200ResponseResponse.md)
 - [DescribeProjectMemberPrincipalsRequest](docs/DescribeProjectMemberPrincipalsRequest.md)
 - [DescribeProjectMembers200Response](docs/DescribeProjectMembers200Response.md)
 - [DescribeProjectMembers200ResponseResponse](docs/DescribeProjectMembers200ResponseResponse.md)
 - [DescribeProjectMembersRequest](docs/DescribeProjectMembersRequest.md)
 - [DescribeProjectMergeRequestsRequest](docs/DescribeProjectMergeRequestsRequest.md)
 - [DescribeProjectRoles200Response](docs/DescribeProjectRoles200Response.md)
 - [DescribeProjectRoles200ResponseResponse](docs/DescribeProjectRoles200ResponseResponse.md)
 - [DescribeProjectRolesRequest](docs/DescribeProjectRolesRequest.md)
 - [DescribeProjectsByFeature200Response](docs/DescribeProjectsByFeature200Response.md)
 - [DescribeProjectsByFeature200ResponseResponse](docs/DescribeProjectsByFeature200ResponseResponse.md)
 - [DescribeProjectsByFeatureRequest](docs/DescribeProjectsByFeatureRequest.md)
 - [DescribeProtectedBranch200Response](docs/DescribeProtectedBranch200Response.md)
 - [DescribeProtectedBranch200ResponseResponse](docs/DescribeProtectedBranch200ResponseResponse.md)
 - [DescribeProtectedBranchMembers200Response](docs/DescribeProtectedBranchMembers200Response.md)
 - [DescribeProtectedBranchMembers200ResponseResponse](docs/DescribeProtectedBranchMembers200ResponseResponse.md)
 - [DescribeProtectedBranchMembersRequest](docs/DescribeProtectedBranchMembersRequest.md)
 - [DescribeProtectedBranchRequest](docs/DescribeProtectedBranchRequest.md)
 - [DescribeProtectedBranches200Response](docs/DescribeProtectedBranches200Response.md)
 - [DescribeProtectedBranches200ResponseResponse](docs/DescribeProtectedBranches200ResponseResponse.md)
 - [DescribeProtectedBranchesRequest](docs/DescribeProtectedBranchesRequest.md)
 - [DescribeRelatedCaseList200Response](docs/DescribeRelatedCaseList200Response.md)
 - [DescribeRelatedCaseList200ResponseResponse](docs/DescribeRelatedCaseList200ResponseResponse.md)
 - [DescribeRelatedCaseListRequest](docs/DescribeRelatedCaseListRequest.md)
 - [DescribeReleaseIssueList200Response](docs/DescribeReleaseIssueList200Response.md)
 - [DescribeReleaseIssueList200ResponseResponse](docs/DescribeReleaseIssueList200ResponseResponse.md)
 - [DescribeReleaseIssueListRequest](docs/DescribeReleaseIssueListRequest.md)
 - [DescribeReleaseList200Response](docs/DescribeReleaseList200Response.md)
 - [DescribeReleaseList200ResponseResponse](docs/DescribeReleaseList200ResponseResponse.md)
 - [DescribeReleaseListRequest](docs/DescribeReleaseListRequest.md)
 - [DescribeReleaseRequest](docs/DescribeReleaseRequest.md)
 - [DescribeReport200Response](docs/DescribeReport200Response.md)
 - [DescribeReport200ResponseResponse](docs/DescribeReport200ResponseResponse.md)
 - [DescribeReportList200Response](docs/DescribeReportList200Response.md)
 - [DescribeReportList200ResponseResponse](docs/DescribeReportList200ResponseResponse.md)
 - [DescribeReportListRequest](docs/DescribeReportListRequest.md)
 - [DescribeRequirementDefectRelation200Response](docs/DescribeRequirementDefectRelation200Response.md)
 - [DescribeRequirementDefectRelation200ResponseResponse](docs/DescribeRequirementDefectRelation200ResponseResponse.md)
 - [DescribeRequirementDefectRelationRequest](docs/DescribeRequirementDefectRelationRequest.md)
 - [DescribeRequirementTestCaseListRequest](docs/DescribeRequirementTestCaseListRequest.md)
 - [DescribeResourceReferences200Response](docs/DescribeResourceReferences200Response.md)
 - [DescribeResourceReferences200ResponseResponse](docs/DescribeResourceReferences200ResponseResponse.md)
 - [DescribeResourceReferencesCited200Response](docs/DescribeResourceReferencesCited200Response.md)
 - [DescribeResourceReferencesCited200ResponseResponse](docs/DescribeResourceReferencesCited200ResponseResponse.md)
 - [DescribeResourceReferencesCiting200Response](docs/DescribeResourceReferencesCiting200Response.md)
 - [DescribeResourceReferencesCiting200ResponseResponse](docs/DescribeResourceReferencesCiting200ResponseResponse.md)
 - [DescribeResourceReferencesCitingRequest](docs/DescribeResourceReferencesCitingRequest.md)
 - [DescribeResourceReferencesRequest](docs/DescribeResourceReferencesRequest.md)
 - [DescribeResourceScopeListOnPolicy200Response](docs/DescribeResourceScopeListOnPolicy200Response.md)
 - [DescribeResourceScopeListOnPolicy200ResponseResponse](docs/DescribeResourceScopeListOnPolicy200ResponseResponse.md)
 - [DescribeResourceScopeListOnPolicyRequest](docs/DescribeResourceScopeListOnPolicyRequest.md)
 - [DescribeResourceScopeListOnPolicyRequestFilter](docs/DescribeResourceScopeListOnPolicyRequestFilter.md)
 - [DescribeResourceScopeListOnPolicyResponseData](docs/DescribeResourceScopeListOnPolicyResponseData.md)
 - [DescribeSelfMergeRequests200Response](docs/DescribeSelfMergeRequests200Response.md)
 - [DescribeSelfMergeRequestsRequest](docs/DescribeSelfMergeRequestsRequest.md)
 - [DescribeServiceHook200Response](docs/DescribeServiceHook200Response.md)
 - [DescribeServiceHookLogs200Response](docs/DescribeServiceHookLogs200Response.md)
 - [DescribeServiceHookLogs200ResponseResponse](docs/DescribeServiceHookLogs200ResponseResponse.md)
 - [DescribeServiceHookLogsRequest](docs/DescribeServiceHookLogsRequest.md)
 - [DescribeServiceHookRequest](docs/DescribeServiceHookRequest.md)
 - [DescribeServiceHooks200Response](docs/DescribeServiceHooks200Response.md)
 - [DescribeServiceHooks200ResponseResponse](docs/DescribeServiceHooks200ResponseResponse.md)
 - [DescribeServiceHooksRequest](docs/DescribeServiceHooksRequest.md)
 - [DescribeSingeMergeRequestNotes200Response](docs/DescribeSingeMergeRequestNotes200Response.md)
 - [DescribeSingeMergeRequestNotes200ResponseResponse](docs/DescribeSingeMergeRequestNotes200ResponseResponse.md)
 - [DescribeSingeMergeRequestNotesRequest](docs/DescribeSingeMergeRequestNotesRequest.md)
 - [DescribeSshKey200Response](docs/DescribeSshKey200Response.md)
 - [DescribeSubIssueList200Response](docs/DescribeSubIssueList200Response.md)
 - [DescribeSubIssueList200ResponseResponse](docs/DescribeSubIssueList200ResponseResponse.md)
 - [DescribeSubIssueListRequest](docs/DescribeSubIssueListRequest.md)
 - [DescribeTeam200Response](docs/DescribeTeam200Response.md)
 - [DescribeTeam200ResponseResponse](docs/DescribeTeam200ResponseResponse.md)
 - [DescribeTeamAdminMembers200Response](docs/DescribeTeamAdminMembers200Response.md)
 - [DescribeTeamAdminMembers200ResponseResponse](docs/DescribeTeamAdminMembers200ResponseResponse.md)
 - [DescribeTeamAdminMembersRequest](docs/DescribeTeamAdminMembersRequest.md)
 - [DescribeTeamArtifacts200Response](docs/DescribeTeamArtifacts200Response.md)
 - [DescribeTeamArtifacts200ResponseResponse](docs/DescribeTeamArtifacts200ResponseResponse.md)
 - [DescribeTeamArtifactsRequest](docs/DescribeTeamArtifactsRequest.md)
 - [DescribeTeamDepotInfoList200Response](docs/DescribeTeamDepotInfoList200Response.md)
 - [DescribeTeamDepotInfoListRequest](docs/DescribeTeamDepotInfoListRequest.md)
 - [DescribeTeamDepots200Response](docs/DescribeTeamDepots200Response.md)
 - [DescribeTeamDepots200ResponseResponse](docs/DescribeTeamDepots200ResponseResponse.md)
 - [DescribeTeamDepotsData](docs/DescribeTeamDepotsData.md)
 - [DescribeTeamDepotsRequest](docs/DescribeTeamDepotsRequest.md)
 - [DescribeTeamIssueTypeList200Response](docs/DescribeTeamIssueTypeList200Response.md)
 - [DescribeTeamIssueTypeList200ResponseResponse](docs/DescribeTeamIssueTypeList200ResponseResponse.md)
 - [DescribeTeamMember200Response](docs/DescribeTeamMember200Response.md)
 - [DescribeTeamMemberByEmail200Response](docs/DescribeTeamMemberByEmail200Response.md)
 - [DescribeTeamMemberByEmail200ResponseResponse](docs/DescribeTeamMemberByEmail200ResponseResponse.md)
 - [DescribeTeamMemberByEmailRequest](docs/DescribeTeamMemberByEmailRequest.md)
 - [DescribeTeamMemberRequest](docs/DescribeTeamMemberRequest.md)
 - [DescribeTeamMembers200Response](docs/DescribeTeamMembers200Response.md)
 - [DescribeTeamMembers200ResponseResponse](docs/DescribeTeamMembers200ResponseResponse.md)
 - [DescribeTeamMembersRequest](docs/DescribeTeamMembersRequest.md)
 - [DescribeTest200Response](docs/DescribeTest200Response.md)
 - [DescribeTest200ResponseResponse](docs/DescribeTest200ResponseResponse.md)
 - [DescribeTestCaseList200Response](docs/DescribeTestCaseList200Response.md)
 - [DescribeTestCaseList200ResponseResponse](docs/DescribeTestCaseList200ResponseResponse.md)
 - [DescribeTestCaseListRequest](docs/DescribeTestCaseListRequest.md)
 - [DescribeTestCaseSectionList200Response](docs/DescribeTestCaseSectionList200Response.md)
 - [DescribeTestCaseSectionList200ResponseResponse](docs/DescribeTestCaseSectionList200ResponseResponse.md)
 - [DescribeTestCaseSectionListRequest](docs/DescribeTestCaseSectionListRequest.md)
 - [DescribeTestDefectList200Response](docs/DescribeTestDefectList200Response.md)
 - [DescribeTestDefectList200ResponseResponse](docs/DescribeTestDefectList200ResponseResponse.md)
 - [DescribeTestDefectListRequest](docs/DescribeTestDefectListRequest.md)
 - [DescribeTestList200Response](docs/DescribeTestList200Response.md)
 - [DescribeTestList200ResponseResponse](docs/DescribeTestList200ResponseResponse.md)
 - [DescribeTestListRequest](docs/DescribeTestListRequest.md)
 - [DescribeTestRequest](docs/DescribeTestRequest.md)
 - [DescribeTestRun200Response](docs/DescribeTestRun200Response.md)
 - [DescribeTestRunList200Response](docs/DescribeTestRunList200Response.md)
 - [DescribeTestRunList200ResponseResponse](docs/DescribeTestRunList200ResponseResponse.md)
 - [DescribeTestRunListRequest](docs/DescribeTestRunListRequest.md)
 - [DescribeUserGroups200Response](docs/DescribeUserGroups200Response.md)
 - [DescribeUserGroups200ResponseResponse](docs/DescribeUserGroups200ResponseResponse.md)
 - [DescribeUserGroupsRequest](docs/DescribeUserGroupsRequest.md)
 - [DescribeUserGroupsRequestFilter](docs/DescribeUserGroupsRequestFilter.md)
 - [DescribeUserGroupsResponsePageData](docs/DescribeUserGroupsResponsePageData.md)
 - [DescribeUserProjectsRequest](docs/DescribeUserProjectsRequest.md)
 - [DescribeUsersByGroupId200Response](docs/DescribeUsersByGroupId200Response.md)
 - [DescribeUsersByGroupId200ResponseResponse](docs/DescribeUsersByGroupId200ResponseResponse.md)
 - [DescribeUsersByGroupIdRequest](docs/DescribeUsersByGroupIdRequest.md)
 - [DescribeUsersByGroupIdResponsePageData](docs/DescribeUsersByGroupIdResponsePageData.md)
 - [DescribeUsersOnResourceAndGrantObject200Response](docs/DescribeUsersOnResourceAndGrantObject200Response.md)
 - [DescribeUsersOnResourceAndGrantObject200ResponseResponse](docs/DescribeUsersOnResourceAndGrantObject200ResponseResponse.md)
 - [DescribeUsersOnResourceAndGrantObjectGrantInfo](docs/DescribeUsersOnResourceAndGrantObjectGrantInfo.md)
 - [DescribeUsersOnResourceAndGrantObjectPageData](docs/DescribeUsersOnResourceAndGrantObjectPageData.md)
 - [DescribeUsersOnResourceAndGrantObjectRequest](docs/DescribeUsersOnResourceAndGrantObjectRequest.md)
 - [DescribeWikiList200Response](docs/DescribeWikiList200Response.md)
 - [DescribeWikiList200ResponseResponse](docs/DescribeWikiList200ResponseResponse.md)
 - [DescribeWikiRequest](docs/DescribeWikiRequest.md)
 - [DescribeWorkItemSalvage200Response](docs/DescribeWorkItemSalvage200Response.md)
 - [DescribeWorkItemSalvage200ResponseResponse](docs/DescribeWorkItemSalvage200ResponseResponse.md)
 - [DescribeWorkItemSalvageRequest](docs/DescribeWorkItemSalvageRequest.md)
 - [DescribeWorkbenchIssueList200Response](docs/DescribeWorkbenchIssueList200Response.md)
 - [DescribeWorkbenchIssueList200ResponseResponse](docs/DescribeWorkbenchIssueList200ResponseResponse.md)
 - [DescribeWorkbenchIssueListRequest](docs/DescribeWorkbenchIssueListRequest.md)
 - [DetachFromResourceRequest](docs/DetachFromResourceRequest.md)
 - [DetachResourceScopeOnPolicyRequest](docs/DetachResourceScopeOnPolicyRequest.md)
 - [DiffFileInfo](docs/DiffFileInfo.md)
 - [DifferentLine](docs/DifferentLine.md)
 - [DifferentOfCommit](docs/DifferentOfCommit.md)
 - [DifferentOfCommitDetail](docs/DifferentOfCommitDetail.md)
 - [EnabledServiceHook200Response](docs/EnabledServiceHook200Response.md)
 - [EnabledServiceHook200ResponseResponse](docs/EnabledServiceHook200ResponseResponse.md)
 - [EnabledServiceHookRequest](docs/EnabledServiceHookRequest.md)
 - [Epic](docs/Epic.md)
 - [FileDiff](docs/FileDiff.md)
 - [Filter](docs/Filter.md)
 - [ForbiddenArtifactVersionRequest](docs/ForbiddenArtifactVersionRequest.md)
 - [GitAllTagCommit](docs/GitAllTagCommit.md)
 - [GitBranch](docs/GitBranch.md)
 - [GitBranchInfo](docs/GitBranchInfo.md)
 - [GitBranchesData](docs/GitBranchesData.md)
 - [GitCommit](docs/GitCommit.md)
 - [GitCommitComment](docs/GitCommitComment.md)
 - [GitCommitFilePath](docs/GitCommitFilePath.md)
 - [GitDiff](docs/GitDiff.md)
 - [GitFile](docs/GitFile.md)
 - [GitFileContent](docs/GitFileContent.md)
 - [GitFileItem](docs/GitFileItem.md)
 - [GitFilePushRule](docs/GitFilePushRule.md)
 - [GitFilePushRulePrivilege](docs/GitFilePushRulePrivilege.md)
 - [GitFilePushRuleRole](docs/GitFilePushRuleRole.md)
 - [GitFilePushRuleUser](docs/GitFilePushRuleUser.md)
 - [GitMergeRequest](docs/GitMergeRequest.md)
 - [GitRef](docs/GitRef.md)
 - [GitReference](docs/GitReference.md)
 - [GitTag](docs/GitTag.md)
 - [GitTree](docs/GitTree.md)
 - [GitTreeItem](docs/GitTreeItem.md)
 - [GrantInfo](docs/GrantInfo.md)
 - [GrantObjectInfo](docs/GrantObjectInfo.md)
 - [GrepLineData](docs/GrepLineData.md)
 - [GrepLineInfo](docs/GrepLineInfo.md)
 - [HostServerGroup](docs/HostServerGroup.md)
 - [HostServerGroupDetail](docs/HostServerGroupDetail.md)
 - [HostServerGroupLabel](docs/HostServerGroupLabel.md)
 - [IssueComment](docs/IssueComment.md)
 - [IssueCondition](docs/IssueCondition.md)
 - [IssueCustomField](docs/IssueCustomField.md)
 - [IssueCustomFieldForm](docs/IssueCustomFieldForm.md)
 - [IssueDetail](docs/IssueDetail.md)
 - [IssueDetailListDataWithPage](docs/IssueDetailListDataWithPage.md)
 - [IssueField](docs/IssueField.md)
 - [IssueFieldOption](docs/IssueFieldOption.md)
 - [IssueFile](docs/IssueFile.md)
 - [IssueFilter](docs/IssueFilter.md)
 - [IssueFilterListData](docs/IssueFilterListData.md)
 - [IssueListData](docs/IssueListData.md)
 - [IssueModule](docs/IssueModule.md)
 - [IssueProjectLabel](docs/IssueProjectLabel.md)
 - [IssueProjectModule](docs/IssueProjectModule.md)
 - [IssueRecordHourForm](docs/IssueRecordHourForm.md)
 - [IssueSimpleData](docs/IssueSimpleData.md)
 - [IssueStatus](docs/IssueStatus.md)
 - [IssueStatusChangeLog](docs/IssueStatusChangeLog.md)
 - [IssueStatusChangeLogList](docs/IssueStatusChangeLogList.md)
 - [IssueThirdLink](docs/IssueThirdLink.md)
 - [IssueThirdLinkForm](docs/IssueThirdLinkForm.md)
 - [IssueTypeDetail](docs/IssueTypeDetail.md)
 - [IssueTypeDetailWithSplit](docs/IssueTypeDetailWithSplit.md)
 - [IssueWorkLog](docs/IssueWorkLog.md)
 - [Iteration](docs/Iteration.md)
 - [IterationSimple](docs/IterationSimple.md)
 - [JobCreateJobData](docs/JobCreateJobData.md)
 - [KubeConfigForm](docs/KubeConfigForm.md)
 - [Line](docs/Line.md)
 - [Maven](docs/Maven.md)
 - [MergeInfo](docs/MergeInfo.md)
 - [MergeReqInfo](docs/MergeReqInfo.md)
 - [MergeRequestData](docs/MergeRequestData.md)
 - [MergeRequestDetail](docs/MergeRequestDetail.md)
 - [MergeRequestDiff](docs/MergeRequestDiff.md)
 - [MergeRequestDiffFile](docs/MergeRequestDiffFile.md)
 - [MergeRequestDiffNoteForm](docs/MergeRequestDiffNoteForm.md)
 - [MergeRequestFileDiff](docs/MergeRequestFileDiff.md)
 - [MergeRequestInfo](docs/MergeRequestInfo.md)
 - [MergeRequestLog](docs/MergeRequestLog.md)
 - [MergeRequestNote](docs/MergeRequestNote.md)
 - [MergeRequestNoteList](docs/MergeRequestNoteList.md)
 - [Mission](docs/Mission.md)
 - [ModifyAPIDocBaseInfoRequest](docs/ModifyAPIDocBaseInfoRequest.md)
 - [ModifyAPIDocContent200Response](docs/ModifyAPIDocContent200Response.md)
 - [ModifyAPIDocContent200ResponseResponse](docs/ModifyAPIDocContent200ResponseResponse.md)
 - [ModifyAPIDocContentRequest](docs/ModifyAPIDocContentRequest.md)
 - [ModifyArtifactCredit200Response](docs/ModifyArtifactCredit200Response.md)
 - [ModifyArtifactCreditRequest](docs/ModifyArtifactCreditRequest.md)
 - [ModifyArtifactPropertiesRequest](docs/ModifyArtifactPropertiesRequest.md)
 - [ModifyBranchProtection200Response](docs/ModifyBranchProtection200Response.md)
 - [ModifyBranchProtection200ResponseResponse](docs/ModifyBranchProtection200ResponseResponse.md)
 - [ModifyBranchProtectionMemberPermissionRequest](docs/ModifyBranchProtectionMemberPermissionRequest.md)
 - [ModifyBranchProtectionRequest](docs/ModifyBranchProtectionRequest.md)
 - [ModifyCdCloudAccount200Response](docs/ModifyCdCloudAccount200Response.md)
 - [ModifyCdCloudAccount200ResponseResponse](docs/ModifyCdCloudAccount200ResponseResponse.md)
 - [ModifyCdCloudAccountRequest](docs/ModifyCdCloudAccountRequest.md)
 - [ModifyCdCloudAccountResponseData](docs/ModifyCdCloudAccountResponseData.md)
 - [ModifyCdHostServerGroup200Response](docs/ModifyCdHostServerGroup200Response.md)
 - [ModifyCdHostServerGroup200ResponseResponse](docs/ModifyCdHostServerGroup200ResponseResponse.md)
 - [ModifyCdHostServerGroupRequest](docs/ModifyCdHostServerGroupRequest.md)
 - [ModifyCdHostServerGroupResponseData](docs/ModifyCdHostServerGroupResponseData.md)
 - [ModifyCdPipeline200Response](docs/ModifyCdPipeline200Response.md)
 - [ModifyCdPipeline200ResponseResponse](docs/ModifyCdPipeline200ResponseResponse.md)
 - [ModifyCdPipelineRequest](docs/ModifyCdPipelineRequest.md)
 - [ModifyCdPipelineResponseData](docs/ModifyCdPipelineResponseData.md)
 - [ModifyChooseDepotSpec200Response](docs/ModifyChooseDepotSpec200Response.md)
 - [ModifyChooseDepotSpec200ResponseResponse](docs/ModifyChooseDepotSpec200ResponseResponse.md)
 - [ModifyChooseDepotSpecRequest](docs/ModifyChooseDepotSpecRequest.md)
 - [ModifyCloseMRRequest](docs/ModifyCloseMRRequest.md)
 - [ModifyCodingCIAgentEnableRequest](docs/ModifyCodingCIAgentEnableRequest.md)
 - [ModifyCodingCIJobRequest](docs/ModifyCodingCIJobRequest.md)
 - [ModifyDefaultBranchRequest](docs/ModifyDefaultBranchRequest.md)
 - [ModifyDefectRelatedRequirementRequest](docs/ModifyDefectRelatedRequirementRequest.md)
 - [ModifyDepartment200Response](docs/ModifyDepartment200Response.md)
 - [ModifyDepartment200ResponseResponse](docs/ModifyDepartment200ResponseResponse.md)
 - [ModifyDepartmentAssigneeRequest](docs/ModifyDepartmentAssigneeRequest.md)
 - [ModifyDepartmentMemberRequest](docs/ModifyDepartmentMemberRequest.md)
 - [ModifyDepartmentRequest](docs/ModifyDepartmentRequest.md)
 - [ModifyDepotDescriptionRequest](docs/ModifyDepotDescriptionRequest.md)
 - [ModifyDepotFilePushRule200Response](docs/ModifyDepotFilePushRule200Response.md)
 - [ModifyDepotFilePushRuleDenyPrivilege200Response](docs/ModifyDepotFilePushRuleDenyPrivilege200Response.md)
 - [ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse](docs/ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse.md)
 - [ModifyDepotFilePushRuleDenyPrivilegeRequest](docs/ModifyDepotFilePushRuleDenyPrivilegeRequest.md)
 - [ModifyDepotFilePushRuleRequest](docs/ModifyDepotFilePushRuleRequest.md)
 - [ModifyDepotLevelDepotSpec200Response](docs/ModifyDepotLevelDepotSpec200Response.md)
 - [ModifyDepotLevelDepotSpecRequest](docs/ModifyDepotLevelDepotSpecRequest.md)
 - [ModifyDepotName200Response](docs/ModifyDepotName200Response.md)
 - [ModifyDepotName200ResponseResponse](docs/ModifyDepotName200ResponseResponse.md)
 - [ModifyDepotNameRequest](docs/ModifyDepotNameRequest.md)
 - [ModifyDepotPushSetting200Response](docs/ModifyDepotPushSetting200Response.md)
 - [ModifyDepotPushSetting200ResponseResponse](docs/ModifyDepotPushSetting200ResponseResponse.md)
 - [ModifyDepotPushSettingRequest](docs/ModifyDepotPushSettingRequest.md)
 - [ModifyDepotQuotaRequest](docs/ModifyDepotQuotaRequest.md)
 - [ModifyDepotSettingsRequest](docs/ModifyDepotSettingsRequest.md)
 - [ModifyDepotSharedSettingRequest](docs/ModifyDepotSharedSettingRequest.md)
 - [ModifyGitCherryPickRequest](docs/ModifyGitCherryPickRequest.md)
 - [ModifyGitCommitRevertRequest](docs/ModifyGitCommitRevertRequest.md)
 - [ModifyGitCommitStatusRequest](docs/ModifyGitCommitStatusRequest.md)
 - [ModifyGitDepotArchiveRequest](docs/ModifyGitDepotArchiveRequest.md)
 - [ModifyGitDepotUnarchiveRequest](docs/ModifyGitDepotUnarchiveRequest.md)
 - [ModifyGitFiles200Response](docs/ModifyGitFiles200Response.md)
 - [ModifyGitFiles200ResponseResponse](docs/ModifyGitFiles200ResponseResponse.md)
 - [ModifyGitFilesRequest](docs/ModifyGitFilesRequest.md)
 - [ModifyGitMergeBranch200Response](docs/ModifyGitMergeBranch200Response.md)
 - [ModifyGitMergeBranch200ResponseResponse](docs/ModifyGitMergeBranch200ResponseResponse.md)
 - [ModifyGitMergeBranchData](docs/ModifyGitMergeBranchData.md)
 - [ModifyGitMergeBranchRequest](docs/ModifyGitMergeBranchRequest.md)
 - [ModifyGitMergeRequest200Response](docs/ModifyGitMergeRequest200Response.md)
 - [ModifyGitMergeRequestRebaseRequest](docs/ModifyGitMergeRequestRebaseRequest.md)
 - [ModifyGitMergeRequestRequest](docs/ModifyGitMergeRequestRequest.md)
 - [ModifyGitRebaseRequest](docs/ModifyGitRebaseRequest.md)
 - [ModifyGitReleaseRequest](docs/ModifyGitReleaseRequest.md)
 - [ModifyGitTransfer200Response](docs/ModifyGitTransfer200Response.md)
 - [ModifyGitTransfer200ResponseResponse](docs/ModifyGitTransfer200ResponseResponse.md)
 - [ModifyGitTransferRequest](docs/ModifyGitTransferRequest.md)
 - [ModifyIssue200Response](docs/ModifyIssue200Response.md)
 - [ModifyIssue200ResponseResponse](docs/ModifyIssue200ResponseResponse.md)
 - [ModifyIssueCommentRequest](docs/ModifyIssueCommentRequest.md)
 - [ModifyIssueDescriptionRequest](docs/ModifyIssueDescriptionRequest.md)
 - [ModifyIssueParentRequirementRequest](docs/ModifyIssueParentRequirementRequest.md)
 - [ModifyIssueRequest](docs/ModifyIssueRequest.md)
 - [ModifyIteration200Response](docs/ModifyIteration200Response.md)
 - [ModifyIteration200ResponseResponse](docs/ModifyIteration200ResponseResponse.md)
 - [ModifyIterationRequest](docs/ModifyIterationRequest.md)
 - [ModifyMergeMR200Response](docs/ModifyMergeMR200Response.md)
 - [ModifyMergeMR200ResponseResponse](docs/ModifyMergeMR200ResponseResponse.md)
 - [ModifyMergeMRRequest](docs/ModifyMergeMRRequest.md)
 - [ModifyMergeRequestBasicSettingsRequest](docs/ModifyMergeRequestBasicSettingsRequest.md)
 - [ModifyMergeRequestSquashCommitMessageTemplateRequest](docs/ModifyMergeRequestSquashCommitMessageTemplateRequest.md)
 - [ModifyPolicy200Response](docs/ModifyPolicy200Response.md)
 - [ModifyPolicy200ResponseResponse](docs/ModifyPolicy200ResponseResponse.md)
 - [ModifyPolicyRequest](docs/ModifyPolicyRequest.md)
 - [ModifyProjectAnnouncement200Response](docs/ModifyProjectAnnouncement200Response.md)
 - [ModifyProjectAnnouncementRequest](docs/ModifyProjectAnnouncementRequest.md)
 - [ModifyProjectLabel200Response](docs/ModifyProjectLabel200Response.md)
 - [ModifyProjectLabel200ResponseResponse](docs/ModifyProjectLabel200ResponseResponse.md)
 - [ModifyProjectLabelRequest](docs/ModifyProjectLabelRequest.md)
 - [ModifyProjectPermissionRequest](docs/ModifyProjectPermissionRequest.md)
 - [ModifyProjectRequest](docs/ModifyProjectRequest.md)
 - [ModifyRelease200Response](docs/ModifyRelease200Response.md)
 - [ModifyRelease200ResponseResponse](docs/ModifyRelease200ResponseResponse.md)
 - [ModifyReleaseRequest](docs/ModifyReleaseRequest.md)
 - [ModifyServiceHook200Response](docs/ModifyServiceHook200Response.md)
 - [ModifyServiceHook200ResponseResponse](docs/ModifyServiceHook200ResponseResponse.md)
 - [ModifyServiceHookRequest](docs/ModifyServiceHookRequest.md)
 - [ModifyTeamLevelDepotSpec200Response](docs/ModifyTeamLevelDepotSpec200Response.md)
 - [ModifyTeamLevelDepotSpecRequest](docs/ModifyTeamLevelDepotSpecRequest.md)
 - [ModifyTeamMemberUnlockedRequest](docs/ModifyTeamMemberUnlockedRequest.md)
 - [ModifyTestCase200Response](docs/ModifyTestCase200Response.md)
 - [ModifyTestCase200ResponseResponse](docs/ModifyTestCase200ResponseResponse.md)
 - [ModifyTestCaseRequest](docs/ModifyTestCaseRequest.md)
 - [ModifyTestCaseSection200Response](docs/ModifyTestCaseSection200Response.md)
 - [ModifyTestCaseSection200ResponseResponse](docs/ModifyTestCaseSection200ResponseResponse.md)
 - [ModifyTestCaseSectionRequest](docs/ModifyTestCaseSectionRequest.md)
 - [ModifyTestRun200Response](docs/ModifyTestRun200Response.md)
 - [ModifyTestRun200ResponseResponse](docs/ModifyTestRun200ResponseResponse.md)
 - [ModifyTestRunRequest](docs/ModifyTestRunRequest.md)
 - [ModifyWiki200Response](docs/ModifyWiki200Response.md)
 - [ModifyWiki200ResponseResponse](docs/ModifyWiki200ResponseResponse.md)
 - [ModifyWikiByZipRequest](docs/ModifyWikiByZipRequest.md)
 - [ModifyWikiOrderRequest](docs/ModifyWikiOrderRequest.md)
 - [ModifyWikiRequest](docs/ModifyWikiRequest.md)
 - [ModifyWikiTitle200Response](docs/ModifyWikiTitle200Response.md)
 - [ModifyWikiTitleRequest](docs/ModifyWikiTitleRequest.md)
 - [ModifyWorkItemSplitIssuesRequest](docs/ModifyWorkItemSplitIssuesRequest.md)
 - [OpenApiIssueListDataWithPage](docs/OpenApiIssueListDataWithPage.md)
 - [OpenApiRelease](docs/OpenApiRelease.md)
 - [OpenApiReleaseListDataWithPage](docs/OpenApiReleaseListDataWithPage.md)
 - [OpenApiTeamIssueData](docs/OpenApiTeamIssueData.md)
 - [OpenApiWorkbenchIssue](docs/OpenApiWorkbenchIssue.md)
 - [PageInfo](docs/PageInfo.md)
 - [PingServiceHookRequest](docs/PingServiceHookRequest.md)
 - [PinyinPinyinArray](docs/PinyinPinyinArray.md)
 - [PinyinPinyinResult](docs/PinyinPinyinResult.md)
 - [PipelineConfig](docs/PipelineConfig.md)
 - [PipelineIdIndex](docs/PipelineIdIndex.md)
 - [PlanIterationIssueRequest](docs/PlanIterationIssueRequest.md)
 - [Policy](docs/Policy.md)
 - [PolicyDetail](docs/PolicyDetail.md)
 - [PolicyDocument](docs/PolicyDocument.md)
 - [PolicyInfo](docs/PolicyInfo.md)
 - [PolicyResourceScopeInfo](docs/PolicyResourceScopeInfo.md)
 - [PolicyStatement](docs/PolicyStatement.md)
 - [PredicatePolicy](docs/PredicatePolicy.md)
 - [PrincipalAdd](docs/PrincipalAdd.md)
 - [PrincipalData](docs/PrincipalData.md)
 - [PrincipalDel](docs/PrincipalDel.md)
 - [PrincipalResp](docs/PrincipalResp.md)
 - [Program](docs/Program.md)
 - [ProgramData](docs/ProgramData.md)
 - [ProgramProgram](docs/ProgramProgram.md)
 - [ProgramProgramMemberPolicy](docs/ProgramProgramMemberPolicy.md)
 - [Project](docs/Project.md)
 - [ProjectAnnouncementProjectAnnouncement](docs/ProjectAnnouncementProjectAnnouncement.md)
 - [ProjectIssueField](docs/ProjectIssueField.md)
 - [ProjectIssueStatus](docs/ProjectIssueStatus.md)
 - [ProjectLabelLabels](docs/ProjectLabelLabels.md)
 - [ProjectMemberData](docs/ProjectMemberData.md)
 - [ProjectMemberDepartmentMember](docs/ProjectMemberDepartmentMember.md)
 - [ProjectMemberMemberRef](docs/ProjectMemberMemberRef.md)
 - [ProjectMemberUserData](docs/ProjectMemberUserData.md)
 - [ProjectsData](docs/ProjectsData.md)
 - [ProtectedBranch](docs/ProtectedBranch.md)
 - [ProtectedBranchMember](docs/ProtectedBranchMember.md)
 - [QcloudApiGitDepotDiffFileInfo](docs/QcloudApiGitDepotDiffFileInfo.md)
 - [QcloudApiGitDepotDifferentLine](docs/QcloudApiGitDepotDifferentLine.md)
 - [RamGrantResourceInfoRequest](docs/RamGrantResourceInfoRequest.md)
 - [RefInfo](docs/RefInfo.md)
 - [RelatedCase](docs/RelatedCase.md)
 - [Release](docs/Release.md)
 - [ReleaseListPage](docs/ReleaseListPage.md)
 - [ReorderCdPipelines200Response](docs/ReorderCdPipelines200Response.md)
 - [ReorderCdPipelines200ResponseResponse](docs/ReorderCdPipelines200ResponseResponse.md)
 - [ReorderCdPipelinesRequest](docs/ReorderCdPipelinesRequest.md)
 - [ReorderCdPipelinesResponseData](docs/ReorderCdPipelinesResponseData.md)
 - [Report](docs/Report.md)
 - [ReportData](docs/ReportData.md)
 - [ReportLittle](docs/ReportLittle.md)
 - [ReportLittleData](docs/ReportLittleData.md)
 - [ReportOverview](docs/ReportOverview.md)
 - [ReportsLittleData](docs/ReportsLittleData.md)
 - [RequirementType](docs/RequirementType.md)
 - [ResendServiceHookLog200Response](docs/ResendServiceHookLog200Response.md)
 - [ResendServiceHookLog200ResponseResponse](docs/ResendServiceHookLog200ResponseResponse.md)
 - [ResendServiceHookLogRequest](docs/ResendServiceHookLogRequest.md)
 - [ResourceInfo](docs/ResourceInfo.md)
 - [ResourceInfoOfPolicyScope](docs/ResourceInfoOfPolicyScope.md)
 - [ResourceReference](docs/ResourceReference.md)
 - [ResourceReferenceItem](docs/ResourceReferenceItem.md)
 - [ResourceReferenceResource](docs/ResourceReferenceResource.md)
 - [ResourceReferenceResourceInfo](docs/ResourceReferenceResourceInfo.md)
 - [ResourceReferenceResourceScope](docs/ResourceReferenceResourceScope.md)
 - [Role](docs/Role.md)
 - [Run](docs/Run.md)
 - [RunData](docs/RunData.md)
 - [RunsData](docs/RunsData.md)
 - [ScanInfo](docs/ScanInfo.md)
 - [ScanSeverityProtoDTO](docs/ScanSeverityProtoDTO.md)
 - [Section](docs/Section.md)
 - [SectionData](docs/SectionData.md)
 - [SectionsData](docs/SectionsData.md)
 - [ServiceAccountForm](docs/ServiceAccountForm.md)
 - [ServiceHook](docs/ServiceHook.md)
 - [ServiceHookEvent](docs/ServiceHookEvent.md)
 - [ServiceHookLog](docs/ServiceHookLog.md)
 - [ServiceHookLogPage](docs/ServiceHookLogPage.md)
 - [ServiceHookPage](docs/ServiceHookPage.md)
 - [ServiceHookResendServiceHookResult](docs/ServiceHookResendServiceHookResult.md)
 - [ServiceHookUser](docs/ServiceHookUser.md)
 - [SetGrantToResourceRequest](docs/SetGrantToResourceRequest.md)
 - [SetPredicatePolicyOnResourceRequest](docs/SetPredicatePolicyOnResourceRequest.md)
 - [SpecifiedArtifact](docs/SpecifiedArtifact.md)
 - [SshKeyInfo](docs/SshKeyInfo.md)
 - [StatusCheckResult](docs/StatusCheckResult.md)
 - [SubTask](docs/SubTask.md)
 - [TKEConfigForm](docs/TKEConfigForm.md)
 - [TeamAdminMember](docs/TeamAdminMember.md)
 - [TeamAdminMemberData](docs/TeamAdminMemberData.md)
 - [TeamArtifact](docs/TeamArtifact.md)
 - [TeamArtifactPage](docs/TeamArtifactPage.md)
 - [TeamData](docs/TeamData.md)
 - [TeamMemberData](docs/TeamMemberData.md)
 - [TencentCloudAccountForm](docs/TencentCloudAccountForm.md)
 - [Test](docs/Test.md)
 - [TestData](docs/TestData.md)
 - [TestDefect](docs/TestDefect.md)
 - [TestDefectsData](docs/TestDefectsData.md)
 - [TestFull](docs/TestFull.md)
 - [TestsData](docs/TestsData.md)
 - [Token](docs/Token.md)
 - [TriggerCdPipeline200Response](docs/TriggerCdPipeline200Response.md)
 - [TriggerCdPipeline200ResponseResponse](docs/TriggerCdPipeline200ResponseResponse.md)
 - [TriggerCdPipelineRequest](docs/TriggerCdPipelineRequest.md)
 - [TriggerCdPipelineResponseData](docs/TriggerCdPipelineResponseData.md)
 - [TriggerCodingCIBuild200Response](docs/TriggerCodingCIBuild200Response.md)
 - [TriggerCodingCIBuild200ResponseResponse](docs/TriggerCodingCIBuild200ResponseResponse.md)
 - [TriggerCodingCIBuildData](docs/TriggerCodingCIBuildData.md)
 - [TriggerCodingCIBuildRequest](docs/TriggerCodingCIBuildRequest.md)
 - [TriggerCodingScan200Response](docs/TriggerCodingScan200Response.md)
 - [TriggerCodingScan200ResponseResponse](docs/TriggerCodingScan200ResponseResponse.md)
 - [TriggerCodingScanData](docs/TriggerCodingScanData.md)
 - [TriggerCodingScanRequest](docs/TriggerCodingScanRequest.md)
 - [UpdateUserGroupByIdRequest](docs/UpdateUserGroupByIdRequest.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)
 - [UserGroup](docs/UserGroup.md)
 - [UserGroupUserInfos](docs/UserGroupUserInfos.md)
 - [WikiChildrenData](docs/WikiChildrenData.md)
 - [WikiData](docs/WikiData.md)
 - [WikiJobStatus](docs/WikiJobStatus.md)
 - [WikiListData](docs/WikiListData.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



