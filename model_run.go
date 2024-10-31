/*
CODING OPEN API

CODING 提供了丰富的 API 接口，注册应用即可使用，无需审核，支持两种认证方式：[OAuth 2.0 协议](#oauth-认证)、[个人访问令牌](#个人令牌认证)、[项目令牌](#项目令牌认证)。  # [OAuth 认证](#oauth)  #### [生态令牌权限](#ecology-scope)  令牌权限说明 > Scope的权限分为只读（ro）和读写(rw) 俩种，用户可通过权限点+权限方式获取完整的权限 Scope 信息。例如，用户信息授权为只读时，用户的令牌权限 Scope 为 user:profile:ro  | 名称            | 描述信息                 | Scope 信息                  | 权限范围  | 示例                           | |---------------|----------------------|---------------------------|-------|------------------------------| | 用户信息          | 管理用户的基础信息。           | user:profile              | ro    | user:profile:ro              | | 用户邮箱          | 管理用户的电子邮件地址。         | user:email                | ro    | user:email:ro                | | 用户通知          | 管理用户的站内通知。           | user:notification         | ro、rw | user:notification:rw         | | 用户公钥          | 管理用户配置的个人公钥和部署公钥信息。  | user:public-key           | ro、rw | user:public-key:rw           | | 团队信息          | 管理团队基本信息。            | team:profile              | ro    | team:profile:ro              | | 团队成员          | 管理团队成员信息以及团队成员相关操作。  | team:member               | ro、rw | team:member:rw               | | 项目信息          | 管理项目基本信息。            | project:profile           | ro、rw | project:profile:rw           | | 项目成员          | 管理项目成员。              | project:member            | ro、rw | project:member:rw            | | 项目令牌          | 管理项目令牌。              | project:token             | ro、rw | project:token:rw             | | 项目公告          | 管理项目公告。              | project:notice            | ro、rw | project:notice:rw            | | 项目标签          | 管理项目标签。              | project:label             | ro、rw | project:label:rw             | | 项目集信息         | 管理项目集基本信息。           | program:profile           | ro、rw | program:profile:rw           | | 项目集项目         | 管理项目集下的项目列表。         | program:project           | ro、rw | program:project:rw           | | 项目集成员         | 管理项目集下的成员列表。         | program:member            | ro、rw | program:member:rw            | | 关联资源          | 管理团队和项目资源关联关系。       | related-resource:resource | ro、rw | related-resource:resource:rw | | 凭据信息          | 管理团队凭据。              | credential:profile        | ro、rw | credential:profile:rw        | | Service Hooks | 管理和配置 Service Hooks。 | service-hook:profile      | ro、rw | service-hook:profile:rw      | | 权限组           | 管理权限组。               | ram:policy                | ro、rw | ram:policy:rw                | | 授权            | 配置权限授权。              | ram:grant                 | ro、rw | ram:grant:rw                 | | 用户组           | 管理权限用户组。             | ram:user-group            | ro    | ram:user-group:ro            | | 研发度量数据集       | 研发度量数据集              | performance:dataset       | ro    | performance:dataset:ro       | | 项目协同          | 配置和使用项目协同功能。         | collaboration:issue       | ro、rw | collaboration:issue:rw       | | 知识管理          | 管理知识空间和撰写知识文档。       | document:knowledge        | ro、rw | document:knowledge:rw        | | 文件网盘          | 管理上传、分享和下载文件等。       | document:file             | ro、rw | document:file:rw             | | API 文档        | 发布、授权发布 API 文档。      | document:api-doc          | ro、rw | document:api-doc:rw          | | 代码仓库          | 管理仓库                 | vcs:repository            | ro、rw | vcs:repository:rw            | | 合并请求          | 管理代码仓库的合并请求。         | vcs:merge                 | ro、rw | vcs:merge:rw                 | | 部署公钥          | 管理代码仓库的部署公钥。         | vcs:ssh-key               | ro、rw | vcs:ssh-key:rw               | | 版本管理          | 管理代码仓库的版本信息。         | vcs:release               | ro、rw | vcs:release:rw               | | 外部仓库          | 管理关联的外部仓库信息。         | depot:external-repository | ro、rw | depot:external-repository:rw | | 测试管理          | 管理测试计、测试用例和测试报告等。    | testing:profile           | ro、rw | testing:profile:rw           | | 持续部署数据统计      | 持续部署发布数据统计。          | cd:statistics             | ro    | cd:statistics:ro             | | 持续部署主机组       | 管理持续部署主机组。           | cd:host-server            | ro、rw | cd:host-server:rw            | | 持续部署云账号       | 管理持续部署云账号。           | cd:cloud-account          | ro、rw | cd:cloud-account:rw          | | 持续部署应用        | 管理和配置持续部署应用。         | cd:application            | ro、rw | cd:application:rw            | | 持续部署流程        | 管理和配置持续部署流程。         | cd:pipeline               | ro、rw | cd:pipeline:rw               | | 制品库仓库         | 管理制品库仓库。             | artifact:repository       | ro、rw | artifact:repository:rw       | | 制品库版本         | 管理制品版本信息。            | artifact:version          | ro、rw | artifact:version:rw          | | 制品库配置         | 管理制品库配置。             | artifact:properties       | ro、rw | artifact:properties:rw       | | 制品库包          | 管理制品库包。              | artifact:package          | ro    | artifact:package:ro          | | 资产列表          | 管理资产列表               | assets:list               | ro、rw | assets:list:rw               | | 资产属性          | 管理资产属性               | assets:attribute          | ro、rw | assets:attribute:rw          |  #### [创建 CODING 应用](#创建-CODING-应用)  ##### 1、新建应用  点击【团队设置】->【生态能力】->【发布应用】->【新建应用】，填写信息。「回调地址」处填写回调服务地址，这里以codesign为例。  ![](https://help-assets.codehub.cn/enterprise/202309201515397.png)  ![](https://help-assets.codehub.cn/enterprise/202309201519877.png)  ##### 2、获取 Client ID 和 Client Secret  点击创建好的应用，点击客户端秘钥，即可看到客户端ID与客户端秘钥  ![](https://help-assets.codehub.cn/enterprise/202309201520614.png)  ##### 注意事项： - 更换令牌时，先新建令牌，再删除旧令牌 - 最多同时可创建 5 个令牌 - 令牌只有创建时可查看，创建后，任何人无法查看，请妥善保管。  #### 3、修改令牌权限  根据需求修改令牌权限，这里设置了用户信息只读权限与项目信息读写权限作为示例  ![](https://help-assets.codehub.cn/enterprise/202309201520611.png)  #### [用户授权](#oauth-scopes)  浏览器访问以下链接，进入到授权登录页面：  ```shell GET https://{your-team}.coding.net/oauth_authorize.html?client_id={client_id}&redirect_uri={redirect_uri}&response_type=code&state=${state}&scope={scope}  ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - redirect_uri：应用详情页填写的回调地址；  - response_type：返回类型，固定为 code；  - state: 业务侧可以传任何值，这个值会原封不动传递回去，用来给业务识别一些场景用的。  - scope：授权范围，多个权限间以逗号分隔，如：user:profile:ro,project:profile:rw  ![](https://help-assets.codehub.cn/enterprise/202309201520497.png)  点击授权后，浏览器将带着授权码（code）参数和业务状态码（state）跳转到回调地址，如：  ```shell https://codesign.qq.com/?code={code}&state={state}&team={teamGk}&scope=user%3Aprofile%3Aro,project%3Aprofile%3Arw ```  #### [获取 access_token](#oauth-access-token)  获取授权码（code）后，开发者的后端程序向 CODING 发送请求，获取 access_token。  请求链接：  ```shell POST https://{your-team}.coding.net/api/oauth/access_token Content-Type: application/x-www-form-urlencoded  client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&code={code} ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - client_secret：应用详情页的 Client Secret；  - code：上一步获取的授权码，须注意每个 code 仅能使用一次，且有效期仅5分钟；  - grant_type：授权类型，固定为 authorization_code；  返回值：  ```json {   \"access_token\": \"RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975\",   \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",   \"scope\": \"user:profile:ro project:profile:rw\",   \"team\": \"jackwhu-test\",   \"token_type\": \"Bearer\",   \"expires_in\": \"7200\" } ```  返回值： - access_token： access_token值，可用于调用OpenAPI接口，建议按expires_in保存access_token - refresh_token： 刷新时使用的token，有效期90天。access_token过期后可用于刷新access_token - scope：令牌权限范围 - team： 团队gk - token_type：token类型 - expires_in：过期时间，单位秒   #### [刷新 access_token](#oauth-access-token)  通过上面获取到的 refresh_token，开发者的后端程序可以向 CODING 发送请求，刷新 access_token。   注意：调用刷新接口后，即使 access_token 未过期，原 access_token 也将失效  请求链接：  ```shell POST https://{your-team}.coding.net/api/oauth/access_token Content-Type: application/x-www-form-urlencoded  client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&refresh_token={refresh_token} ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - client_secret：应用详情页的 Client Secret；  - refresh_token：获取 access_token接口返回的refresh_token字段  - grant_type：授权类型，固定为 refresh_token；  返回值：  ```json {   \"access_token\": \"q4qIkUGhJ2qfcdSV3bWx0YfQj-WjLqXG7LSdP9Oo3sOAjmuY-Bb_QJ6ORt-By-bc7WFFT7PyH7RXEvPLBM5lFU9PBOofgzXN9Lh5zp3FWRdyV_4XGno4U_S7zMkixWnv\",   \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",   \"scope\": \"user:profile:ro project:profile:rw\",   \"team\": \"jackwhu-test\",   \"token_type\": \"Bearer\",   \"expires_in\": \"7200\" } ```  #### [获取当前用户信息](#oauth-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: Bearer RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975 Content-Type: application/json  {     \"Action\": \"DescribeCodingCurrentUser\" } ```  Headers说明：  - Authorization：\"Bearer {access_token}\"  参数说明：  - Action：固定为DescribeCodingCurrentUser，令牌须打开user:profile:ro权限  返回值：  ```json {   \"Response\": {     \"User\": {       \"Id\": 183478,       \"Status\": 1,       \"Email\": \"test@coding.com\",       \"GlobalKey\": \"anywhere\",       \"Avatar\": \"https://e.coding.net/static/fruit_avatar/Fruit-20.png\",       \"Gravatar\": \"\",       \"Name\": \"qqq\",       \"NamePinYin\": \"qqq\",       \"Phone\": \"18800000000\",       \"PhoneValidation\": 1,       \"EmailValidation\": 1,       \"PhoneRegionCode\": \"+86\",       \"Region\": \"cn\",       \"TeamId\": 1,       \"WeComId\": \"woG7VgCgAAov2F-sAQkD_YPsLNJiP1Vg\"     },     \"RequestId\": \"133e152f-8852-4d99-b704-c7ff245a1640\"   } }  ```  # [个人令牌认证](#personal_token)  #### [获取个人令牌](#personal-token-create)  - 点击左下角的【个人账户设置】>【访问令牌】>【新建访问令牌】，勾选相关权限后会生成「个人访问令牌」。若刷新页面令牌会消失，需输入账号密码后重新生成。 - 令牌权限与[OAuth令牌权限](#生态令牌权限)一样  ![](https://help-assets.codehub.cn/enterprise/202309201521630.png)  #### [获取当前用户信息](#personal-token-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8 Content-Type: application/json  {     \"Action\": \"DescribeCodingCurrentUser\" } ```  header：  ```text Authorization: token {访问令牌}  ```  请求示例：  ```shell curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\ --header 'Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"Action\": \"DescribeCodingCurrentUser\" }' ```  # [项目令牌认证](#personal_token)  #### [项目令牌Scope](#project-token-scope)  | 名称      | 描述                        | Scope 信息                  | 权限范围 | |---------|---------------------------|---------------------------|------| | 项目协同    | 读取与操作项目协同                 | project_issue_rw          | 读写   | | 文件      | 新建、查询、删除、编辑文件             | file_rw                   | 读写   | | WIKI    | 新建、查询、删除、编辑Wiki           | wiki_rw                   | 读写   | | 项目公告    | 新建、查询、删除、编辑项目公告           | project_tweet_rw          | 读写   | | API 文档  | 发布 API 文档                 | api_doc_release           | 读写   | | 关联资源    | 新建、查询、删除关联资源              | resource_reference_rw     | 读写   | | 项目成员    | 读取与操作项目成员                 | project_member_rw         | 读写   | | 项目权限    | 读取与操作项目权限                 | project_permission_rw     | 读写   | | 项目标签    | 新建、查询、删除、编辑项目标签           | project_label_rw          | 读写   | | 测试协同    | 执行 API 自动化测试              | ifbook_run_task           | 读写   | | 测试协同    | API 文档导入与导出               | ifbook_import_export      | 读写   | | 读取代码仓库  | 读取代码仓库                    | depot_read                | 只读   | | 推送至代码仓库 | 推送至代码仓库                   | depot_write               | 读写   | | MR      | 新建、查询、删除、编辑合并请求           | merge_request_rw          | 读写   | | 版本发布    | 新建、查询、删除、编辑版本发布           | release_rw                | 读写   | | 制品库     | 拉取制品库                     | artifact_r                | 只读   | | 制品库     | 拉取、推送制品库                  | artifact_rw               | 读写   | | 制品属性    | 新建、查询、删除、编辑制品属性           | artifact_version_props_rw | 读写   | | 构建节点    | 允许构建节点接入                  | ci_agent_register         | 读写   | | API触发   | 构建计划管理/触发构建               | ci_manager                | 读写   | | 构建计划    | 构建计划管理/触发构建（仅用于 Open API） | open_ci_manager           | 读写   |  #### [获取项目令牌](#project-token-create)  1. 点击【项目设置】>【开发者选项】>【项目令牌】>【新建项目令牌】，勾选相关权限后会生成「项目令牌」。点击查看密码即可获取密码信息  ![](https://help-assets.codehub.cn/enterprise/202309201843081.png)  2. Basic认证：将用户名与密码通过”用户名:密码“方式拼接后，使用Base64进行编码，获取凭证  #### [获取当前项目信息](#project-token-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE Content-Type: application/json  {     \"Action\": \"DescribeOneProject\",     \"ProjectId\": 11595365 } ```  header：  ```text Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE ```  请求示例：  ```shell curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\ --header 'Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"Action\": \"DescribeOneProject\",     \"ProjectId\": 11595365 }' ```  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Run type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Run{}

// Run 测试计划详情
type Run struct {
	// 处理人ID
	AssignedToId NullableInt32 `json:"AssignedToId,omitempty"`
	// 计划内阻塞测试数量
	BlockedCount NullableInt32 `json:"BlockedCount,omitempty"`
	// 归档时间
	CompletedAt NullableString `json:"CompletedAt,omitempty"`
	// 环境标识
	ConfigEnvironmentId NullableInt32 `json:"ConfigEnvironmentId,omitempty"`
	// 创建时间
	CreatedAt NullableString `json:"CreatedAt,omitempty"`
	// 创建人ID
	CreatedBy NullableInt32 `json:"CreatedBy,omitempty"`
	// 持续天数
	Days NullableInt32 `json:"Days,omitempty"`
	// 描述
	Description NullableString `json:"Description,omitempty"`
	// 执行方式: 1-手动执行 2-自动化流水线执行
	ExecuteType NullableInt32 `json:"ExecuteType,omitempty"`
	// 计划内失败测试数量
	FailedCount NullableInt32 `json:"FailedCount,omitempty"`
	// 代码仓库 ID
	GitDepotId NullableInt32 `json:"GitDepotId,omitempty"`
	// 代码仓库名名称
	GitDepotName NullableString `json:"GitDepotName,omitempty"`
	// 发布版本 ID（资源 ID）
	GitReleaseId NullableInt32 `json:"GitReleaseId,omitempty"`
	// 发布版本名称
	GitReleaseName NullableString `json:"GitReleaseName,omitempty"`
	// 发布版本名称状态：0-未发布 1-已发布
	GitReleaseState NullableInt32 `json:"GitReleaseState,omitempty"`
	// ID 主键
	Id NullableInt32 `json:"Id,omitempty"`
	// 是否包含全部用例
	IncludeAll NullableBool `json:"IncludeAll,omitempty"`
	// 是否归档
	IsCompleted NullableBool `json:"IsCompleted,omitempty"`
	// 所属迭代 ID
	IterationId NullableInt64 `json:"IterationId,omitempty"`
	// 迭代名称
	IterationName NullableString `json:"IterationName,omitempty"`
	// 名称
	Name NullableString `json:"Name,omitempty"`
	// 计划内通过测试数量
	PassedCount NullableInt32 `json:"PassedCount,omitempty"`
	// 计划内重新测试数量
	RetestCount NullableInt32 `json:"RetestCount,omitempty"`
	// 分组 ID
	SectionId NullableInt64 `json:"SectionId,omitempty"`
	// 分组名
	SectionName NullableString `json:"SectionName,omitempty"`
	// 状态: 0-未开始 1-进行中 2-已测完
	State NullableInt32 `json:"State,omitempty"`
	// 计划内未测试数量
	UntestedCount NullableInt32 `json:"UntestedCount,omitempty"`
}

// NewRun instantiates a new Run object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRun() *Run {
	this := Run{}
	var completedAt string = ""
	this.CompletedAt = *NewNullableString(&completedAt)
	var createdAt string = ""
	this.CreatedAt = *NewNullableString(&createdAt)
	var description string = ""
	this.Description = *NewNullableString(&description)
	var gitDepotName string = ""
	this.GitDepotName = *NewNullableString(&gitDepotName)
	var gitReleaseName string = ""
	this.GitReleaseName = *NewNullableString(&gitReleaseName)
	var includeAll bool = false
	this.IncludeAll = *NewNullableBool(&includeAll)
	var isCompleted bool = false
	this.IsCompleted = *NewNullableBool(&isCompleted)
	var iterationName string = ""
	this.IterationName = *NewNullableString(&iterationName)
	var name string = ""
	this.Name = *NewNullableString(&name)
	var sectionName string = ""
	this.SectionName = *NewNullableString(&sectionName)
	return &this
}

// NewRunWithDefaults instantiates a new Run object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunWithDefaults() *Run {
	this := Run{}
	var completedAt string = ""
	this.CompletedAt = *NewNullableString(&completedAt)
	var createdAt string = ""
	this.CreatedAt = *NewNullableString(&createdAt)
	var description string = ""
	this.Description = *NewNullableString(&description)
	var gitDepotName string = ""
	this.GitDepotName = *NewNullableString(&gitDepotName)
	var gitReleaseName string = ""
	this.GitReleaseName = *NewNullableString(&gitReleaseName)
	var includeAll bool = false
	this.IncludeAll = *NewNullableBool(&includeAll)
	var isCompleted bool = false
	this.IsCompleted = *NewNullableBool(&isCompleted)
	var iterationName string = ""
	this.IterationName = *NewNullableString(&iterationName)
	var name string = ""
	this.Name = *NewNullableString(&name)
	var sectionName string = ""
	this.SectionName = *NewNullableString(&sectionName)
	return &this
}

// GetAssignedToId returns the AssignedToId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetAssignedToId() int32 {
	if o == nil || IsNil(o.AssignedToId.Get()) {
		var ret int32
		return ret
	}
	return *o.AssignedToId.Get()
}

// GetAssignedToIdOk returns a tuple with the AssignedToId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetAssignedToIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedToId.Get(), o.AssignedToId.IsSet()
}

// HasAssignedToId returns a boolean if a field has been set.
func (o *Run) HasAssignedToId() bool {
	if o != nil && o.AssignedToId.IsSet() {
		return true
	}

	return false
}

// SetAssignedToId gets a reference to the given NullableInt32 and assigns it to the AssignedToId field.
func (o *Run) SetAssignedToId(v int32) {
	o.AssignedToId.Set(&v)
}
// SetAssignedToIdNil sets the value for AssignedToId to be an explicit nil
func (o *Run) SetAssignedToIdNil() {
	o.AssignedToId.Set(nil)
}

// UnsetAssignedToId ensures that no value is present for AssignedToId, not even an explicit nil
func (o *Run) UnsetAssignedToId() {
	o.AssignedToId.Unset()
}

// GetBlockedCount returns the BlockedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetBlockedCount() int32 {
	if o == nil || IsNil(o.BlockedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.BlockedCount.Get()
}

// GetBlockedCountOk returns a tuple with the BlockedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetBlockedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockedCount.Get(), o.BlockedCount.IsSet()
}

// HasBlockedCount returns a boolean if a field has been set.
func (o *Run) HasBlockedCount() bool {
	if o != nil && o.BlockedCount.IsSet() {
		return true
	}

	return false
}

// SetBlockedCount gets a reference to the given NullableInt32 and assigns it to the BlockedCount field.
func (o *Run) SetBlockedCount(v int32) {
	o.BlockedCount.Set(&v)
}
// SetBlockedCountNil sets the value for BlockedCount to be an explicit nil
func (o *Run) SetBlockedCountNil() {
	o.BlockedCount.Set(nil)
}

// UnsetBlockedCount ensures that no value is present for BlockedCount, not even an explicit nil
func (o *Run) UnsetBlockedCount() {
	o.BlockedCount.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetCompletedAt() string {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetCompletedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *Run) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableString and assigns it to the CompletedAt field.
func (o *Run) SetCompletedAt(v string) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *Run) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *Run) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetConfigEnvironmentId returns the ConfigEnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetConfigEnvironmentId() int32 {
	if o == nil || IsNil(o.ConfigEnvironmentId.Get()) {
		var ret int32
		return ret
	}
	return *o.ConfigEnvironmentId.Get()
}

// GetConfigEnvironmentIdOk returns a tuple with the ConfigEnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetConfigEnvironmentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigEnvironmentId.Get(), o.ConfigEnvironmentId.IsSet()
}

// HasConfigEnvironmentId returns a boolean if a field has been set.
func (o *Run) HasConfigEnvironmentId() bool {
	if o != nil && o.ConfigEnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetConfigEnvironmentId gets a reference to the given NullableInt32 and assigns it to the ConfigEnvironmentId field.
func (o *Run) SetConfigEnvironmentId(v int32) {
	o.ConfigEnvironmentId.Set(&v)
}
// SetConfigEnvironmentIdNil sets the value for ConfigEnvironmentId to be an explicit nil
func (o *Run) SetConfigEnvironmentIdNil() {
	o.ConfigEnvironmentId.Set(nil)
}

// UnsetConfigEnvironmentId ensures that no value is present for ConfigEnvironmentId, not even an explicit nil
func (o *Run) UnsetConfigEnvironmentId() {
	o.ConfigEnvironmentId.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Run) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Run) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Run) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Run) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetCreatedBy() int32 {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret int32
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Run) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableInt32 and assigns it to the CreatedBy field.
func (o *Run) SetCreatedBy(v int32) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *Run) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *Run) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetDays returns the Days field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetDays() int32 {
	if o == nil || IsNil(o.Days.Get()) {
		var ret int32
		return ret
	}
	return *o.Days.Get()
}

// GetDaysOk returns a tuple with the Days field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Days.Get(), o.Days.IsSet()
}

// HasDays returns a boolean if a field has been set.
func (o *Run) HasDays() bool {
	if o != nil && o.Days.IsSet() {
		return true
	}

	return false
}

// SetDays gets a reference to the given NullableInt32 and assigns it to the Days field.
func (o *Run) SetDays(v int32) {
	o.Days.Set(&v)
}
// SetDaysNil sets the value for Days to be an explicit nil
func (o *Run) SetDaysNil() {
	o.Days.Set(nil)
}

// UnsetDays ensures that no value is present for Days, not even an explicit nil
func (o *Run) UnsetDays() {
	o.Days.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Run) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Run) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Run) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Run) UnsetDescription() {
	o.Description.Unset()
}

// GetExecuteType returns the ExecuteType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetExecuteType() int32 {
	if o == nil || IsNil(o.ExecuteType.Get()) {
		var ret int32
		return ret
	}
	return *o.ExecuteType.Get()
}

// GetExecuteTypeOk returns a tuple with the ExecuteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetExecuteTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecuteType.Get(), o.ExecuteType.IsSet()
}

// HasExecuteType returns a boolean if a field has been set.
func (o *Run) HasExecuteType() bool {
	if o != nil && o.ExecuteType.IsSet() {
		return true
	}

	return false
}

// SetExecuteType gets a reference to the given NullableInt32 and assigns it to the ExecuteType field.
func (o *Run) SetExecuteType(v int32) {
	o.ExecuteType.Set(&v)
}
// SetExecuteTypeNil sets the value for ExecuteType to be an explicit nil
func (o *Run) SetExecuteTypeNil() {
	o.ExecuteType.Set(nil)
}

// UnsetExecuteType ensures that no value is present for ExecuteType, not even an explicit nil
func (o *Run) UnsetExecuteType() {
	o.ExecuteType.Unset()
}

// GetFailedCount returns the FailedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetFailedCount() int32 {
	if o == nil || IsNil(o.FailedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FailedCount.Get()
}

// GetFailedCountOk returns a tuple with the FailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetFailedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedCount.Get(), o.FailedCount.IsSet()
}

// HasFailedCount returns a boolean if a field has been set.
func (o *Run) HasFailedCount() bool {
	if o != nil && o.FailedCount.IsSet() {
		return true
	}

	return false
}

// SetFailedCount gets a reference to the given NullableInt32 and assigns it to the FailedCount field.
func (o *Run) SetFailedCount(v int32) {
	o.FailedCount.Set(&v)
}
// SetFailedCountNil sets the value for FailedCount to be an explicit nil
func (o *Run) SetFailedCountNil() {
	o.FailedCount.Set(nil)
}

// UnsetFailedCount ensures that no value is present for FailedCount, not even an explicit nil
func (o *Run) UnsetFailedCount() {
	o.FailedCount.Unset()
}

// GetGitDepotId returns the GitDepotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitDepotId() int32 {
	if o == nil || IsNil(o.GitDepotId.Get()) {
		var ret int32
		return ret
	}
	return *o.GitDepotId.Get()
}

// GetGitDepotIdOk returns a tuple with the GitDepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitDepotIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitDepotId.Get(), o.GitDepotId.IsSet()
}

// HasGitDepotId returns a boolean if a field has been set.
func (o *Run) HasGitDepotId() bool {
	if o != nil && o.GitDepotId.IsSet() {
		return true
	}

	return false
}

// SetGitDepotId gets a reference to the given NullableInt32 and assigns it to the GitDepotId field.
func (o *Run) SetGitDepotId(v int32) {
	o.GitDepotId.Set(&v)
}
// SetGitDepotIdNil sets the value for GitDepotId to be an explicit nil
func (o *Run) SetGitDepotIdNil() {
	o.GitDepotId.Set(nil)
}

// UnsetGitDepotId ensures that no value is present for GitDepotId, not even an explicit nil
func (o *Run) UnsetGitDepotId() {
	o.GitDepotId.Unset()
}

// GetGitDepotName returns the GitDepotName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitDepotName() string {
	if o == nil || IsNil(o.GitDepotName.Get()) {
		var ret string
		return ret
	}
	return *o.GitDepotName.Get()
}

// GetGitDepotNameOk returns a tuple with the GitDepotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitDepotNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitDepotName.Get(), o.GitDepotName.IsSet()
}

// HasGitDepotName returns a boolean if a field has been set.
func (o *Run) HasGitDepotName() bool {
	if o != nil && o.GitDepotName.IsSet() {
		return true
	}

	return false
}

// SetGitDepotName gets a reference to the given NullableString and assigns it to the GitDepotName field.
func (o *Run) SetGitDepotName(v string) {
	o.GitDepotName.Set(&v)
}
// SetGitDepotNameNil sets the value for GitDepotName to be an explicit nil
func (o *Run) SetGitDepotNameNil() {
	o.GitDepotName.Set(nil)
}

// UnsetGitDepotName ensures that no value is present for GitDepotName, not even an explicit nil
func (o *Run) UnsetGitDepotName() {
	o.GitDepotName.Unset()
}

// GetGitReleaseId returns the GitReleaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitReleaseId() int32 {
	if o == nil || IsNil(o.GitReleaseId.Get()) {
		var ret int32
		return ret
	}
	return *o.GitReleaseId.Get()
}

// GetGitReleaseIdOk returns a tuple with the GitReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitReleaseIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitReleaseId.Get(), o.GitReleaseId.IsSet()
}

// HasGitReleaseId returns a boolean if a field has been set.
func (o *Run) HasGitReleaseId() bool {
	if o != nil && o.GitReleaseId.IsSet() {
		return true
	}

	return false
}

// SetGitReleaseId gets a reference to the given NullableInt32 and assigns it to the GitReleaseId field.
func (o *Run) SetGitReleaseId(v int32) {
	o.GitReleaseId.Set(&v)
}
// SetGitReleaseIdNil sets the value for GitReleaseId to be an explicit nil
func (o *Run) SetGitReleaseIdNil() {
	o.GitReleaseId.Set(nil)
}

// UnsetGitReleaseId ensures that no value is present for GitReleaseId, not even an explicit nil
func (o *Run) UnsetGitReleaseId() {
	o.GitReleaseId.Unset()
}

// GetGitReleaseName returns the GitReleaseName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitReleaseName() string {
	if o == nil || IsNil(o.GitReleaseName.Get()) {
		var ret string
		return ret
	}
	return *o.GitReleaseName.Get()
}

// GetGitReleaseNameOk returns a tuple with the GitReleaseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitReleaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitReleaseName.Get(), o.GitReleaseName.IsSet()
}

// HasGitReleaseName returns a boolean if a field has been set.
func (o *Run) HasGitReleaseName() bool {
	if o != nil && o.GitReleaseName.IsSet() {
		return true
	}

	return false
}

// SetGitReleaseName gets a reference to the given NullableString and assigns it to the GitReleaseName field.
func (o *Run) SetGitReleaseName(v string) {
	o.GitReleaseName.Set(&v)
}
// SetGitReleaseNameNil sets the value for GitReleaseName to be an explicit nil
func (o *Run) SetGitReleaseNameNil() {
	o.GitReleaseName.Set(nil)
}

// UnsetGitReleaseName ensures that no value is present for GitReleaseName, not even an explicit nil
func (o *Run) UnsetGitReleaseName() {
	o.GitReleaseName.Unset()
}

// GetGitReleaseState returns the GitReleaseState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetGitReleaseState() int32 {
	if o == nil || IsNil(o.GitReleaseState.Get()) {
		var ret int32
		return ret
	}
	return *o.GitReleaseState.Get()
}

// GetGitReleaseStateOk returns a tuple with the GitReleaseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetGitReleaseStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitReleaseState.Get(), o.GitReleaseState.IsSet()
}

// HasGitReleaseState returns a boolean if a field has been set.
func (o *Run) HasGitReleaseState() bool {
	if o != nil && o.GitReleaseState.IsSet() {
		return true
	}

	return false
}

// SetGitReleaseState gets a reference to the given NullableInt32 and assigns it to the GitReleaseState field.
func (o *Run) SetGitReleaseState(v int32) {
	o.GitReleaseState.Set(&v)
}
// SetGitReleaseStateNil sets the value for GitReleaseState to be an explicit nil
func (o *Run) SetGitReleaseStateNil() {
	o.GitReleaseState.Set(nil)
}

// UnsetGitReleaseState ensures that no value is present for GitReleaseState, not even an explicit nil
func (o *Run) UnsetGitReleaseState() {
	o.GitReleaseState.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Run) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Run) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Run) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Run) UnsetId() {
	o.Id.Unset()
}

// GetIncludeAll returns the IncludeAll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIncludeAll() bool {
	if o == nil || IsNil(o.IncludeAll.Get()) {
		var ret bool
		return ret
	}
	return *o.IncludeAll.Get()
}

// GetIncludeAllOk returns a tuple with the IncludeAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIncludeAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludeAll.Get(), o.IncludeAll.IsSet()
}

// HasIncludeAll returns a boolean if a field has been set.
func (o *Run) HasIncludeAll() bool {
	if o != nil && o.IncludeAll.IsSet() {
		return true
	}

	return false
}

// SetIncludeAll gets a reference to the given NullableBool and assigns it to the IncludeAll field.
func (o *Run) SetIncludeAll(v bool) {
	o.IncludeAll.Set(&v)
}
// SetIncludeAllNil sets the value for IncludeAll to be an explicit nil
func (o *Run) SetIncludeAllNil() {
	o.IncludeAll.Set(nil)
}

// UnsetIncludeAll ensures that no value is present for IncludeAll, not even an explicit nil
func (o *Run) UnsetIncludeAll() {
	o.IncludeAll.Unset()
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCompleted.Get()
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIsCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCompleted.Get(), o.IsCompleted.IsSet()
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *Run) HasIsCompleted() bool {
	if o != nil && o.IsCompleted.IsSet() {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given NullableBool and assigns it to the IsCompleted field.
func (o *Run) SetIsCompleted(v bool) {
	o.IsCompleted.Set(&v)
}
// SetIsCompletedNil sets the value for IsCompleted to be an explicit nil
func (o *Run) SetIsCompletedNil() {
	o.IsCompleted.Set(nil)
}

// UnsetIsCompleted ensures that no value is present for IsCompleted, not even an explicit nil
func (o *Run) UnsetIsCompleted() {
	o.IsCompleted.Unset()
}

// GetIterationId returns the IterationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId.Get()) {
		var ret int64
		return ret
	}
	return *o.IterationId.Get()
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIterationIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationId.Get(), o.IterationId.IsSet()
}

// HasIterationId returns a boolean if a field has been set.
func (o *Run) HasIterationId() bool {
	if o != nil && o.IterationId.IsSet() {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given NullableInt64 and assigns it to the IterationId field.
func (o *Run) SetIterationId(v int64) {
	o.IterationId.Set(&v)
}
// SetIterationIdNil sets the value for IterationId to be an explicit nil
func (o *Run) SetIterationIdNil() {
	o.IterationId.Set(nil)
}

// UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
func (o *Run) UnsetIterationId() {
	o.IterationId.Unset()
}

// GetIterationName returns the IterationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetIterationName() string {
	if o == nil || IsNil(o.IterationName.Get()) {
		var ret string
		return ret
	}
	return *o.IterationName.Get()
}

// GetIterationNameOk returns a tuple with the IterationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetIterationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationName.Get(), o.IterationName.IsSet()
}

// HasIterationName returns a boolean if a field has been set.
func (o *Run) HasIterationName() bool {
	if o != nil && o.IterationName.IsSet() {
		return true
	}

	return false
}

// SetIterationName gets a reference to the given NullableString and assigns it to the IterationName field.
func (o *Run) SetIterationName(v string) {
	o.IterationName.Set(&v)
}
// SetIterationNameNil sets the value for IterationName to be an explicit nil
func (o *Run) SetIterationNameNil() {
	o.IterationName.Set(nil)
}

// UnsetIterationName ensures that no value is present for IterationName, not even an explicit nil
func (o *Run) UnsetIterationName() {
	o.IterationName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Run) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Run) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Run) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Run) UnsetName() {
	o.Name.Unset()
}

// GetPassedCount returns the PassedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetPassedCount() int32 {
	if o == nil || IsNil(o.PassedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PassedCount.Get()
}

// GetPassedCountOk returns a tuple with the PassedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetPassedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassedCount.Get(), o.PassedCount.IsSet()
}

// HasPassedCount returns a boolean if a field has been set.
func (o *Run) HasPassedCount() bool {
	if o != nil && o.PassedCount.IsSet() {
		return true
	}

	return false
}

// SetPassedCount gets a reference to the given NullableInt32 and assigns it to the PassedCount field.
func (o *Run) SetPassedCount(v int32) {
	o.PassedCount.Set(&v)
}
// SetPassedCountNil sets the value for PassedCount to be an explicit nil
func (o *Run) SetPassedCountNil() {
	o.PassedCount.Set(nil)
}

// UnsetPassedCount ensures that no value is present for PassedCount, not even an explicit nil
func (o *Run) UnsetPassedCount() {
	o.PassedCount.Unset()
}

// GetRetestCount returns the RetestCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetRetestCount() int32 {
	if o == nil || IsNil(o.RetestCount.Get()) {
		var ret int32
		return ret
	}
	return *o.RetestCount.Get()
}

// GetRetestCountOk returns a tuple with the RetestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetRetestCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetestCount.Get(), o.RetestCount.IsSet()
}

// HasRetestCount returns a boolean if a field has been set.
func (o *Run) HasRetestCount() bool {
	if o != nil && o.RetestCount.IsSet() {
		return true
	}

	return false
}

// SetRetestCount gets a reference to the given NullableInt32 and assigns it to the RetestCount field.
func (o *Run) SetRetestCount(v int32) {
	o.RetestCount.Set(&v)
}
// SetRetestCountNil sets the value for RetestCount to be an explicit nil
func (o *Run) SetRetestCountNil() {
	o.RetestCount.Set(nil)
}

// UnsetRetestCount ensures that no value is present for RetestCount, not even an explicit nil
func (o *Run) UnsetRetestCount() {
	o.RetestCount.Unset()
}

// GetSectionId returns the SectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetSectionId() int64 {
	if o == nil || IsNil(o.SectionId.Get()) {
		var ret int64
		return ret
	}
	return *o.SectionId.Get()
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetSectionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionId.Get(), o.SectionId.IsSet()
}

// HasSectionId returns a boolean if a field has been set.
func (o *Run) HasSectionId() bool {
	if o != nil && o.SectionId.IsSet() {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given NullableInt64 and assigns it to the SectionId field.
func (o *Run) SetSectionId(v int64) {
	o.SectionId.Set(&v)
}
// SetSectionIdNil sets the value for SectionId to be an explicit nil
func (o *Run) SetSectionIdNil() {
	o.SectionId.Set(nil)
}

// UnsetSectionId ensures that no value is present for SectionId, not even an explicit nil
func (o *Run) UnsetSectionId() {
	o.SectionId.Unset()
}

// GetSectionName returns the SectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetSectionName() string {
	if o == nil || IsNil(o.SectionName.Get()) {
		var ret string
		return ret
	}
	return *o.SectionName.Get()
}

// GetSectionNameOk returns a tuple with the SectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetSectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SectionName.Get(), o.SectionName.IsSet()
}

// HasSectionName returns a boolean if a field has been set.
func (o *Run) HasSectionName() bool {
	if o != nil && o.SectionName.IsSet() {
		return true
	}

	return false
}

// SetSectionName gets a reference to the given NullableString and assigns it to the SectionName field.
func (o *Run) SetSectionName(v string) {
	o.SectionName.Set(&v)
}
// SetSectionNameNil sets the value for SectionName to be an explicit nil
func (o *Run) SetSectionNameNil() {
	o.SectionName.Set(nil)
}

// UnsetSectionName ensures that no value is present for SectionName, not even an explicit nil
func (o *Run) UnsetSectionName() {
	o.SectionName.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetState() int32 {
	if o == nil || IsNil(o.State.Get()) {
		var ret int32
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetStateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *Run) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableInt32 and assigns it to the State field.
func (o *Run) SetState(v int32) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *Run) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *Run) UnsetState() {
	o.State.Unset()
}

// GetUntestedCount returns the UntestedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Run) GetUntestedCount() int32 {
	if o == nil || IsNil(o.UntestedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UntestedCount.Get()
}

// GetUntestedCountOk returns a tuple with the UntestedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Run) GetUntestedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UntestedCount.Get(), o.UntestedCount.IsSet()
}

// HasUntestedCount returns a boolean if a field has been set.
func (o *Run) HasUntestedCount() bool {
	if o != nil && o.UntestedCount.IsSet() {
		return true
	}

	return false
}

// SetUntestedCount gets a reference to the given NullableInt32 and assigns it to the UntestedCount field.
func (o *Run) SetUntestedCount(v int32) {
	o.UntestedCount.Set(&v)
}
// SetUntestedCountNil sets the value for UntestedCount to be an explicit nil
func (o *Run) SetUntestedCountNil() {
	o.UntestedCount.Set(nil)
}

// UnsetUntestedCount ensures that no value is present for UntestedCount, not even an explicit nil
func (o *Run) UnsetUntestedCount() {
	o.UntestedCount.Unset()
}

func (o Run) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Run) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AssignedToId.IsSet() {
		toSerialize["AssignedToId"] = o.AssignedToId.Get()
	}
	if o.BlockedCount.IsSet() {
		toSerialize["BlockedCount"] = o.BlockedCount.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["CompletedAt"] = o.CompletedAt.Get()
	}
	if o.ConfigEnvironmentId.IsSet() {
		toSerialize["ConfigEnvironmentId"] = o.ConfigEnvironmentId.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["CreatedBy"] = o.CreatedBy.Get()
	}
	if o.Days.IsSet() {
		toSerialize["Days"] = o.Days.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.ExecuteType.IsSet() {
		toSerialize["ExecuteType"] = o.ExecuteType.Get()
	}
	if o.FailedCount.IsSet() {
		toSerialize["FailedCount"] = o.FailedCount.Get()
	}
	if o.GitDepotId.IsSet() {
		toSerialize["GitDepotId"] = o.GitDepotId.Get()
	}
	if o.GitDepotName.IsSet() {
		toSerialize["GitDepotName"] = o.GitDepotName.Get()
	}
	if o.GitReleaseId.IsSet() {
		toSerialize["GitReleaseId"] = o.GitReleaseId.Get()
	}
	if o.GitReleaseName.IsSet() {
		toSerialize["GitReleaseName"] = o.GitReleaseName.Get()
	}
	if o.GitReleaseState.IsSet() {
		toSerialize["GitReleaseState"] = o.GitReleaseState.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.IncludeAll.IsSet() {
		toSerialize["IncludeAll"] = o.IncludeAll.Get()
	}
	if o.IsCompleted.IsSet() {
		toSerialize["IsCompleted"] = o.IsCompleted.Get()
	}
	if o.IterationId.IsSet() {
		toSerialize["IterationId"] = o.IterationId.Get()
	}
	if o.IterationName.IsSet() {
		toSerialize["IterationName"] = o.IterationName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.PassedCount.IsSet() {
		toSerialize["PassedCount"] = o.PassedCount.Get()
	}
	if o.RetestCount.IsSet() {
		toSerialize["RetestCount"] = o.RetestCount.Get()
	}
	if o.SectionId.IsSet() {
		toSerialize["SectionId"] = o.SectionId.Get()
	}
	if o.SectionName.IsSet() {
		toSerialize["SectionName"] = o.SectionName.Get()
	}
	if o.State.IsSet() {
		toSerialize["State"] = o.State.Get()
	}
	if o.UntestedCount.IsSet() {
		toSerialize["UntestedCount"] = o.UntestedCount.Get()
	}
	return toSerialize, nil
}

type NullableRun struct {
	value *Run
	isSet bool
}

func (v NullableRun) Get() *Run {
	return v.value
}

func (v *NullableRun) Set(val *Run) {
	v.value = val
	v.isSet = true
}

func (v NullableRun) IsSet() bool {
	return v.isSet
}

func (v *NullableRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRun(val *Run) *NullableRun {
	return &NullableRun{value: val, isSet: true}
}

func (v NullableRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


