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

// checks if the WikiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WikiData{}

// WikiData wiki的具体信息
type WikiData struct {
	// 是否为维护者
	CanMaintain *bool `json:"CanMaintain,omitempty"`
	// 是否可以阅读
	CanRead NullableBool `json:"CanRead,omitempty"`
	// 内容
	Content *string `json:"Content,omitempty"`
	// 创建时间
	CreatedAt map[string]interface{} `json:"CreatedAt,omitempty"`
	Creator *User `json:"Creator,omitempty"`
	// 当前versionId
	CurrentVersion *int64 `json:"CurrentVersion,omitempty"`
	Editor *User `json:"Editor,omitempty"`
	// 修改次数
	HistoriesCount *int64 `json:"HistoriesCount,omitempty"`
	// wiki历史Id
	HistoryId *int64 `json:"HistoryId,omitempty"`
	// 内容转成的html
	Html *string `json:"Html,omitempty"`
	// wikiId
	Id *int64 `json:"Id,omitempty"`
	// wik的code
	Iid *int64 `json:"Iid,omitempty"`
	// 最新versionId
	LastVersion *int64 `json:"LastVersion,omitempty"`
	// 提交说明
	Msg NullableString `json:"Msg,omitempty"`
	// 所处顺序位置
	Order *int64 `json:"Order,omitempty"`
	// 父级 IiD
	ParentIid *int64 `json:"ParentIid,omitempty"`
	// 是否父级分享
	ParentShared *bool `json:"ParentShared,omitempty"`
	// 父级可见范围
	ParentVisibleRange NullableString `json:"ParentVisibleRange,omitempty"`
	// 路径
	Path NullableString `json:"Path,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty"`
	// 修改时间
	UpdatedAt map[string]interface{} `json:"UpdatedAt,omitempty"`
	// 可见范围
	VisibleRange NullableString `json:"VisibleRange,omitempty"`
}

// NewWikiData instantiates a new WikiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiData() *WikiData {
	this := WikiData{}
	var canMaintain bool = false
	this.CanMaintain = &canMaintain
	var canRead bool = false
	this.CanRead = *NewNullableBool(&canRead)
	var content string = ""
	this.Content = &content
	var html string = ""
	this.Html = &html
	var msg string = ""
	this.Msg = *NewNullableString(&msg)
	var parentShared bool = false
	this.ParentShared = &parentShared
	var parentVisibleRange string = ""
	this.ParentVisibleRange = *NewNullableString(&parentVisibleRange)
	var path string = ""
	this.Path = *NewNullableString(&path)
	var title string = ""
	this.Title = &title
	var visibleRange string = ""
	this.VisibleRange = *NewNullableString(&visibleRange)
	return &this
}

// NewWikiDataWithDefaults instantiates a new WikiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiDataWithDefaults() *WikiData {
	this := WikiData{}
	var canMaintain bool = false
	this.CanMaintain = &canMaintain
	var canRead bool = false
	this.CanRead = *NewNullableBool(&canRead)
	var content string = ""
	this.Content = &content
	var html string = ""
	this.Html = &html
	var msg string = ""
	this.Msg = *NewNullableString(&msg)
	var parentShared bool = false
	this.ParentShared = &parentShared
	var parentVisibleRange string = ""
	this.ParentVisibleRange = *NewNullableString(&parentVisibleRange)
	var path string = ""
	this.Path = *NewNullableString(&path)
	var title string = ""
	this.Title = &title
	var visibleRange string = ""
	this.VisibleRange = *NewNullableString(&visibleRange)
	return &this
}

// GetCanMaintain returns the CanMaintain field value if set, zero value otherwise.
func (o *WikiData) GetCanMaintain() bool {
	if o == nil || IsNil(o.CanMaintain) {
		var ret bool
		return ret
	}
	return *o.CanMaintain
}

// GetCanMaintainOk returns a tuple with the CanMaintain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetCanMaintainOk() (*bool, bool) {
	if o == nil || IsNil(o.CanMaintain) {
		return nil, false
	}
	return o.CanMaintain, true
}

// HasCanMaintain returns a boolean if a field has been set.
func (o *WikiData) HasCanMaintain() bool {
	if o != nil && !IsNil(o.CanMaintain) {
		return true
	}

	return false
}

// SetCanMaintain gets a reference to the given bool and assigns it to the CanMaintain field.
func (o *WikiData) SetCanMaintain(v bool) {
	o.CanMaintain = &v
}

// GetCanRead returns the CanRead field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiData) GetCanRead() bool {
	if o == nil || IsNil(o.CanRead.Get()) {
		var ret bool
		return ret
	}
	return *o.CanRead.Get()
}

// GetCanReadOk returns a tuple with the CanRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiData) GetCanReadOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanRead.Get(), o.CanRead.IsSet()
}

// HasCanRead returns a boolean if a field has been set.
func (o *WikiData) HasCanRead() bool {
	if o != nil && o.CanRead.IsSet() {
		return true
	}

	return false
}

// SetCanRead gets a reference to the given NullableBool and assigns it to the CanRead field.
func (o *WikiData) SetCanRead(v bool) {
	o.CanRead.Set(&v)
}
// SetCanReadNil sets the value for CanRead to be an explicit nil
func (o *WikiData) SetCanReadNil() {
	o.CanRead.Set(nil)
}

// UnsetCanRead ensures that no value is present for CanRead, not even an explicit nil
func (o *WikiData) UnsetCanRead() {
	o.CanRead.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *WikiData) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *WikiData) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *WikiData) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WikiData) GetCreatedAt() map[string]interface{} {
	if o == nil || IsNil(o.CreatedAt) {
		var ret map[string]interface{}
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetCreatedAtOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return map[string]interface{}{}, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WikiData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given map[string]interface{} and assigns it to the CreatedAt field.
func (o *WikiData) SetCreatedAt(v map[string]interface{}) {
	o.CreatedAt = v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *WikiData) GetCreator() User {
	if o == nil || IsNil(o.Creator) {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetCreatorOk() (*User, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *WikiData) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *WikiData) SetCreator(v User) {
	o.Creator = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *WikiData) GetCurrentVersion() int64 {
	if o == nil || IsNil(o.CurrentVersion) {
		var ret int64
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetCurrentVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *WikiData) HasCurrentVersion() bool {
	if o != nil && !IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given int64 and assigns it to the CurrentVersion field.
func (o *WikiData) SetCurrentVersion(v int64) {
	o.CurrentVersion = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *WikiData) GetEditor() User {
	if o == nil || IsNil(o.Editor) {
		var ret User
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetEditorOk() (*User, bool) {
	if o == nil || IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *WikiData) HasEditor() bool {
	if o != nil && !IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given User and assigns it to the Editor field.
func (o *WikiData) SetEditor(v User) {
	o.Editor = &v
}

// GetHistoriesCount returns the HistoriesCount field value if set, zero value otherwise.
func (o *WikiData) GetHistoriesCount() int64 {
	if o == nil || IsNil(o.HistoriesCount) {
		var ret int64
		return ret
	}
	return *o.HistoriesCount
}

// GetHistoriesCountOk returns a tuple with the HistoriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetHistoriesCountOk() (*int64, bool) {
	if o == nil || IsNil(o.HistoriesCount) {
		return nil, false
	}
	return o.HistoriesCount, true
}

// HasHistoriesCount returns a boolean if a field has been set.
func (o *WikiData) HasHistoriesCount() bool {
	if o != nil && !IsNil(o.HistoriesCount) {
		return true
	}

	return false
}

// SetHistoriesCount gets a reference to the given int64 and assigns it to the HistoriesCount field.
func (o *WikiData) SetHistoriesCount(v int64) {
	o.HistoriesCount = &v
}

// GetHistoryId returns the HistoryId field value if set, zero value otherwise.
func (o *WikiData) GetHistoryId() int64 {
	if o == nil || IsNil(o.HistoryId) {
		var ret int64
		return ret
	}
	return *o.HistoryId
}

// GetHistoryIdOk returns a tuple with the HistoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetHistoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.HistoryId) {
		return nil, false
	}
	return o.HistoryId, true
}

// HasHistoryId returns a boolean if a field has been set.
func (o *WikiData) HasHistoryId() bool {
	if o != nil && !IsNil(o.HistoryId) {
		return true
	}

	return false
}

// SetHistoryId gets a reference to the given int64 and assigns it to the HistoryId field.
func (o *WikiData) SetHistoryId(v int64) {
	o.HistoryId = &v
}

// GetHtml returns the Html field value if set, zero value otherwise.
func (o *WikiData) GetHtml() string {
	if o == nil || IsNil(o.Html) {
		var ret string
		return ret
	}
	return *o.Html
}

// GetHtmlOk returns a tuple with the Html field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetHtmlOk() (*string, bool) {
	if o == nil || IsNil(o.Html) {
		return nil, false
	}
	return o.Html, true
}

// HasHtml returns a boolean if a field has been set.
func (o *WikiData) HasHtml() bool {
	if o != nil && !IsNil(o.Html) {
		return true
	}

	return false
}

// SetHtml gets a reference to the given string and assigns it to the Html field.
func (o *WikiData) SetHtml(v string) {
	o.Html = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WikiData) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WikiData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WikiData) SetId(v int64) {
	o.Id = &v
}

// GetIid returns the Iid field value if set, zero value otherwise.
func (o *WikiData) GetIid() int64 {
	if o == nil || IsNil(o.Iid) {
		var ret int64
		return ret
	}
	return *o.Iid
}

// GetIidOk returns a tuple with the Iid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetIidOk() (*int64, bool) {
	if o == nil || IsNil(o.Iid) {
		return nil, false
	}
	return o.Iid, true
}

// HasIid returns a boolean if a field has been set.
func (o *WikiData) HasIid() bool {
	if o != nil && !IsNil(o.Iid) {
		return true
	}

	return false
}

// SetIid gets a reference to the given int64 and assigns it to the Iid field.
func (o *WikiData) SetIid(v int64) {
	o.Iid = &v
}

// GetLastVersion returns the LastVersion field value if set, zero value otherwise.
func (o *WikiData) GetLastVersion() int64 {
	if o == nil || IsNil(o.LastVersion) {
		var ret int64
		return ret
	}
	return *o.LastVersion
}

// GetLastVersionOk returns a tuple with the LastVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetLastVersionOk() (*int64, bool) {
	if o == nil || IsNil(o.LastVersion) {
		return nil, false
	}
	return o.LastVersion, true
}

// HasLastVersion returns a boolean if a field has been set.
func (o *WikiData) HasLastVersion() bool {
	if o != nil && !IsNil(o.LastVersion) {
		return true
	}

	return false
}

// SetLastVersion gets a reference to the given int64 and assigns it to the LastVersion field.
func (o *WikiData) SetLastVersion(v int64) {
	o.LastVersion = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiData) GetMsg() string {
	if o == nil || IsNil(o.Msg.Get()) {
		var ret string
		return ret
	}
	return *o.Msg.Get()
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiData) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Msg.Get(), o.Msg.IsSet()
}

// HasMsg returns a boolean if a field has been set.
func (o *WikiData) HasMsg() bool {
	if o != nil && o.Msg.IsSet() {
		return true
	}

	return false
}

// SetMsg gets a reference to the given NullableString and assigns it to the Msg field.
func (o *WikiData) SetMsg(v string) {
	o.Msg.Set(&v)
}
// SetMsgNil sets the value for Msg to be an explicit nil
func (o *WikiData) SetMsgNil() {
	o.Msg.Set(nil)
}

// UnsetMsg ensures that no value is present for Msg, not even an explicit nil
func (o *WikiData) UnsetMsg() {
	o.Msg.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *WikiData) GetOrder() int64 {
	if o == nil || IsNil(o.Order) {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *WikiData) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *WikiData) SetOrder(v int64) {
	o.Order = &v
}

// GetParentIid returns the ParentIid field value if set, zero value otherwise.
func (o *WikiData) GetParentIid() int64 {
	if o == nil || IsNil(o.ParentIid) {
		var ret int64
		return ret
	}
	return *o.ParentIid
}

// GetParentIidOk returns a tuple with the ParentIid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetParentIidOk() (*int64, bool) {
	if o == nil || IsNil(o.ParentIid) {
		return nil, false
	}
	return o.ParentIid, true
}

// HasParentIid returns a boolean if a field has been set.
func (o *WikiData) HasParentIid() bool {
	if o != nil && !IsNil(o.ParentIid) {
		return true
	}

	return false
}

// SetParentIid gets a reference to the given int64 and assigns it to the ParentIid field.
func (o *WikiData) SetParentIid(v int64) {
	o.ParentIid = &v
}

// GetParentShared returns the ParentShared field value if set, zero value otherwise.
func (o *WikiData) GetParentShared() bool {
	if o == nil || IsNil(o.ParentShared) {
		var ret bool
		return ret
	}
	return *o.ParentShared
}

// GetParentSharedOk returns a tuple with the ParentShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetParentSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.ParentShared) {
		return nil, false
	}
	return o.ParentShared, true
}

// HasParentShared returns a boolean if a field has been set.
func (o *WikiData) HasParentShared() bool {
	if o != nil && !IsNil(o.ParentShared) {
		return true
	}

	return false
}

// SetParentShared gets a reference to the given bool and assigns it to the ParentShared field.
func (o *WikiData) SetParentShared(v bool) {
	o.ParentShared = &v
}

// GetParentVisibleRange returns the ParentVisibleRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiData) GetParentVisibleRange() string {
	if o == nil || IsNil(o.ParentVisibleRange.Get()) {
		var ret string
		return ret
	}
	return *o.ParentVisibleRange.Get()
}

// GetParentVisibleRangeOk returns a tuple with the ParentVisibleRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiData) GetParentVisibleRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentVisibleRange.Get(), o.ParentVisibleRange.IsSet()
}

// HasParentVisibleRange returns a boolean if a field has been set.
func (o *WikiData) HasParentVisibleRange() bool {
	if o != nil && o.ParentVisibleRange.IsSet() {
		return true
	}

	return false
}

// SetParentVisibleRange gets a reference to the given NullableString and assigns it to the ParentVisibleRange field.
func (o *WikiData) SetParentVisibleRange(v string) {
	o.ParentVisibleRange.Set(&v)
}
// SetParentVisibleRangeNil sets the value for ParentVisibleRange to be an explicit nil
func (o *WikiData) SetParentVisibleRangeNil() {
	o.ParentVisibleRange.Set(nil)
}

// UnsetParentVisibleRange ensures that no value is present for ParentVisibleRange, not even an explicit nil
func (o *WikiData) UnsetParentVisibleRange() {
	o.ParentVisibleRange.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiData) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiData) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *WikiData) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *WikiData) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *WikiData) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *WikiData) UnsetPath() {
	o.Path.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WikiData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WikiData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WikiData) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WikiData) GetUpdatedAt() map[string]interface{} {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret map[string]interface{}
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiData) GetUpdatedAtOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return map[string]interface{}{}, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WikiData) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given map[string]interface{} and assigns it to the UpdatedAt field.
func (o *WikiData) SetUpdatedAt(v map[string]interface{}) {
	o.UpdatedAt = v
}

// GetVisibleRange returns the VisibleRange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiData) GetVisibleRange() string {
	if o == nil || IsNil(o.VisibleRange.Get()) {
		var ret string
		return ret
	}
	return *o.VisibleRange.Get()
}

// GetVisibleRangeOk returns a tuple with the VisibleRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiData) GetVisibleRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisibleRange.Get(), o.VisibleRange.IsSet()
}

// HasVisibleRange returns a boolean if a field has been set.
func (o *WikiData) HasVisibleRange() bool {
	if o != nil && o.VisibleRange.IsSet() {
		return true
	}

	return false
}

// SetVisibleRange gets a reference to the given NullableString and assigns it to the VisibleRange field.
func (o *WikiData) SetVisibleRange(v string) {
	o.VisibleRange.Set(&v)
}
// SetVisibleRangeNil sets the value for VisibleRange to be an explicit nil
func (o *WikiData) SetVisibleRangeNil() {
	o.VisibleRange.Set(nil)
}

// UnsetVisibleRange ensures that no value is present for VisibleRange, not even an explicit nil
func (o *WikiData) UnsetVisibleRange() {
	o.VisibleRange.Unset()
}

func (o WikiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WikiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanMaintain) {
		toSerialize["CanMaintain"] = o.CanMaintain
	}
	if o.CanRead.IsSet() {
		toSerialize["CanRead"] = o.CanRead.Get()
	}
	if !IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.Creator) {
		toSerialize["Creator"] = o.Creator
	}
	if !IsNil(o.CurrentVersion) {
		toSerialize["CurrentVersion"] = o.CurrentVersion
	}
	if !IsNil(o.Editor) {
		toSerialize["Editor"] = o.Editor
	}
	if !IsNil(o.HistoriesCount) {
		toSerialize["HistoriesCount"] = o.HistoriesCount
	}
	if !IsNil(o.HistoryId) {
		toSerialize["HistoryId"] = o.HistoryId
	}
	if !IsNil(o.Html) {
		toSerialize["Html"] = o.Html
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Iid) {
		toSerialize["Iid"] = o.Iid
	}
	if !IsNil(o.LastVersion) {
		toSerialize["LastVersion"] = o.LastVersion
	}
	if o.Msg.IsSet() {
		toSerialize["Msg"] = o.Msg.Get()
	}
	if !IsNil(o.Order) {
		toSerialize["Order"] = o.Order
	}
	if !IsNil(o.ParentIid) {
		toSerialize["ParentIid"] = o.ParentIid
	}
	if !IsNil(o.ParentShared) {
		toSerialize["ParentShared"] = o.ParentShared
	}
	if o.ParentVisibleRange.IsSet() {
		toSerialize["ParentVisibleRange"] = o.ParentVisibleRange.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if o.VisibleRange.IsSet() {
		toSerialize["VisibleRange"] = o.VisibleRange.Get()
	}
	return toSerialize, nil
}

type NullableWikiData struct {
	value *WikiData
	isSet bool
}

func (v NullableWikiData) Get() *WikiData {
	return v.value
}

func (v *NullableWikiData) Set(val *WikiData) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiData) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiData(val *WikiData) *NullableWikiData {
	return &NullableWikiData{value: val, isSet: true}
}

func (v NullableWikiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


