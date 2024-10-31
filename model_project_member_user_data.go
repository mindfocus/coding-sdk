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

// checks if the ProjectMemberUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMemberUserData{}

// ProjectMemberUserData 用户成员信息
type ProjectMemberUserData struct {
	// 头像
	Avatar *string `json:"Avatar,omitempty"`
	// 邮箱
	Email *string `json:"Email,omitempty"`
	// 邮箱是否验证 0 否 /1 是
	EmailValidation *int32 `json:"EmailValidation,omitempty"`
	// 用户 GK
	GlobalKey *string `json:"GlobalKey,omitempty"`
	// 用户Id
	Id *int32 `json:"Id,omitempty"`
	// 用户名
	Name *string `json:"Name,omitempty"`
	// 用户名拼音
	NamePinYin *string `json:"NamePinYin,omitempty"`
	// 手机号
	Phone *string `json:"Phone,omitempty"`
	// 手机是否验证 0 否 /1 是
	PhoneValidation *int32 `json:"PhoneValidation,omitempty"`
	// 用户组
	Roles []Role `json:"Roles,omitempty"`
	// 用户状态 0不活跃，1活跃，-1被锁定，-2锁定登录，-3退出团队
	Status *int64 `json:"Status,omitempty"`
	// 团队Id
	TeamId *int32 `json:"TeamId,omitempty"`
	// 拓展字段
	UniqueExtField *string `json:"UniqueExtField,omitempty"`
	DepartmentMember ProjectMemberDepartmentMember `json:"DepartmentMember,omitempty"`
}

// NewProjectMemberUserData instantiates a new ProjectMemberUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMemberUserData() *ProjectMemberUserData {
	this := ProjectMemberUserData{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var phone string = ""
	this.Phone = &phone
	var teamId int32 = 0
	this.TeamId = &teamId
	var uniqueExtField string = ""
	this.UniqueExtField = &uniqueExtField
	return &this
}

// NewProjectMemberUserDataWithDefaults instantiates a new ProjectMemberUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMemberUserDataWithDefaults() *ProjectMemberUserData {
	this := ProjectMemberUserData{}
	var avatar string = ""
	this.Avatar = &avatar
	var email string = ""
	this.Email = &email
	var globalKey string = ""
	this.GlobalKey = &globalKey
	var name string = ""
	this.Name = &name
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var phone string = ""
	this.Phone = &phone
	var teamId int32 = 0
	this.TeamId = &teamId
	var uniqueExtField string = ""
	this.UniqueExtField = &uniqueExtField
	return &this
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetAvatar() string {
	if o == nil || IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetAvatarOk() (*string, bool) {
	if o == nil || IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasAvatar() bool {
	if o != nil && !IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *ProjectMemberUserData) SetAvatar(v string) {
	o.Avatar = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ProjectMemberUserData) SetEmail(v string) {
	o.Email = &v
}

// GetEmailValidation returns the EmailValidation field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetEmailValidation() int32 {
	if o == nil || IsNil(o.EmailValidation) {
		var ret int32
		return ret
	}
	return *o.EmailValidation
}

// GetEmailValidationOk returns a tuple with the EmailValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetEmailValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.EmailValidation) {
		return nil, false
	}
	return o.EmailValidation, true
}

// HasEmailValidation returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasEmailValidation() bool {
	if o != nil && !IsNil(o.EmailValidation) {
		return true
	}

	return false
}

// SetEmailValidation gets a reference to the given int32 and assigns it to the EmailValidation field.
func (o *ProjectMemberUserData) SetEmailValidation(v int32) {
	o.EmailValidation = &v
}

// GetGlobalKey returns the GlobalKey field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetGlobalKey() string {
	if o == nil || IsNil(o.GlobalKey) {
		var ret string
		return ret
	}
	return *o.GlobalKey
}

// GetGlobalKeyOk returns a tuple with the GlobalKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetGlobalKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalKey) {
		return nil, false
	}
	return o.GlobalKey, true
}

// HasGlobalKey returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasGlobalKey() bool {
	if o != nil && !IsNil(o.GlobalKey) {
		return true
	}

	return false
}

// SetGlobalKey gets a reference to the given string and assigns it to the GlobalKey field.
func (o *ProjectMemberUserData) SetGlobalKey(v string) {
	o.GlobalKey = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProjectMemberUserData) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectMemberUserData) SetName(v string) {
	o.Name = &v
}

// GetNamePinYin returns the NamePinYin field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetNamePinYin() string {
	if o == nil || IsNil(o.NamePinYin) {
		var ret string
		return ret
	}
	return *o.NamePinYin
}

// GetNamePinYinOk returns a tuple with the NamePinYin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetNamePinYinOk() (*string, bool) {
	if o == nil || IsNil(o.NamePinYin) {
		return nil, false
	}
	return o.NamePinYin, true
}

// HasNamePinYin returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasNamePinYin() bool {
	if o != nil && !IsNil(o.NamePinYin) {
		return true
	}

	return false
}

// SetNamePinYin gets a reference to the given string and assigns it to the NamePinYin field.
func (o *ProjectMemberUserData) SetNamePinYin(v string) {
	o.NamePinYin = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *ProjectMemberUserData) SetPhone(v string) {
	o.Phone = &v
}

// GetPhoneValidation returns the PhoneValidation field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetPhoneValidation() int32 {
	if o == nil || IsNil(o.PhoneValidation) {
		var ret int32
		return ret
	}
	return *o.PhoneValidation
}

// GetPhoneValidationOk returns a tuple with the PhoneValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetPhoneValidationOk() (*int32, bool) {
	if o == nil || IsNil(o.PhoneValidation) {
		return nil, false
	}
	return o.PhoneValidation, true
}

// HasPhoneValidation returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasPhoneValidation() bool {
	if o != nil && !IsNil(o.PhoneValidation) {
		return true
	}

	return false
}

// SetPhoneValidation gets a reference to the given int32 and assigns it to the PhoneValidation field.
func (o *ProjectMemberUserData) SetPhoneValidation(v int32) {
	o.PhoneValidation = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetRoles() []Role {
	if o == nil || IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetRolesOk() ([]Role, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *ProjectMemberUserData) SetRoles(v []Role) {
	o.Roles = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetStatus() int64 {
	if o == nil || IsNil(o.Status) {
		var ret int64
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetStatusOk() (*int64, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int64 and assigns it to the Status field.
func (o *ProjectMemberUserData) SetStatus(v int64) {
	o.Status = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetTeamId() int32 {
	if o == nil || IsNil(o.TeamId) {
		var ret int32
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetTeamIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given int32 and assigns it to the TeamId field.
func (o *ProjectMemberUserData) SetTeamId(v int32) {
	o.TeamId = &v
}

// GetUniqueExtField returns the UniqueExtField field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetUniqueExtField() string {
	if o == nil || IsNil(o.UniqueExtField) {
		var ret string
		return ret
	}
	return *o.UniqueExtField
}

// GetUniqueExtFieldOk returns a tuple with the UniqueExtField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetUniqueExtFieldOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueExtField) {
		return nil, false
	}
	return o.UniqueExtField, true
}

// HasUniqueExtField returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasUniqueExtField() bool {
	if o != nil && !IsNil(o.UniqueExtField) {
		return true
	}

	return false
}

// SetUniqueExtField gets a reference to the given string and assigns it to the UniqueExtField field.
func (o *ProjectMemberUserData) SetUniqueExtField(v string) {
	o.UniqueExtField = &v
}

// GetDepartmentMember returns the DepartmentMember field value if set, zero value otherwise.
func (o *ProjectMemberUserData) GetDepartmentMember() ProjectMemberDepartmentMember {
	if o == nil || IsNil(o.DepartmentMember) {
		var ret ProjectMemberDepartmentMember
		return ret
	}
	return o.DepartmentMember
}

// GetDepartmentMemberOk returns a tuple with the DepartmentMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMemberUserData) GetDepartmentMemberOk() (ProjectMemberDepartmentMember, bool) {
	if o == nil || IsNil(o.DepartmentMember) {
		return ProjectMemberDepartmentMember{}, false
	}
	return o.DepartmentMember, true
}

// HasDepartmentMember returns a boolean if a field has been set.
func (o *ProjectMemberUserData) HasDepartmentMember() bool {
	if o != nil && !IsNil(o.DepartmentMember) {
		return true
	}

	return false
}

// SetDepartmentMember gets a reference to the given ProjectMemberDepartmentMember and assigns it to the DepartmentMember field.
func (o *ProjectMemberUserData) SetDepartmentMember(v ProjectMemberDepartmentMember) {
	o.DepartmentMember = v
}

func (o ProjectMemberUserData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMemberUserData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
	}
	if !IsNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	if !IsNil(o.EmailValidation) {
		toSerialize["EmailValidation"] = o.EmailValidation
	}
	if !IsNil(o.GlobalKey) {
		toSerialize["GlobalKey"] = o.GlobalKey
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.NamePinYin) {
		toSerialize["NamePinYin"] = o.NamePinYin
	}
	if !IsNil(o.Phone) {
		toSerialize["Phone"] = o.Phone
	}
	if !IsNil(o.PhoneValidation) {
		toSerialize["PhoneValidation"] = o.PhoneValidation
	}
	if !IsNil(o.Roles) {
		toSerialize["Roles"] = o.Roles
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.TeamId) {
		toSerialize["TeamId"] = o.TeamId
	}
	if !IsNil(o.UniqueExtField) {
		toSerialize["UniqueExtField"] = o.UniqueExtField
	}
	if !IsNil(o.DepartmentMember) {
		toSerialize["DepartmentMember"] = o.DepartmentMember
	}
	return toSerialize, nil
}

type NullableProjectMemberUserData struct {
	value *ProjectMemberUserData
	isSet bool
}

func (v NullableProjectMemberUserData) Get() *ProjectMemberUserData {
	return v.value
}

func (v *NullableProjectMemberUserData) Set(val *ProjectMemberUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMemberUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMemberUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMemberUserData(val *ProjectMemberUserData) *NullableProjectMemberUserData {
	return &NullableProjectMemberUserData{value: val, isSet: true}
}

func (v NullableProjectMemberUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMemberUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


