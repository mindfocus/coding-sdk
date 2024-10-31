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

// checks if the DepotInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepotInfo{}

// DepotInfo 仓库信息
type DepotInfo struct {
	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 默认分支
	DefaultBranch NullableString `json:"DefaultBranch,omitempty"`
	// 仓库描述
	Description *string `json:"Description,omitempty"`
	// 项目组Id
	GroupId NullableInt64 `json:"GroupId,omitempty"`
	// 项目名称
	GroupName NullableString `json:"GroupName,omitempty"`
	// 项目类型
	GroupType NullableString `json:"GroupType,omitempty"`
	// 仓库的https地址
	HttpsUrl *string `json:"HttpsUrl,omitempty"`
	// 仓库id
	Id *int64 `json:"Id,omitempty"`
	// 最终push时间
	LastPushAt *int64 `json:"LastPushAt,omitempty"`
	// 仓库名称
	Name *string `json:"Name,omitempty"`
	// 项目Id
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty"`
	// 仓库类型,具体值为git或者svn
	RepoType *string `json:"RepoType,omitempty"`
	// 仓库的ssh地址
	SshUrl *string `json:"SshUrl,omitempty"`
	// 仓库类型,具体值为git或者svn
	VcsType *string `json:"VcsType,omitempty"`
	// 仓库webURL
	WebUrl *string `json:"WebUrl,omitempty"`
}

// NewDepotInfo instantiates a new DepotInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepotInfo() *DepotInfo {
	this := DepotInfo{}
	var defaultBranch string = ""
	this.DefaultBranch = *NewNullableString(&defaultBranch)
	var description string = ""
	this.Description = &description
	var groupName string = ""
	this.GroupName = *NewNullableString(&groupName)
	var groupType string = ""
	this.GroupType = *NewNullableString(&groupType)
	var httpsUrl string = ""
	this.HttpsUrl = &httpsUrl
	var name string = ""
	this.Name = &name
	var projectName string = ""
	this.ProjectName = &projectName
	var repoType string = ""
	this.RepoType = &repoType
	var sshUrl string = ""
	this.SshUrl = &sshUrl
	var vcsType string = ""
	this.VcsType = &vcsType
	var webUrl string = ""
	this.WebUrl = &webUrl
	return &this
}

// NewDepotInfoWithDefaults instantiates a new DepotInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepotInfoWithDefaults() *DepotInfo {
	this := DepotInfo{}
	var defaultBranch string = ""
	this.DefaultBranch = *NewNullableString(&defaultBranch)
	var description string = ""
	this.Description = &description
	var groupName string = ""
	this.GroupName = *NewNullableString(&groupName)
	var groupType string = ""
	this.GroupType = *NewNullableString(&groupType)
	var httpsUrl string = ""
	this.HttpsUrl = &httpsUrl
	var name string = ""
	this.Name = &name
	var projectName string = ""
	this.ProjectName = &projectName
	var repoType string = ""
	this.RepoType = &repoType
	var sshUrl string = ""
	this.SshUrl = &sshUrl
	var vcsType string = ""
	this.VcsType = &vcsType
	var webUrl string = ""
	this.WebUrl = &webUrl
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DepotInfo) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DepotInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *DepotInfo) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetDefaultBranch() string {
	if o == nil || IsNil(o.DefaultBranch.Get()) {
		var ret string
		return ret
	}
	return *o.DefaultBranch.Get()
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultBranch.Get(), o.DefaultBranch.IsSet()
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *DepotInfo) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch.IsSet() {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given NullableString and assigns it to the DefaultBranch field.
func (o *DepotInfo) SetDefaultBranch(v string) {
	o.DefaultBranch.Set(&v)
}
// SetDefaultBranchNil sets the value for DefaultBranch to be an explicit nil
func (o *DepotInfo) SetDefaultBranchNil() {
	o.DefaultBranch.Set(nil)
}

// UnsetDefaultBranch ensures that no value is present for DefaultBranch, not even an explicit nil
func (o *DepotInfo) UnsetDefaultBranch() {
	o.DefaultBranch.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DepotInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DepotInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DepotInfo) SetDescription(v string) {
	o.Description = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableInt64 and assigns it to the GroupId field.
func (o *DepotInfo) SetGroupId(v int64) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *DepotInfo) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *DepotInfo) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *DepotInfo) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *DepotInfo) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *DepotInfo) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetGroupType returns the GroupType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepotInfo) GetGroupType() string {
	if o == nil || IsNil(o.GroupType.Get()) {
		var ret string
		return ret
	}
	return *o.GroupType.Get()
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepotInfo) GetGroupTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupType.Get(), o.GroupType.IsSet()
}

// HasGroupType returns a boolean if a field has been set.
func (o *DepotInfo) HasGroupType() bool {
	if o != nil && o.GroupType.IsSet() {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given NullableString and assigns it to the GroupType field.
func (o *DepotInfo) SetGroupType(v string) {
	o.GroupType.Set(&v)
}
// SetGroupTypeNil sets the value for GroupType to be an explicit nil
func (o *DepotInfo) SetGroupTypeNil() {
	o.GroupType.Set(nil)
}

// UnsetGroupType ensures that no value is present for GroupType, not even an explicit nil
func (o *DepotInfo) UnsetGroupType() {
	o.GroupType.Unset()
}

// GetHttpsUrl returns the HttpsUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetHttpsUrl() string {
	if o == nil || IsNil(o.HttpsUrl) {
		var ret string
		return ret
	}
	return *o.HttpsUrl
}

// GetHttpsUrlOk returns a tuple with the HttpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetHttpsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.HttpsUrl) {
		return nil, false
	}
	return o.HttpsUrl, true
}

// HasHttpsUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasHttpsUrl() bool {
	if o != nil && !IsNil(o.HttpsUrl) {
		return true
	}

	return false
}

// SetHttpsUrl gets a reference to the given string and assigns it to the HttpsUrl field.
func (o *DepotInfo) SetHttpsUrl(v string) {
	o.HttpsUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DepotInfo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DepotInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DepotInfo) SetId(v int64) {
	o.Id = &v
}

// GetLastPushAt returns the LastPushAt field value if set, zero value otherwise.
func (o *DepotInfo) GetLastPushAt() int64 {
	if o == nil || IsNil(o.LastPushAt) {
		var ret int64
		return ret
	}
	return *o.LastPushAt
}

// GetLastPushAtOk returns a tuple with the LastPushAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetLastPushAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastPushAt) {
		return nil, false
	}
	return o.LastPushAt, true
}

// HasLastPushAt returns a boolean if a field has been set.
func (o *DepotInfo) HasLastPushAt() bool {
	if o != nil && !IsNil(o.LastPushAt) {
		return true
	}

	return false
}

// SetLastPushAt gets a reference to the given int64 and assigns it to the LastPushAt field.
func (o *DepotInfo) SetLastPushAt(v int64) {
	o.LastPushAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DepotInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DepotInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DepotInfo) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *DepotInfo) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetProjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *DepotInfo) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *DepotInfo) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *DepotInfo) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *DepotInfo) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *DepotInfo) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetRepoType returns the RepoType field value if set, zero value otherwise.
func (o *DepotInfo) GetRepoType() string {
	if o == nil || IsNil(o.RepoType) {
		var ret string
		return ret
	}
	return *o.RepoType
}

// GetRepoTypeOk returns a tuple with the RepoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetRepoTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RepoType) {
		return nil, false
	}
	return o.RepoType, true
}

// HasRepoType returns a boolean if a field has been set.
func (o *DepotInfo) HasRepoType() bool {
	if o != nil && !IsNil(o.RepoType) {
		return true
	}

	return false
}

// SetRepoType gets a reference to the given string and assigns it to the RepoType field.
func (o *DepotInfo) SetRepoType(v string) {
	o.RepoType = &v
}

// GetSshUrl returns the SshUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetSshUrl() string {
	if o == nil || IsNil(o.SshUrl) {
		var ret string
		return ret
	}
	return *o.SshUrl
}

// GetSshUrlOk returns a tuple with the SshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetSshUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SshUrl) {
		return nil, false
	}
	return o.SshUrl, true
}

// HasSshUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasSshUrl() bool {
	if o != nil && !IsNil(o.SshUrl) {
		return true
	}

	return false
}

// SetSshUrl gets a reference to the given string and assigns it to the SshUrl field.
func (o *DepotInfo) SetSshUrl(v string) {
	o.SshUrl = &v
}

// GetVcsType returns the VcsType field value if set, zero value otherwise.
func (o *DepotInfo) GetVcsType() string {
	if o == nil || IsNil(o.VcsType) {
		var ret string
		return ret
	}
	return *o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetVcsTypeOk() (*string, bool) {
	if o == nil || IsNil(o.VcsType) {
		return nil, false
	}
	return o.VcsType, true
}

// HasVcsType returns a boolean if a field has been set.
func (o *DepotInfo) HasVcsType() bool {
	if o != nil && !IsNil(o.VcsType) {
		return true
	}

	return false
}

// SetVcsType gets a reference to the given string and assigns it to the VcsType field.
func (o *DepotInfo) SetVcsType(v string) {
	o.VcsType = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise.
func (o *DepotInfo) GetWebUrl() string {
	if o == nil || IsNil(o.WebUrl) {
		var ret string
		return ret
	}
	return *o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepotInfo) GetWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebUrl) {
		return nil, false
	}
	return o.WebUrl, true
}

// HasWebUrl returns a boolean if a field has been set.
func (o *DepotInfo) HasWebUrl() bool {
	if o != nil && !IsNil(o.WebUrl) {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given string and assigns it to the WebUrl field.
func (o *DepotInfo) SetWebUrl(v string) {
	o.WebUrl = &v
}

func (o DepotInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepotInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if o.DefaultBranch.IsSet() {
		toSerialize["DefaultBranch"] = o.DefaultBranch.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.GroupId.IsSet() {
		toSerialize["GroupId"] = o.GroupId.Get()
	}
	if o.GroupName.IsSet() {
		toSerialize["GroupName"] = o.GroupName.Get()
	}
	if o.GroupType.IsSet() {
		toSerialize["GroupType"] = o.GroupType.Get()
	}
	if !IsNil(o.HttpsUrl) {
		toSerialize["HttpsUrl"] = o.HttpsUrl
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.LastPushAt) {
		toSerialize["LastPushAt"] = o.LastPushAt
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if !IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	if !IsNil(o.RepoType) {
		toSerialize["RepoType"] = o.RepoType
	}
	if !IsNil(o.SshUrl) {
		toSerialize["SshUrl"] = o.SshUrl
	}
	if !IsNil(o.VcsType) {
		toSerialize["VcsType"] = o.VcsType
	}
	if !IsNil(o.WebUrl) {
		toSerialize["WebUrl"] = o.WebUrl
	}
	return toSerialize, nil
}

type NullableDepotInfo struct {
	value *DepotInfo
	isSet bool
}

func (v NullableDepotInfo) Get() *DepotInfo {
	return v.value
}

func (v *NullableDepotInfo) Set(val *DepotInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDepotInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDepotInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepotInfo(val *DepotInfo) *NullableDepotInfo {
	return &NullableDepotInfo{value: val, isSet: true}
}

func (v NullableDepotInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepotInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


