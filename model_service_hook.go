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

// checks if the ServiceHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceHook{}

// ServiceHook ServiceHook 对象
type ServiceHook struct {
	// 发送类型
	Action *string `json:"Action,omitempty"`
	// 发送类型描述
	ActionLabel *string `json:"ActionLabel,omitempty"`
	// 发送行为属性
	ActionProperties *string `json:"ActionProperties,omitempty"`
	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 创建者编号
	CreatorBy *int64 `json:"CreatorBy,omitempty"`
	CreatorByUser *ServiceHookUser `json:"CreatorByUser,omitempty"`
	// 事件开关
	Enabled *bool `json:"Enabled,omitempty"`
	// 事件列表
	Event []string `json:"Event,omitempty"`
	// 事件描述列表
	EventLabel []string `json:"EventLabel,omitempty"`
	// 过滤器属性
	FilterProperties *string `json:"FilterProperties,omitempty"`
	// Service Hook 编号
	Id *string `json:"Id,omitempty"`
	// 最近发送成功时间
	LastSucceedAt *int64 `json:"LastSucceedAt,omitempty"`
	// Service Hook 名称
	Name *string `json:"Name,omitempty"`
	// 服务类型
	Service *string `json:"Service,omitempty"`
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty"`
	// 发送状态
	Status *int64 `json:"Status,omitempty"`
	// 目标数据编号
	TargetId *int64 `json:"TargetId,omitempty"`
	// 目标数据类型：Project、Team,Space,Program
	TargetType *string `json:"TargetType,omitempty"`
	// 更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
	// 更新者编号
	UpdatedBy *int64 `json:"UpdatedBy,omitempty"`
	UpdatedByUser *ServiceHookUser `json:"UpdatedByUser,omitempty"`
	// 版本
	Version *int64 `json:"Version,omitempty"`
}

// NewServiceHook instantiates a new ServiceHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceHook() *ServiceHook {
	this := ServiceHook{}
	var action string = ""
	this.Action = &action
	var actionLabel string = ""
	this.ActionLabel = &actionLabel
	var actionProperties string = ""
	this.ActionProperties = &actionProperties
	var enabled bool = false
	this.Enabled = &enabled
	var filterProperties string = ""
	this.FilterProperties = &filterProperties
	var id string = ""
	this.Id = &id
	var name string = ""
	this.Name = &name
	var service string = ""
	this.Service = &service
	var serviceName string = ""
	this.ServiceName = &serviceName
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// NewServiceHookWithDefaults instantiates a new ServiceHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceHookWithDefaults() *ServiceHook {
	this := ServiceHook{}
	var action string = ""
	this.Action = &action
	var actionLabel string = ""
	this.ActionLabel = &actionLabel
	var actionProperties string = ""
	this.ActionProperties = &actionProperties
	var enabled bool = false
	this.Enabled = &enabled
	var filterProperties string = ""
	this.FilterProperties = &filterProperties
	var id string = ""
	this.Id = &id
	var name string = ""
	this.Name = &name
	var service string = ""
	this.Service = &service
	var serviceName string = ""
	this.ServiceName = &serviceName
	var targetType string = ""
	this.TargetType = &targetType
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ServiceHook) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ServiceHook) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ServiceHook) SetAction(v string) {
	o.Action = &v
}

// GetActionLabel returns the ActionLabel field value if set, zero value otherwise.
func (o *ServiceHook) GetActionLabel() string {
	if o == nil || IsNil(o.ActionLabel) {
		var ret string
		return ret
	}
	return *o.ActionLabel
}

// GetActionLabelOk returns a tuple with the ActionLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionLabelOk() (*string, bool) {
	if o == nil || IsNil(o.ActionLabel) {
		return nil, false
	}
	return o.ActionLabel, true
}

// HasActionLabel returns a boolean if a field has been set.
func (o *ServiceHook) HasActionLabel() bool {
	if o != nil && !IsNil(o.ActionLabel) {
		return true
	}

	return false
}

// SetActionLabel gets a reference to the given string and assigns it to the ActionLabel field.
func (o *ServiceHook) SetActionLabel(v string) {
	o.ActionLabel = &v
}

// GetActionProperties returns the ActionProperties field value if set, zero value otherwise.
func (o *ServiceHook) GetActionProperties() string {
	if o == nil || IsNil(o.ActionProperties) {
		var ret string
		return ret
	}
	return *o.ActionProperties
}

// GetActionPropertiesOk returns a tuple with the ActionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetActionPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.ActionProperties) {
		return nil, false
	}
	return o.ActionProperties, true
}

// HasActionProperties returns a boolean if a field has been set.
func (o *ServiceHook) HasActionProperties() bool {
	if o != nil && !IsNil(o.ActionProperties) {
		return true
	}

	return false
}

// SetActionProperties gets a reference to the given string and assigns it to the ActionProperties field.
func (o *ServiceHook) SetActionProperties(v string) {
	o.ActionProperties = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *ServiceHook) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetCreatorBy returns the CreatorBy field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatorBy() int64 {
	if o == nil || IsNil(o.CreatorBy) {
		var ret int64
		return ret
	}
	return *o.CreatorBy
}

// GetCreatorByOk returns a tuple with the CreatorBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatorByOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatorBy) {
		return nil, false
	}
	return o.CreatorBy, true
}

// HasCreatorBy returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatorBy() bool {
	if o != nil && !IsNil(o.CreatorBy) {
		return true
	}

	return false
}

// SetCreatorBy gets a reference to the given int64 and assigns it to the CreatorBy field.
func (o *ServiceHook) SetCreatorBy(v int64) {
	o.CreatorBy = &v
}

// GetCreatorByUser returns the CreatorByUser field value if set, zero value otherwise.
func (o *ServiceHook) GetCreatorByUser() ServiceHookUser {
	if o == nil || IsNil(o.CreatorByUser) {
		var ret ServiceHookUser
		return ret
	}
	return *o.CreatorByUser
}

// GetCreatorByUserOk returns a tuple with the CreatorByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetCreatorByUserOk() (*ServiceHookUser, bool) {
	if o == nil || IsNil(o.CreatorByUser) {
		return nil, false
	}
	return o.CreatorByUser, true
}

// HasCreatorByUser returns a boolean if a field has been set.
func (o *ServiceHook) HasCreatorByUser() bool {
	if o != nil && !IsNil(o.CreatorByUser) {
		return true
	}

	return false
}

// SetCreatorByUser gets a reference to the given ServiceHookUser and assigns it to the CreatorByUser field.
func (o *ServiceHook) SetCreatorByUser(v ServiceHookUser) {
	o.CreatorByUser = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceHook) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceHook) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceHook) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *ServiceHook) GetEvent() []string {
	if o == nil || IsNil(o.Event) {
		var ret []string
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEventOk() ([]string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *ServiceHook) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given []string and assigns it to the Event field.
func (o *ServiceHook) SetEvent(v []string) {
	o.Event = v
}

// GetEventLabel returns the EventLabel field value if set, zero value otherwise.
func (o *ServiceHook) GetEventLabel() []string {
	if o == nil || IsNil(o.EventLabel) {
		var ret []string
		return ret
	}
	return o.EventLabel
}

// GetEventLabelOk returns a tuple with the EventLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetEventLabelOk() ([]string, bool) {
	if o == nil || IsNil(o.EventLabel) {
		return nil, false
	}
	return o.EventLabel, true
}

// HasEventLabel returns a boolean if a field has been set.
func (o *ServiceHook) HasEventLabel() bool {
	if o != nil && !IsNil(o.EventLabel) {
		return true
	}

	return false
}

// SetEventLabel gets a reference to the given []string and assigns it to the EventLabel field.
func (o *ServiceHook) SetEventLabel(v []string) {
	o.EventLabel = v
}

// GetFilterProperties returns the FilterProperties field value if set, zero value otherwise.
func (o *ServiceHook) GetFilterProperties() string {
	if o == nil || IsNil(o.FilterProperties) {
		var ret string
		return ret
	}
	return *o.FilterProperties
}

// GetFilterPropertiesOk returns a tuple with the FilterProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetFilterPropertiesOk() (*string, bool) {
	if o == nil || IsNil(o.FilterProperties) {
		return nil, false
	}
	return o.FilterProperties, true
}

// HasFilterProperties returns a boolean if a field has been set.
func (o *ServiceHook) HasFilterProperties() bool {
	if o != nil && !IsNil(o.FilterProperties) {
		return true
	}

	return false
}

// SetFilterProperties gets a reference to the given string and assigns it to the FilterProperties field.
func (o *ServiceHook) SetFilterProperties(v string) {
	o.FilterProperties = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceHook) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceHook) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceHook) SetId(v string) {
	o.Id = &v
}

// GetLastSucceedAt returns the LastSucceedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetLastSucceedAt() int64 {
	if o == nil || IsNil(o.LastSucceedAt) {
		var ret int64
		return ret
	}
	return *o.LastSucceedAt
}

// GetLastSucceedAtOk returns a tuple with the LastSucceedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetLastSucceedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.LastSucceedAt) {
		return nil, false
	}
	return o.LastSucceedAt, true
}

// HasLastSucceedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasLastSucceedAt() bool {
	if o != nil && !IsNil(o.LastSucceedAt) {
		return true
	}

	return false
}

// SetLastSucceedAt gets a reference to the given int64 and assigns it to the LastSucceedAt field.
func (o *ServiceHook) SetLastSucceedAt(v int64) {
	o.LastSucceedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceHook) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceHook) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceHook) SetName(v string) {
	o.Name = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceHook) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceHook) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ServiceHook) SetService(v string) {
	o.Service = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceHook) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceHook) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceHook) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceHook) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceHook) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ServiceHook) SetStatus(v int64) {
	o.Status = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *ServiceHook) GetTargetId() int64 {
	if o == nil || IsNil(o.TargetId) {
		var ret int64
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetTargetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *ServiceHook) HasTargetId() bool {
	if o != nil && !IsNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given int64 and assigns it to the TargetId field.
func (o *ServiceHook) SetTargetId(v int64) {
	o.TargetId = &v
}

// GetTargetType returns the TargetType field value if set, zero value otherwise.
func (o *ServiceHook) GetTargetType() string {
	if o == nil || IsNil(o.TargetType) {
		var ret string
		return ret
	}
	return *o.TargetType
}

// GetTargetTypeOk returns a tuple with the TargetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetTargetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TargetType) {
		return nil, false
	}
	return o.TargetType, true
}

// HasTargetType returns a boolean if a field has been set.
func (o *ServiceHook) HasTargetType() bool {
	if o != nil && !IsNil(o.TargetType) {
		return true
	}

	return false
}

// SetTargetType gets a reference to the given string and assigns it to the TargetType field.
func (o *ServiceHook) SetTargetType(v string) {
	o.TargetType = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *ServiceHook) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedBy() int64 {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret int64
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedByOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given int64 and assigns it to the UpdatedBy field.
func (o *ServiceHook) SetUpdatedBy(v int64) {
	o.UpdatedBy = &v
}

// GetUpdatedByUser returns the UpdatedByUser field value if set, zero value otherwise.
func (o *ServiceHook) GetUpdatedByUser() ServiceHookUser {
	if o == nil || IsNil(o.UpdatedByUser) {
		var ret ServiceHookUser
		return ret
	}
	return *o.UpdatedByUser
}

// GetUpdatedByUserOk returns a tuple with the UpdatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetUpdatedByUserOk() (*ServiceHookUser, bool) {
	if o == nil || IsNil(o.UpdatedByUser) {
		return nil, false
	}
	return o.UpdatedByUser, true
}

// HasUpdatedByUser returns a boolean if a field has been set.
func (o *ServiceHook) HasUpdatedByUser() bool {
	if o != nil && !IsNil(o.UpdatedByUser) {
		return true
	}

	return false
}

// SetUpdatedByUser gets a reference to the given ServiceHookUser and assigns it to the UpdatedByUser field.
func (o *ServiceHook) SetUpdatedByUser(v ServiceHookUser) {
	o.UpdatedByUser = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceHook) GetVersion() int64 {
	if o == nil || IsNil(o.Version) {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceHook) GetVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceHook) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *ServiceHook) SetVersion(v int64) {
	o.Version = &v
}

func (o ServiceHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if !IsNil(o.ActionLabel) {
		toSerialize["ActionLabel"] = o.ActionLabel
	}
	if !IsNil(o.ActionProperties) {
		toSerialize["ActionProperties"] = o.ActionProperties
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.CreatorBy) {
		toSerialize["CreatorBy"] = o.CreatorBy
	}
	if !IsNil(o.CreatorByUser) {
		toSerialize["CreatorByUser"] = o.CreatorByUser
	}
	if !IsNil(o.Enabled) {
		toSerialize["Enabled"] = o.Enabled
	}
	if !IsNil(o.Event) {
		toSerialize["Event"] = o.Event
	}
	if !IsNil(o.EventLabel) {
		toSerialize["EventLabel"] = o.EventLabel
	}
	if !IsNil(o.FilterProperties) {
		toSerialize["FilterProperties"] = o.FilterProperties
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.LastSucceedAt) {
		toSerialize["LastSucceedAt"] = o.LastSucceedAt
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Service) {
		toSerialize["Service"] = o.Service
	}
	if !IsNil(o.ServiceName) {
		toSerialize["ServiceName"] = o.ServiceName
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TargetId) {
		toSerialize["TargetId"] = o.TargetId
	}
	if !IsNil(o.TargetType) {
		toSerialize["TargetType"] = o.TargetType
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["UpdatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedByUser) {
		toSerialize["UpdatedByUser"] = o.UpdatedByUser
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	return toSerialize, nil
}

type NullableServiceHook struct {
	value *ServiceHook
	isSet bool
}

func (v NullableServiceHook) Get() *ServiceHook {
	return v.value
}

func (v *NullableServiceHook) Set(val *ServiceHook) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceHook) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceHook(val *ServiceHook) *NullableServiceHook {
	return &NullableServiceHook{value: val, isSet: true}
}

func (v NullableServiceHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


