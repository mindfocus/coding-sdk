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

// checks if the IssueDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueDetail{}

// IssueDetail 事项详情
type IssueDetail struct {
	Assignee *User `json:"Assignee,omitempty"`
	// 事项 Code
	Code *int64 `json:"Code,omitempty"`
	// 完成时间戳
	CompletedAt *int64 `json:"CompletedAt,omitempty"`
	// 创建时间戳
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	Creator *User `json:"Creator,omitempty"`
	// 自定义属性列表
	CustomFields []CustomFields `json:"CustomFields,omitempty"`
	DefectType *DefectType `json:"DefectType,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 截止日期时间戳
	DueDate *int64 `json:"DueDate,omitempty"`
	Epic *Epic `json:"Epic,omitempty"`
	// 附件列表
	Files []IssueFile `json:"Files,omitempty"`
	// 事项状态 Id
	IssueStatusId *int64 `json:"IssueStatusId,omitempty"`
	// 事项状态名称
	IssueStatusName *string `json:"IssueStatusName,omitempty"`
	// 事项状态类型
	IssueStatusType *string `json:"IssueStatusType,omitempty"`
	IssueTypeDetail *IssueTypeDetail `json:"IssueTypeDetail,omitempty"`
	// 事项类型 ID
	IssueTypeId *int64 `json:"IssueTypeId,omitempty"`
	Iteration *IterationSimple `json:"Iteration,omitempty"`
	// 迭代 Id
	IterationId *int64 `json:"IterationId,omitempty"`
	// 标签列表
	Labels []IssueProjectLabel `json:"Labels,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	Parent *IssueSimpleData `json:"Parent,omitempty"`
	// 父事项类型
	ParentType *string `json:"ParentType,omitempty"`
	// 优先级：  0 - 低，  1 - 中，  2 - 高，  3 - 紧急，  \"\" - 未指定
	Priority *string `json:"Priority,omitempty"`
	Project *Project `json:"Project,omitempty"`
	ProjectModule *IssueProjectModule `json:"ProjectModule,omitempty"`
	RequirementType *RequirementType `json:"RequirementType,omitempty"`
	// 开始日期时间戳
	StartDate *int64 `json:"StartDate,omitempty"`
	// 故事点，例如：0.5、0、1  空字符串 \"\" 表示未指定。
	StoryPoint *string `json:"StoryPoint,omitempty"`
	// 子工作项列表
	SubTasks []SubTask `json:"SubTasks,omitempty"`
	// 第三方链接列表
	ThirdLinks []IssueThirdLink `json:"ThirdLinks,omitempty"`
	// 事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗，  SUB_TASK - 子工作项
	Type *string `json:"Type,omitempty"`
	// 修改时间戳
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
	// 关注人列表
	Watchers []User `json:"Watchers,omitempty"`
	// 工时（小时数）
	WorkingHours *float32 `json:"WorkingHours,omitempty"`
	// 经办人列表
	Assignees []User `json:"Assignees,omitempty"`
}

// NewIssueDetail instantiates a new IssueDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueDetail() *IssueDetail {
	this := IssueDetail{}
	var description string = ""
	this.Description = &description
	var issueStatusName string = ""
	this.IssueStatusName = &issueStatusName
	var issueStatusType string = ""
	this.IssueStatusType = &issueStatusType
	var name string = ""
	this.Name = &name
	var parentType string = ""
	this.ParentType = &parentType
	var priority string = ""
	this.Priority = &priority
	var storyPoint string = ""
	this.StoryPoint = &storyPoint
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIssueDetailWithDefaults instantiates a new IssueDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueDetailWithDefaults() *IssueDetail {
	this := IssueDetail{}
	var description string = ""
	this.Description = &description
	var issueStatusName string = ""
	this.IssueStatusName = &issueStatusName
	var issueStatusType string = ""
	this.IssueStatusType = &issueStatusType
	var name string = ""
	this.Name = &name
	var parentType string = ""
	this.ParentType = &parentType
	var priority string = ""
	this.Priority = &priority
	var storyPoint string = ""
	this.StoryPoint = &storyPoint
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *IssueDetail) GetAssignee() User {
	if o == nil || IsNil(o.Assignee) {
		var ret User
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetAssigneeOk() (*User, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *IssueDetail) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given User and assigns it to the Assignee field.
func (o *IssueDetail) SetAssignee(v User) {
	o.Assignee = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IssueDetail) GetCode() int64 {
	if o == nil || IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IssueDetail) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *IssueDetail) SetCode(v int64) {
	o.Code = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise.
func (o *IssueDetail) GetCompletedAt() int64 {
	if o == nil || IsNil(o.CompletedAt) {
		var ret int64
		return ret
	}
	return *o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetCompletedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CompletedAt) {
		return nil, false
	}
	return o.CompletedAt, true
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *IssueDetail) HasCompletedAt() bool {
	if o != nil && !IsNil(o.CompletedAt) {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given int64 and assigns it to the CompletedAt field.
func (o *IssueDetail) SetCompletedAt(v int64) {
	o.CompletedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IssueDetail) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IssueDetail) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *IssueDetail) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *IssueDetail) GetCreator() User {
	if o == nil || IsNil(o.Creator) {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetCreatorOk() (*User, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *IssueDetail) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *IssueDetail) SetCreator(v User) {
	o.Creator = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IssueDetail) GetCustomFields() []CustomFields {
	if o == nil || IsNil(o.CustomFields) {
		var ret []CustomFields
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetCustomFieldsOk() ([]CustomFields, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IssueDetail) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []CustomFields and assigns it to the CustomFields field.
func (o *IssueDetail) SetCustomFields(v []CustomFields) {
	o.CustomFields = v
}

// GetDefectType returns the DefectType field value if set, zero value otherwise.
func (o *IssueDetail) GetDefectType() DefectType {
	if o == nil || IsNil(o.DefectType) {
		var ret DefectType
		return ret
	}
	return *o.DefectType
}

// GetDefectTypeOk returns a tuple with the DefectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetDefectTypeOk() (*DefectType, bool) {
	if o == nil || IsNil(o.DefectType) {
		return nil, false
	}
	return o.DefectType, true
}

// HasDefectType returns a boolean if a field has been set.
func (o *IssueDetail) HasDefectType() bool {
	if o != nil && !IsNil(o.DefectType) {
		return true
	}

	return false
}

// SetDefectType gets a reference to the given DefectType and assigns it to the DefectType field.
func (o *IssueDetail) SetDefectType(v DefectType) {
	o.DefectType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssueDetail) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueDetail) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssueDetail) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *IssueDetail) GetDueDate() int64 {
	if o == nil || IsNil(o.DueDate) {
		var ret int64
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetDueDateOk() (*int64, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *IssueDetail) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given int64 and assigns it to the DueDate field.
func (o *IssueDetail) SetDueDate(v int64) {
	o.DueDate = &v
}

// GetEpic returns the Epic field value if set, zero value otherwise.
func (o *IssueDetail) GetEpic() Epic {
	if o == nil || IsNil(o.Epic) {
		var ret Epic
		return ret
	}
	return *o.Epic
}

// GetEpicOk returns a tuple with the Epic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetEpicOk() (*Epic, bool) {
	if o == nil || IsNil(o.Epic) {
		return nil, false
	}
	return o.Epic, true
}

// HasEpic returns a boolean if a field has been set.
func (o *IssueDetail) HasEpic() bool {
	if o != nil && !IsNil(o.Epic) {
		return true
	}

	return false
}

// SetEpic gets a reference to the given Epic and assigns it to the Epic field.
func (o *IssueDetail) SetEpic(v Epic) {
	o.Epic = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *IssueDetail) GetFiles() []IssueFile {
	if o == nil || IsNil(o.Files) {
		var ret []IssueFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetFilesOk() ([]IssueFile, bool) {
	if o == nil || IsNil(o.Files) {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *IssueDetail) HasFiles() bool {
	if o != nil && !IsNil(o.Files) {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []IssueFile and assigns it to the Files field.
func (o *IssueDetail) SetFiles(v []IssueFile) {
	o.Files = v
}

// GetIssueStatusId returns the IssueStatusId field value if set, zero value otherwise.
func (o *IssueDetail) GetIssueStatusId() int64 {
	if o == nil || IsNil(o.IssueStatusId) {
		var ret int64
		return ret
	}
	return *o.IssueStatusId
}

// GetIssueStatusIdOk returns a tuple with the IssueStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIssueStatusIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IssueStatusId) {
		return nil, false
	}
	return o.IssueStatusId, true
}

// HasIssueStatusId returns a boolean if a field has been set.
func (o *IssueDetail) HasIssueStatusId() bool {
	if o != nil && !IsNil(o.IssueStatusId) {
		return true
	}

	return false
}

// SetIssueStatusId gets a reference to the given int64 and assigns it to the IssueStatusId field.
func (o *IssueDetail) SetIssueStatusId(v int64) {
	o.IssueStatusId = &v
}

// GetIssueStatusName returns the IssueStatusName field value if set, zero value otherwise.
func (o *IssueDetail) GetIssueStatusName() string {
	if o == nil || IsNil(o.IssueStatusName) {
		var ret string
		return ret
	}
	return *o.IssueStatusName
}

// GetIssueStatusNameOk returns a tuple with the IssueStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIssueStatusNameOk() (*string, bool) {
	if o == nil || IsNil(o.IssueStatusName) {
		return nil, false
	}
	return o.IssueStatusName, true
}

// HasIssueStatusName returns a boolean if a field has been set.
func (o *IssueDetail) HasIssueStatusName() bool {
	if o != nil && !IsNil(o.IssueStatusName) {
		return true
	}

	return false
}

// SetIssueStatusName gets a reference to the given string and assigns it to the IssueStatusName field.
func (o *IssueDetail) SetIssueStatusName(v string) {
	o.IssueStatusName = &v
}

// GetIssueStatusType returns the IssueStatusType field value if set, zero value otherwise.
func (o *IssueDetail) GetIssueStatusType() string {
	if o == nil || IsNil(o.IssueStatusType) {
		var ret string
		return ret
	}
	return *o.IssueStatusType
}

// GetIssueStatusTypeOk returns a tuple with the IssueStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIssueStatusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IssueStatusType) {
		return nil, false
	}
	return o.IssueStatusType, true
}

// HasIssueStatusType returns a boolean if a field has been set.
func (o *IssueDetail) HasIssueStatusType() bool {
	if o != nil && !IsNil(o.IssueStatusType) {
		return true
	}

	return false
}

// SetIssueStatusType gets a reference to the given string and assigns it to the IssueStatusType field.
func (o *IssueDetail) SetIssueStatusType(v string) {
	o.IssueStatusType = &v
}

// GetIssueTypeDetail returns the IssueTypeDetail field value if set, zero value otherwise.
func (o *IssueDetail) GetIssueTypeDetail() IssueTypeDetail {
	if o == nil || IsNil(o.IssueTypeDetail) {
		var ret IssueTypeDetail
		return ret
	}
	return *o.IssueTypeDetail
}

// GetIssueTypeDetailOk returns a tuple with the IssueTypeDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIssueTypeDetailOk() (*IssueTypeDetail, bool) {
	if o == nil || IsNil(o.IssueTypeDetail) {
		return nil, false
	}
	return o.IssueTypeDetail, true
}

// HasIssueTypeDetail returns a boolean if a field has been set.
func (o *IssueDetail) HasIssueTypeDetail() bool {
	if o != nil && !IsNil(o.IssueTypeDetail) {
		return true
	}

	return false
}

// SetIssueTypeDetail gets a reference to the given IssueTypeDetail and assigns it to the IssueTypeDetail field.
func (o *IssueDetail) SetIssueTypeDetail(v IssueTypeDetail) {
	o.IssueTypeDetail = &v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise.
func (o *IssueDetail) GetIssueTypeId() int64 {
	if o == nil || IsNil(o.IssueTypeId) {
		var ret int64
		return ret
	}
	return *o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIssueTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IssueTypeId) {
		return nil, false
	}
	return o.IssueTypeId, true
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *IssueDetail) HasIssueTypeId() bool {
	if o != nil && !IsNil(o.IssueTypeId) {
		return true
	}

	return false
}

// SetIssueTypeId gets a reference to the given int64 and assigns it to the IssueTypeId field.
func (o *IssueDetail) SetIssueTypeId(v int64) {
	o.IssueTypeId = &v
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *IssueDetail) GetIteration() IterationSimple {
	if o == nil || IsNil(o.Iteration) {
		var ret IterationSimple
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIterationOk() (*IterationSimple, bool) {
	if o == nil || IsNil(o.Iteration) {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *IssueDetail) HasIteration() bool {
	if o != nil && !IsNil(o.Iteration) {
		return true
	}

	return false
}

// SetIteration gets a reference to the given IterationSimple and assigns it to the Iteration field.
func (o *IssueDetail) SetIteration(v IterationSimple) {
	o.Iteration = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise.
func (o *IssueDetail) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId) {
		var ret int64
		return ret
	}
	return *o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetIterationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IterationId) {
		return nil, false
	}
	return o.IterationId, true
}

// HasIterationId returns a boolean if a field has been set.
func (o *IssueDetail) HasIterationId() bool {
	if o != nil && !IsNil(o.IterationId) {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given int64 and assigns it to the IterationId field.
func (o *IssueDetail) SetIterationId(v int64) {
	o.IterationId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *IssueDetail) GetLabels() []IssueProjectLabel {
	if o == nil || IsNil(o.Labels) {
		var ret []IssueProjectLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetLabelsOk() ([]IssueProjectLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *IssueDetail) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []IssueProjectLabel and assigns it to the Labels field.
func (o *IssueDetail) SetLabels(v []IssueProjectLabel) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueDetail) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueDetail) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueDetail) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *IssueDetail) GetParent() IssueSimpleData {
	if o == nil || IsNil(o.Parent) {
		var ret IssueSimpleData
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetParentOk() (*IssueSimpleData, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *IssueDetail) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given IssueSimpleData and assigns it to the Parent field.
func (o *IssueDetail) SetParent(v IssueSimpleData) {
	o.Parent = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *IssueDetail) GetParentType() string {
	if o == nil || IsNil(o.ParentType) {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetParentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentType) {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *IssueDetail) HasParentType() bool {
	if o != nil && !IsNil(o.ParentType) {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *IssueDetail) SetParentType(v string) {
	o.ParentType = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IssueDetail) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IssueDetail) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *IssueDetail) SetPriority(v string) {
	o.Priority = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *IssueDetail) GetProject() Project {
	if o == nil || IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetProjectOk() (*Project, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *IssueDetail) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *IssueDetail) SetProject(v Project) {
	o.Project = &v
}

// GetProjectModule returns the ProjectModule field value if set, zero value otherwise.
func (o *IssueDetail) GetProjectModule() IssueProjectModule {
	if o == nil || IsNil(o.ProjectModule) {
		var ret IssueProjectModule
		return ret
	}
	return *o.ProjectModule
}

// GetProjectModuleOk returns a tuple with the ProjectModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetProjectModuleOk() (*IssueProjectModule, bool) {
	if o == nil || IsNil(o.ProjectModule) {
		return nil, false
	}
	return o.ProjectModule, true
}

// HasProjectModule returns a boolean if a field has been set.
func (o *IssueDetail) HasProjectModule() bool {
	if o != nil && !IsNil(o.ProjectModule) {
		return true
	}

	return false
}

// SetProjectModule gets a reference to the given IssueProjectModule and assigns it to the ProjectModule field.
func (o *IssueDetail) SetProjectModule(v IssueProjectModule) {
	o.ProjectModule = &v
}

// GetRequirementType returns the RequirementType field value if set, zero value otherwise.
func (o *IssueDetail) GetRequirementType() RequirementType {
	if o == nil || IsNil(o.RequirementType) {
		var ret RequirementType
		return ret
	}
	return *o.RequirementType
}

// GetRequirementTypeOk returns a tuple with the RequirementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetRequirementTypeOk() (*RequirementType, bool) {
	if o == nil || IsNil(o.RequirementType) {
		return nil, false
	}
	return o.RequirementType, true
}

// HasRequirementType returns a boolean if a field has been set.
func (o *IssueDetail) HasRequirementType() bool {
	if o != nil && !IsNil(o.RequirementType) {
		return true
	}

	return false
}

// SetRequirementType gets a reference to the given RequirementType and assigns it to the RequirementType field.
func (o *IssueDetail) SetRequirementType(v RequirementType) {
	o.RequirementType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *IssueDetail) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate) {
		var ret int64
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetStartDateOk() (*int64, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *IssueDetail) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given int64 and assigns it to the StartDate field.
func (o *IssueDetail) SetStartDate(v int64) {
	o.StartDate = &v
}

// GetStoryPoint returns the StoryPoint field value if set, zero value otherwise.
func (o *IssueDetail) GetStoryPoint() string {
	if o == nil || IsNil(o.StoryPoint) {
		var ret string
		return ret
	}
	return *o.StoryPoint
}

// GetStoryPointOk returns a tuple with the StoryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetStoryPointOk() (*string, bool) {
	if o == nil || IsNil(o.StoryPoint) {
		return nil, false
	}
	return o.StoryPoint, true
}

// HasStoryPoint returns a boolean if a field has been set.
func (o *IssueDetail) HasStoryPoint() bool {
	if o != nil && !IsNil(o.StoryPoint) {
		return true
	}

	return false
}

// SetStoryPoint gets a reference to the given string and assigns it to the StoryPoint field.
func (o *IssueDetail) SetStoryPoint(v string) {
	o.StoryPoint = &v
}

// GetSubTasks returns the SubTasks field value if set, zero value otherwise.
func (o *IssueDetail) GetSubTasks() []SubTask {
	if o == nil || IsNil(o.SubTasks) {
		var ret []SubTask
		return ret
	}
	return o.SubTasks
}

// GetSubTasksOk returns a tuple with the SubTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetSubTasksOk() ([]SubTask, bool) {
	if o == nil || IsNil(o.SubTasks) {
		return nil, false
	}
	return o.SubTasks, true
}

// HasSubTasks returns a boolean if a field has been set.
func (o *IssueDetail) HasSubTasks() bool {
	if o != nil && !IsNil(o.SubTasks) {
		return true
	}

	return false
}

// SetSubTasks gets a reference to the given []SubTask and assigns it to the SubTasks field.
func (o *IssueDetail) SetSubTasks(v []SubTask) {
	o.SubTasks = v
}

// GetThirdLinks returns the ThirdLinks field value if set, zero value otherwise.
func (o *IssueDetail) GetThirdLinks() []IssueThirdLink {
	if o == nil || IsNil(o.ThirdLinks) {
		var ret []IssueThirdLink
		return ret
	}
	return o.ThirdLinks
}

// GetThirdLinksOk returns a tuple with the ThirdLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetThirdLinksOk() ([]IssueThirdLink, bool) {
	if o == nil || IsNil(o.ThirdLinks) {
		return nil, false
	}
	return o.ThirdLinks, true
}

// HasThirdLinks returns a boolean if a field has been set.
func (o *IssueDetail) HasThirdLinks() bool {
	if o != nil && !IsNil(o.ThirdLinks) {
		return true
	}

	return false
}

// SetThirdLinks gets a reference to the given []IssueThirdLink and assigns it to the ThirdLinks field.
func (o *IssueDetail) SetThirdLinks(v []IssueThirdLink) {
	o.ThirdLinks = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueDetail) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueDetail) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueDetail) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IssueDetail) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IssueDetail) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *IssueDetail) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetWatchers returns the Watchers field value if set, zero value otherwise.
func (o *IssueDetail) GetWatchers() []User {
	if o == nil || IsNil(o.Watchers) {
		var ret []User
		return ret
	}
	return o.Watchers
}

// GetWatchersOk returns a tuple with the Watchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetWatchersOk() ([]User, bool) {
	if o == nil || IsNil(o.Watchers) {
		return nil, false
	}
	return o.Watchers, true
}

// HasWatchers returns a boolean if a field has been set.
func (o *IssueDetail) HasWatchers() bool {
	if o != nil && !IsNil(o.Watchers) {
		return true
	}

	return false
}

// SetWatchers gets a reference to the given []User and assigns it to the Watchers field.
func (o *IssueDetail) SetWatchers(v []User) {
	o.Watchers = v
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *IssueDetail) GetWorkingHours() float32 {
	if o == nil || IsNil(o.WorkingHours) {
		var ret float32
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetWorkingHoursOk() (*float32, bool) {
	if o == nil || IsNil(o.WorkingHours) {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *IssueDetail) HasWorkingHours() bool {
	if o != nil && !IsNil(o.WorkingHours) {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given float32 and assigns it to the WorkingHours field.
func (o *IssueDetail) SetWorkingHours(v float32) {
	o.WorkingHours = &v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *IssueDetail) GetAssignees() []User {
	if o == nil || IsNil(o.Assignees) {
		var ret []User
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueDetail) GetAssigneesOk() ([]User, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *IssueDetail) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []User and assigns it to the Assignees field.
func (o *IssueDetail) SetAssignees(v []User) {
	o.Assignees = v
}

func (o IssueDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignee) {
		toSerialize["Assignee"] = o.Assignee
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.CompletedAt) {
		toSerialize["CompletedAt"] = o.CompletedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.Creator) {
		toSerialize["Creator"] = o.Creator
	}
	if !IsNil(o.CustomFields) {
		toSerialize["CustomFields"] = o.CustomFields
	}
	if !IsNil(o.DefectType) {
		toSerialize["DefectType"] = o.DefectType
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !IsNil(o.Epic) {
		toSerialize["Epic"] = o.Epic
	}
	if !IsNil(o.Files) {
		toSerialize["Files"] = o.Files
	}
	if !IsNil(o.IssueStatusId) {
		toSerialize["IssueStatusId"] = o.IssueStatusId
	}
	if !IsNil(o.IssueStatusName) {
		toSerialize["IssueStatusName"] = o.IssueStatusName
	}
	if !IsNil(o.IssueStatusType) {
		toSerialize["IssueStatusType"] = o.IssueStatusType
	}
	if !IsNil(o.IssueTypeDetail) {
		toSerialize["IssueTypeDetail"] = o.IssueTypeDetail
	}
	if !IsNil(o.IssueTypeId) {
		toSerialize["IssueTypeId"] = o.IssueTypeId
	}
	if !IsNil(o.Iteration) {
		toSerialize["Iteration"] = o.Iteration
	}
	if !IsNil(o.IterationId) {
		toSerialize["IterationId"] = o.IterationId
	}
	if !IsNil(o.Labels) {
		toSerialize["Labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["Parent"] = o.Parent
	}
	if !IsNil(o.ParentType) {
		toSerialize["ParentType"] = o.ParentType
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if !IsNil(o.Project) {
		toSerialize["Project"] = o.Project
	}
	if !IsNil(o.ProjectModule) {
		toSerialize["ProjectModule"] = o.ProjectModule
	}
	if !IsNil(o.RequirementType) {
		toSerialize["RequirementType"] = o.RequirementType
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.StoryPoint) {
		toSerialize["StoryPoint"] = o.StoryPoint
	}
	if !IsNil(o.SubTasks) {
		toSerialize["SubTasks"] = o.SubTasks
	}
	if !IsNil(o.ThirdLinks) {
		toSerialize["ThirdLinks"] = o.ThirdLinks
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Watchers) {
		toSerialize["Watchers"] = o.Watchers
	}
	if !IsNil(o.WorkingHours) {
		toSerialize["WorkingHours"] = o.WorkingHours
	}
	if !IsNil(o.Assignees) {
		toSerialize["Assignees"] = o.Assignees
	}
	return toSerialize, nil
}

type NullableIssueDetail struct {
	value *IssueDetail
	isSet bool
}

func (v NullableIssueDetail) Get() *IssueDetail {
	return v.value
}

func (v *NullableIssueDetail) Set(val *IssueDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueDetail(val *IssueDetail) *NullableIssueDetail {
	return &NullableIssueDetail{value: val, isSet: true}
}

func (v NullableIssueDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


