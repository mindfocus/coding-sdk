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

// checks if the MergeRequestDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeRequestDetail{}

// MergeRequestDetail 合并请求详细数据
type MergeRequestDetail struct {
	// 是否允许合并时间
	ActionAt NullableInt64 `json:"ActionAt,omitempty"`
	ActionAuthor *CodingUser `json:"ActionAuthor,omitempty"`
	Author *CodingUser `json:"Author,omitempty"`
	// 合并请求描述
	Content NullableString `json:"Content,omitempty"`
	// 合并请求描述（html 格式）
	ContentHtml NullableString `json:"ContentHtml,omitempty"`
	// 创建时间
	CreateAt NullableInt64 `json:"CreateAt,omitempty"`
	// 仓库 Id
	DepotId NullableInt64 `json:"DepotId,omitempty"`
	// 合并请求 Id
	Id NullableInt64 `json:"Id,omitempty"`
	// 合并请求 Iid
	MergeId NullableInt64 `json:"MergeId,omitempty"`
	// 合并请求状态
	MergeStatus NullableString `json:"MergeStatus,omitempty"`
	// 合并 Sha
	MergedSha NullableString `json:"MergedSha,omitempty"`
	// 源分支名
	SourceBranch NullableString `json:"SourceBranch,omitempty"`
	// 源分支 Sha
	SourceSha NullableString `json:"SourceSha,omitempty"`
	// 目标分支名
	TargetBranch NullableString `json:"TargetBranch,omitempty"`
	// 目标分支 Sha
	TargetSha NullableString `json:"TargetSha,omitempty"`
	// 合并请求标题
	Title NullableString `json:"Title,omitempty"`
	// 更新时间
	UpdateAt NullableInt64 `json:"UpdateAt,omitempty"`
}

// NewMergeRequestDetail instantiates a new MergeRequestDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeRequestDetail() *MergeRequestDetail {
	this := MergeRequestDetail{}
	var content string = ""
	this.Content = *NewNullableString(&content)
	var contentHtml string = ""
	this.ContentHtml = *NewNullableString(&contentHtml)
	var mergeStatus string = ""
	this.MergeStatus = *NewNullableString(&mergeStatus)
	var mergedSha string = ""
	this.MergedSha = *NewNullableString(&mergedSha)
	var sourceBranch string = ""
	this.SourceBranch = *NewNullableString(&sourceBranch)
	var sourceSha string = ""
	this.SourceSha = *NewNullableString(&sourceSha)
	var targetBranch string = ""
	this.TargetBranch = *NewNullableString(&targetBranch)
	var targetSha string = ""
	this.TargetSha = *NewNullableString(&targetSha)
	var title string = ""
	this.Title = *NewNullableString(&title)
	return &this
}

// NewMergeRequestDetailWithDefaults instantiates a new MergeRequestDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeRequestDetailWithDefaults() *MergeRequestDetail {
	this := MergeRequestDetail{}
	var content string = ""
	this.Content = *NewNullableString(&content)
	var contentHtml string = ""
	this.ContentHtml = *NewNullableString(&contentHtml)
	var mergeStatus string = ""
	this.MergeStatus = *NewNullableString(&mergeStatus)
	var mergedSha string = ""
	this.MergedSha = *NewNullableString(&mergedSha)
	var sourceBranch string = ""
	this.SourceBranch = *NewNullableString(&sourceBranch)
	var sourceSha string = ""
	this.SourceSha = *NewNullableString(&sourceSha)
	var targetBranch string = ""
	this.TargetBranch = *NewNullableString(&targetBranch)
	var targetSha string = ""
	this.TargetSha = *NewNullableString(&targetSha)
	var title string = ""
	this.Title = *NewNullableString(&title)
	return &this
}

// GetActionAt returns the ActionAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetActionAt() int64 {
	if o == nil || IsNil(o.ActionAt.Get()) {
		var ret int64
		return ret
	}
	return *o.ActionAt.Get()
}

// GetActionAtOk returns a tuple with the ActionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetActionAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionAt.Get(), o.ActionAt.IsSet()
}

// HasActionAt returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasActionAt() bool {
	if o != nil && o.ActionAt.IsSet() {
		return true
	}

	return false
}

// SetActionAt gets a reference to the given NullableInt64 and assigns it to the ActionAt field.
func (o *MergeRequestDetail) SetActionAt(v int64) {
	o.ActionAt.Set(&v)
}
// SetActionAtNil sets the value for ActionAt to be an explicit nil
func (o *MergeRequestDetail) SetActionAtNil() {
	o.ActionAt.Set(nil)
}

// UnsetActionAt ensures that no value is present for ActionAt, not even an explicit nil
func (o *MergeRequestDetail) UnsetActionAt() {
	o.ActionAt.Unset()
}

// GetActionAuthor returns the ActionAuthor field value if set, zero value otherwise.
func (o *MergeRequestDetail) GetActionAuthor() CodingUser {
	if o == nil || IsNil(o.ActionAuthor) {
		var ret CodingUser
		return ret
	}
	return *o.ActionAuthor
}

// GetActionAuthorOk returns a tuple with the ActionAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestDetail) GetActionAuthorOk() (*CodingUser, bool) {
	if o == nil || IsNil(o.ActionAuthor) {
		return nil, false
	}
	return o.ActionAuthor, true
}

// HasActionAuthor returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasActionAuthor() bool {
	if o != nil && !IsNil(o.ActionAuthor) {
		return true
	}

	return false
}

// SetActionAuthor gets a reference to the given CodingUser and assigns it to the ActionAuthor field.
func (o *MergeRequestDetail) SetActionAuthor(v CodingUser) {
	o.ActionAuthor = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MergeRequestDetail) GetAuthor() CodingUser {
	if o == nil || IsNil(o.Author) {
		var ret CodingUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestDetail) GetAuthorOk() (*CodingUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CodingUser and assigns it to the Author field.
func (o *MergeRequestDetail) SetAuthor(v CodingUser) {
	o.Author = &v
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetContent() string {
	if o == nil || IsNil(o.Content.Get()) {
		var ret string
		return ret
	}
	return *o.Content.Get()
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Content.Get(), o.Content.IsSet()
}

// HasContent returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasContent() bool {
	if o != nil && o.Content.IsSet() {
		return true
	}

	return false
}

// SetContent gets a reference to the given NullableString and assigns it to the Content field.
func (o *MergeRequestDetail) SetContent(v string) {
	o.Content.Set(&v)
}
// SetContentNil sets the value for Content to be an explicit nil
func (o *MergeRequestDetail) SetContentNil() {
	o.Content.Set(nil)
}

// UnsetContent ensures that no value is present for Content, not even an explicit nil
func (o *MergeRequestDetail) UnsetContent() {
	o.Content.Unset()
}

// GetContentHtml returns the ContentHtml field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetContentHtml() string {
	if o == nil || IsNil(o.ContentHtml.Get()) {
		var ret string
		return ret
	}
	return *o.ContentHtml.Get()
}

// GetContentHtmlOk returns a tuple with the ContentHtml field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetContentHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentHtml.Get(), o.ContentHtml.IsSet()
}

// HasContentHtml returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasContentHtml() bool {
	if o != nil && o.ContentHtml.IsSet() {
		return true
	}

	return false
}

// SetContentHtml gets a reference to the given NullableString and assigns it to the ContentHtml field.
func (o *MergeRequestDetail) SetContentHtml(v string) {
	o.ContentHtml.Set(&v)
}
// SetContentHtmlNil sets the value for ContentHtml to be an explicit nil
func (o *MergeRequestDetail) SetContentHtmlNil() {
	o.ContentHtml.Set(nil)
}

// UnsetContentHtml ensures that no value is present for ContentHtml, not even an explicit nil
func (o *MergeRequestDetail) UnsetContentHtml() {
	o.ContentHtml.Unset()
}

// GetCreateAt returns the CreateAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetCreateAt() int64 {
	if o == nil || IsNil(o.CreateAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreateAt.Get()
}

// GetCreateAtOk returns a tuple with the CreateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetCreateAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateAt.Get(), o.CreateAt.IsSet()
}

// HasCreateAt returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasCreateAt() bool {
	if o != nil && o.CreateAt.IsSet() {
		return true
	}

	return false
}

// SetCreateAt gets a reference to the given NullableInt64 and assigns it to the CreateAt field.
func (o *MergeRequestDetail) SetCreateAt(v int64) {
	o.CreateAt.Set(&v)
}
// SetCreateAtNil sets the value for CreateAt to be an explicit nil
func (o *MergeRequestDetail) SetCreateAtNil() {
	o.CreateAt.Set(nil)
}

// UnsetCreateAt ensures that no value is present for CreateAt, not even an explicit nil
func (o *MergeRequestDetail) UnsetCreateAt() {
	o.CreateAt.Unset()
}

// GetDepotId returns the DepotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetDepotId() int64 {
	if o == nil || IsNil(o.DepotId.Get()) {
		var ret int64
		return ret
	}
	return *o.DepotId.Get()
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotId.Get(), o.DepotId.IsSet()
}

// HasDepotId returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasDepotId() bool {
	if o != nil && o.DepotId.IsSet() {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given NullableInt64 and assigns it to the DepotId field.
func (o *MergeRequestDetail) SetDepotId(v int64) {
	o.DepotId.Set(&v)
}
// SetDepotIdNil sets the value for DepotId to be an explicit nil
func (o *MergeRequestDetail) SetDepotIdNil() {
	o.DepotId.Set(nil)
}

// UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
func (o *MergeRequestDetail) UnsetDepotId() {
	o.DepotId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *MergeRequestDetail) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MergeRequestDetail) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MergeRequestDetail) UnsetId() {
	o.Id.Unset()
}

// GetMergeId returns the MergeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetMergeId() int64 {
	if o == nil || IsNil(o.MergeId.Get()) {
		var ret int64
		return ret
	}
	return *o.MergeId.Get()
}

// GetMergeIdOk returns a tuple with the MergeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetMergeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergeId.Get(), o.MergeId.IsSet()
}

// HasMergeId returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasMergeId() bool {
	if o != nil && o.MergeId.IsSet() {
		return true
	}

	return false
}

// SetMergeId gets a reference to the given NullableInt64 and assigns it to the MergeId field.
func (o *MergeRequestDetail) SetMergeId(v int64) {
	o.MergeId.Set(&v)
}
// SetMergeIdNil sets the value for MergeId to be an explicit nil
func (o *MergeRequestDetail) SetMergeIdNil() {
	o.MergeId.Set(nil)
}

// UnsetMergeId ensures that no value is present for MergeId, not even an explicit nil
func (o *MergeRequestDetail) UnsetMergeId() {
	o.MergeId.Unset()
}

// GetMergeStatus returns the MergeStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetMergeStatus() string {
	if o == nil || IsNil(o.MergeStatus.Get()) {
		var ret string
		return ret
	}
	return *o.MergeStatus.Get()
}

// GetMergeStatusOk returns a tuple with the MergeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetMergeStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergeStatus.Get(), o.MergeStatus.IsSet()
}

// HasMergeStatus returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasMergeStatus() bool {
	if o != nil && o.MergeStatus.IsSet() {
		return true
	}

	return false
}

// SetMergeStatus gets a reference to the given NullableString and assigns it to the MergeStatus field.
func (o *MergeRequestDetail) SetMergeStatus(v string) {
	o.MergeStatus.Set(&v)
}
// SetMergeStatusNil sets the value for MergeStatus to be an explicit nil
func (o *MergeRequestDetail) SetMergeStatusNil() {
	o.MergeStatus.Set(nil)
}

// UnsetMergeStatus ensures that no value is present for MergeStatus, not even an explicit nil
func (o *MergeRequestDetail) UnsetMergeStatus() {
	o.MergeStatus.Unset()
}

// GetMergedSha returns the MergedSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetMergedSha() string {
	if o == nil || IsNil(o.MergedSha.Get()) {
		var ret string
		return ret
	}
	return *o.MergedSha.Get()
}

// GetMergedShaOk returns a tuple with the MergedSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetMergedShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MergedSha.Get(), o.MergedSha.IsSet()
}

// HasMergedSha returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasMergedSha() bool {
	if o != nil && o.MergedSha.IsSet() {
		return true
	}

	return false
}

// SetMergedSha gets a reference to the given NullableString and assigns it to the MergedSha field.
func (o *MergeRequestDetail) SetMergedSha(v string) {
	o.MergedSha.Set(&v)
}
// SetMergedShaNil sets the value for MergedSha to be an explicit nil
func (o *MergeRequestDetail) SetMergedShaNil() {
	o.MergedSha.Set(nil)
}

// UnsetMergedSha ensures that no value is present for MergedSha, not even an explicit nil
func (o *MergeRequestDetail) UnsetMergedSha() {
	o.MergedSha.Unset()
}

// GetSourceBranch returns the SourceBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetSourceBranch() string {
	if o == nil || IsNil(o.SourceBranch.Get()) {
		var ret string
		return ret
	}
	return *o.SourceBranch.Get()
}

// GetSourceBranchOk returns a tuple with the SourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetSourceBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceBranch.Get(), o.SourceBranch.IsSet()
}

// HasSourceBranch returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasSourceBranch() bool {
	if o != nil && o.SourceBranch.IsSet() {
		return true
	}

	return false
}

// SetSourceBranch gets a reference to the given NullableString and assigns it to the SourceBranch field.
func (o *MergeRequestDetail) SetSourceBranch(v string) {
	o.SourceBranch.Set(&v)
}
// SetSourceBranchNil sets the value for SourceBranch to be an explicit nil
func (o *MergeRequestDetail) SetSourceBranchNil() {
	o.SourceBranch.Set(nil)
}

// UnsetSourceBranch ensures that no value is present for SourceBranch, not even an explicit nil
func (o *MergeRequestDetail) UnsetSourceBranch() {
	o.SourceBranch.Unset()
}

// GetSourceSha returns the SourceSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetSourceSha() string {
	if o == nil || IsNil(o.SourceSha.Get()) {
		var ret string
		return ret
	}
	return *o.SourceSha.Get()
}

// GetSourceShaOk returns a tuple with the SourceSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetSourceShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceSha.Get(), o.SourceSha.IsSet()
}

// HasSourceSha returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasSourceSha() bool {
	if o != nil && o.SourceSha.IsSet() {
		return true
	}

	return false
}

// SetSourceSha gets a reference to the given NullableString and assigns it to the SourceSha field.
func (o *MergeRequestDetail) SetSourceSha(v string) {
	o.SourceSha.Set(&v)
}
// SetSourceShaNil sets the value for SourceSha to be an explicit nil
func (o *MergeRequestDetail) SetSourceShaNil() {
	o.SourceSha.Set(nil)
}

// UnsetSourceSha ensures that no value is present for SourceSha, not even an explicit nil
func (o *MergeRequestDetail) UnsetSourceSha() {
	o.SourceSha.Unset()
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetTargetBranch() string {
	if o == nil || IsNil(o.TargetBranch.Get()) {
		var ret string
		return ret
	}
	return *o.TargetBranch.Get()
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetTargetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetBranch.Get(), o.TargetBranch.IsSet()
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasTargetBranch() bool {
	if o != nil && o.TargetBranch.IsSet() {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given NullableString and assigns it to the TargetBranch field.
func (o *MergeRequestDetail) SetTargetBranch(v string) {
	o.TargetBranch.Set(&v)
}
// SetTargetBranchNil sets the value for TargetBranch to be an explicit nil
func (o *MergeRequestDetail) SetTargetBranchNil() {
	o.TargetBranch.Set(nil)
}

// UnsetTargetBranch ensures that no value is present for TargetBranch, not even an explicit nil
func (o *MergeRequestDetail) UnsetTargetBranch() {
	o.TargetBranch.Unset()
}

// GetTargetSha returns the TargetSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetTargetSha() string {
	if o == nil || IsNil(o.TargetSha.Get()) {
		var ret string
		return ret
	}
	return *o.TargetSha.Get()
}

// GetTargetShaOk returns a tuple with the TargetSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetTargetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetSha.Get(), o.TargetSha.IsSet()
}

// HasTargetSha returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasTargetSha() bool {
	if o != nil && o.TargetSha.IsSet() {
		return true
	}

	return false
}

// SetTargetSha gets a reference to the given NullableString and assigns it to the TargetSha field.
func (o *MergeRequestDetail) SetTargetSha(v string) {
	o.TargetSha.Set(&v)
}
// SetTargetShaNil sets the value for TargetSha to be an explicit nil
func (o *MergeRequestDetail) SetTargetShaNil() {
	o.TargetSha.Set(nil)
}

// UnsetTargetSha ensures that no value is present for TargetSha, not even an explicit nil
func (o *MergeRequestDetail) UnsetTargetSha() {
	o.TargetSha.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *MergeRequestDetail) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *MergeRequestDetail) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *MergeRequestDetail) UnsetTitle() {
	o.Title.Unset()
}

// GetUpdateAt returns the UpdateAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestDetail) GetUpdateAt() int64 {
	if o == nil || IsNil(o.UpdateAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdateAt.Get()
}

// GetUpdateAtOk returns a tuple with the UpdateAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestDetail) GetUpdateAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateAt.Get(), o.UpdateAt.IsSet()
}

// HasUpdateAt returns a boolean if a field has been set.
func (o *MergeRequestDetail) HasUpdateAt() bool {
	if o != nil && o.UpdateAt.IsSet() {
		return true
	}

	return false
}

// SetUpdateAt gets a reference to the given NullableInt64 and assigns it to the UpdateAt field.
func (o *MergeRequestDetail) SetUpdateAt(v int64) {
	o.UpdateAt.Set(&v)
}
// SetUpdateAtNil sets the value for UpdateAt to be an explicit nil
func (o *MergeRequestDetail) SetUpdateAtNil() {
	o.UpdateAt.Set(nil)
}

// UnsetUpdateAt ensures that no value is present for UpdateAt, not even an explicit nil
func (o *MergeRequestDetail) UnsetUpdateAt() {
	o.UpdateAt.Unset()
}

func (o MergeRequestDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeRequestDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionAt.IsSet() {
		toSerialize["ActionAt"] = o.ActionAt.Get()
	}
	if !IsNil(o.ActionAuthor) {
		toSerialize["ActionAuthor"] = o.ActionAuthor
	}
	if !IsNil(o.Author) {
		toSerialize["Author"] = o.Author
	}
	if o.Content.IsSet() {
		toSerialize["Content"] = o.Content.Get()
	}
	if o.ContentHtml.IsSet() {
		toSerialize["ContentHtml"] = o.ContentHtml.Get()
	}
	if o.CreateAt.IsSet() {
		toSerialize["CreateAt"] = o.CreateAt.Get()
	}
	if o.DepotId.IsSet() {
		toSerialize["DepotId"] = o.DepotId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.MergeId.IsSet() {
		toSerialize["MergeId"] = o.MergeId.Get()
	}
	if o.MergeStatus.IsSet() {
		toSerialize["MergeStatus"] = o.MergeStatus.Get()
	}
	if o.MergedSha.IsSet() {
		toSerialize["MergedSha"] = o.MergedSha.Get()
	}
	if o.SourceBranch.IsSet() {
		toSerialize["SourceBranch"] = o.SourceBranch.Get()
	}
	if o.SourceSha.IsSet() {
		toSerialize["SourceSha"] = o.SourceSha.Get()
	}
	if o.TargetBranch.IsSet() {
		toSerialize["TargetBranch"] = o.TargetBranch.Get()
	}
	if o.TargetSha.IsSet() {
		toSerialize["TargetSha"] = o.TargetSha.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.UpdateAt.IsSet() {
		toSerialize["UpdateAt"] = o.UpdateAt.Get()
	}
	return toSerialize, nil
}

type NullableMergeRequestDetail struct {
	value *MergeRequestDetail
	isSet bool
}

func (v NullableMergeRequestDetail) Get() *MergeRequestDetail {
	return v.value
}

func (v *NullableMergeRequestDetail) Set(val *MergeRequestDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeRequestDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeRequestDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeRequestDetail(val *MergeRequestDetail) *NullableMergeRequestDetail {
	return &NullableMergeRequestDetail{value: val, isSet: true}
}

func (v NullableMergeRequestDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeRequestDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


