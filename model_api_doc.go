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

// checks if the APIDoc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIDoc{}

// APIDoc API 文档实体
type APIDoc struct {
	// 接口数量
	APICount *int32 `json:"APICount,omitempty"`
	// API 文档编号
	Code *int32 `json:"Code,omitempty"`
	// 创建时间
	CreatedAt *int32 `json:"CreatedAt,omitempty"`
	// 创建者
	CreatedBy *int32 `json:"CreatedBy,omitempty"`
	// API 文档描述
	Description *string `json:"Description,omitempty"`
	// API 文档简介
	Introduction *string `json:"Introduction,omitempty"`
	// 上次发布编号，编号为 0 代表没有发布过
	LatestReleaseCode *int32 `json:"LatestReleaseCode,omitempty"`
	// 上次发布状态
	LatestReleaseStatus *string `json:"LatestReleaseStatus,omitempty"`
	// 上次发布的文档类型
	LatestSourceType *string `json:"LatestSourceType,omitempty"`
	// 上次发布时间
	LatestUpdatedAt NullableInt32 `json:"LatestUpdatedAt,omitempty"`
	// API 文档路由前缀
	PrefixUri *string `json:"PrefixUri,omitempty"`
	// API 文档是否私有访问
	PrivateAccess *bool `json:"PrivateAccess,omitempty"`
	// 发布次数
	ReleaseCount *int32 `json:"ReleaseCount,omitempty"`
	// 分享密码
	SharePassword NullableString `json:"SharePassword,omitempty"`
	// API 文档标题
	Title *string `json:"Title,omitempty"`
	// 更新时间
	UpdatedAt *int32 `json:"UpdatedAt,omitempty"`
	// API 文档路由
	Uri *string `json:"Uri,omitempty"`
	// API 文档是否使用文件中的描述
	UseFileDes *bool `json:"UseFileDes,omitempty"`
	// API 文档是否使用文件中的配置
	UseFileSetting *bool `json:"UseFileSetting,omitempty"`
	// 查看次数
	ViewCount *int32 `json:"ViewCount,omitempty"`
}

// NewAPIDoc instantiates a new APIDoc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIDoc() *APIDoc {
	this := APIDoc{}
	var description string = ""
	this.Description = &description
	var introduction string = ""
	this.Introduction = &introduction
	var latestReleaseStatus string = ""
	this.LatestReleaseStatus = &latestReleaseStatus
	var latestSourceType string = ""
	this.LatestSourceType = &latestSourceType
	var prefixUri string = ""
	this.PrefixUri = &prefixUri
	var privateAccess bool = false
	this.PrivateAccess = &privateAccess
	var sharePassword string = ""
	this.SharePassword = *NewNullableString(&sharePassword)
	var title string = ""
	this.Title = &title
	var uri string = ""
	this.Uri = &uri
	var useFileDes bool = false
	this.UseFileDes = &useFileDes
	var useFileSetting bool = false
	this.UseFileSetting = &useFileSetting
	return &this
}

// NewAPIDocWithDefaults instantiates a new APIDoc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIDocWithDefaults() *APIDoc {
	this := APIDoc{}
	var description string = ""
	this.Description = &description
	var introduction string = ""
	this.Introduction = &introduction
	var latestReleaseStatus string = ""
	this.LatestReleaseStatus = &latestReleaseStatus
	var latestSourceType string = ""
	this.LatestSourceType = &latestSourceType
	var prefixUri string = ""
	this.PrefixUri = &prefixUri
	var privateAccess bool = false
	this.PrivateAccess = &privateAccess
	var sharePassword string = ""
	this.SharePassword = *NewNullableString(&sharePassword)
	var title string = ""
	this.Title = &title
	var uri string = ""
	this.Uri = &uri
	var useFileDes bool = false
	this.UseFileDes = &useFileDes
	var useFileSetting bool = false
	this.UseFileSetting = &useFileSetting
	return &this
}

// GetAPICount returns the APICount field value if set, zero value otherwise.
func (o *APIDoc) GetAPICount() int32 {
	if o == nil || IsNil(o.APICount) {
		var ret int32
		return ret
	}
	return *o.APICount
}

// GetAPICountOk returns a tuple with the APICount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetAPICountOk() (*int32, bool) {
	if o == nil || IsNil(o.APICount) {
		return nil, false
	}
	return o.APICount, true
}

// HasAPICount returns a boolean if a field has been set.
func (o *APIDoc) HasAPICount() bool {
	if o != nil && !IsNil(o.APICount) {
		return true
	}

	return false
}

// SetAPICount gets a reference to the given int32 and assigns it to the APICount field.
func (o *APIDoc) SetAPICount(v int32) {
	o.APICount = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *APIDoc) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *APIDoc) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *APIDoc) SetCode(v int32) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *APIDoc) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *APIDoc) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *APIDoc) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *APIDoc) GetCreatedBy() int32 {
	if o == nil || IsNil(o.CreatedBy) {
		var ret int32
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetCreatedByOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *APIDoc) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.
func (o *APIDoc) SetCreatedBy(v int32) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *APIDoc) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *APIDoc) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *APIDoc) SetDescription(v string) {
	o.Description = &v
}

// GetIntroduction returns the Introduction field value if set, zero value otherwise.
func (o *APIDoc) GetIntroduction() string {
	if o == nil || IsNil(o.Introduction) {
		var ret string
		return ret
	}
	return *o.Introduction
}

// GetIntroductionOk returns a tuple with the Introduction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetIntroductionOk() (*string, bool) {
	if o == nil || IsNil(o.Introduction) {
		return nil, false
	}
	return o.Introduction, true
}

// HasIntroduction returns a boolean if a field has been set.
func (o *APIDoc) HasIntroduction() bool {
	if o != nil && !IsNil(o.Introduction) {
		return true
	}

	return false
}

// SetIntroduction gets a reference to the given string and assigns it to the Introduction field.
func (o *APIDoc) SetIntroduction(v string) {
	o.Introduction = &v
}

// GetLatestReleaseCode returns the LatestReleaseCode field value if set, zero value otherwise.
func (o *APIDoc) GetLatestReleaseCode() int32 {
	if o == nil || IsNil(o.LatestReleaseCode) {
		var ret int32
		return ret
	}
	return *o.LatestReleaseCode
}

// GetLatestReleaseCodeOk returns a tuple with the LatestReleaseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetLatestReleaseCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.LatestReleaseCode) {
		return nil, false
	}
	return o.LatestReleaseCode, true
}

// HasLatestReleaseCode returns a boolean if a field has been set.
func (o *APIDoc) HasLatestReleaseCode() bool {
	if o != nil && !IsNil(o.LatestReleaseCode) {
		return true
	}

	return false
}

// SetLatestReleaseCode gets a reference to the given int32 and assigns it to the LatestReleaseCode field.
func (o *APIDoc) SetLatestReleaseCode(v int32) {
	o.LatestReleaseCode = &v
}

// GetLatestReleaseStatus returns the LatestReleaseStatus field value if set, zero value otherwise.
func (o *APIDoc) GetLatestReleaseStatus() string {
	if o == nil || IsNil(o.LatestReleaseStatus) {
		var ret string
		return ret
	}
	return *o.LatestReleaseStatus
}

// GetLatestReleaseStatusOk returns a tuple with the LatestReleaseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetLatestReleaseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LatestReleaseStatus) {
		return nil, false
	}
	return o.LatestReleaseStatus, true
}

// HasLatestReleaseStatus returns a boolean if a field has been set.
func (o *APIDoc) HasLatestReleaseStatus() bool {
	if o != nil && !IsNil(o.LatestReleaseStatus) {
		return true
	}

	return false
}

// SetLatestReleaseStatus gets a reference to the given string and assigns it to the LatestReleaseStatus field.
func (o *APIDoc) SetLatestReleaseStatus(v string) {
	o.LatestReleaseStatus = &v
}

// GetLatestSourceType returns the LatestSourceType field value if set, zero value otherwise.
func (o *APIDoc) GetLatestSourceType() string {
	if o == nil || IsNil(o.LatestSourceType) {
		var ret string
		return ret
	}
	return *o.LatestSourceType
}

// GetLatestSourceTypeOk returns a tuple with the LatestSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetLatestSourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LatestSourceType) {
		return nil, false
	}
	return o.LatestSourceType, true
}

// HasLatestSourceType returns a boolean if a field has been set.
func (o *APIDoc) HasLatestSourceType() bool {
	if o != nil && !IsNil(o.LatestSourceType) {
		return true
	}

	return false
}

// SetLatestSourceType gets a reference to the given string and assigns it to the LatestSourceType field.
func (o *APIDoc) SetLatestSourceType(v string) {
	o.LatestSourceType = &v
}

// GetLatestUpdatedAt returns the LatestUpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *APIDoc) GetLatestUpdatedAt() int32 {
	if o == nil || IsNil(o.LatestUpdatedAt.Get()) {
		var ret int32
		return ret
	}
	return *o.LatestUpdatedAt.Get()
}

// GetLatestUpdatedAtOk returns a tuple with the LatestUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *APIDoc) GetLatestUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestUpdatedAt.Get(), o.LatestUpdatedAt.IsSet()
}

// HasLatestUpdatedAt returns a boolean if a field has been set.
func (o *APIDoc) HasLatestUpdatedAt() bool {
	if o != nil && o.LatestUpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetLatestUpdatedAt gets a reference to the given NullableInt32 and assigns it to the LatestUpdatedAt field.
func (o *APIDoc) SetLatestUpdatedAt(v int32) {
	o.LatestUpdatedAt.Set(&v)
}
// SetLatestUpdatedAtNil sets the value for LatestUpdatedAt to be an explicit nil
func (o *APIDoc) SetLatestUpdatedAtNil() {
	o.LatestUpdatedAt.Set(nil)
}

// UnsetLatestUpdatedAt ensures that no value is present for LatestUpdatedAt, not even an explicit nil
func (o *APIDoc) UnsetLatestUpdatedAt() {
	o.LatestUpdatedAt.Unset()
}

// GetPrefixUri returns the PrefixUri field value if set, zero value otherwise.
func (o *APIDoc) GetPrefixUri() string {
	if o == nil || IsNil(o.PrefixUri) {
		var ret string
		return ret
	}
	return *o.PrefixUri
}

// GetPrefixUriOk returns a tuple with the PrefixUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetPrefixUriOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixUri) {
		return nil, false
	}
	return o.PrefixUri, true
}

// HasPrefixUri returns a boolean if a field has been set.
func (o *APIDoc) HasPrefixUri() bool {
	if o != nil && !IsNil(o.PrefixUri) {
		return true
	}

	return false
}

// SetPrefixUri gets a reference to the given string and assigns it to the PrefixUri field.
func (o *APIDoc) SetPrefixUri(v string) {
	o.PrefixUri = &v
}

// GetPrivateAccess returns the PrivateAccess field value if set, zero value otherwise.
func (o *APIDoc) GetPrivateAccess() bool {
	if o == nil || IsNil(o.PrivateAccess) {
		var ret bool
		return ret
	}
	return *o.PrivateAccess
}

// GetPrivateAccessOk returns a tuple with the PrivateAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetPrivateAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateAccess) {
		return nil, false
	}
	return o.PrivateAccess, true
}

// HasPrivateAccess returns a boolean if a field has been set.
func (o *APIDoc) HasPrivateAccess() bool {
	if o != nil && !IsNil(o.PrivateAccess) {
		return true
	}

	return false
}

// SetPrivateAccess gets a reference to the given bool and assigns it to the PrivateAccess field.
func (o *APIDoc) SetPrivateAccess(v bool) {
	o.PrivateAccess = &v
}

// GetReleaseCount returns the ReleaseCount field value if set, zero value otherwise.
func (o *APIDoc) GetReleaseCount() int32 {
	if o == nil || IsNil(o.ReleaseCount) {
		var ret int32
		return ret
	}
	return *o.ReleaseCount
}

// GetReleaseCountOk returns a tuple with the ReleaseCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetReleaseCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleaseCount) {
		return nil, false
	}
	return o.ReleaseCount, true
}

// HasReleaseCount returns a boolean if a field has been set.
func (o *APIDoc) HasReleaseCount() bool {
	if o != nil && !IsNil(o.ReleaseCount) {
		return true
	}

	return false
}

// SetReleaseCount gets a reference to the given int32 and assigns it to the ReleaseCount field.
func (o *APIDoc) SetReleaseCount(v int32) {
	o.ReleaseCount = &v
}

// GetSharePassword returns the SharePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *APIDoc) GetSharePassword() string {
	if o == nil || IsNil(o.SharePassword.Get()) {
		var ret string
		return ret
	}
	return *o.SharePassword.Get()
}

// GetSharePasswordOk returns a tuple with the SharePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *APIDoc) GetSharePasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SharePassword.Get(), o.SharePassword.IsSet()
}

// HasSharePassword returns a boolean if a field has been set.
func (o *APIDoc) HasSharePassword() bool {
	if o != nil && o.SharePassword.IsSet() {
		return true
	}

	return false
}

// SetSharePassword gets a reference to the given NullableString and assigns it to the SharePassword field.
func (o *APIDoc) SetSharePassword(v string) {
	o.SharePassword.Set(&v)
}
// SetSharePasswordNil sets the value for SharePassword to be an explicit nil
func (o *APIDoc) SetSharePasswordNil() {
	o.SharePassword.Set(nil)
}

// UnsetSharePassword ensures that no value is present for SharePassword, not even an explicit nil
func (o *APIDoc) UnsetSharePassword() {
	o.SharePassword.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *APIDoc) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *APIDoc) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *APIDoc) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *APIDoc) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *APIDoc) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *APIDoc) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *APIDoc) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *APIDoc) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *APIDoc) SetUri(v string) {
	o.Uri = &v
}

// GetUseFileDes returns the UseFileDes field value if set, zero value otherwise.
func (o *APIDoc) GetUseFileDes() bool {
	if o == nil || IsNil(o.UseFileDes) {
		var ret bool
		return ret
	}
	return *o.UseFileDes
}

// GetUseFileDesOk returns a tuple with the UseFileDes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetUseFileDesOk() (*bool, bool) {
	if o == nil || IsNil(o.UseFileDes) {
		return nil, false
	}
	return o.UseFileDes, true
}

// HasUseFileDes returns a boolean if a field has been set.
func (o *APIDoc) HasUseFileDes() bool {
	if o != nil && !IsNil(o.UseFileDes) {
		return true
	}

	return false
}

// SetUseFileDes gets a reference to the given bool and assigns it to the UseFileDes field.
func (o *APIDoc) SetUseFileDes(v bool) {
	o.UseFileDes = &v
}

// GetUseFileSetting returns the UseFileSetting field value if set, zero value otherwise.
func (o *APIDoc) GetUseFileSetting() bool {
	if o == nil || IsNil(o.UseFileSetting) {
		var ret bool
		return ret
	}
	return *o.UseFileSetting
}

// GetUseFileSettingOk returns a tuple with the UseFileSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetUseFileSettingOk() (*bool, bool) {
	if o == nil || IsNil(o.UseFileSetting) {
		return nil, false
	}
	return o.UseFileSetting, true
}

// HasUseFileSetting returns a boolean if a field has been set.
func (o *APIDoc) HasUseFileSetting() bool {
	if o != nil && !IsNil(o.UseFileSetting) {
		return true
	}

	return false
}

// SetUseFileSetting gets a reference to the given bool and assigns it to the UseFileSetting field.
func (o *APIDoc) SetUseFileSetting(v bool) {
	o.UseFileSetting = &v
}

// GetViewCount returns the ViewCount field value if set, zero value otherwise.
func (o *APIDoc) GetViewCount() int32 {
	if o == nil || IsNil(o.ViewCount) {
		var ret int32
		return ret
	}
	return *o.ViewCount
}

// GetViewCountOk returns a tuple with the ViewCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIDoc) GetViewCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ViewCount) {
		return nil, false
	}
	return o.ViewCount, true
}

// HasViewCount returns a boolean if a field has been set.
func (o *APIDoc) HasViewCount() bool {
	if o != nil && !IsNil(o.ViewCount) {
		return true
	}

	return false
}

// SetViewCount gets a reference to the given int32 and assigns it to the ViewCount field.
func (o *APIDoc) SetViewCount(v int32) {
	o.ViewCount = &v
}

func (o APIDoc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIDoc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.APICount) {
		toSerialize["APICount"] = o.APICount
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["CreatedBy"] = o.CreatedBy
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Introduction) {
		toSerialize["Introduction"] = o.Introduction
	}
	if !IsNil(o.LatestReleaseCode) {
		toSerialize["LatestReleaseCode"] = o.LatestReleaseCode
	}
	if !IsNil(o.LatestReleaseStatus) {
		toSerialize["LatestReleaseStatus"] = o.LatestReleaseStatus
	}
	if !IsNil(o.LatestSourceType) {
		toSerialize["LatestSourceType"] = o.LatestSourceType
	}
	if o.LatestUpdatedAt.IsSet() {
		toSerialize["LatestUpdatedAt"] = o.LatestUpdatedAt.Get()
	}
	if !IsNil(o.PrefixUri) {
		toSerialize["PrefixUri"] = o.PrefixUri
	}
	if !IsNil(o.PrivateAccess) {
		toSerialize["PrivateAccess"] = o.PrivateAccess
	}
	if !IsNil(o.ReleaseCount) {
		toSerialize["ReleaseCount"] = o.ReleaseCount
	}
	if o.SharePassword.IsSet() {
		toSerialize["SharePassword"] = o.SharePassword.Get()
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Uri) {
		toSerialize["Uri"] = o.Uri
	}
	if !IsNil(o.UseFileDes) {
		toSerialize["UseFileDes"] = o.UseFileDes
	}
	if !IsNil(o.UseFileSetting) {
		toSerialize["UseFileSetting"] = o.UseFileSetting
	}
	if !IsNil(o.ViewCount) {
		toSerialize["ViewCount"] = o.ViewCount
	}
	return toSerialize, nil
}

type NullableAPIDoc struct {
	value *APIDoc
	isSet bool
}

func (v NullableAPIDoc) Get() *APIDoc {
	return v.value
}

func (v *NullableAPIDoc) Set(val *APIDoc) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIDoc) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIDoc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIDoc(val *APIDoc) *NullableAPIDoc {
	return &NullableAPIDoc{value: val, isSet: true}
}

func (v NullableAPIDoc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIDoc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


