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

// checks if the IssueListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueListData{}

// IssueListData 事项列表信息
type IssueListData struct {
	// 处理人 Id
	AssigneeId *int64 `json:"AssigneeId,omitempty"`
	// 事项 Code
	Code *int64 `json:"Code,omitempty"`
	// 完成时间时间戳
	CompletedAt NullableInt64 `json:"CompletedAt,omitempty"`
	// 创建时间时间戳
	CreatedAt NullableInt64 `json:"CreatedAt,omitempty"`
	// 创建人 Id
	CreatorId *int64 `json:"CreatorId,omitempty"`
	// 自定义属性
	CustomFields []IssueCustomField `json:"CustomFields,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 截止日期时间戳
	DueDate NullableInt64 `json:"DueDate,omitempty"`
	// 状态 Id
	IssueStatusId *int64 `json:"IssueStatusId,omitempty"`
	// 状态名称
	IssueStatusName *string `json:"IssueStatusName,omitempty"`
	// 状态类型：  TODO｜PROCESSING｜COMPLETED
	IssueStatusType *string `json:"IssueStatusType,omitempty"`
	IssueTypeDetail *IssueTypeDetail `json:"IssueTypeDetail,omitempty"`
	// 事项类型 ID
	IssueTypeId NullableInt64 `json:"IssueTypeId,omitempty"`
	Iteration *IterationSimple `json:"Iteration,omitempty"`
	// 迭代 Id
	IterationId *int64 `json:"IterationId,omitempty"`
	// 事项标签
	Labels []IssueProjectLabel `json:"Labels,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	// 父事项code
	ParentCode NullableInt64 `json:"ParentCode,omitempty"`
	// 父事项ID
	ParentId NullableInt64 `json:"ParentId,omitempty"`
	// 父事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗
	ParentType *string `json:"ParentType,omitempty"`
	// 优先级：  \"0\" - 低，  \"1\" - 中，  \"2\" - 高，  \"3\" - 紧急，  \"\" - 未指定
	Priority *string `json:"Priority,omitempty"`
	// 开始日期时间戳
	StartDate NullableInt64 `json:"StartDate,omitempty"`
	// 故事点，例如：\"0.5\"、\"0\"，  空字符串 \"\" 表示未指定
	StoryPoint NullableString `json:"StoryPoint,omitempty"`
	// 事项类型：  DEFECT - 缺陷，  REQUIREMENT - 需求，  MISSION - 任务，  EPIC - 史诗，  SUB_TASK - 子工作项
	Type *string `json:"Type,omitempty"`
	// 修改时间时间戳
	UpdatedAt NullableInt64 `json:"UpdatedAt,omitempty"`
	// 工时（小时）
	WorkingHours *float32 `json:"WorkingHours,omitempty"`
	// 经办人列表
	Assignees []User `json:"Assignees,omitempty"`
}

// NewIssueListData instantiates a new IssueListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueListData() *IssueListData {
	this := IssueListData{}
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
	this.StoryPoint = *NewNullableString(&storyPoint)
	var type_ string = ""
	this.Type = &type_
	return &this
}

// NewIssueListDataWithDefaults instantiates a new IssueListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueListDataWithDefaults() *IssueListData {
	this := IssueListData{}
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
	this.StoryPoint = *NewNullableString(&storyPoint)
	var type_ string = ""
	this.Type = &type_
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *IssueListData) GetAssigneeId() int64 {
	if o == nil || IsNil(o.AssigneeId) {
		var ret int64
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetAssigneeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssigneeId) {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *IssueListData) HasAssigneeId() bool {
	if o != nil && !IsNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given int64 and assigns it to the AssigneeId field.
func (o *IssueListData) SetAssigneeId(v int64) {
	o.AssigneeId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *IssueListData) GetCode() int64 {
	if o == nil || IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *IssueListData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *IssueListData) SetCode(v int64) {
	o.Code = &v
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetCompletedAt() int64 {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetCompletedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *IssueListData) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableInt64 and assigns it to the CompletedAt field.
func (o *IssueListData) SetCompletedAt(v int64) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *IssueListData) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *IssueListData) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IssueListData) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *IssueListData) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *IssueListData) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *IssueListData) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *IssueListData) GetCreatorId() int64 {
	if o == nil || IsNil(o.CreatorId) {
		var ret int64
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetCreatorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *IssueListData) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given int64 and assigns it to the CreatorId field.
func (o *IssueListData) SetCreatorId(v int64) {
	o.CreatorId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetCustomFields() []IssueCustomField {
	if o == nil {
		var ret []IssueCustomField
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetCustomFieldsOk() ([]IssueCustomField, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IssueListData) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []IssueCustomField and assigns it to the CustomFields field.
func (o *IssueListData) SetCustomFields(v []IssueCustomField) {
	o.CustomFields = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IssueListData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IssueListData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IssueListData) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetDueDate() int64 {
	if o == nil || IsNil(o.DueDate.Get()) {
		var ret int64
		return ret
	}
	return *o.DueDate.Get()
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetDueDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DueDate.Get(), o.DueDate.IsSet()
}

// HasDueDate returns a boolean if a field has been set.
func (o *IssueListData) HasDueDate() bool {
	if o != nil && o.DueDate.IsSet() {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given NullableInt64 and assigns it to the DueDate field.
func (o *IssueListData) SetDueDate(v int64) {
	o.DueDate.Set(&v)
}
// SetDueDateNil sets the value for DueDate to be an explicit nil
func (o *IssueListData) SetDueDateNil() {
	o.DueDate.Set(nil)
}

// UnsetDueDate ensures that no value is present for DueDate, not even an explicit nil
func (o *IssueListData) UnsetDueDate() {
	o.DueDate.Unset()
}

// GetIssueStatusId returns the IssueStatusId field value if set, zero value otherwise.
func (o *IssueListData) GetIssueStatusId() int64 {
	if o == nil || IsNil(o.IssueStatusId) {
		var ret int64
		return ret
	}
	return *o.IssueStatusId
}

// GetIssueStatusIdOk returns a tuple with the IssueStatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIssueStatusIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IssueStatusId) {
		return nil, false
	}
	return o.IssueStatusId, true
}

// HasIssueStatusId returns a boolean if a field has been set.
func (o *IssueListData) HasIssueStatusId() bool {
	if o != nil && !IsNil(o.IssueStatusId) {
		return true
	}

	return false
}

// SetIssueStatusId gets a reference to the given int64 and assigns it to the IssueStatusId field.
func (o *IssueListData) SetIssueStatusId(v int64) {
	o.IssueStatusId = &v
}

// GetIssueStatusName returns the IssueStatusName field value if set, zero value otherwise.
func (o *IssueListData) GetIssueStatusName() string {
	if o == nil || IsNil(o.IssueStatusName) {
		var ret string
		return ret
	}
	return *o.IssueStatusName
}

// GetIssueStatusNameOk returns a tuple with the IssueStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIssueStatusNameOk() (*string, bool) {
	if o == nil || IsNil(o.IssueStatusName) {
		return nil, false
	}
	return o.IssueStatusName, true
}

// HasIssueStatusName returns a boolean if a field has been set.
func (o *IssueListData) HasIssueStatusName() bool {
	if o != nil && !IsNil(o.IssueStatusName) {
		return true
	}

	return false
}

// SetIssueStatusName gets a reference to the given string and assigns it to the IssueStatusName field.
func (o *IssueListData) SetIssueStatusName(v string) {
	o.IssueStatusName = &v
}

// GetIssueStatusType returns the IssueStatusType field value if set, zero value otherwise.
func (o *IssueListData) GetIssueStatusType() string {
	if o == nil || IsNil(o.IssueStatusType) {
		var ret string
		return ret
	}
	return *o.IssueStatusType
}

// GetIssueStatusTypeOk returns a tuple with the IssueStatusType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIssueStatusTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IssueStatusType) {
		return nil, false
	}
	return o.IssueStatusType, true
}

// HasIssueStatusType returns a boolean if a field has been set.
func (o *IssueListData) HasIssueStatusType() bool {
	if o != nil && !IsNil(o.IssueStatusType) {
		return true
	}

	return false
}

// SetIssueStatusType gets a reference to the given string and assigns it to the IssueStatusType field.
func (o *IssueListData) SetIssueStatusType(v string) {
	o.IssueStatusType = &v
}

// GetIssueTypeDetail returns the IssueTypeDetail field value if set, zero value otherwise.
func (o *IssueListData) GetIssueTypeDetail() IssueTypeDetail {
	if o == nil || IsNil(o.IssueTypeDetail) {
		var ret IssueTypeDetail
		return ret
	}
	return *o.IssueTypeDetail
}

// GetIssueTypeDetailOk returns a tuple with the IssueTypeDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIssueTypeDetailOk() (*IssueTypeDetail, bool) {
	if o == nil || IsNil(o.IssueTypeDetail) {
		return nil, false
	}
	return o.IssueTypeDetail, true
}

// HasIssueTypeDetail returns a boolean if a field has been set.
func (o *IssueListData) HasIssueTypeDetail() bool {
	if o != nil && !IsNil(o.IssueTypeDetail) {
		return true
	}

	return false
}

// SetIssueTypeDetail gets a reference to the given IssueTypeDetail and assigns it to the IssueTypeDetail field.
func (o *IssueListData) SetIssueTypeDetail(v IssueTypeDetail) {
	o.IssueTypeDetail = &v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetIssueTypeId() int64 {
	if o == nil || IsNil(o.IssueTypeId.Get()) {
		var ret int64
		return ret
	}
	return *o.IssueTypeId.Get()
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetIssueTypeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssueTypeId.Get(), o.IssueTypeId.IsSet()
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *IssueListData) HasIssueTypeId() bool {
	if o != nil && o.IssueTypeId.IsSet() {
		return true
	}

	return false
}

// SetIssueTypeId gets a reference to the given NullableInt64 and assigns it to the IssueTypeId field.
func (o *IssueListData) SetIssueTypeId(v int64) {
	o.IssueTypeId.Set(&v)
}
// SetIssueTypeIdNil sets the value for IssueTypeId to be an explicit nil
func (o *IssueListData) SetIssueTypeIdNil() {
	o.IssueTypeId.Set(nil)
}

// UnsetIssueTypeId ensures that no value is present for IssueTypeId, not even an explicit nil
func (o *IssueListData) UnsetIssueTypeId() {
	o.IssueTypeId.Unset()
}

// GetIteration returns the Iteration field value if set, zero value otherwise.
func (o *IssueListData) GetIteration() IterationSimple {
	if o == nil || IsNil(o.Iteration) {
		var ret IterationSimple
		return ret
	}
	return *o.Iteration
}

// GetIterationOk returns a tuple with the Iteration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIterationOk() (*IterationSimple, bool) {
	if o == nil || IsNil(o.Iteration) {
		return nil, false
	}
	return o.Iteration, true
}

// HasIteration returns a boolean if a field has been set.
func (o *IssueListData) HasIteration() bool {
	if o != nil && !IsNil(o.Iteration) {
		return true
	}

	return false
}

// SetIteration gets a reference to the given IterationSimple and assigns it to the Iteration field.
func (o *IssueListData) SetIteration(v IterationSimple) {
	o.Iteration = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise.
func (o *IssueListData) GetIterationId() int64 {
	if o == nil || IsNil(o.IterationId) {
		var ret int64
		return ret
	}
	return *o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetIterationIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IterationId) {
		return nil, false
	}
	return o.IterationId, true
}

// HasIterationId returns a boolean if a field has been set.
func (o *IssueListData) HasIterationId() bool {
	if o != nil && !IsNil(o.IterationId) {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given int64 and assigns it to the IterationId field.
func (o *IssueListData) SetIterationId(v int64) {
	o.IterationId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetLabels() []IssueProjectLabel {
	if o == nil {
		var ret []IssueProjectLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetLabelsOk() ([]IssueProjectLabel, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *IssueListData) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []IssueProjectLabel and assigns it to the Labels field.
func (o *IssueListData) SetLabels(v []IssueProjectLabel) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IssueListData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IssueListData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IssueListData) SetName(v string) {
	o.Name = &v
}

// GetParentCode returns the ParentCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetParentCode() int64 {
	if o == nil || IsNil(o.ParentCode.Get()) {
		var ret int64
		return ret
	}
	return *o.ParentCode.Get()
}

// GetParentCodeOk returns a tuple with the ParentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetParentCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentCode.Get(), o.ParentCode.IsSet()
}

// HasParentCode returns a boolean if a field has been set.
func (o *IssueListData) HasParentCode() bool {
	if o != nil && o.ParentCode.IsSet() {
		return true
	}

	return false
}

// SetParentCode gets a reference to the given NullableInt64 and assigns it to the ParentCode field.
func (o *IssueListData) SetParentCode(v int64) {
	o.ParentCode.Set(&v)
}
// SetParentCodeNil sets the value for ParentCode to be an explicit nil
func (o *IssueListData) SetParentCodeNil() {
	o.ParentCode.Set(nil)
}

// UnsetParentCode ensures that no value is present for ParentCode, not even an explicit nil
func (o *IssueListData) UnsetParentCode() {
	o.ParentCode.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetParentId() int64 {
	if o == nil || IsNil(o.ParentId.Get()) {
		var ret int64
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *IssueListData) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableInt64 and assigns it to the ParentId field.
func (o *IssueListData) SetParentId(v int64) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *IssueListData) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *IssueListData) UnsetParentId() {
	o.ParentId.Unset()
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *IssueListData) GetParentType() string {
	if o == nil || IsNil(o.ParentType) {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetParentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentType) {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *IssueListData) HasParentType() bool {
	if o != nil && !IsNil(o.ParentType) {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *IssueListData) SetParentType(v string) {
	o.ParentType = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IssueListData) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IssueListData) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *IssueListData) SetPriority(v string) {
	o.Priority = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetStartDate() int64 {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret int64
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *IssueListData) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableInt64 and assigns it to the StartDate field.
func (o *IssueListData) SetStartDate(v int64) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *IssueListData) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *IssueListData) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetStoryPoint returns the StoryPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetStoryPoint() string {
	if o == nil || IsNil(o.StoryPoint.Get()) {
		var ret string
		return ret
	}
	return *o.StoryPoint.Get()
}

// GetStoryPointOk returns a tuple with the StoryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetStoryPointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StoryPoint.Get(), o.StoryPoint.IsSet()
}

// HasStoryPoint returns a boolean if a field has been set.
func (o *IssueListData) HasStoryPoint() bool {
	if o != nil && o.StoryPoint.IsSet() {
		return true
	}

	return false
}

// SetStoryPoint gets a reference to the given NullableString and assigns it to the StoryPoint field.
func (o *IssueListData) SetStoryPoint(v string) {
	o.StoryPoint.Set(&v)
}
// SetStoryPointNil sets the value for StoryPoint to be an explicit nil
func (o *IssueListData) SetStoryPointNil() {
	o.StoryPoint.Set(nil)
}

// UnsetStoryPoint ensures that no value is present for StoryPoint, not even an explicit nil
func (o *IssueListData) UnsetStoryPoint() {
	o.StoryPoint.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IssueListData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IssueListData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IssueListData) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueListData) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueListData) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IssueListData) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *IssueListData) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *IssueListData) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *IssueListData) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *IssueListData) GetWorkingHours() float32 {
	if o == nil || IsNil(o.WorkingHours) {
		var ret float32
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetWorkingHoursOk() (*float32, bool) {
	if o == nil || IsNil(o.WorkingHours) {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *IssueListData) HasWorkingHours() bool {
	if o != nil && !IsNil(o.WorkingHours) {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given float32 and assigns it to the WorkingHours field.
func (o *IssueListData) SetWorkingHours(v float32) {
	o.WorkingHours = &v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *IssueListData) GetAssignees() []User {
	if o == nil || IsNil(o.Assignees) {
		var ret []User
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueListData) GetAssigneesOk() ([]User, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *IssueListData) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []User and assigns it to the Assignees field.
func (o *IssueListData) SetAssignees(v []User) {
	o.Assignees = v
}

func (o IssueListData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssigneeId) {
		toSerialize["AssigneeId"] = o.AssigneeId
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if o.CompletedAt.IsSet() {
		toSerialize["CompletedAt"] = o.CompletedAt.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if !IsNil(o.CreatorId) {
		toSerialize["CreatorId"] = o.CreatorId
	}
	if o.CustomFields != nil {
		toSerialize["CustomFields"] = o.CustomFields
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if o.DueDate.IsSet() {
		toSerialize["DueDate"] = o.DueDate.Get()
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
	if o.IssueTypeId.IsSet() {
		toSerialize["IssueTypeId"] = o.IssueTypeId.Get()
	}
	if !IsNil(o.Iteration) {
		toSerialize["Iteration"] = o.Iteration
	}
	if !IsNil(o.IterationId) {
		toSerialize["IterationId"] = o.IterationId
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.ParentCode.IsSet() {
		toSerialize["ParentCode"] = o.ParentCode.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["ParentId"] = o.ParentId.Get()
	}
	if !IsNil(o.ParentType) {
		toSerialize["ParentType"] = o.ParentType
	}
	if !IsNil(o.Priority) {
		toSerialize["Priority"] = o.Priority
	}
	if o.StartDate.IsSet() {
		toSerialize["StartDate"] = o.StartDate.Get()
	}
	if o.StoryPoint.IsSet() {
		toSerialize["StoryPoint"] = o.StoryPoint.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	if !IsNil(o.WorkingHours) {
		toSerialize["WorkingHours"] = o.WorkingHours
	}
	if !IsNil(o.Assignees) {
		toSerialize["Assignees"] = o.Assignees
	}
	return toSerialize, nil
}

type NullableIssueListData struct {
	value *IssueListData
	isSet bool
}

func (v NullableIssueListData) Get() *IssueListData {
	return v.value
}

func (v *NullableIssueListData) Set(val *IssueListData) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueListData) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueListData(val *IssueListData) *NullableIssueListData {
	return &NullableIssueListData{value: val, isSet: true}
}

func (v NullableIssueListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


