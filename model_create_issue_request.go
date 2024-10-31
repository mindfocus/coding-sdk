/*
CODING OPEN API

CODING 提供了丰富的 API 接口，注册应用即可使用，无需审核，支持两种认证方式：[OAuth 2.0 协议](#oauth-认证)、[个人访问令牌](#个人令牌认证)、[项目令牌](#项目令牌认证)。  # [OAuth 认证](#oauth)  #### [生态令牌权限](#ecology-scope)  令牌权限说明 > Scope的权限分为只读（ro）和读写(rw) 俩种，用户可通过权限点+权限方式获取完整的权限 Scope 信息。例如，用户信息授权为只读时，用户的令牌权限 Scope 为 user:profile:ro  | 名称            | 描述信息                 | Scope 信息                  | 权限范围  | 示例                           | |---------------|----------------------|---------------------------|-------|------------------------------| | 用户信息          | 管理用户的基础信息。           | user:profile              | ro    | user:profile:ro              | | 用户邮箱          | 管理用户的电子邮件地址。         | user:email                | ro    | user:email:ro                | | 用户通知          | 管理用户的站内通知。           | user:notification         | ro、rw | user:notification:rw         | | 用户公钥          | 管理用户配置的个人公钥和部署公钥信息。  | user:public-key           | ro、rw | user:public-key:rw           | | 团队信息          | 管理团队基本信息。            | team:profile              | ro    | team:profile:ro              | | 团队成员          | 管理团队成员信息以及团队成员相关操作。  | team:member               | ro、rw | team:member:rw               | | 项目信息          | 管理项目基本信息。            | project:profile           | ro、rw | project:profile:rw           | | 项目成员          | 管理项目成员。              | project:member            | ro、rw | project:member:rw            | | 项目令牌          | 管理项目令牌。              | project:token             | ro、rw | project:token:rw             | | 项目公告          | 管理项目公告。              | project:notice            | ro、rw | project:notice:rw            | | 项目标签          | 管理项目标签。              | project:label             | ro、rw | project:label:rw             | | 项目集信息         | 管理项目集基本信息。           | program:profile           | ro、rw | program:profile:rw           | | 项目集项目         | 管理项目集下的项目列表。         | program:project           | ro、rw | program:project:rw           | | 项目集成员         | 管理项目集下的成员列表。         | program:member            | ro、rw | program:member:rw            | | 关联资源          | 管理团队和项目资源关联关系。       | related-resource:resource | ro、rw | related-resource:resource:rw | | 凭据信息          | 管理团队凭据。              | credential:profile        | ro、rw | credential:profile:rw        | | Service Hooks | 管理和配置 Service Hooks。 | service-hook:profile      | ro、rw | service-hook:profile:rw      | | 权限组           | 管理权限组。               | ram:policy                | ro、rw | ram:policy:rw                | | 授权            | 配置权限授权。              | ram:grant                 | ro、rw | ram:grant:rw                 | | 用户组           | 管理权限用户组。             | ram:user-group            | ro    | ram:user-group:ro            | | 研发度量数据集       | 研发度量数据集              | performance:dataset       | ro    | performance:dataset:ro       | | 项目协同          | 配置和使用项目协同功能。         | collaboration:issue       | ro、rw | collaboration:issue:rw       | | 知识管理          | 管理知识空间和撰写知识文档。       | document:knowledge        | ro、rw | document:knowledge:rw        | | 文件网盘          | 管理上传、分享和下载文件等。       | document:file             | ro、rw | document:file:rw             | | API 文档        | 发布、授权发布 API 文档。      | document:api-doc          | ro、rw | document:api-doc:rw          | | 代码仓库          | 管理仓库                 | vcs:repository            | ro、rw | vcs:repository:rw            | | 合并请求          | 管理代码仓库的合并请求。         | vcs:merge                 | ro、rw | vcs:merge:rw                 | | 部署公钥          | 管理代码仓库的部署公钥。         | vcs:ssh-key               | ro、rw | vcs:ssh-key:rw               | | 版本管理          | 管理代码仓库的版本信息。         | vcs:release               | ro、rw | vcs:release:rw               | | 外部仓库          | 管理关联的外部仓库信息。         | depot:external-repository | ro、rw | depot:external-repository:rw | | 测试管理          | 管理测试计、测试用例和测试报告等。    | testing:profile           | ro、rw | testing:profile:rw           | | 持续部署数据统计      | 持续部署发布数据统计。          | cd:statistics             | ro    | cd:statistics:ro             | | 持续部署主机组       | 管理持续部署主机组。           | cd:host-server            | ro、rw | cd:host-server:rw            | | 持续部署云账号       | 管理持续部署云账号。           | cd:cloud-account          | ro、rw | cd:cloud-account:rw          | | 持续部署应用        | 管理和配置持续部署应用。         | cd:application            | ro、rw | cd:application:rw            | | 持续部署流程        | 管理和配置持续部署流程。         | cd:pipeline               | ro、rw | cd:pipeline:rw               | | 制品库仓库         | 管理制品库仓库。             | artifact:repository       | ro、rw | artifact:repository:rw       | | 制品库版本         | 管理制品版本信息。            | artifact:version          | ro、rw | artifact:version:rw          | | 制品库配置         | 管理制品库配置。             | artifact:properties       | ro、rw | artifact:properties:rw       | | 制品库包          | 管理制品库包。              | artifact:package          | ro    | artifact:package:ro          | | 资产列表          | 管理资产列表               | assets:list               | ro、rw | assets:list:rw               | | 资产属性          | 管理资产属性               | assets:attribute          | ro、rw | assets:attribute:rw          |  #### [创建 CODING 应用](#创建-CODING-应用)  ##### 1、新建应用  点击【团队设置】->【生态能力】->【发布应用】->【新建应用】，填写信息。「回调地址」处填写回调服务地址，这里以codesign为例。  ![](https://help-assets.codehub.cn/enterprise/202309201515397.png)  ![](https://help-assets.codehub.cn/enterprise/202309201519877.png)  ##### 2、获取 Client ID 和 Client Secret  点击创建好的应用，点击客户端秘钥，即可看到客户端ID与客户端秘钥  ![](https://help-assets.codehub.cn/enterprise/202309201520614.png)  ##### 注意事项： - 更换令牌时，先新建令牌，再删除旧令牌 - 最多同时可创建 5 个令牌 - 令牌只有创建时可查看，创建后，任何人无法查看，请妥善保管。  #### 3、修改令牌权限  根据需求修改令牌权限，这里设置了用户信息只读权限与项目信息读写权限作为示例  ![](https://help-assets.codehub.cn/enterprise/202309201520611.png)  #### [用户授权](#oauth-scopes)  浏览器访问以下链接，进入到授权登录页面：  ```shell GET https://{your-team}.coding.net/oauth_authorize.html?client_id={client_id}&redirect_uri={redirect_uri}&response_type=code&state=${state}&scope={scope}  ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - redirect_uri：应用详情页填写的回调地址；  - response_type：返回类型，固定为 code；  - state: 业务侧可以传任何值，这个值会原封不动传递回去，用来给业务识别一些场景用的。  - scope：授权范围，多个权限间以逗号分隔，如：user:profile:ro,project:profile:rw  ![](https://help-assets.codehub.cn/enterprise/202309201520497.png)  点击授权后，浏览器将带着授权码（code）参数和业务状态码（state）跳转到回调地址，如：  ```shell https://codesign.qq.com/?code={code}&state={state}&team={teamGk}&scope=user%3Aprofile%3Aro,project%3Aprofile%3Arw ```  #### [获取 access_token](#oauth-access-token)  获取授权码（code）后，开发者的后端程序向 CODING 发送请求，获取 access_token。  请求链接：  ```shell POST https://{your-team}.coding.net/api/oauth/access_token Content-Type: application/x-www-form-urlencoded  client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&code={code} ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - client_secret：应用详情页的 Client Secret；  - code：上一步获取的授权码，须注意每个 code 仅能使用一次，且有效期仅5分钟；  - grant_type：授权类型，固定为 authorization_code；  返回值：  ```json {   \"access_token\": \"RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975\",   \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",   \"scope\": \"user:profile:ro project:profile:rw\",   \"team\": \"jackwhu-test\",   \"token_type\": \"Bearer\",   \"expires_in\": \"7200\" } ```  返回值： - access_token： access_token值，可用于调用OpenAPI接口，建议按expires_in保存access_token - refresh_token： 刷新时使用的token，有效期90天。access_token过期后可用于刷新access_token - scope：令牌权限范围 - team： 团队gk - token_type：token类型 - expires_in：过期时间，单位秒   #### [刷新 access_token](#oauth-access-token)  通过上面获取到的 refresh_token，开发者的后端程序可以向 CODING 发送请求，刷新 access_token。   注意：调用刷新接口后，即使 access_token 未过期，原 access_token 也将失效  请求链接：  ```shell POST https://{your-team}.coding.net/api/oauth/access_token Content-Type: application/x-www-form-urlencoded  client_id={client_id}&client_secret={client_secret}&grant_type=authorization_code&refresh_token={refresh_token} ```  参数说明：  - your-team：团队域名前缀；  - client_id：应用详情页的 Client ID；  - client_secret：应用详情页的 Client Secret；  - refresh_token：获取 access_token接口返回的refresh_token字段  - grant_type：授权类型，固定为 refresh_token；  返回值：  ```json {   \"access_token\": \"q4qIkUGhJ2qfcdSV3bWx0YfQj-WjLqXG7LSdP9Oo3sOAjmuY-Bb_QJ6ORt-By-bc7WFFT7PyH7RXEvPLBM5lFU9PBOofgzXN9Lh5zp3FWRdyV_4XGno4U_S7zMkixWnv\",   \"refresh_token\": \"9pqtQ6QUQlX9cdJCJ726PbeEVI7wCxNWflb16vz59QxSHQlFv0nnlxr8CpKf4gwPhaijKsDFmPWdR9ryiPMdzESKaMAqIhfCYvBjzmpsqiSHxqLpCXGV1amOtqjx9eyJ\",   \"scope\": \"user:profile:ro project:profile:rw\",   \"team\": \"jackwhu-test\",   \"token_type\": \"Bearer\",   \"expires_in\": \"7200\" } ```  #### [获取当前用户信息](#oauth-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: Bearer RtdlB8UgzKZ7BJdXKC5dQRiUAzwFcBFZbB2NPaNX3DUkJMhXqi5Cb_k-V7QxbwzFmo9oBkelu95xSoxualndQm9e1oc6F3H77z97aSE6neRrjf5oOu8Owzew_J1yC975 Content-Type: application/json  {     \"Action\": \"DescribeCodingCurrentUser\" } ```  Headers说明：  - Authorization：\"Bearer {access_token}\"  参数说明：  - Action：固定为DescribeCodingCurrentUser，令牌须打开user:profile:ro权限  返回值：  ```json {   \"Response\": {     \"User\": {       \"Id\": 183478,       \"Status\": 1,       \"Email\": \"test@coding.com\",       \"GlobalKey\": \"anywhere\",       \"Avatar\": \"https://e.coding.net/static/fruit_avatar/Fruit-20.png\",       \"Gravatar\": \"\",       \"Name\": \"qqq\",       \"NamePinYin\": \"qqq\",       \"Phone\": \"18800000000\",       \"PhoneValidation\": 1,       \"EmailValidation\": 1,       \"PhoneRegionCode\": \"+86\",       \"Region\": \"cn\",       \"TeamId\": 1,       \"WeComId\": \"woG7VgCgAAov2F-sAQkD_YPsLNJiP1Vg\"     },     \"RequestId\": \"133e152f-8852-4d99-b704-c7ff245a1640\"   } }  ```  # [个人令牌认证](#personal_token)  #### [获取个人令牌](#personal-token-create)  - 点击左下角的【个人账户设置】>【访问令牌】>【新建访问令牌】，勾选相关权限后会生成「个人访问令牌」。若刷新页面令牌会消失，需输入账号密码后重新生成。 - 令牌权限与[OAuth令牌权限](#生态令牌权限)一样  ![](https://help-assets.codehub.cn/enterprise/202309201521630.png)  #### [获取当前用户信息](#personal-token-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8 Content-Type: application/json  {     \"Action\": \"DescribeCodingCurrentUser\" } ```  header：  ```text Authorization: token {访问令牌}  ```  请求示例：  ```shell curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\ --header 'Authorization: token 70823e19fb3f2aa7d1ef08727f6ab5d4dcd20af8' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"Action\": \"DescribeCodingCurrentUser\" }' ```  # [项目令牌认证](#personal_token)  #### [项目令牌Scope](#project-token-scope)  | 名称      | 描述                        | Scope 信息                  | 权限范围 | |---------|---------------------------|---------------------------|------| | 项目协同    | 读取与操作项目协同                 | project_issue_rw          | 读写   | | 文件      | 新建、查询、删除、编辑文件             | file_rw                   | 读写   | | WIKI    | 新建、查询、删除、编辑Wiki           | wiki_rw                   | 读写   | | 项目公告    | 新建、查询、删除、编辑项目公告           | project_tweet_rw          | 读写   | | API 文档  | 发布 API 文档                 | api_doc_release           | 读写   | | 关联资源    | 新建、查询、删除关联资源              | resource_reference_rw     | 读写   | | 项目成员    | 读取与操作项目成员                 | project_member_rw         | 读写   | | 项目权限    | 读取与操作项目权限                 | project_permission_rw     | 读写   | | 项目标签    | 新建、查询、删除、编辑项目标签           | project_label_rw          | 读写   | | 测试协同    | 执行 API 自动化测试              | ifbook_run_task           | 读写   | | 测试协同    | API 文档导入与导出               | ifbook_import_export      | 读写   | | 读取代码仓库  | 读取代码仓库                    | depot_read                | 只读   | | 推送至代码仓库 | 推送至代码仓库                   | depot_write               | 读写   | | MR      | 新建、查询、删除、编辑合并请求           | merge_request_rw          | 读写   | | 版本发布    | 新建、查询、删除、编辑版本发布           | release_rw                | 读写   | | 制品库     | 拉取制品库                     | artifact_r                | 只读   | | 制品库     | 拉取、推送制品库                  | artifact_rw               | 读写   | | 制品属性    | 新建、查询、删除、编辑制品属性           | artifact_version_props_rw | 读写   | | 构建节点    | 允许构建节点接入                  | ci_agent_register         | 读写   | | API触发   | 构建计划管理/触发构建               | ci_manager                | 读写   | | 构建计划    | 构建计划管理/触发构建（仅用于 Open API） | open_ci_manager           | 读写   |  #### [获取项目令牌](#project-token-create)  1. 点击【项目设置】>【开发者选项】>【项目令牌】>【新建项目令牌】，勾选相关权限后会生成「项目令牌」。点击查看密码即可获取密码信息  ![](https://help-assets.codehub.cn/enterprise/202309201843081.png)  2. Basic认证：将用户名与密码通过”用户名:密码“方式拼接后，使用Base64进行编码，获取凭证  #### [获取当前项目信息](#project-token-get-user)  请求链接：  ```shell POST https://{your-team}.coding.net/open-api Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE Content-Type: application/json  {     \"Action\": \"DescribeOneProject\",     \"ProjectId\": 11595365 } ```  header：  ```text Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE ```  请求示例：  ```shell curl --location --request POST 'https://jackwhu-test.coding.net/open-api' \\ --header 'Authorization: Basic cHRib2x5enBpYzB4OjNlZmYzOGY2MzU3MzhkYTNlMzAxYjcwZmI2ZGZhYzlhZjQ3MTQyZmE' \\ --header 'Content-Type: application/json' \\ --data-raw '{     \"Action\": \"DescribeOneProject\",     \"ProjectId\": 11595365 }' ```  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateIssueRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateIssueRequest{}

// CreateIssueRequest struct for CreateIssueRequest
type CreateIssueRequest struct {
	// 处理人 Id
	AssigneeId *int64 `json:"AssigneeId,omitempty"`
	// 自定义属性值列表
	CustomFieldValues []IssueCustomFieldForm `json:"CustomFieldValues,omitempty"`
	// 缺陷类型 Id
	DefectTypeId *int64 `json:"DefectTypeId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 截止日期，格式：2021-01-01
	DueDate map[string]interface{} `json:"DueDate,omitempty"`
	// 史诗 Code，Type 为 EPIC 或 SUB_TASK时，不传该值
	EpicCode *int64 `json:"EpicCode,omitempty"`
	// 附件的文件 Id 列表
	FileIds []int64 `json:"FileIds,omitempty"`
	// 事项类型 Id
	IssueTypeId *int64 `json:"IssueTypeId,omitempty"`
	// 迭代 Code，Type 为 EPIC 或 SUB_TASK时，忽略该值
	IterationCode *int64 `json:"IterationCode,omitempty"`
	// 标签 Id 列表
	LabelIds []int64 `json:"LabelIds,omitempty"`
	// 事项名称
	Name string `json:"Name"`
	// 父事项 Code  敏捷模式：Type 为 SUB_TASK 时必须指定  经典模式：Type 为 MISSION、REQUIREMENT 时可指定
	ParentCode *int64 `json:"ParentCode,omitempty"`
	// 紧急程度  \"0\" - 低  \"1\" - 中  \"2\" - 高  \"3\" - 紧急
	Priority string `json:"Priority"`
	// 项目模块 Id
	ProjectModuleId *int64 `json:"ProjectModuleId,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 版本code列表
	ReleaseCodes []int64 `json:"ReleaseCodes,omitempty"`
	// 需求类型 Id
	RequirementTypeId *int64 `json:"RequirementTypeId,omitempty"`
	// 开始日期，格式：2021-01-01
	StartDate map[string]interface{} `json:"StartDate,omitempty"`
	// 事项状态 Id
	StatusId *int64 `json:"StatusId,omitempty"`
	// 故事点，例如：0.5、1
	StoryPoint *string `json:"StoryPoint,omitempty"`
	// 排序目标位置的事项 code  可不填，排在底位
	TargetSortCode *int64 `json:"TargetSortCode,omitempty"`
	// 第三方链接列表
	ThirdLinks []IssueThirdLinkForm `json:"ThirdLinks,omitempty"`
	// 事项类型  DEFECT - 缺陷  REQUIREMENT - 需求  MISSION - 任务  EPIC - 史诗  SUB_TASK - 子任务
	Type string `json:"Type"`
	// 关注人 Id 列表
	WatcherIds []int64 `json:"WatcherIds,omitempty"`
	// 工时（小时数）
	WorkingHours *float32 `json:"WorkingHours,omitempty"`
}

type _CreateIssueRequest CreateIssueRequest

// NewCreateIssueRequest instantiates a new CreateIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIssueRequest(name string, priority string, projectName string, type_ string) *CreateIssueRequest {
	this := CreateIssueRequest{}
	this.Name = name
	this.Priority = priority
	this.ProjectName = projectName
	this.Type = type_
	return &this
}

// NewCreateIssueRequestWithDefaults instantiates a new CreateIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIssueRequestWithDefaults() *CreateIssueRequest {
	this := CreateIssueRequest{}
	return &this
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetAssigneeId() int64 {
	if o == nil || IsNil(o.AssigneeId) {
		var ret int64
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetAssigneeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AssigneeId) {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasAssigneeId() bool {
	if o != nil && !IsNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given int64 and assigns it to the AssigneeId field.
func (o *CreateIssueRequest) SetAssigneeId(v int64) {
	o.AssigneeId = &v
}

// GetCustomFieldValues returns the CustomFieldValues field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetCustomFieldValues() []IssueCustomFieldForm {
	if o == nil || IsNil(o.CustomFieldValues) {
		var ret []IssueCustomFieldForm
		return ret
	}
	return o.CustomFieldValues
}

// GetCustomFieldValuesOk returns a tuple with the CustomFieldValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetCustomFieldValuesOk() ([]IssueCustomFieldForm, bool) {
	if o == nil || IsNil(o.CustomFieldValues) {
		return nil, false
	}
	return o.CustomFieldValues, true
}

// HasCustomFieldValues returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasCustomFieldValues() bool {
	if o != nil && !IsNil(o.CustomFieldValues) {
		return true
	}

	return false
}

// SetCustomFieldValues gets a reference to the given []IssueCustomFieldForm and assigns it to the CustomFieldValues field.
func (o *CreateIssueRequest) SetCustomFieldValues(v []IssueCustomFieldForm) {
	o.CustomFieldValues = v
}

// GetDefectTypeId returns the DefectTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDefectTypeId() int64 {
	if o == nil || IsNil(o.DefectTypeId) {
		var ret int64
		return ret
	}
	return *o.DefectTypeId
}

// GetDefectTypeIdOk returns a tuple with the DefectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDefectTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DefectTypeId) {
		return nil, false
	}
	return o.DefectTypeId, true
}

// HasDefectTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDefectTypeId() bool {
	if o != nil && !IsNil(o.DefectTypeId) {
		return true
	}

	return false
}

// SetDefectTypeId gets a reference to the given int64 and assigns it to the DefectTypeId field.
func (o *CreateIssueRequest) SetDefectTypeId(v int64) {
	o.DefectTypeId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateIssueRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetDueDate() map[string]interface{} {
	if o == nil || IsNil(o.DueDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetDueDateOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DueDate) {
		return map[string]interface{}{}, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given map[string]interface{} and assigns it to the DueDate field.
func (o *CreateIssueRequest) SetDueDate(v map[string]interface{}) {
	o.DueDate = v
}

// GetEpicCode returns the EpicCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetEpicCode() int64 {
	if o == nil || IsNil(o.EpicCode) {
		var ret int64
		return ret
	}
	return *o.EpicCode
}

// GetEpicCodeOk returns a tuple with the EpicCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetEpicCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.EpicCode) {
		return nil, false
	}
	return o.EpicCode, true
}

// HasEpicCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasEpicCode() bool {
	if o != nil && !IsNil(o.EpicCode) {
		return true
	}

	return false
}

// SetEpicCode gets a reference to the given int64 and assigns it to the EpicCode field.
func (o *CreateIssueRequest) SetEpicCode(v int64) {
	o.EpicCode = &v
}

// GetFileIds returns the FileIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetFileIds() []int64 {
	if o == nil || IsNil(o.FileIds) {
		var ret []int64
		return ret
	}
	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetFileIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.FileIds) {
		return nil, false
	}
	return o.FileIds, true
}

// HasFileIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasFileIds() bool {
	if o != nil && !IsNil(o.FileIds) {
		return true
	}

	return false
}

// SetFileIds gets a reference to the given []int64 and assigns it to the FileIds field.
func (o *CreateIssueRequest) SetFileIds(v []int64) {
	o.FileIds = v
}

// GetIssueTypeId returns the IssueTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetIssueTypeId() int64 {
	if o == nil || IsNil(o.IssueTypeId) {
		var ret int64
		return ret
	}
	return *o.IssueTypeId
}

// GetIssueTypeIdOk returns a tuple with the IssueTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetIssueTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.IssueTypeId) {
		return nil, false
	}
	return o.IssueTypeId, true
}

// HasIssueTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasIssueTypeId() bool {
	if o != nil && !IsNil(o.IssueTypeId) {
		return true
	}

	return false
}

// SetIssueTypeId gets a reference to the given int64 and assigns it to the IssueTypeId field.
func (o *CreateIssueRequest) SetIssueTypeId(v int64) {
	o.IssueTypeId = &v
}

// GetIterationCode returns the IterationCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetIterationCode() int64 {
	if o == nil || IsNil(o.IterationCode) {
		var ret int64
		return ret
	}
	return *o.IterationCode
}

// GetIterationCodeOk returns a tuple with the IterationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetIterationCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.IterationCode) {
		return nil, false
	}
	return o.IterationCode, true
}

// HasIterationCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasIterationCode() bool {
	if o != nil && !IsNil(o.IterationCode) {
		return true
	}

	return false
}

// SetIterationCode gets a reference to the given int64 and assigns it to the IterationCode field.
func (o *CreateIssueRequest) SetIterationCode(v int64) {
	o.IterationCode = &v
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetLabelIds() []int64 {
	if o == nil || IsNil(o.LabelIds) {
		var ret []int64
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetLabelIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LabelIds) {
		return nil, false
	}
	return o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasLabelIds() bool {
	if o != nil && !IsNil(o.LabelIds) {
		return true
	}

	return false
}

// SetLabelIds gets a reference to the given []int64 and assigns it to the LabelIds field.
func (o *CreateIssueRequest) SetLabelIds(v []int64) {
	o.LabelIds = v
}

// GetName returns the Name field value
func (o *CreateIssueRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateIssueRequest) SetName(v string) {
	o.Name = v
}

// GetParentCode returns the ParentCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetParentCode() int64 {
	if o == nil || IsNil(o.ParentCode) {
		var ret int64
		return ret
	}
	return *o.ParentCode
}

// GetParentCodeOk returns a tuple with the ParentCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetParentCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.ParentCode) {
		return nil, false
	}
	return o.ParentCode, true
}

// HasParentCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasParentCode() bool {
	if o != nil && !IsNil(o.ParentCode) {
		return true
	}

	return false
}

// SetParentCode gets a reference to the given int64 and assigns it to the ParentCode field.
func (o *CreateIssueRequest) SetParentCode(v int64) {
	o.ParentCode = &v
}

// GetPriority returns the Priority field value
func (o *CreateIssueRequest) GetPriority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetPriorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateIssueRequest) SetPriority(v string) {
	o.Priority = v
}

// GetProjectModuleId returns the ProjectModuleId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetProjectModuleId() int64 {
	if o == nil || IsNil(o.ProjectModuleId) {
		var ret int64
		return ret
	}
	return *o.ProjectModuleId
}

// GetProjectModuleIdOk returns a tuple with the ProjectModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetProjectModuleIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProjectModuleId) {
		return nil, false
	}
	return o.ProjectModuleId, true
}

// HasProjectModuleId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasProjectModuleId() bool {
	if o != nil && !IsNil(o.ProjectModuleId) {
		return true
	}

	return false
}

// SetProjectModuleId gets a reference to the given int64 and assigns it to the ProjectModuleId field.
func (o *CreateIssueRequest) SetProjectModuleId(v int64) {
	o.ProjectModuleId = &v
}

// GetProjectName returns the ProjectName field value
func (o *CreateIssueRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateIssueRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetReleaseCodes returns the ReleaseCodes field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetReleaseCodes() []int64 {
	if o == nil || IsNil(o.ReleaseCodes) {
		var ret []int64
		return ret
	}
	return o.ReleaseCodes
}

// GetReleaseCodesOk returns a tuple with the ReleaseCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetReleaseCodesOk() ([]int64, bool) {
	if o == nil || IsNil(o.ReleaseCodes) {
		return nil, false
	}
	return o.ReleaseCodes, true
}

// HasReleaseCodes returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasReleaseCodes() bool {
	if o != nil && !IsNil(o.ReleaseCodes) {
		return true
	}

	return false
}

// SetReleaseCodes gets a reference to the given []int64 and assigns it to the ReleaseCodes field.
func (o *CreateIssueRequest) SetReleaseCodes(v []int64) {
	o.ReleaseCodes = v
}

// GetRequirementTypeId returns the RequirementTypeId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetRequirementTypeId() int64 {
	if o == nil || IsNil(o.RequirementTypeId) {
		var ret int64
		return ret
	}
	return *o.RequirementTypeId
}

// GetRequirementTypeIdOk returns a tuple with the RequirementTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetRequirementTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RequirementTypeId) {
		return nil, false
	}
	return o.RequirementTypeId, true
}

// HasRequirementTypeId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasRequirementTypeId() bool {
	if o != nil && !IsNil(o.RequirementTypeId) {
		return true
	}

	return false
}

// SetRequirementTypeId gets a reference to the given int64 and assigns it to the RequirementTypeId field.
func (o *CreateIssueRequest) SetRequirementTypeId(v int64) {
	o.RequirementTypeId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStartDate() map[string]interface{} {
	if o == nil || IsNil(o.StartDate) {
		var ret map[string]interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStartDateOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StartDate) {
		return map[string]interface{}{}, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given map[string]interface{} and assigns it to the StartDate field.
func (o *CreateIssueRequest) SetStartDate(v map[string]interface{}) {
	o.StartDate = v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStatusId() int64 {
	if o == nil || IsNil(o.StatusId) {
		var ret int64
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStatusIdOk() (*int64, bool) {
	if o == nil || IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStatusId() bool {
	if o != nil && !IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given int64 and assigns it to the StatusId field.
func (o *CreateIssueRequest) SetStatusId(v int64) {
	o.StatusId = &v
}

// GetStoryPoint returns the StoryPoint field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetStoryPoint() string {
	if o == nil || IsNil(o.StoryPoint) {
		var ret string
		return ret
	}
	return *o.StoryPoint
}

// GetStoryPointOk returns a tuple with the StoryPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetStoryPointOk() (*string, bool) {
	if o == nil || IsNil(o.StoryPoint) {
		return nil, false
	}
	return o.StoryPoint, true
}

// HasStoryPoint returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasStoryPoint() bool {
	if o != nil && !IsNil(o.StoryPoint) {
		return true
	}

	return false
}

// SetStoryPoint gets a reference to the given string and assigns it to the StoryPoint field.
func (o *CreateIssueRequest) SetStoryPoint(v string) {
	o.StoryPoint = &v
}

// GetTargetSortCode returns the TargetSortCode field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetTargetSortCode() int64 {
	if o == nil || IsNil(o.TargetSortCode) {
		var ret int64
		return ret
	}
	return *o.TargetSortCode
}

// GetTargetSortCodeOk returns a tuple with the TargetSortCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetTargetSortCodeOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetSortCode) {
		return nil, false
	}
	return o.TargetSortCode, true
}

// HasTargetSortCode returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasTargetSortCode() bool {
	if o != nil && !IsNil(o.TargetSortCode) {
		return true
	}

	return false
}

// SetTargetSortCode gets a reference to the given int64 and assigns it to the TargetSortCode field.
func (o *CreateIssueRequest) SetTargetSortCode(v int64) {
	o.TargetSortCode = &v
}

// GetThirdLinks returns the ThirdLinks field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetThirdLinks() []IssueThirdLinkForm {
	if o == nil || IsNil(o.ThirdLinks) {
		var ret []IssueThirdLinkForm
		return ret
	}
	return o.ThirdLinks
}

// GetThirdLinksOk returns a tuple with the ThirdLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetThirdLinksOk() ([]IssueThirdLinkForm, bool) {
	if o == nil || IsNil(o.ThirdLinks) {
		return nil, false
	}
	return o.ThirdLinks, true
}

// HasThirdLinks returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasThirdLinks() bool {
	if o != nil && !IsNil(o.ThirdLinks) {
		return true
	}

	return false
}

// SetThirdLinks gets a reference to the given []IssueThirdLinkForm and assigns it to the ThirdLinks field.
func (o *CreateIssueRequest) SetThirdLinks(v []IssueThirdLinkForm) {
	o.ThirdLinks = v
}

// GetType returns the Type field value
func (o *CreateIssueRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateIssueRequest) SetType(v string) {
	o.Type = v
}

// GetWatcherIds returns the WatcherIds field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetWatcherIds() []int64 {
	if o == nil || IsNil(o.WatcherIds) {
		var ret []int64
		return ret
	}
	return o.WatcherIds
}

// GetWatcherIdsOk returns a tuple with the WatcherIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetWatcherIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.WatcherIds) {
		return nil, false
	}
	return o.WatcherIds, true
}

// HasWatcherIds returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasWatcherIds() bool {
	if o != nil && !IsNil(o.WatcherIds) {
		return true
	}

	return false
}

// SetWatcherIds gets a reference to the given []int64 and assigns it to the WatcherIds field.
func (o *CreateIssueRequest) SetWatcherIds(v []int64) {
	o.WatcherIds = v
}

// GetWorkingHours returns the WorkingHours field value if set, zero value otherwise.
func (o *CreateIssueRequest) GetWorkingHours() float32 {
	if o == nil || IsNil(o.WorkingHours) {
		var ret float32
		return ret
	}
	return *o.WorkingHours
}

// GetWorkingHoursOk returns a tuple with the WorkingHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIssueRequest) GetWorkingHoursOk() (*float32, bool) {
	if o == nil || IsNil(o.WorkingHours) {
		return nil, false
	}
	return o.WorkingHours, true
}

// HasWorkingHours returns a boolean if a field has been set.
func (o *CreateIssueRequest) HasWorkingHours() bool {
	if o != nil && !IsNil(o.WorkingHours) {
		return true
	}

	return false
}

// SetWorkingHours gets a reference to the given float32 and assigns it to the WorkingHours field.
func (o *CreateIssueRequest) SetWorkingHours(v float32) {
	o.WorkingHours = &v
}

func (o CreateIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssigneeId) {
		toSerialize["AssigneeId"] = o.AssigneeId
	}
	if !IsNil(o.CustomFieldValues) {
		toSerialize["CustomFieldValues"] = o.CustomFieldValues
	}
	if !IsNil(o.DefectTypeId) {
		toSerialize["DefectTypeId"] = o.DefectTypeId
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.DueDate) {
		toSerialize["DueDate"] = o.DueDate
	}
	if !IsNil(o.EpicCode) {
		toSerialize["EpicCode"] = o.EpicCode
	}
	if !IsNil(o.FileIds) {
		toSerialize["FileIds"] = o.FileIds
	}
	if !IsNil(o.IssueTypeId) {
		toSerialize["IssueTypeId"] = o.IssueTypeId
	}
	if !IsNil(o.IterationCode) {
		toSerialize["IterationCode"] = o.IterationCode
	}
	if !IsNil(o.LabelIds) {
		toSerialize["LabelIds"] = o.LabelIds
	}
	toSerialize["Name"] = o.Name
	if !IsNil(o.ParentCode) {
		toSerialize["ParentCode"] = o.ParentCode
	}
	toSerialize["Priority"] = o.Priority
	if !IsNil(o.ProjectModuleId) {
		toSerialize["ProjectModuleId"] = o.ProjectModuleId
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !IsNil(o.ReleaseCodes) {
		toSerialize["ReleaseCodes"] = o.ReleaseCodes
	}
	if !IsNil(o.RequirementTypeId) {
		toSerialize["RequirementTypeId"] = o.RequirementTypeId
	}
	if !IsNil(o.StartDate) {
		toSerialize["StartDate"] = o.StartDate
	}
	if !IsNil(o.StatusId) {
		toSerialize["StatusId"] = o.StatusId
	}
	if !IsNil(o.StoryPoint) {
		toSerialize["StoryPoint"] = o.StoryPoint
	}
	if !IsNil(o.TargetSortCode) {
		toSerialize["TargetSortCode"] = o.TargetSortCode
	}
	if !IsNil(o.ThirdLinks) {
		toSerialize["ThirdLinks"] = o.ThirdLinks
	}
	toSerialize["Type"] = o.Type
	if !IsNil(o.WatcherIds) {
		toSerialize["WatcherIds"] = o.WatcherIds
	}
	if !IsNil(o.WorkingHours) {
		toSerialize["WorkingHours"] = o.WorkingHours
	}
	return toSerialize, nil
}

func (o *CreateIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Name",
		"Priority",
		"ProjectName",
		"Type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateIssueRequest := _CreateIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateIssueRequest)

	if err != nil {
		return err
	}

	*o = CreateIssueRequest(varCreateIssueRequest)

	return err
}

type NullableCreateIssueRequest struct {
	value *CreateIssueRequest
	isSet bool
}

func (v NullableCreateIssueRequest) Get() *CreateIssueRequest {
	return v.value
}

func (v *NullableCreateIssueRequest) Set(val *CreateIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIssueRequest(val *CreateIssueRequest) *NullableCreateIssueRequest {
	return &NullableCreateIssueRequest{value: val, isSet: true}
}

func (v NullableCreateIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


