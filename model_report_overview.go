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

// checks if the ReportOverview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportOverview{}

// ReportOverview 测试报告概览
type ReportOverview struct {
	// 自动化覆盖率：百分比
	AutomationPercent NullableInt32 `json:"AutomationPercent,omitempty"`
	// 平均关闭时长
	AvgClosedSeconds NullableInt64 `json:"AvgClosedSeconds,omitempty"`
	// 用例总数
	CaseSum NullableInt32 `json:"CaseSum,omitempty"`
	// 已完成数量
	CompletedSum NullableInt32 `json:"CompletedSum,omitempty"`
	// 缺陷修复率：百分比
	DefectFixPercent NullableInt32 `json:"DefectFixPercent,omitempty"`
	// 重新激活率：百分比
	DefectReopenPercent NullableInt32 `json:"DefectReopenPercent,omitempty"`
	// 缺陷总数
	DefectSum NullableInt32 `json:"DefectSum,omitempty"`
	// 85%解决时长
	DurationFixed NullableInt64 `json:"DurationFixed,omitempty"`
	// 执行率：百分比
	ExecPercent NullableInt32 `json:"ExecPercent,omitempty"`
	// 需求总数
	IssuesSum NullableInt32 `json:"IssuesSum,omitempty"`
	// 通过率：百分比
	PassPercent NullableInt32 `json:"PassPercent,omitempty"`
	// 处理中数量
	ProcessingSum NullableInt32 `json:"ProcessingSum,omitempty"`
	// 需求覆盖率：百分比
	RequirementCoverPercent NullableInt32 `json:"RequirementCoverPercent,omitempty"`
	// 未开始数量
	TodoSum NullableInt32 `json:"TodoSum,omitempty"`
}

// NewReportOverview instantiates a new ReportOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportOverview() *ReportOverview {
	this := ReportOverview{}
	return &this
}

// NewReportOverviewWithDefaults instantiates a new ReportOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportOverviewWithDefaults() *ReportOverview {
	this := ReportOverview{}
	return &this
}

// GetAutomationPercent returns the AutomationPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetAutomationPercent() int32 {
	if o == nil || IsNil(o.AutomationPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.AutomationPercent.Get()
}

// GetAutomationPercentOk returns a tuple with the AutomationPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetAutomationPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomationPercent.Get(), o.AutomationPercent.IsSet()
}

// HasAutomationPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasAutomationPercent() bool {
	if o != nil && o.AutomationPercent.IsSet() {
		return true
	}

	return false
}

// SetAutomationPercent gets a reference to the given NullableInt32 and assigns it to the AutomationPercent field.
func (o *ReportOverview) SetAutomationPercent(v int32) {
	o.AutomationPercent.Set(&v)
}
// SetAutomationPercentNil sets the value for AutomationPercent to be an explicit nil
func (o *ReportOverview) SetAutomationPercentNil() {
	o.AutomationPercent.Set(nil)
}

// UnsetAutomationPercent ensures that no value is present for AutomationPercent, not even an explicit nil
func (o *ReportOverview) UnsetAutomationPercent() {
	o.AutomationPercent.Unset()
}

// GetAvgClosedSeconds returns the AvgClosedSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetAvgClosedSeconds() int64 {
	if o == nil || IsNil(o.AvgClosedSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.AvgClosedSeconds.Get()
}

// GetAvgClosedSecondsOk returns a tuple with the AvgClosedSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetAvgClosedSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgClosedSeconds.Get(), o.AvgClosedSeconds.IsSet()
}

// HasAvgClosedSeconds returns a boolean if a field has been set.
func (o *ReportOverview) HasAvgClosedSeconds() bool {
	if o != nil && o.AvgClosedSeconds.IsSet() {
		return true
	}

	return false
}

// SetAvgClosedSeconds gets a reference to the given NullableInt64 and assigns it to the AvgClosedSeconds field.
func (o *ReportOverview) SetAvgClosedSeconds(v int64) {
	o.AvgClosedSeconds.Set(&v)
}
// SetAvgClosedSecondsNil sets the value for AvgClosedSeconds to be an explicit nil
func (o *ReportOverview) SetAvgClosedSecondsNil() {
	o.AvgClosedSeconds.Set(nil)
}

// UnsetAvgClosedSeconds ensures that no value is present for AvgClosedSeconds, not even an explicit nil
func (o *ReportOverview) UnsetAvgClosedSeconds() {
	o.AvgClosedSeconds.Unset()
}

// GetCaseSum returns the CaseSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetCaseSum() int32 {
	if o == nil || IsNil(o.CaseSum.Get()) {
		var ret int32
		return ret
	}
	return *o.CaseSum.Get()
}

// GetCaseSumOk returns a tuple with the CaseSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetCaseSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseSum.Get(), o.CaseSum.IsSet()
}

// HasCaseSum returns a boolean if a field has been set.
func (o *ReportOverview) HasCaseSum() bool {
	if o != nil && o.CaseSum.IsSet() {
		return true
	}

	return false
}

// SetCaseSum gets a reference to the given NullableInt32 and assigns it to the CaseSum field.
func (o *ReportOverview) SetCaseSum(v int32) {
	o.CaseSum.Set(&v)
}
// SetCaseSumNil sets the value for CaseSum to be an explicit nil
func (o *ReportOverview) SetCaseSumNil() {
	o.CaseSum.Set(nil)
}

// UnsetCaseSum ensures that no value is present for CaseSum, not even an explicit nil
func (o *ReportOverview) UnsetCaseSum() {
	o.CaseSum.Unset()
}

// GetCompletedSum returns the CompletedSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetCompletedSum() int32 {
	if o == nil || IsNil(o.CompletedSum.Get()) {
		var ret int32
		return ret
	}
	return *o.CompletedSum.Get()
}

// GetCompletedSumOk returns a tuple with the CompletedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetCompletedSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedSum.Get(), o.CompletedSum.IsSet()
}

// HasCompletedSum returns a boolean if a field has been set.
func (o *ReportOverview) HasCompletedSum() bool {
	if o != nil && o.CompletedSum.IsSet() {
		return true
	}

	return false
}

// SetCompletedSum gets a reference to the given NullableInt32 and assigns it to the CompletedSum field.
func (o *ReportOverview) SetCompletedSum(v int32) {
	o.CompletedSum.Set(&v)
}
// SetCompletedSumNil sets the value for CompletedSum to be an explicit nil
func (o *ReportOverview) SetCompletedSumNil() {
	o.CompletedSum.Set(nil)
}

// UnsetCompletedSum ensures that no value is present for CompletedSum, not even an explicit nil
func (o *ReportOverview) UnsetCompletedSum() {
	o.CompletedSum.Unset()
}

// GetDefectFixPercent returns the DefectFixPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectFixPercent() int32 {
	if o == nil || IsNil(o.DefectFixPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectFixPercent.Get()
}

// GetDefectFixPercentOk returns a tuple with the DefectFixPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectFixPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectFixPercent.Get(), o.DefectFixPercent.IsSet()
}

// HasDefectFixPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectFixPercent() bool {
	if o != nil && o.DefectFixPercent.IsSet() {
		return true
	}

	return false
}

// SetDefectFixPercent gets a reference to the given NullableInt32 and assigns it to the DefectFixPercent field.
func (o *ReportOverview) SetDefectFixPercent(v int32) {
	o.DefectFixPercent.Set(&v)
}
// SetDefectFixPercentNil sets the value for DefectFixPercent to be an explicit nil
func (o *ReportOverview) SetDefectFixPercentNil() {
	o.DefectFixPercent.Set(nil)
}

// UnsetDefectFixPercent ensures that no value is present for DefectFixPercent, not even an explicit nil
func (o *ReportOverview) UnsetDefectFixPercent() {
	o.DefectFixPercent.Unset()
}

// GetDefectReopenPercent returns the DefectReopenPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectReopenPercent() int32 {
	if o == nil || IsNil(o.DefectReopenPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectReopenPercent.Get()
}

// GetDefectReopenPercentOk returns a tuple with the DefectReopenPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectReopenPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectReopenPercent.Get(), o.DefectReopenPercent.IsSet()
}

// HasDefectReopenPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectReopenPercent() bool {
	if o != nil && o.DefectReopenPercent.IsSet() {
		return true
	}

	return false
}

// SetDefectReopenPercent gets a reference to the given NullableInt32 and assigns it to the DefectReopenPercent field.
func (o *ReportOverview) SetDefectReopenPercent(v int32) {
	o.DefectReopenPercent.Set(&v)
}
// SetDefectReopenPercentNil sets the value for DefectReopenPercent to be an explicit nil
func (o *ReportOverview) SetDefectReopenPercentNil() {
	o.DefectReopenPercent.Set(nil)
}

// UnsetDefectReopenPercent ensures that no value is present for DefectReopenPercent, not even an explicit nil
func (o *ReportOverview) UnsetDefectReopenPercent() {
	o.DefectReopenPercent.Unset()
}

// GetDefectSum returns the DefectSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectSum() int32 {
	if o == nil || IsNil(o.DefectSum.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectSum.Get()
}

// GetDefectSumOk returns a tuple with the DefectSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectSum.Get(), o.DefectSum.IsSet()
}

// HasDefectSum returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectSum() bool {
	if o != nil && o.DefectSum.IsSet() {
		return true
	}

	return false
}

// SetDefectSum gets a reference to the given NullableInt32 and assigns it to the DefectSum field.
func (o *ReportOverview) SetDefectSum(v int32) {
	o.DefectSum.Set(&v)
}
// SetDefectSumNil sets the value for DefectSum to be an explicit nil
func (o *ReportOverview) SetDefectSumNil() {
	o.DefectSum.Set(nil)
}

// UnsetDefectSum ensures that no value is present for DefectSum, not even an explicit nil
func (o *ReportOverview) UnsetDefectSum() {
	o.DefectSum.Unset()
}

// GetDurationFixed returns the DurationFixed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDurationFixed() int64 {
	if o == nil || IsNil(o.DurationFixed.Get()) {
		var ret int64
		return ret
	}
	return *o.DurationFixed.Get()
}

// GetDurationFixedOk returns a tuple with the DurationFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDurationFixedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationFixed.Get(), o.DurationFixed.IsSet()
}

// HasDurationFixed returns a boolean if a field has been set.
func (o *ReportOverview) HasDurationFixed() bool {
	if o != nil && o.DurationFixed.IsSet() {
		return true
	}

	return false
}

// SetDurationFixed gets a reference to the given NullableInt64 and assigns it to the DurationFixed field.
func (o *ReportOverview) SetDurationFixed(v int64) {
	o.DurationFixed.Set(&v)
}
// SetDurationFixedNil sets the value for DurationFixed to be an explicit nil
func (o *ReportOverview) SetDurationFixedNil() {
	o.DurationFixed.Set(nil)
}

// UnsetDurationFixed ensures that no value is present for DurationFixed, not even an explicit nil
func (o *ReportOverview) UnsetDurationFixed() {
	o.DurationFixed.Unset()
}

// GetExecPercent returns the ExecPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetExecPercent() int32 {
	if o == nil || IsNil(o.ExecPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.ExecPercent.Get()
}

// GetExecPercentOk returns a tuple with the ExecPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetExecPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecPercent.Get(), o.ExecPercent.IsSet()
}

// HasExecPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasExecPercent() bool {
	if o != nil && o.ExecPercent.IsSet() {
		return true
	}

	return false
}

// SetExecPercent gets a reference to the given NullableInt32 and assigns it to the ExecPercent field.
func (o *ReportOverview) SetExecPercent(v int32) {
	o.ExecPercent.Set(&v)
}
// SetExecPercentNil sets the value for ExecPercent to be an explicit nil
func (o *ReportOverview) SetExecPercentNil() {
	o.ExecPercent.Set(nil)
}

// UnsetExecPercent ensures that no value is present for ExecPercent, not even an explicit nil
func (o *ReportOverview) UnsetExecPercent() {
	o.ExecPercent.Unset()
}

// GetIssuesSum returns the IssuesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetIssuesSum() int32 {
	if o == nil || IsNil(o.IssuesSum.Get()) {
		var ret int32
		return ret
	}
	return *o.IssuesSum.Get()
}

// GetIssuesSumOk returns a tuple with the IssuesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetIssuesSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuesSum.Get(), o.IssuesSum.IsSet()
}

// HasIssuesSum returns a boolean if a field has been set.
func (o *ReportOverview) HasIssuesSum() bool {
	if o != nil && o.IssuesSum.IsSet() {
		return true
	}

	return false
}

// SetIssuesSum gets a reference to the given NullableInt32 and assigns it to the IssuesSum field.
func (o *ReportOverview) SetIssuesSum(v int32) {
	o.IssuesSum.Set(&v)
}
// SetIssuesSumNil sets the value for IssuesSum to be an explicit nil
func (o *ReportOverview) SetIssuesSumNil() {
	o.IssuesSum.Set(nil)
}

// UnsetIssuesSum ensures that no value is present for IssuesSum, not even an explicit nil
func (o *ReportOverview) UnsetIssuesSum() {
	o.IssuesSum.Unset()
}

// GetPassPercent returns the PassPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetPassPercent() int32 {
	if o == nil || IsNil(o.PassPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.PassPercent.Get()
}

// GetPassPercentOk returns a tuple with the PassPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetPassPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassPercent.Get(), o.PassPercent.IsSet()
}

// HasPassPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasPassPercent() bool {
	if o != nil && o.PassPercent.IsSet() {
		return true
	}

	return false
}

// SetPassPercent gets a reference to the given NullableInt32 and assigns it to the PassPercent field.
func (o *ReportOverview) SetPassPercent(v int32) {
	o.PassPercent.Set(&v)
}
// SetPassPercentNil sets the value for PassPercent to be an explicit nil
func (o *ReportOverview) SetPassPercentNil() {
	o.PassPercent.Set(nil)
}

// UnsetPassPercent ensures that no value is present for PassPercent, not even an explicit nil
func (o *ReportOverview) UnsetPassPercent() {
	o.PassPercent.Unset()
}

// GetProcessingSum returns the ProcessingSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetProcessingSum() int32 {
	if o == nil || IsNil(o.ProcessingSum.Get()) {
		var ret int32
		return ret
	}
	return *o.ProcessingSum.Get()
}

// GetProcessingSumOk returns a tuple with the ProcessingSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetProcessingSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessingSum.Get(), o.ProcessingSum.IsSet()
}

// HasProcessingSum returns a boolean if a field has been set.
func (o *ReportOverview) HasProcessingSum() bool {
	if o != nil && o.ProcessingSum.IsSet() {
		return true
	}

	return false
}

// SetProcessingSum gets a reference to the given NullableInt32 and assigns it to the ProcessingSum field.
func (o *ReportOverview) SetProcessingSum(v int32) {
	o.ProcessingSum.Set(&v)
}
// SetProcessingSumNil sets the value for ProcessingSum to be an explicit nil
func (o *ReportOverview) SetProcessingSumNil() {
	o.ProcessingSum.Set(nil)
}

// UnsetProcessingSum ensures that no value is present for ProcessingSum, not even an explicit nil
func (o *ReportOverview) UnsetProcessingSum() {
	o.ProcessingSum.Unset()
}

// GetRequirementCoverPercent returns the RequirementCoverPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetRequirementCoverPercent() int32 {
	if o == nil || IsNil(o.RequirementCoverPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.RequirementCoverPercent.Get()
}

// GetRequirementCoverPercentOk returns a tuple with the RequirementCoverPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetRequirementCoverPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequirementCoverPercent.Get(), o.RequirementCoverPercent.IsSet()
}

// HasRequirementCoverPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasRequirementCoverPercent() bool {
	if o != nil && o.RequirementCoverPercent.IsSet() {
		return true
	}

	return false
}

// SetRequirementCoverPercent gets a reference to the given NullableInt32 and assigns it to the RequirementCoverPercent field.
func (o *ReportOverview) SetRequirementCoverPercent(v int32) {
	o.RequirementCoverPercent.Set(&v)
}
// SetRequirementCoverPercentNil sets the value for RequirementCoverPercent to be an explicit nil
func (o *ReportOverview) SetRequirementCoverPercentNil() {
	o.RequirementCoverPercent.Set(nil)
}

// UnsetRequirementCoverPercent ensures that no value is present for RequirementCoverPercent, not even an explicit nil
func (o *ReportOverview) UnsetRequirementCoverPercent() {
	o.RequirementCoverPercent.Unset()
}

// GetTodoSum returns the TodoSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetTodoSum() int32 {
	if o == nil || IsNil(o.TodoSum.Get()) {
		var ret int32
		return ret
	}
	return *o.TodoSum.Get()
}

// GetTodoSumOk returns a tuple with the TodoSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetTodoSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TodoSum.Get(), o.TodoSum.IsSet()
}

// HasTodoSum returns a boolean if a field has been set.
func (o *ReportOverview) HasTodoSum() bool {
	if o != nil && o.TodoSum.IsSet() {
		return true
	}

	return false
}

// SetTodoSum gets a reference to the given NullableInt32 and assigns it to the TodoSum field.
func (o *ReportOverview) SetTodoSum(v int32) {
	o.TodoSum.Set(&v)
}
// SetTodoSumNil sets the value for TodoSum to be an explicit nil
func (o *ReportOverview) SetTodoSumNil() {
	o.TodoSum.Set(nil)
}

// UnsetTodoSum ensures that no value is present for TodoSum, not even an explicit nil
func (o *ReportOverview) UnsetTodoSum() {
	o.TodoSum.Unset()
}

func (o ReportOverview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportOverview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutomationPercent.IsSet() {
		toSerialize["AutomationPercent"] = o.AutomationPercent.Get()
	}
	if o.AvgClosedSeconds.IsSet() {
		toSerialize["AvgClosedSeconds"] = o.AvgClosedSeconds.Get()
	}
	if o.CaseSum.IsSet() {
		toSerialize["CaseSum"] = o.CaseSum.Get()
	}
	if o.CompletedSum.IsSet() {
		toSerialize["CompletedSum"] = o.CompletedSum.Get()
	}
	if o.DefectFixPercent.IsSet() {
		toSerialize["DefectFixPercent"] = o.DefectFixPercent.Get()
	}
	if o.DefectReopenPercent.IsSet() {
		toSerialize["DefectReopenPercent"] = o.DefectReopenPercent.Get()
	}
	if o.DefectSum.IsSet() {
		toSerialize["DefectSum"] = o.DefectSum.Get()
	}
	if o.DurationFixed.IsSet() {
		toSerialize["DurationFixed"] = o.DurationFixed.Get()
	}
	if o.ExecPercent.IsSet() {
		toSerialize["ExecPercent"] = o.ExecPercent.Get()
	}
	if o.IssuesSum.IsSet() {
		toSerialize["IssuesSum"] = o.IssuesSum.Get()
	}
	if o.PassPercent.IsSet() {
		toSerialize["PassPercent"] = o.PassPercent.Get()
	}
	if o.ProcessingSum.IsSet() {
		toSerialize["ProcessingSum"] = o.ProcessingSum.Get()
	}
	if o.RequirementCoverPercent.IsSet() {
		toSerialize["RequirementCoverPercent"] = o.RequirementCoverPercent.Get()
	}
	if o.TodoSum.IsSet() {
		toSerialize["TodoSum"] = o.TodoSum.Get()
	}
	return toSerialize, nil
}

type NullableReportOverview struct {
	value *ReportOverview
	isSet bool
}

func (v NullableReportOverview) Get() *ReportOverview {
	return v.value
}

func (v *NullableReportOverview) Set(val *ReportOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableReportOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableReportOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportOverview(val *ReportOverview) *NullableReportOverview {
	return &NullableReportOverview{value: val, isSet: true}
}

func (v NullableReportOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


