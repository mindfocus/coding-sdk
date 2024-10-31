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

// checks if the Report type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Report{}

// Report 测试报告信息
type Report struct {
	// 附件列表
	Attachments []Attachment `json:"Attachments,omitempty"`
	// 创建时间
	CreatedAt NullableString `json:"CreatedAt,omitempty"`
	// 创建人
	CreatedBy NullableInt32 `json:"CreatedBy,omitempty"`
	// ID 主键
	Id NullableInt32 `json:"Id,omitempty"`
	// 迭代 ID
	IterationId NullableString `json:"IterationId,omitempty"`
	// 迭代名称
	IterationName NullableString `json:"IterationName,omitempty"`
	// 报告名称
	Name NullableString `json:"Name,omitempty"`
	// 项目名称
	ProjectName NullableString `json:"ProjectName,omitempty"`
	ReportOverview *ReportOverview `json:"ReportOverview,omitempty"`
	// 测试计划 ID
	RunIds []string `json:"RunIds,omitempty"`
	// 测试计划名称
	RunNames []string `json:"RunNames,omitempty"`
	// 数据统计结束时间
	StatisticsEndTime NullableString `json:"StatisticsEndTime,omitempty"`
	// 数据统计开始时间
	StatisticsStartTime NullableString `json:"StatisticsStartTime,omitempty"`
	// 报告状态：CREATING 创建中，AVAILABLE 可用，UNAVAILABLE 不可用
	Status NullableString `json:"Status,omitempty"`
	// 报告总结
	Summary NullableString `json:"Summary,omitempty"`
	// 模板 ID
	TemplateId NullableInt32 `json:"TemplateId,omitempty"`
}

// NewReport instantiates a new Report object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReport() *Report {
	this := Report{}
	var createdAt string = ""
	this.CreatedAt = *NewNullableString(&createdAt)
	var iterationId string = ""
	this.IterationId = *NewNullableString(&iterationId)
	var iterationName string = ""
	this.IterationName = *NewNullableString(&iterationName)
	var name string = ""
	this.Name = *NewNullableString(&name)
	var projectName string = ""
	this.ProjectName = *NewNullableString(&projectName)
	var statisticsEndTime string = ""
	this.StatisticsEndTime = *NewNullableString(&statisticsEndTime)
	var statisticsStartTime string = ""
	this.StatisticsStartTime = *NewNullableString(&statisticsStartTime)
	var status string = ""
	this.Status = *NewNullableString(&status)
	var summary string = ""
	this.Summary = *NewNullableString(&summary)
	return &this
}

// NewReportWithDefaults instantiates a new Report object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportWithDefaults() *Report {
	this := Report{}
	var createdAt string = ""
	this.CreatedAt = *NewNullableString(&createdAt)
	var iterationId string = ""
	this.IterationId = *NewNullableString(&iterationId)
	var iterationName string = ""
	this.IterationName = *NewNullableString(&iterationName)
	var name string = ""
	this.Name = *NewNullableString(&name)
	var projectName string = ""
	this.ProjectName = *NewNullableString(&projectName)
	var statisticsEndTime string = ""
	this.StatisticsEndTime = *NewNullableString(&statisticsEndTime)
	var statisticsStartTime string = ""
	this.StatisticsStartTime = *NewNullableString(&statisticsStartTime)
	var status string = ""
	this.Status = *NewNullableString(&status)
	var summary string = ""
	this.Summary = *NewNullableString(&summary)
	return &this
}

// GetAttachments returns the Attachments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetAttachments() []Attachment {
	if o == nil {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Report) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *Report) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret string
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Report) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableString and assigns it to the CreatedAt field.
func (o *Report) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Report) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Report) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetCreatedBy() int32 {
	if o == nil || IsNil(o.CreatedBy.Get()) {
		var ret int32
		return ret
	}
	return *o.CreatedBy.Get()
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetCreatedByOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedBy.Get(), o.CreatedBy.IsSet()
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Report) HasCreatedBy() bool {
	if o != nil && o.CreatedBy.IsSet() {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given NullableInt32 and assigns it to the CreatedBy field.
func (o *Report) SetCreatedBy(v int32) {
	o.CreatedBy.Set(&v)
}
// SetCreatedByNil sets the value for CreatedBy to be an explicit nil
func (o *Report) SetCreatedByNil() {
	o.CreatedBy.Set(nil)
}

// UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
func (o *Report) UnsetCreatedBy() {
	o.CreatedBy.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Report) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *Report) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Report) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Report) UnsetId() {
	o.Id.Unset()
}

// GetIterationId returns the IterationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetIterationId() string {
	if o == nil || IsNil(o.IterationId.Get()) {
		var ret string
		return ret
	}
	return *o.IterationId.Get()
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetIterationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationId.Get(), o.IterationId.IsSet()
}

// HasIterationId returns a boolean if a field has been set.
func (o *Report) HasIterationId() bool {
	if o != nil && o.IterationId.IsSet() {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given NullableString and assigns it to the IterationId field.
func (o *Report) SetIterationId(v string) {
	o.IterationId.Set(&v)
}
// SetIterationIdNil sets the value for IterationId to be an explicit nil
func (o *Report) SetIterationIdNil() {
	o.IterationId.Set(nil)
}

// UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
func (o *Report) UnsetIterationId() {
	o.IterationId.Unset()
}

// GetIterationName returns the IterationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetIterationName() string {
	if o == nil || IsNil(o.IterationName.Get()) {
		var ret string
		return ret
	}
	return *o.IterationName.Get()
}

// GetIterationNameOk returns a tuple with the IterationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetIterationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IterationName.Get(), o.IterationName.IsSet()
}

// HasIterationName returns a boolean if a field has been set.
func (o *Report) HasIterationName() bool {
	if o != nil && o.IterationName.IsSet() {
		return true
	}

	return false
}

// SetIterationName gets a reference to the given NullableString and assigns it to the IterationName field.
func (o *Report) SetIterationName(v string) {
	o.IterationName.Set(&v)
}
// SetIterationNameNil sets the value for IterationName to be an explicit nil
func (o *Report) SetIterationNameNil() {
	o.IterationName.Set(nil)
}

// UnsetIterationName ensures that no value is present for IterationName, not even an explicit nil
func (o *Report) UnsetIterationName() {
	o.IterationName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Report) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Report) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Report) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Report) UnsetName() {
	o.Name.Unset()
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName.Get()) {
		var ret string
		return ret
	}
	return *o.ProjectName.Get()
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectName.Get(), o.ProjectName.IsSet()
}

// HasProjectName returns a boolean if a field has been set.
func (o *Report) HasProjectName() bool {
	if o != nil && o.ProjectName.IsSet() {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given NullableString and assigns it to the ProjectName field.
func (o *Report) SetProjectName(v string) {
	o.ProjectName.Set(&v)
}
// SetProjectNameNil sets the value for ProjectName to be an explicit nil
func (o *Report) SetProjectNameNil() {
	o.ProjectName.Set(nil)
}

// UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
func (o *Report) UnsetProjectName() {
	o.ProjectName.Unset()
}

// GetReportOverview returns the ReportOverview field value if set, zero value otherwise.
func (o *Report) GetReportOverview() ReportOverview {
	if o == nil || IsNil(o.ReportOverview) {
		var ret ReportOverview
		return ret
	}
	return *o.ReportOverview
}

// GetReportOverviewOk returns a tuple with the ReportOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Report) GetReportOverviewOk() (*ReportOverview, bool) {
	if o == nil || IsNil(o.ReportOverview) {
		return nil, false
	}
	return o.ReportOverview, true
}

// HasReportOverview returns a boolean if a field has been set.
func (o *Report) HasReportOverview() bool {
	if o != nil && !IsNil(o.ReportOverview) {
		return true
	}

	return false
}

// SetReportOverview gets a reference to the given ReportOverview and assigns it to the ReportOverview field.
func (o *Report) SetReportOverview(v ReportOverview) {
	o.ReportOverview = &v
}

// GetRunIds returns the RunIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetRunIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RunIds
}

// GetRunIdsOk returns a tuple with the RunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetRunIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.RunIds) {
		return nil, false
	}
	return o.RunIds, true
}

// HasRunIds returns a boolean if a field has been set.
func (o *Report) HasRunIds() bool {
	if o != nil && !IsNil(o.RunIds) {
		return true
	}

	return false
}

// SetRunIds gets a reference to the given []string and assigns it to the RunIds field.
func (o *Report) SetRunIds(v []string) {
	o.RunIds = v
}

// GetRunNames returns the RunNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetRunNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RunNames
}

// GetRunNamesOk returns a tuple with the RunNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetRunNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.RunNames) {
		return nil, false
	}
	return o.RunNames, true
}

// HasRunNames returns a boolean if a field has been set.
func (o *Report) HasRunNames() bool {
	if o != nil && !IsNil(o.RunNames) {
		return true
	}

	return false
}

// SetRunNames gets a reference to the given []string and assigns it to the RunNames field.
func (o *Report) SetRunNames(v []string) {
	o.RunNames = v
}

// GetStatisticsEndTime returns the StatisticsEndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetStatisticsEndTime() string {
	if o == nil || IsNil(o.StatisticsEndTime.Get()) {
		var ret string
		return ret
	}
	return *o.StatisticsEndTime.Get()
}

// GetStatisticsEndTimeOk returns a tuple with the StatisticsEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetStatisticsEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatisticsEndTime.Get(), o.StatisticsEndTime.IsSet()
}

// HasStatisticsEndTime returns a boolean if a field has been set.
func (o *Report) HasStatisticsEndTime() bool {
	if o != nil && o.StatisticsEndTime.IsSet() {
		return true
	}

	return false
}

// SetStatisticsEndTime gets a reference to the given NullableString and assigns it to the StatisticsEndTime field.
func (o *Report) SetStatisticsEndTime(v string) {
	o.StatisticsEndTime.Set(&v)
}
// SetStatisticsEndTimeNil sets the value for StatisticsEndTime to be an explicit nil
func (o *Report) SetStatisticsEndTimeNil() {
	o.StatisticsEndTime.Set(nil)
}

// UnsetStatisticsEndTime ensures that no value is present for StatisticsEndTime, not even an explicit nil
func (o *Report) UnsetStatisticsEndTime() {
	o.StatisticsEndTime.Unset()
}

// GetStatisticsStartTime returns the StatisticsStartTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetStatisticsStartTime() string {
	if o == nil || IsNil(o.StatisticsStartTime.Get()) {
		var ret string
		return ret
	}
	return *o.StatisticsStartTime.Get()
}

// GetStatisticsStartTimeOk returns a tuple with the StatisticsStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetStatisticsStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatisticsStartTime.Get(), o.StatisticsStartTime.IsSet()
}

// HasStatisticsStartTime returns a boolean if a field has been set.
func (o *Report) HasStatisticsStartTime() bool {
	if o != nil && o.StatisticsStartTime.IsSet() {
		return true
	}

	return false
}

// SetStatisticsStartTime gets a reference to the given NullableString and assigns it to the StatisticsStartTime field.
func (o *Report) SetStatisticsStartTime(v string) {
	o.StatisticsStartTime.Set(&v)
}
// SetStatisticsStartTimeNil sets the value for StatisticsStartTime to be an explicit nil
func (o *Report) SetStatisticsStartTimeNil() {
	o.StatisticsStartTime.Set(nil)
}

// UnsetStatisticsStartTime ensures that no value is present for StatisticsStartTime, not even an explicit nil
func (o *Report) UnsetStatisticsStartTime() {
	o.StatisticsStartTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Report) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Report) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Report) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Report) UnsetStatus() {
	o.Status.Unset()
}

// GetSummary returns the Summary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetSummary() string {
	if o == nil || IsNil(o.Summary.Get()) {
		var ret string
		return ret
	}
	return *o.Summary.Get()
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Summary.Get(), o.Summary.IsSet()
}

// HasSummary returns a boolean if a field has been set.
func (o *Report) HasSummary() bool {
	if o != nil && o.Summary.IsSet() {
		return true
	}

	return false
}

// SetSummary gets a reference to the given NullableString and assigns it to the Summary field.
func (o *Report) SetSummary(v string) {
	o.Summary.Set(&v)
}
// SetSummaryNil sets the value for Summary to be an explicit nil
func (o *Report) SetSummaryNil() {
	o.Summary.Set(nil)
}

// UnsetSummary ensures that no value is present for Summary, not even an explicit nil
func (o *Report) UnsetSummary() {
	o.Summary.Unset()
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Report) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId.Get()) {
		var ret int32
		return ret
	}
	return *o.TemplateId.Get()
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Report) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemplateId.Get(), o.TemplateId.IsSet()
}

// HasTemplateId returns a boolean if a field has been set.
func (o *Report) HasTemplateId() bool {
	if o != nil && o.TemplateId.IsSet() {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given NullableInt32 and assigns it to the TemplateId field.
func (o *Report) SetTemplateId(v int32) {
	o.TemplateId.Set(&v)
}
// SetTemplateIdNil sets the value for TemplateId to be an explicit nil
func (o *Report) SetTemplateIdNil() {
	o.TemplateId.Set(nil)
}

// UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
func (o *Report) UnsetTemplateId() {
	o.TemplateId.Unset()
}

func (o Report) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Report) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attachments != nil {
		toSerialize["Attachments"] = o.Attachments
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if o.CreatedBy.IsSet() {
		toSerialize["CreatedBy"] = o.CreatedBy.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
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
	if o.ProjectName.IsSet() {
		toSerialize["ProjectName"] = o.ProjectName.Get()
	}
	if !IsNil(o.ReportOverview) {
		toSerialize["ReportOverview"] = o.ReportOverview
	}
	if o.RunIds != nil {
		toSerialize["RunIds"] = o.RunIds
	}
	if o.RunNames != nil {
		toSerialize["RunNames"] = o.RunNames
	}
	if o.StatisticsEndTime.IsSet() {
		toSerialize["StatisticsEndTime"] = o.StatisticsEndTime.Get()
	}
	if o.StatisticsStartTime.IsSet() {
		toSerialize["StatisticsStartTime"] = o.StatisticsStartTime.Get()
	}
	if o.Status.IsSet() {
		toSerialize["Status"] = o.Status.Get()
	}
	if o.Summary.IsSet() {
		toSerialize["Summary"] = o.Summary.Get()
	}
	if o.TemplateId.IsSet() {
		toSerialize["TemplateId"] = o.TemplateId.Get()
	}
	return toSerialize, nil
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


