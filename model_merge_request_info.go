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

// checks if the MergeRequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeRequestInfo{}

// MergeRequestInfo 合并请求详情
type MergeRequestInfo struct {
	// 操作时间
	ActionAt *int64 `json:"ActionAt,omitempty"`
	ActionAuthor *CodingUser `json:"ActionAuthor,omitempty"`
	Author *CodingUser `json:"Author,omitempty"`
	// 基础Sha
	BaseSha *string `json:"BaseSha,omitempty"`
	// 评论数
	CommentCount *int64 `json:"CommentCount,omitempty"`
	// 冲突文件列表
	Conflicts []string `json:"Conflicts,omitempty"`
	// 创建时间
	CreatedAt *int64 `json:"CreatedAt,omitempty"`
	// 仓库ID
	DepotId *int64 `json:"DepotId,omitempty"`
	// 描述,为 markdown 格式
	Describe *string `json:"Describe,omitempty"`
	// 是否授权
	Granted *int64 `json:"Granted,omitempty"`
	// 合并请求ID
	Id *int64 `json:"Id,omitempty"`
	// 合并请求标签列表
	Labels []string `json:"Labels,omitempty"`
	// 合并Commit Sha
	MergeCommitSha *string `json:"MergeCommitSha,omitempty"`
	// 合并请求iid
	MergeId *int64 `json:"MergeId,omitempty"`
	Mission *Mission `json:"Mission,omitempty"`
	// 合并请求 web 页面路径
	Path *string `json:"Path,omitempty"`
	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 是否提醒
	Reminded NullableBool `json:"Reminded,omitempty"`
	// MR评审者
	Reviewers []CodingUser `json:"Reviewers,omitempty"`
	// 源分支
	SourceBranch *string `json:"SourceBranch,omitempty"`
	// 源分支Commit Sha
	SourceBranchSha *string `json:"SourceBranchSha,omitempty"`
	// 合并状态,CANMERGE(状态可自动合并),ACCEPTED(状态已接受), CANNOTMERGE(状态不可自动合并), REFUSED(状态已拒绝(关闭)), CANCEL(取消), MERGING(正在合并中), ABNORMAL(状态异常)
	Status *string `json:"Status,omitempty"`
	// MR阻塞点
	StickingPoint *string `json:"StickingPoint,omitempty"`
	// 目标分支
	TargetBranch *string `json:"TargetBranch,omitempty"`
	// 目标分支是否为保护分支
	TargetBranchProtected *bool `json:"TargetBranchProtected,omitempty"`
	// 目标分支Commit Sha
	TargetBranchSha *string `json:"TargetBranchSha,omitempty"`
	// 合并标题
	Title *string `json:"Title,omitempty"`
	// 更新时间
	UpdatedAt *int64 `json:"UpdatedAt,omitempty"`
}

// NewMergeRequestInfo instantiates a new MergeRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeRequestInfo() *MergeRequestInfo {
	this := MergeRequestInfo{}
	var baseSha string = ""
	this.BaseSha = &baseSha
	var describe string = ""
	this.Describe = &describe
	var mergeCommitSha string = ""
	this.MergeCommitSha = &mergeCommitSha
	var path string = ""
	this.Path = &path
	var reminded bool = false
	this.Reminded = *NewNullableBool(&reminded)
	var sourceBranch string = ""
	this.SourceBranch = &sourceBranch
	var sourceBranchSha string = ""
	this.SourceBranchSha = &sourceBranchSha
	var status string = ""
	this.Status = &status
	var stickingPoint string = ""
	this.StickingPoint = &stickingPoint
	var targetBranch string = ""
	this.TargetBranch = &targetBranch
	var targetBranchProtected bool = false
	this.TargetBranchProtected = &targetBranchProtected
	var targetBranchSha string = ""
	this.TargetBranchSha = &targetBranchSha
	var title string = ""
	this.Title = &title
	return &this
}

// NewMergeRequestInfoWithDefaults instantiates a new MergeRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeRequestInfoWithDefaults() *MergeRequestInfo {
	this := MergeRequestInfo{}
	var baseSha string = ""
	this.BaseSha = &baseSha
	var describe string = ""
	this.Describe = &describe
	var mergeCommitSha string = ""
	this.MergeCommitSha = &mergeCommitSha
	var path string = ""
	this.Path = &path
	var reminded bool = false
	this.Reminded = *NewNullableBool(&reminded)
	var sourceBranch string = ""
	this.SourceBranch = &sourceBranch
	var sourceBranchSha string = ""
	this.SourceBranchSha = &sourceBranchSha
	var status string = ""
	this.Status = &status
	var stickingPoint string = ""
	this.StickingPoint = &stickingPoint
	var targetBranch string = ""
	this.TargetBranch = &targetBranch
	var targetBranchProtected bool = false
	this.TargetBranchProtected = &targetBranchProtected
	var targetBranchSha string = ""
	this.TargetBranchSha = &targetBranchSha
	var title string = ""
	this.Title = &title
	return &this
}

// GetActionAt returns the ActionAt field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetActionAt() int64 {
	if o == nil || IsNil(o.ActionAt) {
		var ret int64
		return ret
	}
	return *o.ActionAt
}

// GetActionAtOk returns a tuple with the ActionAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetActionAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ActionAt) {
		return nil, false
	}
	return o.ActionAt, true
}

// HasActionAt returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasActionAt() bool {
	if o != nil && !IsNil(o.ActionAt) {
		return true
	}

	return false
}

// SetActionAt gets a reference to the given int64 and assigns it to the ActionAt field.
func (o *MergeRequestInfo) SetActionAt(v int64) {
	o.ActionAt = &v
}

// GetActionAuthor returns the ActionAuthor field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetActionAuthor() CodingUser {
	if o == nil || IsNil(o.ActionAuthor) {
		var ret CodingUser
		return ret
	}
	return *o.ActionAuthor
}

// GetActionAuthorOk returns a tuple with the ActionAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetActionAuthorOk() (*CodingUser, bool) {
	if o == nil || IsNil(o.ActionAuthor) {
		return nil, false
	}
	return o.ActionAuthor, true
}

// HasActionAuthor returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasActionAuthor() bool {
	if o != nil && !IsNil(o.ActionAuthor) {
		return true
	}

	return false
}

// SetActionAuthor gets a reference to the given CodingUser and assigns it to the ActionAuthor field.
func (o *MergeRequestInfo) SetActionAuthor(v CodingUser) {
	o.ActionAuthor = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetAuthor() CodingUser {
	if o == nil || IsNil(o.Author) {
		var ret CodingUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetAuthorOk() (*CodingUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CodingUser and assigns it to the Author field.
func (o *MergeRequestInfo) SetAuthor(v CodingUser) {
	o.Author = &v
}

// GetBaseSha returns the BaseSha field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetBaseSha() string {
	if o == nil || IsNil(o.BaseSha) {
		var ret string
		return ret
	}
	return *o.BaseSha
}

// GetBaseShaOk returns a tuple with the BaseSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetBaseShaOk() (*string, bool) {
	if o == nil || IsNil(o.BaseSha) {
		return nil, false
	}
	return o.BaseSha, true
}

// HasBaseSha returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasBaseSha() bool {
	if o != nil && !IsNil(o.BaseSha) {
		return true
	}

	return false
}

// SetBaseSha gets a reference to the given string and assigns it to the BaseSha field.
func (o *MergeRequestInfo) SetBaseSha(v string) {
	o.BaseSha = &v
}

// GetCommentCount returns the CommentCount field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetCommentCount() int64 {
	if o == nil || IsNil(o.CommentCount) {
		var ret int64
		return ret
	}
	return *o.CommentCount
}

// GetCommentCountOk returns a tuple with the CommentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetCommentCountOk() (*int64, bool) {
	if o == nil || IsNil(o.CommentCount) {
		return nil, false
	}
	return o.CommentCount, true
}

// HasCommentCount returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasCommentCount() bool {
	if o != nil && !IsNil(o.CommentCount) {
		return true
	}

	return false
}

// SetCommentCount gets a reference to the given int64 and assigns it to the CommentCount field.
func (o *MergeRequestInfo) SetCommentCount(v int64) {
	o.CommentCount = &v
}

// GetConflicts returns the Conflicts field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetConflicts() []string {
	if o == nil || IsNil(o.Conflicts) {
		var ret []string
		return ret
	}
	return o.Conflicts
}

// GetConflictsOk returns a tuple with the Conflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetConflictsOk() ([]string, bool) {
	if o == nil || IsNil(o.Conflicts) {
		return nil, false
	}
	return o.Conflicts, true
}

// HasConflicts returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasConflicts() bool {
	if o != nil && !IsNil(o.Conflicts) {
		return true
	}

	return false
}

// SetConflicts gets a reference to the given []string and assigns it to the Conflicts field.
func (o *MergeRequestInfo) SetConflicts(v []string) {
	o.Conflicts = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *MergeRequestInfo) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDepotId returns the DepotId field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetDepotId() int64 {
	if o == nil || IsNil(o.DepotId) {
		var ret int64
		return ret
	}
	return *o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetDepotIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DepotId) {
		return nil, false
	}
	return o.DepotId, true
}

// HasDepotId returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasDepotId() bool {
	if o != nil && !IsNil(o.DepotId) {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given int64 and assigns it to the DepotId field.
func (o *MergeRequestInfo) SetDepotId(v int64) {
	o.DepotId = &v
}

// GetDescribe returns the Describe field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetDescribe() string {
	if o == nil || IsNil(o.Describe) {
		var ret string
		return ret
	}
	return *o.Describe
}

// GetDescribeOk returns a tuple with the Describe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetDescribeOk() (*string, bool) {
	if o == nil || IsNil(o.Describe) {
		return nil, false
	}
	return o.Describe, true
}

// HasDescribe returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasDescribe() bool {
	if o != nil && !IsNil(o.Describe) {
		return true
	}

	return false
}

// SetDescribe gets a reference to the given string and assigns it to the Describe field.
func (o *MergeRequestInfo) SetDescribe(v string) {
	o.Describe = &v
}

// GetGranted returns the Granted field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetGranted() int64 {
	if o == nil || IsNil(o.Granted) {
		var ret int64
		return ret
	}
	return *o.Granted
}

// GetGrantedOk returns a tuple with the Granted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetGrantedOk() (*int64, bool) {
	if o == nil || IsNil(o.Granted) {
		return nil, false
	}
	return o.Granted, true
}

// HasGranted returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasGranted() bool {
	if o != nil && !IsNil(o.Granted) {
		return true
	}

	return false
}

// SetGranted gets a reference to the given int64 and assigns it to the Granted field.
func (o *MergeRequestInfo) SetGranted(v int64) {
	o.Granted = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MergeRequestInfo) SetId(v int64) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestInfo) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestInfo) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *MergeRequestInfo) SetLabels(v []string) {
	o.Labels = v
}

// GetMergeCommitSha returns the MergeCommitSha field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetMergeCommitSha() string {
	if o == nil || IsNil(o.MergeCommitSha) {
		var ret string
		return ret
	}
	return *o.MergeCommitSha
}

// GetMergeCommitShaOk returns a tuple with the MergeCommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetMergeCommitShaOk() (*string, bool) {
	if o == nil || IsNil(o.MergeCommitSha) {
		return nil, false
	}
	return o.MergeCommitSha, true
}

// HasMergeCommitSha returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasMergeCommitSha() bool {
	if o != nil && !IsNil(o.MergeCommitSha) {
		return true
	}

	return false
}

// SetMergeCommitSha gets a reference to the given string and assigns it to the MergeCommitSha field.
func (o *MergeRequestInfo) SetMergeCommitSha(v string) {
	o.MergeCommitSha = &v
}

// GetMergeId returns the MergeId field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetMergeId() int64 {
	if o == nil || IsNil(o.MergeId) {
		var ret int64
		return ret
	}
	return *o.MergeId
}

// GetMergeIdOk returns a tuple with the MergeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetMergeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.MergeId) {
		return nil, false
	}
	return o.MergeId, true
}

// HasMergeId returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasMergeId() bool {
	if o != nil && !IsNil(o.MergeId) {
		return true
	}

	return false
}

// SetMergeId gets a reference to the given int64 and assigns it to the MergeId field.
func (o *MergeRequestInfo) SetMergeId(v int64) {
	o.MergeId = &v
}

// GetMission returns the Mission field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetMission() Mission {
	if o == nil || IsNil(o.Mission) {
		var ret Mission
		return ret
	}
	return *o.Mission
}

// GetMissionOk returns a tuple with the Mission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetMissionOk() (*Mission, bool) {
	if o == nil || IsNil(o.Mission) {
		return nil, false
	}
	return o.Mission, true
}

// HasMission returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasMission() bool {
	if o != nil && !IsNil(o.Mission) {
		return true
	}

	return false
}

// SetMission gets a reference to the given Mission and assigns it to the Mission field.
func (o *MergeRequestInfo) SetMission(v Mission) {
	o.Mission = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MergeRequestInfo) SetPath(v string) {
	o.Path = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetProjectId() int64 {
	if o == nil || IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetProjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int64 and assigns it to the ProjectId field.
func (o *MergeRequestInfo) SetProjectId(v int64) {
	o.ProjectId = &v
}

// GetReminded returns the Reminded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestInfo) GetReminded() bool {
	if o == nil || IsNil(o.Reminded.Get()) {
		var ret bool
		return ret
	}
	return *o.Reminded.Get()
}

// GetRemindedOk returns a tuple with the Reminded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestInfo) GetRemindedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reminded.Get(), o.Reminded.IsSet()
}

// HasReminded returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasReminded() bool {
	if o != nil && o.Reminded.IsSet() {
		return true
	}

	return false
}

// SetReminded gets a reference to the given NullableBool and assigns it to the Reminded field.
func (o *MergeRequestInfo) SetReminded(v bool) {
	o.Reminded.Set(&v)
}
// SetRemindedNil sets the value for Reminded to be an explicit nil
func (o *MergeRequestInfo) SetRemindedNil() {
	o.Reminded.Set(nil)
}

// UnsetReminded ensures that no value is present for Reminded, not even an explicit nil
func (o *MergeRequestInfo) UnsetReminded() {
	o.Reminded.Unset()
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MergeRequestInfo) GetReviewers() []CodingUser {
	if o == nil {
		var ret []CodingUser
		return ret
	}
	return o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MergeRequestInfo) GetReviewersOk() ([]CodingUser, bool) {
	if o == nil || IsNil(o.Reviewers) {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasReviewers() bool {
	if o != nil && !IsNil(o.Reviewers) {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []CodingUser and assigns it to the Reviewers field.
func (o *MergeRequestInfo) SetReviewers(v []CodingUser) {
	o.Reviewers = v
}

// GetSourceBranch returns the SourceBranch field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetSourceBranch() string {
	if o == nil || IsNil(o.SourceBranch) {
		var ret string
		return ret
	}
	return *o.SourceBranch
}

// GetSourceBranchOk returns a tuple with the SourceBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetSourceBranchOk() (*string, bool) {
	if o == nil || IsNil(o.SourceBranch) {
		return nil, false
	}
	return o.SourceBranch, true
}

// HasSourceBranch returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasSourceBranch() bool {
	if o != nil && !IsNil(o.SourceBranch) {
		return true
	}

	return false
}

// SetSourceBranch gets a reference to the given string and assigns it to the SourceBranch field.
func (o *MergeRequestInfo) SetSourceBranch(v string) {
	o.SourceBranch = &v
}

// GetSourceBranchSha returns the SourceBranchSha field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetSourceBranchSha() string {
	if o == nil || IsNil(o.SourceBranchSha) {
		var ret string
		return ret
	}
	return *o.SourceBranchSha
}

// GetSourceBranchShaOk returns a tuple with the SourceBranchSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetSourceBranchShaOk() (*string, bool) {
	if o == nil || IsNil(o.SourceBranchSha) {
		return nil, false
	}
	return o.SourceBranchSha, true
}

// HasSourceBranchSha returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasSourceBranchSha() bool {
	if o != nil && !IsNil(o.SourceBranchSha) {
		return true
	}

	return false
}

// SetSourceBranchSha gets a reference to the given string and assigns it to the SourceBranchSha field.
func (o *MergeRequestInfo) SetSourceBranchSha(v string) {
	o.SourceBranchSha = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MergeRequestInfo) SetStatus(v string) {
	o.Status = &v
}

// GetStickingPoint returns the StickingPoint field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetStickingPoint() string {
	if o == nil || IsNil(o.StickingPoint) {
		var ret string
		return ret
	}
	return *o.StickingPoint
}

// GetStickingPointOk returns a tuple with the StickingPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetStickingPointOk() (*string, bool) {
	if o == nil || IsNil(o.StickingPoint) {
		return nil, false
	}
	return o.StickingPoint, true
}

// HasStickingPoint returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasStickingPoint() bool {
	if o != nil && !IsNil(o.StickingPoint) {
		return true
	}

	return false
}

// SetStickingPoint gets a reference to the given string and assigns it to the StickingPoint field.
func (o *MergeRequestInfo) SetStickingPoint(v string) {
	o.StickingPoint = &v
}

// GetTargetBranch returns the TargetBranch field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetTargetBranch() string {
	if o == nil || IsNil(o.TargetBranch) {
		var ret string
		return ret
	}
	return *o.TargetBranch
}

// GetTargetBranchOk returns a tuple with the TargetBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetTargetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.TargetBranch) {
		return nil, false
	}
	return o.TargetBranch, true
}

// HasTargetBranch returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasTargetBranch() bool {
	if o != nil && !IsNil(o.TargetBranch) {
		return true
	}

	return false
}

// SetTargetBranch gets a reference to the given string and assigns it to the TargetBranch field.
func (o *MergeRequestInfo) SetTargetBranch(v string) {
	o.TargetBranch = &v
}

// GetTargetBranchProtected returns the TargetBranchProtected field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetTargetBranchProtected() bool {
	if o == nil || IsNil(o.TargetBranchProtected) {
		var ret bool
		return ret
	}
	return *o.TargetBranchProtected
}

// GetTargetBranchProtectedOk returns a tuple with the TargetBranchProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetTargetBranchProtectedOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetBranchProtected) {
		return nil, false
	}
	return o.TargetBranchProtected, true
}

// HasTargetBranchProtected returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasTargetBranchProtected() bool {
	if o != nil && !IsNil(o.TargetBranchProtected) {
		return true
	}

	return false
}

// SetTargetBranchProtected gets a reference to the given bool and assigns it to the TargetBranchProtected field.
func (o *MergeRequestInfo) SetTargetBranchProtected(v bool) {
	o.TargetBranchProtected = &v
}

// GetTargetBranchSha returns the TargetBranchSha field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetTargetBranchSha() string {
	if o == nil || IsNil(o.TargetBranchSha) {
		var ret string
		return ret
	}
	return *o.TargetBranchSha
}

// GetTargetBranchShaOk returns a tuple with the TargetBranchSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetTargetBranchShaOk() (*string, bool) {
	if o == nil || IsNil(o.TargetBranchSha) {
		return nil, false
	}
	return o.TargetBranchSha, true
}

// HasTargetBranchSha returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasTargetBranchSha() bool {
	if o != nil && !IsNil(o.TargetBranchSha) {
		return true
	}

	return false
}

// SetTargetBranchSha gets a reference to the given string and assigns it to the TargetBranchSha field.
func (o *MergeRequestInfo) SetTargetBranchSha(v string) {
	o.TargetBranchSha = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MergeRequestInfo) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MergeRequestInfo) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeRequestInfo) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MergeRequestInfo) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *MergeRequestInfo) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

func (o MergeRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeRequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionAt) {
		toSerialize["ActionAt"] = o.ActionAt
	}
	if !IsNil(o.ActionAuthor) {
		toSerialize["ActionAuthor"] = o.ActionAuthor
	}
	if !IsNil(o.Author) {
		toSerialize["Author"] = o.Author
	}
	if !IsNil(o.BaseSha) {
		toSerialize["BaseSha"] = o.BaseSha
	}
	if !IsNil(o.CommentCount) {
		toSerialize["CommentCount"] = o.CommentCount
	}
	if !IsNil(o.Conflicts) {
		toSerialize["Conflicts"] = o.Conflicts
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !IsNil(o.DepotId) {
		toSerialize["DepotId"] = o.DepotId
	}
	if !IsNil(o.Describe) {
		toSerialize["Describe"] = o.Describe
	}
	if !IsNil(o.Granted) {
		toSerialize["Granted"] = o.Granted
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Labels != nil {
		toSerialize["Labels"] = o.Labels
	}
	if !IsNil(o.MergeCommitSha) {
		toSerialize["MergeCommitSha"] = o.MergeCommitSha
	}
	if !IsNil(o.MergeId) {
		toSerialize["MergeId"] = o.MergeId
	}
	if !IsNil(o.Mission) {
		toSerialize["Mission"] = o.Mission
	}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if !IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if o.Reminded.IsSet() {
		toSerialize["Reminded"] = o.Reminded.Get()
	}
	if o.Reviewers != nil {
		toSerialize["Reviewers"] = o.Reviewers
	}
	if !IsNil(o.SourceBranch) {
		toSerialize["SourceBranch"] = o.SourceBranch
	}
	if !IsNil(o.SourceBranchSha) {
		toSerialize["SourceBranchSha"] = o.SourceBranchSha
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if !IsNil(o.StickingPoint) {
		toSerialize["StickingPoint"] = o.StickingPoint
	}
	if !IsNil(o.TargetBranch) {
		toSerialize["TargetBranch"] = o.TargetBranch
	}
	if !IsNil(o.TargetBranchProtected) {
		toSerialize["TargetBranchProtected"] = o.TargetBranchProtected
	}
	if !IsNil(o.TargetBranchSha) {
		toSerialize["TargetBranchSha"] = o.TargetBranchSha
	}
	if !IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableMergeRequestInfo struct {
	value *MergeRequestInfo
	isSet bool
}

func (v NullableMergeRequestInfo) Get() *MergeRequestInfo {
	return v.value
}

func (v *NullableMergeRequestInfo) Set(val *MergeRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeRequestInfo(val *MergeRequestInfo) *NullableMergeRequestInfo {
	return &NullableMergeRequestInfo{value: val, isSet: true}
}

func (v NullableMergeRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


