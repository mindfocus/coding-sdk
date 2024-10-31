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

// checks if the Release type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Release{}

// Release git 版本
type Release struct {
	// 内容
	Body NullableString `json:"Body,omitempty"`
	// commit Sha 值
	CommitSha NullableString `json:"CommitSha,omitempty"`
	// 创建时间
	CreatedAt NullableInt64 `json:"CreatedAt,omitempty"`
	// 创建者
	CreatorId NullableInt64 `json:"CreatorId,omitempty"`
	// 仓库 Id
	DepotId NullableInt64 `json:"DepotId,omitempty"`
	// html内容
	Html NullableString `json:"Html,omitempty"`
	// 版本 Id
	Id NullableInt64 `json:"Id,omitempty"`
	// 是否预发布
	Pre NullableBool `json:"Pre,omitempty"`
	// 项目 Id
	ProjectId NullableInt64 `json:"ProjectId,omitempty"`
	// 版本序号Id
	ReleaseId NullableInt64 `json:"ReleaseId,omitempty"`
	// 标签名字
	TagName NullableString `json:"TagName,omitempty"`
	// 目标 commit Sha 值
	TargetCommitBranch NullableString `json:"TargetCommitBranch,omitempty"`
	// 标题
	Title NullableString `json:"Title,omitempty"`
	// 更新时间
	UpdatedAt NullableInt64 `json:"UpdatedAt,omitempty"`
}

// NewRelease instantiates a new Release object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelease() *Release {
	this := Release{}
	var body string = ""
	this.Body = *NewNullableString(&body)
	var commitSha string = ""
	this.CommitSha = *NewNullableString(&commitSha)
	var html string = ""
	this.Html = *NewNullableString(&html)
	var pre bool = false
	this.Pre = *NewNullableBool(&pre)
	var tagName string = ""
	this.TagName = *NewNullableString(&tagName)
	var targetCommitBranch string = ""
	this.TargetCommitBranch = *NewNullableString(&targetCommitBranch)
	var title string = ""
	this.Title = *NewNullableString(&title)
	return &this
}

// NewReleaseWithDefaults instantiates a new Release object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseWithDefaults() *Release {
	this := Release{}
	var body string = ""
	this.Body = *NewNullableString(&body)
	var commitSha string = ""
	this.CommitSha = *NewNullableString(&commitSha)
	var html string = ""
	this.Html = *NewNullableString(&html)
	var pre bool = false
	this.Pre = *NewNullableBool(&pre)
	var tagName string = ""
	this.TagName = *NewNullableString(&tagName)
	var targetCommitBranch string = ""
	this.TargetCommitBranch = *NewNullableString(&targetCommitBranch)
	var title string = ""
	this.Title = *NewNullableString(&title)
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *Release) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *Release) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *Release) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *Release) UnsetBody() {
	o.Body.Unset()
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCommitSha() string {
	if o == nil || IsNil(o.CommitSha.Get()) {
		var ret string
		return ret
	}
	return *o.CommitSha.Get()
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitSha.Get(), o.CommitSha.IsSet()
}

// HasCommitSha returns a boolean if a field has been set.
func (o *Release) HasCommitSha() bool {
	if o != nil && o.CommitSha.IsSet() {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given NullableString and assigns it to the CommitSha field.
func (o *Release) SetCommitSha(v string) {
	o.CommitSha.Set(&v)
}
// SetCommitShaNil sets the value for CommitSha to be an explicit nil
func (o *Release) SetCommitShaNil() {
	o.CommitSha.Set(nil)
}

// UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
func (o *Release) UnsetCommitSha() {
	o.CommitSha.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Release) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *Release) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Release) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Release) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetCreatorId() int64 {
	if o == nil || IsNil(o.CreatorId.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetCreatorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// HasCreatorId returns a boolean if a field has been set.
func (o *Release) HasCreatorId() bool {
	if o != nil && o.CreatorId.IsSet() {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given NullableInt64 and assigns it to the CreatorId field.
func (o *Release) SetCreatorId(v int64) {
	o.CreatorId.Set(&v)
}
// SetCreatorIdNil sets the value for CreatorId to be an explicit nil
func (o *Release) SetCreatorIdNil() {
	o.CreatorId.Set(nil)
}

// UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
func (o *Release) UnsetCreatorId() {
	o.CreatorId.Unset()
}

// GetDepotId returns the DepotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetDepotId() int64 {
	if o == nil || IsNil(o.DepotId.Get()) {
		var ret int64
		return ret
	}
	return *o.DepotId.Get()
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotId.Get(), o.DepotId.IsSet()
}

// HasDepotId returns a boolean if a field has been set.
func (o *Release) HasDepotId() bool {
	if o != nil && o.DepotId.IsSet() {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given NullableInt64 and assigns it to the DepotId field.
func (o *Release) SetDepotId(v int64) {
	o.DepotId.Set(&v)
}
// SetDepotIdNil sets the value for DepotId to be an explicit nil
func (o *Release) SetDepotIdNil() {
	o.DepotId.Set(nil)
}

// UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
func (o *Release) UnsetDepotId() {
	o.DepotId.Unset()
}

// GetHtml returns the Html field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetHtml() string {
	if o == nil || IsNil(o.Html.Get()) {
		var ret string
		return ret
	}
	return *o.Html.Get()
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Html.Get(), o.Html.IsSet()
}

// HasHtml returns a boolean if a field has been set.
func (o *Release) HasHtml() bool {
	if o != nil && o.Html.IsSet() {
		return true
	}

	return false
}

// SetHtml gets a reference to the given NullableString and assigns it to the Html field.
func (o *Release) SetHtml(v string) {
	o.Html.Set(&v)
}
// SetHtmlNil sets the value for Html to be an explicit nil
func (o *Release) SetHtmlNil() {
	o.Html.Set(nil)
}

// UnsetHtml ensures that no value is present for Html, not even an explicit nil
func (o *Release) UnsetHtml() {
	o.Html.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Release) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Release) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Release) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Release) UnsetId() {
	o.Id.Unset()
}

// GetPre returns the Pre field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetPre() bool {
	if o == nil || IsNil(o.Pre.Get()) {
		var ret bool
		return ret
	}
	return *o.Pre.Get()
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetPreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pre.Get(), o.Pre.IsSet()
}

// HasPre returns a boolean if a field has been set.
func (o *Release) HasPre() bool {
	if o != nil && o.Pre.IsSet() {
		return true
	}

	return false
}

// SetPre gets a reference to the given NullableBool and assigns it to the Pre field.
func (o *Release) SetPre(v bool) {
	o.Pre.Set(&v)
}
// SetPreNil sets the value for Pre to be an explicit nil
func (o *Release) SetPreNil() {
	o.Pre.Set(nil)
}

// UnsetPre ensures that no value is present for Pre, not even an explicit nil
func (o *Release) UnsetPre() {
	o.Pre.Unset()
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int64
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *Release) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt64 and assigns it to the ProjectId field.
func (o *Release) SetProjectId(v int64) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *Release) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *Release) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetReleaseId() int64 {
	if o == nil || IsNil(o.ReleaseId.Get()) {
		var ret int64
		return ret
	}
	return *o.ReleaseId.Get()
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetReleaseIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseId.Get(), o.ReleaseId.IsSet()
}

// HasReleaseId returns a boolean if a field has been set.
func (o *Release) HasReleaseId() bool {
	if o != nil && o.ReleaseId.IsSet() {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given NullableInt64 and assigns it to the ReleaseId field.
func (o *Release) SetReleaseId(v int64) {
	o.ReleaseId.Set(&v)
}
// SetReleaseIdNil sets the value for ReleaseId to be an explicit nil
func (o *Release) SetReleaseIdNil() {
	o.ReleaseId.Set(nil)
}

// UnsetReleaseId ensures that no value is present for ReleaseId, not even an explicit nil
func (o *Release) UnsetReleaseId() {
	o.ReleaseId.Unset()
}

// GetTagName returns the TagName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTagName() string {
	if o == nil || IsNil(o.TagName.Get()) {
		var ret string
		return ret
	}
	return *o.TagName.Get()
}

// GetTagNameOk returns a tuple with the TagName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagName.Get(), o.TagName.IsSet()
}

// HasTagName returns a boolean if a field has been set.
func (o *Release) HasTagName() bool {
	if o != nil && o.TagName.IsSet() {
		return true
	}

	return false
}

// SetTagName gets a reference to the given NullableString and assigns it to the TagName field.
func (o *Release) SetTagName(v string) {
	o.TagName.Set(&v)
}
// SetTagNameNil sets the value for TagName to be an explicit nil
func (o *Release) SetTagNameNil() {
	o.TagName.Set(nil)
}

// UnsetTagName ensures that no value is present for TagName, not even an explicit nil
func (o *Release) UnsetTagName() {
	o.TagName.Unset()
}

// GetTargetCommitBranch returns the TargetCommitBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTargetCommitBranch() string {
	if o == nil || IsNil(o.TargetCommitBranch.Get()) {
		var ret string
		return ret
	}
	return *o.TargetCommitBranch.Get()
}

// GetTargetCommitBranchOk returns a tuple with the TargetCommitBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTargetCommitBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetCommitBranch.Get(), o.TargetCommitBranch.IsSet()
}

// HasTargetCommitBranch returns a boolean if a field has been set.
func (o *Release) HasTargetCommitBranch() bool {
	if o != nil && o.TargetCommitBranch.IsSet() {
		return true
	}

	return false
}

// SetTargetCommitBranch gets a reference to the given NullableString and assigns it to the TargetCommitBranch field.
func (o *Release) SetTargetCommitBranch(v string) {
	o.TargetCommitBranch.Set(&v)
}
// SetTargetCommitBranchNil sets the value for TargetCommitBranch to be an explicit nil
func (o *Release) SetTargetCommitBranchNil() {
	o.TargetCommitBranch.Set(nil)
}

// UnsetTargetCommitBranch ensures that no value is present for TargetCommitBranch, not even an explicit nil
func (o *Release) UnsetTargetCommitBranch() {
	o.TargetCommitBranch.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Release) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Release) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Release) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Release) UnsetTitle() {
	o.Title.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Release) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Release) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Release) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *Release) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Release) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Release) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o Release) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Release) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Body.IsSet() {
		toSerialize["Body"] = o.Body.Get()
	}
	if o.CommitSha.IsSet() {
		toSerialize["CommitSha"] = o.CommitSha.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.CreatorId.IsSet() {
		toSerialize["CreatorId"] = o.CreatorId.Get()
	}
	if o.DepotId.IsSet() {
		toSerialize["DepotId"] = o.DepotId.Get()
	}
	if o.Html.IsSet() {
		toSerialize["Html"] = o.Html.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Pre.IsSet() {
		toSerialize["Pre"] = o.Pre.Get()
	}
	if o.ProjectId.IsSet() {
		toSerialize["ProjectId"] = o.ProjectId.Get()
	}
	if o.ReleaseId.IsSet() {
		toSerialize["ReleaseId"] = o.ReleaseId.Get()
	}
	if o.TagName.IsSet() {
		toSerialize["TagName"] = o.TagName.Get()
	}
	if o.TargetCommitBranch.IsSet() {
		toSerialize["TargetCommitBranch"] = o.TargetCommitBranch.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	return toSerialize, nil
}

type NullableRelease struct {
	value *Release
	isSet bool
}

func (v NullableRelease) Get() *Release {
	return v.value
}

func (v *NullableRelease) Set(val *Release) {
	v.value = val
	v.isSet = true
}

func (v NullableRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelease(val *Release) *NullableRelease {
	return &NullableRelease{value: val, isSet: true}
}

func (v NullableRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


