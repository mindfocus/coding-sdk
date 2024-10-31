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

// checks if the DescribeReleaseIssueListRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeReleaseIssueListRequest{}

// DescribeReleaseIssueListRequest struct for DescribeReleaseIssueListRequest
type DescribeReleaseIssueListRequest struct {
	// 处理人ID数组
	Assignees []int64 `json:"Assignees,omitempty"`
	// 事项类型ID数组 
	IssueTypeIds []int64 `json:"IssueTypeIds,omitempty"`
	// 事项类型数组
	IssueTypes []string `json:"IssueTypes,omitempty"`
	// 关键字
	Keywords *string `json:"Keywords,omitempty"`
	// 分页查询中的页数,page从1开始计数 
	Page *int64 `json:"Page,omitempty"`
	// 分页查询中每页的大小 
	PageSize *int64 `json:"PageSize,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 页面上版本ID
	ReleaseCode int64 `json:"ReleaseCode"`
	// 是否展示描述中外部可访问的地址
	ShowImageOutUrl *bool `json:"ShowImageOutUrl,omitempty"`
	// 是否显示字事项，和页面开关对应 
	ShowSubIssues *bool `json:"ShowSubIssues,omitempty"`
	// 排序，取值如\"ID:ASC\" 
	SortBy *string `json:"SortBy,omitempty"`
	// 事项状态类型数组 
	StatusTypes []int64 `json:"StatusTypes,omitempty"`
	// 关注人ID数组 
	Watchers []int64 `json:"Watchers,omitempty"`
}

type _DescribeReleaseIssueListRequest DescribeReleaseIssueListRequest

// NewDescribeReleaseIssueListRequest instantiates a new DescribeReleaseIssueListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReleaseIssueListRequest(projectName string, releaseCode int64) *DescribeReleaseIssueListRequest {
	this := DescribeReleaseIssueListRequest{}
	this.ProjectName = projectName
	this.ReleaseCode = releaseCode
	return &this
}

// NewDescribeReleaseIssueListRequestWithDefaults instantiates a new DescribeReleaseIssueListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReleaseIssueListRequestWithDefaults() *DescribeReleaseIssueListRequest {
	this := DescribeReleaseIssueListRequest{}
	return &this
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetAssignees() []int64 {
	if o == nil || IsNil(o.Assignees) {
		var ret []int64
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetAssigneesOk() ([]int64, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []int64 and assigns it to the Assignees field.
func (o *DescribeReleaseIssueListRequest) SetAssignees(v []int64) {
	o.Assignees = v
}

// GetIssueTypeIds returns the IssueTypeIds field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetIssueTypeIds() []int64 {
	if o == nil || IsNil(o.IssueTypeIds) {
		var ret []int64
		return ret
	}
	return o.IssueTypeIds
}

// GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetIssueTypeIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.IssueTypeIds) {
		return nil, false
	}
	return o.IssueTypeIds, true
}

// HasIssueTypeIds returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasIssueTypeIds() bool {
	if o != nil && !IsNil(o.IssueTypeIds) {
		return true
	}

	return false
}

// SetIssueTypeIds gets a reference to the given []int64 and assigns it to the IssueTypeIds field.
func (o *DescribeReleaseIssueListRequest) SetIssueTypeIds(v []int64) {
	o.IssueTypeIds = v
}

// GetIssueTypes returns the IssueTypes field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetIssueTypes() []string {
	if o == nil || IsNil(o.IssueTypes) {
		var ret []string
		return ret
	}
	return o.IssueTypes
}

// GetIssueTypesOk returns a tuple with the IssueTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetIssueTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.IssueTypes) {
		return nil, false
	}
	return o.IssueTypes, true
}

// HasIssueTypes returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasIssueTypes() bool {
	if o != nil && !IsNil(o.IssueTypes) {
		return true
	}

	return false
}

// SetIssueTypes gets a reference to the given []string and assigns it to the IssueTypes field.
func (o *DescribeReleaseIssueListRequest) SetIssueTypes(v []string) {
	o.IssueTypes = v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetKeywords() string {
	if o == nil || IsNil(o.Keywords) {
		var ret string
		return ret
	}
	return *o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetKeywordsOk() (*string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given string and assigns it to the Keywords field.
func (o *DescribeReleaseIssueListRequest) SetKeywords(v string) {
	o.Keywords = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetPage() int64 {
	if o == nil || IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetPageOk() (*int64, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *DescribeReleaseIssueListRequest) SetPage(v int64) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeReleaseIssueListRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeReleaseIssueListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeReleaseIssueListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetReleaseCode returns the ReleaseCode field value
func (o *DescribeReleaseIssueListRequest) GetReleaseCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ReleaseCode
}

// GetReleaseCodeOk returns a tuple with the ReleaseCode field value
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetReleaseCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseCode, true
}

// SetReleaseCode sets field value
func (o *DescribeReleaseIssueListRequest) SetReleaseCode(v int64) {
	o.ReleaseCode = v
}

// GetShowImageOutUrl returns the ShowImageOutUrl field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetShowImageOutUrl() bool {
	if o == nil || IsNil(o.ShowImageOutUrl) {
		var ret bool
		return ret
	}
	return *o.ShowImageOutUrl
}

// GetShowImageOutUrlOk returns a tuple with the ShowImageOutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetShowImageOutUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowImageOutUrl) {
		return nil, false
	}
	return o.ShowImageOutUrl, true
}

// HasShowImageOutUrl returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasShowImageOutUrl() bool {
	if o != nil && !IsNil(o.ShowImageOutUrl) {
		return true
	}

	return false
}

// SetShowImageOutUrl gets a reference to the given bool and assigns it to the ShowImageOutUrl field.
func (o *DescribeReleaseIssueListRequest) SetShowImageOutUrl(v bool) {
	o.ShowImageOutUrl = &v
}

// GetShowSubIssues returns the ShowSubIssues field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetShowSubIssues() bool {
	if o == nil || IsNil(o.ShowSubIssues) {
		var ret bool
		return ret
	}
	return *o.ShowSubIssues
}

// GetShowSubIssuesOk returns a tuple with the ShowSubIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetShowSubIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowSubIssues) {
		return nil, false
	}
	return o.ShowSubIssues, true
}

// HasShowSubIssues returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasShowSubIssues() bool {
	if o != nil && !IsNil(o.ShowSubIssues) {
		return true
	}

	return false
}

// SetShowSubIssues gets a reference to the given bool and assigns it to the ShowSubIssues field.
func (o *DescribeReleaseIssueListRequest) SetShowSubIssues(v bool) {
	o.ShowSubIssues = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetSortBy() string {
	if o == nil || IsNil(o.SortBy) {
		var ret string
		return ret
	}
	return *o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetSortByOk() (*string, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given string and assigns it to the SortBy field.
func (o *DescribeReleaseIssueListRequest) SetSortBy(v string) {
	o.SortBy = &v
}

// GetStatusTypes returns the StatusTypes field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetStatusTypes() []int64 {
	if o == nil || IsNil(o.StatusTypes) {
		var ret []int64
		return ret
	}
	return o.StatusTypes
}

// GetStatusTypesOk returns a tuple with the StatusTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetStatusTypesOk() ([]int64, bool) {
	if o == nil || IsNil(o.StatusTypes) {
		return nil, false
	}
	return o.StatusTypes, true
}

// HasStatusTypes returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasStatusTypes() bool {
	if o != nil && !IsNil(o.StatusTypes) {
		return true
	}

	return false
}

// SetStatusTypes gets a reference to the given []int64 and assigns it to the StatusTypes field.
func (o *DescribeReleaseIssueListRequest) SetStatusTypes(v []int64) {
	o.StatusTypes = v
}

// GetWatchers returns the Watchers field value if set, zero value otherwise.
func (o *DescribeReleaseIssueListRequest) GetWatchers() []int64 {
	if o == nil || IsNil(o.Watchers) {
		var ret []int64
		return ret
	}
	return o.Watchers
}

// GetWatchersOk returns a tuple with the Watchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseIssueListRequest) GetWatchersOk() ([]int64, bool) {
	if o == nil || IsNil(o.Watchers) {
		return nil, false
	}
	return o.Watchers, true
}

// HasWatchers returns a boolean if a field has been set.
func (o *DescribeReleaseIssueListRequest) HasWatchers() bool {
	if o != nil && !IsNil(o.Watchers) {
		return true
	}

	return false
}

// SetWatchers gets a reference to the given []int64 and assigns it to the Watchers field.
func (o *DescribeReleaseIssueListRequest) SetWatchers(v []int64) {
	o.Watchers = v
}

func (o DescribeReleaseIssueListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReleaseIssueListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignees) {
		toSerialize["Assignees"] = o.Assignees
	}
	if !IsNil(o.IssueTypeIds) {
		toSerialize["IssueTypeIds"] = o.IssueTypeIds
	}
	if !IsNil(o.IssueTypes) {
		toSerialize["IssueTypes"] = o.IssueTypes
	}
	if !IsNil(o.Keywords) {
		toSerialize["Keywords"] = o.Keywords
	}
	if !IsNil(o.Page) {
		toSerialize["Page"] = o.Page
	}
	if !IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["ReleaseCode"] = o.ReleaseCode
	if !IsNil(o.ShowImageOutUrl) {
		toSerialize["ShowImageOutUrl"] = o.ShowImageOutUrl
	}
	if !IsNil(o.ShowSubIssues) {
		toSerialize["ShowSubIssues"] = o.ShowSubIssues
	}
	if !IsNil(o.SortBy) {
		toSerialize["SortBy"] = o.SortBy
	}
	if !IsNil(o.StatusTypes) {
		toSerialize["StatusTypes"] = o.StatusTypes
	}
	if !IsNil(o.Watchers) {
		toSerialize["Watchers"] = o.Watchers
	}
	return toSerialize, nil
}

func (o *DescribeReleaseIssueListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectName",
		"ReleaseCode",
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

	varDescribeReleaseIssueListRequest := _DescribeReleaseIssueListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeReleaseIssueListRequest)

	if err != nil {
		return err
	}

	*o = DescribeReleaseIssueListRequest(varDescribeReleaseIssueListRequest)

	return err
}

type NullableDescribeReleaseIssueListRequest struct {
	value *DescribeReleaseIssueListRequest
	isSet bool
}

func (v NullableDescribeReleaseIssueListRequest) Get() *DescribeReleaseIssueListRequest {
	return v.value
}

func (v *NullableDescribeReleaseIssueListRequest) Set(val *DescribeReleaseIssueListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReleaseIssueListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReleaseIssueListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReleaseIssueListRequest(val *DescribeReleaseIssueListRequest) *NullableDescribeReleaseIssueListRequest {
	return &NullableDescribeReleaseIssueListRequest{value: val, isSet: true}
}

func (v NullableDescribeReleaseIssueListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReleaseIssueListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


