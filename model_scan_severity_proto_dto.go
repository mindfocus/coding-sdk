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

// checks if the ScanSeverityProtoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScanSeverityProtoDTO{}

// ScanSeverityProtoDTO 扫描门禁信息
type ScanSeverityProtoDTO struct {
	// 创建时间
	CreatedTimestamp *int64 `json:"CreatedTimestamp,omitempty"`
	// 创建人
	Creator *string `json:"Creator,omitempty"`
	// 启用错误信息阈值
	EnableErrorThreshold *bool `json:"EnableErrorThreshold,omitempty"`
	// 启用致命信息阈值
	EnableFatalThreshold *bool `json:"EnableFatalThreshold,omitempty"`
	// 启用提示信息阈值
	EnableInfoThreshold *bool `json:"EnableInfoThreshold,omitempty"`
	// 启用警告信息阈值
	EnableWarningThreshold *bool `json:"EnableWarningThreshold,omitempty"`
	// 错误问题
	Error *int64 `json:"Error,omitempty"`
	// 错误信息阈值
	ErrorThreshold *int32 `json:"ErrorThreshold,omitempty"`
	// 致命问题
	Fatal *int64 `json:"Fatal,omitempty"`
	// 致命信息阈值
	FatalThreshold *int64 `json:"FatalThreshold,omitempty"`
	// 全局开关
	GlobalSwitch *bool `json:"GlobalSwitch,omitempty"`
	// 增量扫描
	Increment *bool `json:"Increment,omitempty"`
	// 提示问题
	Info *int64 `json:"Info,omitempty"`
	// 提示信息阈值
	InfoThreshold *int64 `json:"InfoThreshold,omitempty"`
	// 完成状态
	Status *string `json:"Status,omitempty"`
	// 扫描状态
	StatusCheck *string `json:"StatusCheck,omitempty"`
	// 警告问题
	Warning *int64 `json:"Warning,omitempty"`
	// 警告信息阈值
	WarningThreshold *int64 `json:"WarningThreshold,omitempty"`
}

// NewScanSeverityProtoDTO instantiates a new ScanSeverityProtoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScanSeverityProtoDTO() *ScanSeverityProtoDTO {
	this := ScanSeverityProtoDTO{}
	var creator string = ""
	this.Creator = &creator
	var enableErrorThreshold bool = false
	this.EnableErrorThreshold = &enableErrorThreshold
	var enableFatalThreshold bool = false
	this.EnableFatalThreshold = &enableFatalThreshold
	var enableInfoThreshold bool = false
	this.EnableInfoThreshold = &enableInfoThreshold
	var enableWarningThreshold bool = false
	this.EnableWarningThreshold = &enableWarningThreshold
	var globalSwitch bool = false
	this.GlobalSwitch = &globalSwitch
	var increment bool = false
	this.Increment = &increment
	var status string = ""
	this.Status = &status
	var statusCheck string = ""
	this.StatusCheck = &statusCheck
	return &this
}

// NewScanSeverityProtoDTOWithDefaults instantiates a new ScanSeverityProtoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScanSeverityProtoDTOWithDefaults() *ScanSeverityProtoDTO {
	this := ScanSeverityProtoDTO{}
	var creator string = ""
	this.Creator = &creator
	var enableErrorThreshold bool = false
	this.EnableErrorThreshold = &enableErrorThreshold
	var enableFatalThreshold bool = false
	this.EnableFatalThreshold = &enableFatalThreshold
	var enableInfoThreshold bool = false
	this.EnableInfoThreshold = &enableInfoThreshold
	var enableWarningThreshold bool = false
	this.EnableWarningThreshold = &enableWarningThreshold
	var globalSwitch bool = false
	this.GlobalSwitch = &globalSwitch
	var increment bool = false
	this.Increment = &increment
	var status string = ""
	this.Status = &status
	var statusCheck string = ""
	this.StatusCheck = &statusCheck
	return &this
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetCreatedTimestamp() int64 {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret int64
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetCreatedTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given int64 and assigns it to the CreatedTimestamp field.
func (o *ScanSeverityProtoDTO) SetCreatedTimestamp(v int64) {
	o.CreatedTimestamp = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *ScanSeverityProtoDTO) SetCreator(v string) {
	o.Creator = &v
}

// GetEnableErrorThreshold returns the EnableErrorThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetEnableErrorThreshold() bool {
	if o == nil || IsNil(o.EnableErrorThreshold) {
		var ret bool
		return ret
	}
	return *o.EnableErrorThreshold
}

// GetEnableErrorThresholdOk returns a tuple with the EnableErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetEnableErrorThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableErrorThreshold) {
		return nil, false
	}
	return o.EnableErrorThreshold, true
}

// HasEnableErrorThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasEnableErrorThreshold() bool {
	if o != nil && !IsNil(o.EnableErrorThreshold) {
		return true
	}

	return false
}

// SetEnableErrorThreshold gets a reference to the given bool and assigns it to the EnableErrorThreshold field.
func (o *ScanSeverityProtoDTO) SetEnableErrorThreshold(v bool) {
	o.EnableErrorThreshold = &v
}

// GetEnableFatalThreshold returns the EnableFatalThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetEnableFatalThreshold() bool {
	if o == nil || IsNil(o.EnableFatalThreshold) {
		var ret bool
		return ret
	}
	return *o.EnableFatalThreshold
}

// GetEnableFatalThresholdOk returns a tuple with the EnableFatalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetEnableFatalThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableFatalThreshold) {
		return nil, false
	}
	return o.EnableFatalThreshold, true
}

// HasEnableFatalThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasEnableFatalThreshold() bool {
	if o != nil && !IsNil(o.EnableFatalThreshold) {
		return true
	}

	return false
}

// SetEnableFatalThreshold gets a reference to the given bool and assigns it to the EnableFatalThreshold field.
func (o *ScanSeverityProtoDTO) SetEnableFatalThreshold(v bool) {
	o.EnableFatalThreshold = &v
}

// GetEnableInfoThreshold returns the EnableInfoThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetEnableInfoThreshold() bool {
	if o == nil || IsNil(o.EnableInfoThreshold) {
		var ret bool
		return ret
	}
	return *o.EnableInfoThreshold
}

// GetEnableInfoThresholdOk returns a tuple with the EnableInfoThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetEnableInfoThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInfoThreshold) {
		return nil, false
	}
	return o.EnableInfoThreshold, true
}

// HasEnableInfoThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasEnableInfoThreshold() bool {
	if o != nil && !IsNil(o.EnableInfoThreshold) {
		return true
	}

	return false
}

// SetEnableInfoThreshold gets a reference to the given bool and assigns it to the EnableInfoThreshold field.
func (o *ScanSeverityProtoDTO) SetEnableInfoThreshold(v bool) {
	o.EnableInfoThreshold = &v
}

// GetEnableWarningThreshold returns the EnableWarningThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetEnableWarningThreshold() bool {
	if o == nil || IsNil(o.EnableWarningThreshold) {
		var ret bool
		return ret
	}
	return *o.EnableWarningThreshold
}

// GetEnableWarningThresholdOk returns a tuple with the EnableWarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetEnableWarningThresholdOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableWarningThreshold) {
		return nil, false
	}
	return o.EnableWarningThreshold, true
}

// HasEnableWarningThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasEnableWarningThreshold() bool {
	if o != nil && !IsNil(o.EnableWarningThreshold) {
		return true
	}

	return false
}

// SetEnableWarningThreshold gets a reference to the given bool and assigns it to the EnableWarningThreshold field.
func (o *ScanSeverityProtoDTO) SetEnableWarningThreshold(v bool) {
	o.EnableWarningThreshold = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetError() int64 {
	if o == nil || IsNil(o.Error) {
		var ret int64
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetErrorOk() (*int64, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given int64 and assigns it to the Error field.
func (o *ScanSeverityProtoDTO) SetError(v int64) {
	o.Error = &v
}

// GetErrorThreshold returns the ErrorThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetErrorThreshold() int32 {
	if o == nil || IsNil(o.ErrorThreshold) {
		var ret int32
		return ret
	}
	return *o.ErrorThreshold
}

// GetErrorThresholdOk returns a tuple with the ErrorThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetErrorThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorThreshold) {
		return nil, false
	}
	return o.ErrorThreshold, true
}

// HasErrorThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasErrorThreshold() bool {
	if o != nil && !IsNil(o.ErrorThreshold) {
		return true
	}

	return false
}

// SetErrorThreshold gets a reference to the given int32 and assigns it to the ErrorThreshold field.
func (o *ScanSeverityProtoDTO) SetErrorThreshold(v int32) {
	o.ErrorThreshold = &v
}

// GetFatal returns the Fatal field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetFatal() int64 {
	if o == nil || IsNil(o.Fatal) {
		var ret int64
		return ret
	}
	return *o.Fatal
}

// GetFatalOk returns a tuple with the Fatal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetFatalOk() (*int64, bool) {
	if o == nil || IsNil(o.Fatal) {
		return nil, false
	}
	return o.Fatal, true
}

// HasFatal returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasFatal() bool {
	if o != nil && !IsNil(o.Fatal) {
		return true
	}

	return false
}

// SetFatal gets a reference to the given int64 and assigns it to the Fatal field.
func (o *ScanSeverityProtoDTO) SetFatal(v int64) {
	o.Fatal = &v
}

// GetFatalThreshold returns the FatalThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetFatalThreshold() int64 {
	if o == nil || IsNil(o.FatalThreshold) {
		var ret int64
		return ret
	}
	return *o.FatalThreshold
}

// GetFatalThresholdOk returns a tuple with the FatalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetFatalThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.FatalThreshold) {
		return nil, false
	}
	return o.FatalThreshold, true
}

// HasFatalThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasFatalThreshold() bool {
	if o != nil && !IsNil(o.FatalThreshold) {
		return true
	}

	return false
}

// SetFatalThreshold gets a reference to the given int64 and assigns it to the FatalThreshold field.
func (o *ScanSeverityProtoDTO) SetFatalThreshold(v int64) {
	o.FatalThreshold = &v
}

// GetGlobalSwitch returns the GlobalSwitch field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetGlobalSwitch() bool {
	if o == nil || IsNil(o.GlobalSwitch) {
		var ret bool
		return ret
	}
	return *o.GlobalSwitch
}

// GetGlobalSwitchOk returns a tuple with the GlobalSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetGlobalSwitchOk() (*bool, bool) {
	if o == nil || IsNil(o.GlobalSwitch) {
		return nil, false
	}
	return o.GlobalSwitch, true
}

// HasGlobalSwitch returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasGlobalSwitch() bool {
	if o != nil && !IsNil(o.GlobalSwitch) {
		return true
	}

	return false
}

// SetGlobalSwitch gets a reference to the given bool and assigns it to the GlobalSwitch field.
func (o *ScanSeverityProtoDTO) SetGlobalSwitch(v bool) {
	o.GlobalSwitch = &v
}

// GetIncrement returns the Increment field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetIncrement() bool {
	if o == nil || IsNil(o.Increment) {
		var ret bool
		return ret
	}
	return *o.Increment
}

// GetIncrementOk returns a tuple with the Increment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetIncrementOk() (*bool, bool) {
	if o == nil || IsNil(o.Increment) {
		return nil, false
	}
	return o.Increment, true
}

// HasIncrement returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasIncrement() bool {
	if o != nil && !IsNil(o.Increment) {
		return true
	}

	return false
}

// SetIncrement gets a reference to the given bool and assigns it to the Increment field.
func (o *ScanSeverityProtoDTO) SetIncrement(v bool) {
	o.Increment = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetInfo() int64 {
	if o == nil || IsNil(o.Info) {
		var ret int64
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetInfoOk() (*int64, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given int64 and assigns it to the Info field.
func (o *ScanSeverityProtoDTO) SetInfo(v int64) {
	o.Info = &v
}

// GetInfoThreshold returns the InfoThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetInfoThreshold() int64 {
	if o == nil || IsNil(o.InfoThreshold) {
		var ret int64
		return ret
	}
	return *o.InfoThreshold
}

// GetInfoThresholdOk returns a tuple with the InfoThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetInfoThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.InfoThreshold) {
		return nil, false
	}
	return o.InfoThreshold, true
}

// HasInfoThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasInfoThreshold() bool {
	if o != nil && !IsNil(o.InfoThreshold) {
		return true
	}

	return false
}

// SetInfoThreshold gets a reference to the given int64 and assigns it to the InfoThreshold field.
func (o *ScanSeverityProtoDTO) SetInfoThreshold(v int64) {
	o.InfoThreshold = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ScanSeverityProtoDTO) SetStatus(v string) {
	o.Status = &v
}

// GetStatusCheck returns the StatusCheck field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetStatusCheck() string {
	if o == nil || IsNil(o.StatusCheck) {
		var ret string
		return ret
	}
	return *o.StatusCheck
}

// GetStatusCheckOk returns a tuple with the StatusCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetStatusCheckOk() (*string, bool) {
	if o == nil || IsNil(o.StatusCheck) {
		return nil, false
	}
	return o.StatusCheck, true
}

// HasStatusCheck returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasStatusCheck() bool {
	if o != nil && !IsNil(o.StatusCheck) {
		return true
	}

	return false
}

// SetStatusCheck gets a reference to the given string and assigns it to the StatusCheck field.
func (o *ScanSeverityProtoDTO) SetStatusCheck(v string) {
	o.StatusCheck = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetWarning() int64 {
	if o == nil || IsNil(o.Warning) {
		var ret int64
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetWarningOk() (*int64, bool) {
	if o == nil || IsNil(o.Warning) {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasWarning() bool {
	if o != nil && !IsNil(o.Warning) {
		return true
	}

	return false
}

// SetWarning gets a reference to the given int64 and assigns it to the Warning field.
func (o *ScanSeverityProtoDTO) SetWarning(v int64) {
	o.Warning = &v
}

// GetWarningThreshold returns the WarningThreshold field value if set, zero value otherwise.
func (o *ScanSeverityProtoDTO) GetWarningThreshold() int64 {
	if o == nil || IsNil(o.WarningThreshold) {
		var ret int64
		return ret
	}
	return *o.WarningThreshold
}

// GetWarningThresholdOk returns a tuple with the WarningThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScanSeverityProtoDTO) GetWarningThresholdOk() (*int64, bool) {
	if o == nil || IsNil(o.WarningThreshold) {
		return nil, false
	}
	return o.WarningThreshold, true
}

// HasWarningThreshold returns a boolean if a field has been set.
func (o *ScanSeverityProtoDTO) HasWarningThreshold() bool {
	if o != nil && !IsNil(o.WarningThreshold) {
		return true
	}

	return false
}

// SetWarningThreshold gets a reference to the given int64 and assigns it to the WarningThreshold field.
func (o *ScanSeverityProtoDTO) SetWarningThreshold(v int64) {
	o.WarningThreshold = &v
}

func (o ScanSeverityProtoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScanSeverityProtoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["CreatedTimestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.Creator) {
		toSerialize["Creator"] = o.Creator
	}
	if !IsNil(o.EnableErrorThreshold) {
		toSerialize["EnableErrorThreshold"] = o.EnableErrorThreshold
	}
	if !IsNil(o.EnableFatalThreshold) {
		toSerialize["EnableFatalThreshold"] = o.EnableFatalThreshold
	}
	if !IsNil(o.EnableInfoThreshold) {
		toSerialize["EnableInfoThreshold"] = o.EnableInfoThreshold
	}
	if !IsNil(o.EnableWarningThreshold) {
		toSerialize["EnableWarningThreshold"] = o.EnableWarningThreshold
	}
	if !IsNil(o.Error) {
		toSerialize["Error"] = o.Error
	}
	if !IsNil(o.ErrorThreshold) {
		toSerialize["ErrorThreshold"] = o.ErrorThreshold
	}
	if !IsNil(o.Fatal) {
		toSerialize["Fatal"] = o.Fatal
	}
	if !IsNil(o.FatalThreshold) {
		toSerialize["FatalThreshold"] = o.FatalThreshold
	}
	if !IsNil(o.GlobalSwitch) {
		toSerialize["GlobalSwitch"] = o.GlobalSwitch
	}
	if !IsNil(o.Increment) {
		toSerialize["Increment"] = o.Increment
	}
	if !IsNil(o.Info) {
		toSerialize["Info"] = o.Info
	}
	if !IsNil(o.InfoThreshold) {
		toSerialize["InfoThreshold"] = o.InfoThreshold
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StatusCheck) {
		toSerialize["StatusCheck"] = o.StatusCheck
	}
	if !IsNil(o.Warning) {
		toSerialize["Warning"] = o.Warning
	}
	if !IsNil(o.WarningThreshold) {
		toSerialize["WarningThreshold"] = o.WarningThreshold
	}
	return toSerialize, nil
}

type NullableScanSeverityProtoDTO struct {
	value *ScanSeverityProtoDTO
	isSet bool
}

func (v NullableScanSeverityProtoDTO) Get() *ScanSeverityProtoDTO {
	return v.value
}

func (v *NullableScanSeverityProtoDTO) Set(val *ScanSeverityProtoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableScanSeverityProtoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableScanSeverityProtoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScanSeverityProtoDTO(val *ScanSeverityProtoDTO) *NullableScanSeverityProtoDTO {
	return &NullableScanSeverityProtoDTO{value: val, isSet: true}
}

func (v NullableScanSeverityProtoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScanSeverityProtoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


