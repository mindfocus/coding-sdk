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

// checks if the DescribeDepotMergeRequestsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DescribeDepotMergeRequestsRequest{}

// DescribeDepotMergeRequestsRequest struct for DescribeDepotMergeRequestsRequest
type DescribeDepotMergeRequestsRequest struct {
	// mr创建结束时间
	CreatedAtEndDate *string `json:"CreatedAtEndDate,omitempty"`
	// mr创建开始时间
	CreatedAtStartDate *string `json:"CreatedAtStartDate,omitempty"`
	// 合并请求创建者邮箱列表
	CreatorEmails []string `json:"CreatorEmails,omitempty"`
	// 合并请求创建者 Global Key 列表
	CreatorGlobalKeys []string `json:"CreatorGlobalKeys,omitempty"`
	// 仓库 ID
	DepotId int64 `json:"DepotId"`
	// 仓库路径
	DepotPath *string `json:"DepotPath,omitempty"`
	// 是否升序
	IsSortDirectionAsc *bool `json:"IsSortDirectionAsc,omitempty"`
	// 关键词
	KeyWord *string `json:"KeyWord,omitempty"`
	// 关联标签
	Labels []string `json:"Labels,omitempty"`
	// 合并请求合并者邮箱列表
	MergerEmails []string `json:"MergerEmails,omitempty"`
	// 合并请求合并者 Global Key 列表
	MergerGlobalKeys []string `json:"MergerGlobalKeys,omitempty"`
	// 页数 默认为1
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 每页数量 默认为10
	PageSize *int64 `json:"PageSize,omitempty"`
	// 合并请求评审者邮箱列表
	ReviewerEmails []string `json:"ReviewerEmails,omitempty"`
	// 评审者 Global Key 列表
	ReviewerGlobalKeys []string `json:"ReviewerGlobalKeys,omitempty"`
	// 排序 created_at merged_at action_at
	Sort *string `json:"Sort,omitempty"`
	// 源分支
	SourceBranches []string `json:"SourceBranches,omitempty"`
	// 合并请求状态 OPEN CLOSE ALL ACCEPTED
	Status *string `json:"Status,omitempty"`
	// 目标分支
	TargetBranches []string `json:"TargetBranches,omitempty"`
	// mr更新结束时间
	UpdatedAtEndDate *string `json:"UpdatedAtEndDate,omitempty"`
	// mr更新开始时间
	UpdatedAtStartDate *string `json:"UpdatedAtStartDate,omitempty"`
}

type _DescribeDepotMergeRequestsRequest DescribeDepotMergeRequestsRequest

// NewDescribeDepotMergeRequestsRequest instantiates a new DescribeDepotMergeRequestsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotMergeRequestsRequest(depotId int64) *DescribeDepotMergeRequestsRequest {
	this := DescribeDepotMergeRequestsRequest{}
	this.DepotId = depotId
	return &this
}

// NewDescribeDepotMergeRequestsRequestWithDefaults instantiates a new DescribeDepotMergeRequestsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotMergeRequestsRequestWithDefaults() *DescribeDepotMergeRequestsRequest {
	this := DescribeDepotMergeRequestsRequest{}
	return &this
}

// GetCreatedAtEndDate returns the CreatedAtEndDate field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtEndDate() string {
	if o == nil || IsNil(o.CreatedAtEndDate) {
		var ret string
		return ret
	}
	return *o.CreatedAtEndDate
}

// GetCreatedAtEndDateOk returns a tuple with the CreatedAtEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAtEndDate) {
		return nil, false
	}
	return o.CreatedAtEndDate, true
}

// HasCreatedAtEndDate returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasCreatedAtEndDate() bool {
	if o != nil && !IsNil(o.CreatedAtEndDate) {
		return true
	}

	return false
}

// SetCreatedAtEndDate gets a reference to the given string and assigns it to the CreatedAtEndDate field.
func (o *DescribeDepotMergeRequestsRequest) SetCreatedAtEndDate(v string) {
	o.CreatedAtEndDate = &v
}

// GetCreatedAtStartDate returns the CreatedAtStartDate field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtStartDate() string {
	if o == nil || IsNil(o.CreatedAtStartDate) {
		var ret string
		return ret
	}
	return *o.CreatedAtStartDate
}

// GetCreatedAtStartDateOk returns a tuple with the CreatedAtStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetCreatedAtStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAtStartDate) {
		return nil, false
	}
	return o.CreatedAtStartDate, true
}

// HasCreatedAtStartDate returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasCreatedAtStartDate() bool {
	if o != nil && !IsNil(o.CreatedAtStartDate) {
		return true
	}

	return false
}

// SetCreatedAtStartDate gets a reference to the given string and assigns it to the CreatedAtStartDate field.
func (o *DescribeDepotMergeRequestsRequest) SetCreatedAtStartDate(v string) {
	o.CreatedAtStartDate = &v
}

// GetCreatorEmails returns the CreatorEmails field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetCreatorEmails() []string {
	if o == nil || IsNil(o.CreatorEmails) {
		var ret []string
		return ret
	}
	return o.CreatorEmails
}

// GetCreatorEmailsOk returns a tuple with the CreatorEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetCreatorEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatorEmails) {
		return nil, false
	}
	return o.CreatorEmails, true
}

// HasCreatorEmails returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasCreatorEmails() bool {
	if o != nil && !IsNil(o.CreatorEmails) {
		return true
	}

	return false
}

// SetCreatorEmails gets a reference to the given []string and assigns it to the CreatorEmails field.
func (o *DescribeDepotMergeRequestsRequest) SetCreatorEmails(v []string) {
	o.CreatorEmails = v
}

// GetCreatorGlobalKeys returns the CreatorGlobalKeys field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetCreatorGlobalKeys() []string {
	if o == nil || IsNil(o.CreatorGlobalKeys) {
		var ret []string
		return ret
	}
	return o.CreatorGlobalKeys
}

// GetCreatorGlobalKeysOk returns a tuple with the CreatorGlobalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetCreatorGlobalKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.CreatorGlobalKeys) {
		return nil, false
	}
	return o.CreatorGlobalKeys, true
}

// HasCreatorGlobalKeys returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasCreatorGlobalKeys() bool {
	if o != nil && !IsNil(o.CreatorGlobalKeys) {
		return true
	}

	return false
}

// SetCreatorGlobalKeys gets a reference to the given []string and assigns it to the CreatorGlobalKeys field.
func (o *DescribeDepotMergeRequestsRequest) SetCreatorGlobalKeys(v []string) {
	o.CreatorGlobalKeys = v
}

// GetDepotId returns the DepotId field value
func (o *DescribeDepotMergeRequestsRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeDepotMergeRequestsRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetDepotPath() string {
	if o == nil || IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasDepotPath() bool {
	if o != nil && !IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeDepotMergeRequestsRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetIsSortDirectionAsc returns the IsSortDirectionAsc field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetIsSortDirectionAsc() bool {
	if o == nil || IsNil(o.IsSortDirectionAsc) {
		var ret bool
		return ret
	}
	return *o.IsSortDirectionAsc
}

// GetIsSortDirectionAscOk returns a tuple with the IsSortDirectionAsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetIsSortDirectionAscOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSortDirectionAsc) {
		return nil, false
	}
	return o.IsSortDirectionAsc, true
}

// HasIsSortDirectionAsc returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasIsSortDirectionAsc() bool {
	if o != nil && !IsNil(o.IsSortDirectionAsc) {
		return true
	}

	return false
}

// SetIsSortDirectionAsc gets a reference to the given bool and assigns it to the IsSortDirectionAsc field.
func (o *DescribeDepotMergeRequestsRequest) SetIsSortDirectionAsc(v bool) {
	o.IsSortDirectionAsc = &v
}

// GetKeyWord returns the KeyWord field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetKeyWord() string {
	if o == nil || IsNil(o.KeyWord) {
		var ret string
		return ret
	}
	return *o.KeyWord
}

// GetKeyWordOk returns a tuple with the KeyWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetKeyWordOk() (*string, bool) {
	if o == nil || IsNil(o.KeyWord) {
		return nil, false
	}
	return o.KeyWord, true
}

// HasKeyWord returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasKeyWord() bool {
	if o != nil && !IsNil(o.KeyWord) {
		return true
	}

	return false
}

// SetKeyWord gets a reference to the given string and assigns it to the KeyWord field.
func (o *DescribeDepotMergeRequestsRequest) SetKeyWord(v string) {
	o.KeyWord = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *DescribeDepotMergeRequestsRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetMergerEmails returns the MergerEmails field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetMergerEmails() []string {
	if o == nil || IsNil(o.MergerEmails) {
		var ret []string
		return ret
	}
	return o.MergerEmails
}

// GetMergerEmailsOk returns a tuple with the MergerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetMergerEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.MergerEmails) {
		return nil, false
	}
	return o.MergerEmails, true
}

// HasMergerEmails returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasMergerEmails() bool {
	if o != nil && !IsNil(o.MergerEmails) {
		return true
	}

	return false
}

// SetMergerEmails gets a reference to the given []string and assigns it to the MergerEmails field.
func (o *DescribeDepotMergeRequestsRequest) SetMergerEmails(v []string) {
	o.MergerEmails = v
}

// GetMergerGlobalKeys returns the MergerGlobalKeys field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetMergerGlobalKeys() []string {
	if o == nil || IsNil(o.MergerGlobalKeys) {
		var ret []string
		return ret
	}
	return o.MergerGlobalKeys
}

// GetMergerGlobalKeysOk returns a tuple with the MergerGlobalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetMergerGlobalKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.MergerGlobalKeys) {
		return nil, false
	}
	return o.MergerGlobalKeys, true
}

// HasMergerGlobalKeys returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasMergerGlobalKeys() bool {
	if o != nil && !IsNil(o.MergerGlobalKeys) {
		return true
	}

	return false
}

// SetMergerGlobalKeys gets a reference to the given []string and assigns it to the MergerGlobalKeys field.
func (o *DescribeDepotMergeRequestsRequest) SetMergerGlobalKeys(v []string) {
	o.MergerGlobalKeys = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetPageNumber() int64 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeDepotMergeRequestsRequest) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetPageSize() int64 {
	if o == nil || IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeDepotMergeRequestsRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetReviewerEmails returns the ReviewerEmails field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetReviewerEmails() []string {
	if o == nil || IsNil(o.ReviewerEmails) {
		var ret []string
		return ret
	}
	return o.ReviewerEmails
}

// GetReviewerEmailsOk returns a tuple with the ReviewerEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetReviewerEmailsOk() ([]string, bool) {
	if o == nil || IsNil(o.ReviewerEmails) {
		return nil, false
	}
	return o.ReviewerEmails, true
}

// HasReviewerEmails returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasReviewerEmails() bool {
	if o != nil && !IsNil(o.ReviewerEmails) {
		return true
	}

	return false
}

// SetReviewerEmails gets a reference to the given []string and assigns it to the ReviewerEmails field.
func (o *DescribeDepotMergeRequestsRequest) SetReviewerEmails(v []string) {
	o.ReviewerEmails = v
}

// GetReviewerGlobalKeys returns the ReviewerGlobalKeys field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetReviewerGlobalKeys() []string {
	if o == nil || IsNil(o.ReviewerGlobalKeys) {
		var ret []string
		return ret
	}
	return o.ReviewerGlobalKeys
}

// GetReviewerGlobalKeysOk returns a tuple with the ReviewerGlobalKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetReviewerGlobalKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.ReviewerGlobalKeys) {
		return nil, false
	}
	return o.ReviewerGlobalKeys, true
}

// HasReviewerGlobalKeys returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasReviewerGlobalKeys() bool {
	if o != nil && !IsNil(o.ReviewerGlobalKeys) {
		return true
	}

	return false
}

// SetReviewerGlobalKeys gets a reference to the given []string and assigns it to the ReviewerGlobalKeys field.
func (o *DescribeDepotMergeRequestsRequest) SetReviewerGlobalKeys(v []string) {
	o.ReviewerGlobalKeys = v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetSort() string {
	if o == nil || IsNil(o.Sort) {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetSortOk() (*string, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *DescribeDepotMergeRequestsRequest) SetSort(v string) {
	o.Sort = &v
}

// GetSourceBranches returns the SourceBranches field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetSourceBranches() []string {
	if o == nil || IsNil(o.SourceBranches) {
		var ret []string
		return ret
	}
	return o.SourceBranches
}

// GetSourceBranchesOk returns a tuple with the SourceBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetSourceBranchesOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceBranches) {
		return nil, false
	}
	return o.SourceBranches, true
}

// HasSourceBranches returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasSourceBranches() bool {
	if o != nil && !IsNil(o.SourceBranches) {
		return true
	}

	return false
}

// SetSourceBranches gets a reference to the given []string and assigns it to the SourceBranches field.
func (o *DescribeDepotMergeRequestsRequest) SetSourceBranches(v []string) {
	o.SourceBranches = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DescribeDepotMergeRequestsRequest) SetStatus(v string) {
	o.Status = &v
}

// GetTargetBranches returns the TargetBranches field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetTargetBranches() []string {
	if o == nil || IsNil(o.TargetBranches) {
		var ret []string
		return ret
	}
	return o.TargetBranches
}

// GetTargetBranchesOk returns a tuple with the TargetBranches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetTargetBranchesOk() ([]string, bool) {
	if o == nil || IsNil(o.TargetBranches) {
		return nil, false
	}
	return o.TargetBranches, true
}

// HasTargetBranches returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasTargetBranches() bool {
	if o != nil && !IsNil(o.TargetBranches) {
		return true
	}

	return false
}

// SetTargetBranches gets a reference to the given []string and assigns it to the TargetBranches field.
func (o *DescribeDepotMergeRequestsRequest) SetTargetBranches(v []string) {
	o.TargetBranches = v
}

// GetUpdatedAtEndDate returns the UpdatedAtEndDate field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtEndDate() string {
	if o == nil || IsNil(o.UpdatedAtEndDate) {
		var ret string
		return ret
	}
	return *o.UpdatedAtEndDate
}

// GetUpdatedAtEndDateOk returns a tuple with the UpdatedAtEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAtEndDate) {
		return nil, false
	}
	return o.UpdatedAtEndDate, true
}

// HasUpdatedAtEndDate returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasUpdatedAtEndDate() bool {
	if o != nil && !IsNil(o.UpdatedAtEndDate) {
		return true
	}

	return false
}

// SetUpdatedAtEndDate gets a reference to the given string and assigns it to the UpdatedAtEndDate field.
func (o *DescribeDepotMergeRequestsRequest) SetUpdatedAtEndDate(v string) {
	o.UpdatedAtEndDate = &v
}

// GetUpdatedAtStartDate returns the UpdatedAtStartDate field value if set, zero value otherwise.
func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtStartDate() string {
	if o == nil || IsNil(o.UpdatedAtStartDate) {
		var ret string
		return ret
	}
	return *o.UpdatedAtStartDate
}

// GetUpdatedAtStartDateOk returns a tuple with the UpdatedAtStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotMergeRequestsRequest) GetUpdatedAtStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAtStartDate) {
		return nil, false
	}
	return o.UpdatedAtStartDate, true
}

// HasUpdatedAtStartDate returns a boolean if a field has been set.
func (o *DescribeDepotMergeRequestsRequest) HasUpdatedAtStartDate() bool {
	if o != nil && !IsNil(o.UpdatedAtStartDate) {
		return true
	}

	return false
}

// SetUpdatedAtStartDate gets a reference to the given string and assigns it to the UpdatedAtStartDate field.
func (o *DescribeDepotMergeRequestsRequest) SetUpdatedAtStartDate(v string) {
	o.UpdatedAtStartDate = &v
}

func (o DescribeDepotMergeRequestsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotMergeRequestsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAtEndDate) {
		toSerialize["CreatedAtEndDate"] = o.CreatedAtEndDate
	}
	if !IsNil(o.CreatedAtStartDate) {
		toSerialize["CreatedAtStartDate"] = o.CreatedAtStartDate
	}
	if !IsNil(o.CreatorEmails) {
		toSerialize["CreatorEmails"] = o.CreatorEmails
	}
	if !IsNil(o.CreatorGlobalKeys) {
		toSerialize["CreatorGlobalKeys"] = o.CreatorGlobalKeys
	}
	toSerialize["DepotId"] = o.DepotId
	if !IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	if !IsNil(o.IsSortDirectionAsc) {
		toSerialize["IsSortDirectionAsc"] = o.IsSortDirectionAsc
	}
	if !IsNil(o.KeyWord) {
		toSerialize["KeyWord"] = o.KeyWord
	}
	if !IsNil(o.Labels) {
		toSerialize["Labels"] = o.Labels
	}
	if !IsNil(o.MergerEmails) {
		toSerialize["MergerEmails"] = o.MergerEmails
	}
	if !IsNil(o.MergerGlobalKeys) {
		toSerialize["MergerGlobalKeys"] = o.MergerGlobalKeys
	}
	if !IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !IsNil(o.ReviewerEmails) {
		toSerialize["ReviewerEmails"] = o.ReviewerEmails
	}
	if !IsNil(o.ReviewerGlobalKeys) {
		toSerialize["ReviewerGlobalKeys"] = o.ReviewerGlobalKeys
	}
	if !IsNil(o.Sort) {
		toSerialize["Sort"] = o.Sort
	}
	if !IsNil(o.SourceBranches) {
		toSerialize["SourceBranches"] = o.SourceBranches
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TargetBranches) {
		toSerialize["TargetBranches"] = o.TargetBranches
	}
	if !IsNil(o.UpdatedAtEndDate) {
		toSerialize["UpdatedAtEndDate"] = o.UpdatedAtEndDate
	}
	if !IsNil(o.UpdatedAtStartDate) {
		toSerialize["UpdatedAtStartDate"] = o.UpdatedAtStartDate
	}
	return toSerialize, nil
}

func (o *DescribeDepotMergeRequestsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
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

	varDescribeDepotMergeRequestsRequest := _DescribeDepotMergeRequestsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeDepotMergeRequestsRequest)

	if err != nil {
		return err
	}

	*o = DescribeDepotMergeRequestsRequest(varDescribeDepotMergeRequestsRequest)

	return err
}

type NullableDescribeDepotMergeRequestsRequest struct {
	value *DescribeDepotMergeRequestsRequest
	isSet bool
}

func (v NullableDescribeDepotMergeRequestsRequest) Get() *DescribeDepotMergeRequestsRequest {
	return v.value
}

func (v *NullableDescribeDepotMergeRequestsRequest) Set(val *DescribeDepotMergeRequestsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotMergeRequestsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotMergeRequestsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotMergeRequestsRequest(val *DescribeDepotMergeRequestsRequest) *NullableDescribeDepotMergeRequestsRequest {
	return &NullableDescribeDepotMergeRequestsRequest{value: val, isSet: true}
}

func (v NullableDescribeDepotMergeRequestsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotMergeRequestsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


