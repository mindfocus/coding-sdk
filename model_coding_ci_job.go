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

// checks if the CodingCIJob type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodingCIJob{}

// CodingCIJob 构建计划信息
type CodingCIJob struct {
	// 不管构建成功还是失败总是通知的用户
	AlwaysUserIdList []int32 `json:"AlwaysUserIdList"`
	// 自动取消相同 MR
	AutoCancelSameMergeRequest bool `json:"AutoCancelSameMergeRequest"`
	// 自动取消相同版本
	AutoCancelSameRevision bool `json:"AutoCancelSameRevision"`
	// 分支匹配正则
	BranchRegex string `json:"BranchRegex"`
	// 触发构建的分支
	BranchSelector string `json:"BranchSelector"`
	// 仅构建失败时要通知的用户
	BuildFailUserIdList []int32 `json:"BuildFailUserIdList"`
	// 任务缓存目录配置
	CachePathList []CIJobCachePath `json:"CachePathList"`
	// 构建缓存大小
	CacheSize *int32 `json:"CacheSize,omitempty"`
	// 创建时间
	CreatedAt int32 `json:"CreatedAt"`
	// 创建者
	CreatorId int32 `json:"CreatorId"`
	// 仓库库的 Https 地址
	DepotHttpsUrl *string `json:"DepotHttpsUrl,omitempty"`
	// 仓库ID
	DepotId int32 `json:"DepotId"`
	// 仓库名称
	DepotName *string `json:"DepotName,omitempty"`
	// 仓库库的 SSH 地址
	DepotSshUrl *string `json:"DepotSshUrl,omitempty"`
	// 仓库类型
	DepotType string `json:"DepotType"`
	// 仓库库的 Web 页面
	DepotWebUrl *string `json:"DepotWebUrl,omitempty"`
	// 环境变量配置
	EnvList []CIJobEnv `json:"EnvList"`
	// 执行方式
	ExecuteIn string `json:"ExecuteIn"`
	// 自定义构建节点池 ID
	ExecutedAgentPoolId *int32 `json:"ExecutedAgentPoolId,omitempty"`
	// 代码更新触发匹配规则
	HookType string `json:"HookType"`
	// 构建计划ID
	Id *int32 `json:"Id,omitempty"`
	// Jenkinsfile 来源
	JenkinsFileFromType string `json:"JenkinsFileFromType"`
	// Jenkinsfile 在仓库中的文件路径
	JenkinsFilePath string `json:"JenkinsFilePath"`
	// Jenkinsfile 文件内容
	JenkinsFileStaticContent string `json:"JenkinsFileStaticContent"`
	// 构建计划创建来源
	JobFromType string `json:"JobFromType"`
	// 构建计划名称
	Name string `json:"Name"`
	// 项目ID
	ProjectId int32 `json:"ProjectId"`
	// 项目名称
	ProjectName *string `json:"ProjectName,omitempty"`
	// 针对 CRON triggerMethod 的 schedule 规则配置
	ScheduleList []CIJobSchedule `json:"ScheduleList"`
	// 构建计划触发方式
	TriggerMethodList []string `json:"TriggerMethodList"`
	// 构建结果通知触发者机制
	TriggerRemind string `json:"TriggerRemind"`
	// 最后更新时间
	UpdatedAt int32 `json:"UpdatedAt"`
}

type _CodingCIJob CodingCIJob

// NewCodingCIJob instantiates a new CodingCIJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodingCIJob(alwaysUserIdList []int32, autoCancelSameMergeRequest bool, autoCancelSameRevision bool, branchRegex string, branchSelector string, buildFailUserIdList []int32, cachePathList []CIJobCachePath, createdAt int32, creatorId int32, depotId int32, depotType string, envList []CIJobEnv, executeIn string, hookType string, jenkinsFileFromType string, jenkinsFilePath string, jenkinsFileStaticContent string, jobFromType string, name string, projectId int32, scheduleList []CIJobSchedule, triggerMethodList []string, triggerRemind string, updatedAt int32) *CodingCIJob {
	this := CodingCIJob{}
	this.AlwaysUserIdList = alwaysUserIdList
	this.AutoCancelSameMergeRequest = autoCancelSameMergeRequest
	this.AutoCancelSameRevision = autoCancelSameRevision
	this.BranchRegex = branchRegex
	this.BranchSelector = branchSelector
	this.BuildFailUserIdList = buildFailUserIdList
	this.CachePathList = cachePathList
	this.CreatedAt = createdAt
	this.CreatorId = creatorId
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = &depotHttpsUrl
	this.DepotId = depotId
	var depotName string = ""
	this.DepotName = &depotName
	var depotSshUrl string = ""
	this.DepotSshUrl = &depotSshUrl
	this.DepotType = depotType
	var depotWebUrl string = ""
	this.DepotWebUrl = &depotWebUrl
	this.EnvList = envList
	this.ExecuteIn = executeIn
	this.HookType = hookType
	this.JenkinsFileFromType = jenkinsFileFromType
	this.JenkinsFilePath = jenkinsFilePath
	this.JenkinsFileStaticContent = jenkinsFileStaticContent
	this.JobFromType = jobFromType
	this.Name = name
	this.ProjectId = projectId
	var projectName string = ""
	this.ProjectName = &projectName
	this.ScheduleList = scheduleList
	this.TriggerMethodList = triggerMethodList
	this.TriggerRemind = triggerRemind
	this.UpdatedAt = updatedAt
	return &this
}

// NewCodingCIJobWithDefaults instantiates a new CodingCIJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodingCIJobWithDefaults() *CodingCIJob {
	this := CodingCIJob{}
	var autoCancelSameMergeRequest bool = false
	this.AutoCancelSameMergeRequest = autoCancelSameMergeRequest
	var autoCancelSameRevision bool = false
	this.AutoCancelSameRevision = autoCancelSameRevision
	var branchRegex string = ""
	this.BranchRegex = branchRegex
	var branchSelector string = ""
	this.BranchSelector = branchSelector
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = &depotHttpsUrl
	var depotName string = ""
	this.DepotName = &depotName
	var depotSshUrl string = ""
	this.DepotSshUrl = &depotSshUrl
	var depotType string = ""
	this.DepotType = depotType
	var depotWebUrl string = ""
	this.DepotWebUrl = &depotWebUrl
	var executeIn string = ""
	this.ExecuteIn = executeIn
	var hookType string = ""
	this.HookType = hookType
	var jenkinsFileFromType string = ""
	this.JenkinsFileFromType = jenkinsFileFromType
	var jenkinsFilePath string = ""
	this.JenkinsFilePath = jenkinsFilePath
	var jenkinsFileStaticContent string = ""
	this.JenkinsFileStaticContent = jenkinsFileStaticContent
	var jobFromType string = ""
	this.JobFromType = jobFromType
	var name string = ""
	this.Name = name
	var projectName string = ""
	this.ProjectName = &projectName
	var triggerRemind string = ""
	this.TriggerRemind = triggerRemind
	return &this
}

// GetAlwaysUserIdList returns the AlwaysUserIdList field value
func (o *CodingCIJob) GetAlwaysUserIdList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AlwaysUserIdList
}

// GetAlwaysUserIdListOk returns a tuple with the AlwaysUserIdList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAlwaysUserIdListOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlwaysUserIdList, true
}

// SetAlwaysUserIdList sets field value
func (o *CodingCIJob) SetAlwaysUserIdList(v []int32) {
	o.AlwaysUserIdList = v
}

// GetAutoCancelSameMergeRequest returns the AutoCancelSameMergeRequest field value
func (o *CodingCIJob) GetAutoCancelSameMergeRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCancelSameMergeRequest
}

// GetAutoCancelSameMergeRequestOk returns a tuple with the AutoCancelSameMergeRequest field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAutoCancelSameMergeRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCancelSameMergeRequest, true
}

// SetAutoCancelSameMergeRequest sets field value
func (o *CodingCIJob) SetAutoCancelSameMergeRequest(v bool) {
	o.AutoCancelSameMergeRequest = v
}

// GetAutoCancelSameRevision returns the AutoCancelSameRevision field value
func (o *CodingCIJob) GetAutoCancelSameRevision() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCancelSameRevision
}

// GetAutoCancelSameRevisionOk returns a tuple with the AutoCancelSameRevision field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetAutoCancelSameRevisionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCancelSameRevision, true
}

// SetAutoCancelSameRevision sets field value
func (o *CodingCIJob) SetAutoCancelSameRevision(v bool) {
	o.AutoCancelSameRevision = v
}

// GetBranchRegex returns the BranchRegex field value
func (o *CodingCIJob) GetBranchRegex() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchRegex
}

// GetBranchRegexOk returns a tuple with the BranchRegex field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBranchRegexOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchRegex, true
}

// SetBranchRegex sets field value
func (o *CodingCIJob) SetBranchRegex(v string) {
	o.BranchRegex = v
}

// GetBranchSelector returns the BranchSelector field value
func (o *CodingCIJob) GetBranchSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchSelector
}

// GetBranchSelectorOk returns a tuple with the BranchSelector field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBranchSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchSelector, true
}

// SetBranchSelector sets field value
func (o *CodingCIJob) SetBranchSelector(v string) {
	o.BranchSelector = v
}

// GetBuildFailUserIdList returns the BuildFailUserIdList field value
func (o *CodingCIJob) GetBuildFailUserIdList() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.BuildFailUserIdList
}

// GetBuildFailUserIdListOk returns a tuple with the BuildFailUserIdList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetBuildFailUserIdListOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BuildFailUserIdList, true
}

// SetBuildFailUserIdList sets field value
func (o *CodingCIJob) SetBuildFailUserIdList(v []int32) {
	o.BuildFailUserIdList = v
}

// GetCachePathList returns the CachePathList field value
func (o *CodingCIJob) GetCachePathList() []CIJobCachePath {
	if o == nil {
		var ret []CIJobCachePath
		return ret
	}

	return o.CachePathList
}

// GetCachePathListOk returns a tuple with the CachePathList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCachePathListOk() ([]CIJobCachePath, bool) {
	if o == nil {
		return nil, false
	}
	return o.CachePathList, true
}

// SetCachePathList sets field value
func (o *CodingCIJob) SetCachePathList(v []CIJobCachePath) {
	o.CachePathList = v
}

// GetCacheSize returns the CacheSize field value if set, zero value otherwise.
func (o *CodingCIJob) GetCacheSize() int32 {
	if o == nil || IsNil(o.CacheSize) {
		var ret int32
		return ret
	}
	return *o.CacheSize
}

// GetCacheSizeOk returns a tuple with the CacheSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCacheSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheSize) {
		return nil, false
	}
	return o.CacheSize, true
}

// HasCacheSize returns a boolean if a field has been set.
func (o *CodingCIJob) HasCacheSize() bool {
	if o != nil && !IsNil(o.CacheSize) {
		return true
	}

	return false
}

// SetCacheSize gets a reference to the given int32 and assigns it to the CacheSize field.
func (o *CodingCIJob) SetCacheSize(v int32) {
	o.CacheSize = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CodingCIJob) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CodingCIJob) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetCreatorId returns the CreatorId field value
func (o *CodingCIJob) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *CodingCIJob) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetDepotHttpsUrl returns the DepotHttpsUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotHttpsUrl() string {
	if o == nil || IsNil(o.DepotHttpsUrl) {
		var ret string
		return ret
	}
	return *o.DepotHttpsUrl
}

// GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotHttpsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DepotHttpsUrl) {
		return nil, false
	}
	return o.DepotHttpsUrl, true
}

// HasDepotHttpsUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotHttpsUrl() bool {
	if o != nil && !IsNil(o.DepotHttpsUrl) {
		return true
	}

	return false
}

// SetDepotHttpsUrl gets a reference to the given string and assigns it to the DepotHttpsUrl field.
func (o *CodingCIJob) SetDepotHttpsUrl(v string) {
	o.DepotHttpsUrl = &v
}

// GetDepotId returns the DepotId field value
func (o *CodingCIJob) GetDepotId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *CodingCIJob) SetDepotId(v int32) {
	o.DepotId = v
}

// GetDepotName returns the DepotName field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotName() string {
	if o == nil || IsNil(o.DepotName) {
		var ret string
		return ret
	}
	return *o.DepotName
}

// GetDepotNameOk returns a tuple with the DepotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotNameOk() (*string, bool) {
	if o == nil || IsNil(o.DepotName) {
		return nil, false
	}
	return o.DepotName, true
}

// HasDepotName returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotName() bool {
	if o != nil && !IsNil(o.DepotName) {
		return true
	}

	return false
}

// SetDepotName gets a reference to the given string and assigns it to the DepotName field.
func (o *CodingCIJob) SetDepotName(v string) {
	o.DepotName = &v
}

// GetDepotSshUrl returns the DepotSshUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotSshUrl() string {
	if o == nil || IsNil(o.DepotSshUrl) {
		var ret string
		return ret
	}
	return *o.DepotSshUrl
}

// GetDepotSshUrlOk returns a tuple with the DepotSshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotSshUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DepotSshUrl) {
		return nil, false
	}
	return o.DepotSshUrl, true
}

// HasDepotSshUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotSshUrl() bool {
	if o != nil && !IsNil(o.DepotSshUrl) {
		return true
	}

	return false
}

// SetDepotSshUrl gets a reference to the given string and assigns it to the DepotSshUrl field.
func (o *CodingCIJob) SetDepotSshUrl(v string) {
	o.DepotSshUrl = &v
}

// GetDepotType returns the DepotType field value
func (o *CodingCIJob) GetDepotType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotType
}

// GetDepotTypeOk returns a tuple with the DepotType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotType, true
}

// SetDepotType sets field value
func (o *CodingCIJob) SetDepotType(v string) {
	o.DepotType = v
}

// GetDepotWebUrl returns the DepotWebUrl field value if set, zero value otherwise.
func (o *CodingCIJob) GetDepotWebUrl() string {
	if o == nil || IsNil(o.DepotWebUrl) {
		var ret string
		return ret
	}
	return *o.DepotWebUrl
}

// GetDepotWebUrlOk returns a tuple with the DepotWebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetDepotWebUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DepotWebUrl) {
		return nil, false
	}
	return o.DepotWebUrl, true
}

// HasDepotWebUrl returns a boolean if a field has been set.
func (o *CodingCIJob) HasDepotWebUrl() bool {
	if o != nil && !IsNil(o.DepotWebUrl) {
		return true
	}

	return false
}

// SetDepotWebUrl gets a reference to the given string and assigns it to the DepotWebUrl field.
func (o *CodingCIJob) SetDepotWebUrl(v string) {
	o.DepotWebUrl = &v
}

// GetEnvList returns the EnvList field value
func (o *CodingCIJob) GetEnvList() []CIJobEnv {
	if o == nil {
		var ret []CIJobEnv
		return ret
	}

	return o.EnvList
}

// GetEnvListOk returns a tuple with the EnvList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetEnvListOk() ([]CIJobEnv, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvList, true
}

// SetEnvList sets field value
func (o *CodingCIJob) SetEnvList(v []CIJobEnv) {
	o.EnvList = v
}

// GetExecuteIn returns the ExecuteIn field value
func (o *CodingCIJob) GetExecuteIn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecuteIn
}

// GetExecuteInOk returns a tuple with the ExecuteIn field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetExecuteInOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExecuteIn, true
}

// SetExecuteIn sets field value
func (o *CodingCIJob) SetExecuteIn(v string) {
	o.ExecuteIn = v
}

// GetExecutedAgentPoolId returns the ExecutedAgentPoolId field value if set, zero value otherwise.
func (o *CodingCIJob) GetExecutedAgentPoolId() int32 {
	if o == nil || IsNil(o.ExecutedAgentPoolId) {
		var ret int32
		return ret
	}
	return *o.ExecutedAgentPoolId
}

// GetExecutedAgentPoolIdOk returns a tuple with the ExecutedAgentPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetExecutedAgentPoolIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ExecutedAgentPoolId) {
		return nil, false
	}
	return o.ExecutedAgentPoolId, true
}

// HasExecutedAgentPoolId returns a boolean if a field has been set.
func (o *CodingCIJob) HasExecutedAgentPoolId() bool {
	if o != nil && !IsNil(o.ExecutedAgentPoolId) {
		return true
	}

	return false
}

// SetExecutedAgentPoolId gets a reference to the given int32 and assigns it to the ExecutedAgentPoolId field.
func (o *CodingCIJob) SetExecutedAgentPoolId(v int32) {
	o.ExecutedAgentPoolId = &v
}

// GetHookType returns the HookType field value
func (o *CodingCIJob) GetHookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HookType
}

// GetHookTypeOk returns a tuple with the HookType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetHookTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HookType, true
}

// SetHookType sets field value
func (o *CodingCIJob) SetHookType(v string) {
	o.HookType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CodingCIJob) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CodingCIJob) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CodingCIJob) SetId(v int32) {
	o.Id = &v
}

// GetJenkinsFileFromType returns the JenkinsFileFromType field value
func (o *CodingCIJob) GetJenkinsFileFromType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFileFromType
}

// GetJenkinsFileFromTypeOk returns a tuple with the JenkinsFileFromType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFileFromTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFileFromType, true
}

// SetJenkinsFileFromType sets field value
func (o *CodingCIJob) SetJenkinsFileFromType(v string) {
	o.JenkinsFileFromType = v
}

// GetJenkinsFilePath returns the JenkinsFilePath field value
func (o *CodingCIJob) GetJenkinsFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFilePath
}

// GetJenkinsFilePathOk returns a tuple with the JenkinsFilePath field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFilePath, true
}

// SetJenkinsFilePath sets field value
func (o *CodingCIJob) SetJenkinsFilePath(v string) {
	o.JenkinsFilePath = v
}

// GetJenkinsFileStaticContent returns the JenkinsFileStaticContent field value
func (o *CodingCIJob) GetJenkinsFileStaticContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFileStaticContent
}

// GetJenkinsFileStaticContentOk returns a tuple with the JenkinsFileStaticContent field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJenkinsFileStaticContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFileStaticContent, true
}

// SetJenkinsFileStaticContent sets field value
func (o *CodingCIJob) SetJenkinsFileStaticContent(v string) {
	o.JenkinsFileStaticContent = v
}

// GetJobFromType returns the JobFromType field value
func (o *CodingCIJob) GetJobFromType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobFromType
}

// GetJobFromTypeOk returns a tuple with the JobFromType field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetJobFromTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobFromType, true
}

// SetJobFromType sets field value
func (o *CodingCIJob) SetJobFromType(v string) {
	o.JobFromType = v
}

// GetName returns the Name field value
func (o *CodingCIJob) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CodingCIJob) SetName(v string) {
	o.Name = v
}

// GetProjectId returns the ProjectId field value
func (o *CodingCIJob) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CodingCIJob) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *CodingCIJob) GetProjectName() string {
	if o == nil || IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetProjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *CodingCIJob) HasProjectName() bool {
	if o != nil && !IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *CodingCIJob) SetProjectName(v string) {
	o.ProjectName = &v
}

// GetScheduleList returns the ScheduleList field value
func (o *CodingCIJob) GetScheduleList() []CIJobSchedule {
	if o == nil {
		var ret []CIJobSchedule
		return ret
	}

	return o.ScheduleList
}

// GetScheduleListOk returns a tuple with the ScheduleList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetScheduleListOk() ([]CIJobSchedule, bool) {
	if o == nil {
		return nil, false
	}
	return o.ScheduleList, true
}

// SetScheduleList sets field value
func (o *CodingCIJob) SetScheduleList(v []CIJobSchedule) {
	o.ScheduleList = v
}

// GetTriggerMethodList returns the TriggerMethodList field value
func (o *CodingCIJob) GetTriggerMethodList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TriggerMethodList
}

// GetTriggerMethodListOk returns a tuple with the TriggerMethodList field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetTriggerMethodListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TriggerMethodList, true
}

// SetTriggerMethodList sets field value
func (o *CodingCIJob) SetTriggerMethodList(v []string) {
	o.TriggerMethodList = v
}

// GetTriggerRemind returns the TriggerRemind field value
func (o *CodingCIJob) GetTriggerRemind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerRemind
}

// GetTriggerRemindOk returns a tuple with the TriggerRemind field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetTriggerRemindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerRemind, true
}

// SetTriggerRemind sets field value
func (o *CodingCIJob) SetTriggerRemind(v string) {
	o.TriggerRemind = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CodingCIJob) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CodingCIJob) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CodingCIJob) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o CodingCIJob) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodingCIJob) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AlwaysUserIdList"] = o.AlwaysUserIdList
	toSerialize["AutoCancelSameMergeRequest"] = o.AutoCancelSameMergeRequest
	toSerialize["AutoCancelSameRevision"] = o.AutoCancelSameRevision
	toSerialize["BranchRegex"] = o.BranchRegex
	toSerialize["BranchSelector"] = o.BranchSelector
	toSerialize["BuildFailUserIdList"] = o.BuildFailUserIdList
	toSerialize["CachePathList"] = o.CachePathList
	if !IsNil(o.CacheSize) {
		toSerialize["CacheSize"] = o.CacheSize
	}
	toSerialize["CreatedAt"] = o.CreatedAt
	toSerialize["CreatorId"] = o.CreatorId
	if !IsNil(o.DepotHttpsUrl) {
		toSerialize["DepotHttpsUrl"] = o.DepotHttpsUrl
	}
	toSerialize["DepotId"] = o.DepotId
	if !IsNil(o.DepotName) {
		toSerialize["DepotName"] = o.DepotName
	}
	if !IsNil(o.DepotSshUrl) {
		toSerialize["DepotSshUrl"] = o.DepotSshUrl
	}
	toSerialize["DepotType"] = o.DepotType
	if !IsNil(o.DepotWebUrl) {
		toSerialize["DepotWebUrl"] = o.DepotWebUrl
	}
	toSerialize["EnvList"] = o.EnvList
	toSerialize["ExecuteIn"] = o.ExecuteIn
	if !IsNil(o.ExecutedAgentPoolId) {
		toSerialize["ExecutedAgentPoolId"] = o.ExecutedAgentPoolId
	}
	toSerialize["HookType"] = o.HookType
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	toSerialize["JenkinsFileFromType"] = o.JenkinsFileFromType
	toSerialize["JenkinsFilePath"] = o.JenkinsFilePath
	toSerialize["JenkinsFileStaticContent"] = o.JenkinsFileStaticContent
	toSerialize["JobFromType"] = o.JobFromType
	toSerialize["Name"] = o.Name
	toSerialize["ProjectId"] = o.ProjectId
	if !IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	toSerialize["ScheduleList"] = o.ScheduleList
	toSerialize["TriggerMethodList"] = o.TriggerMethodList
	toSerialize["TriggerRemind"] = o.TriggerRemind
	toSerialize["UpdatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *CodingCIJob) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AlwaysUserIdList",
		"AutoCancelSameMergeRequest",
		"AutoCancelSameRevision",
		"BranchRegex",
		"BranchSelector",
		"BuildFailUserIdList",
		"CachePathList",
		"CreatedAt",
		"CreatorId",
		"DepotId",
		"DepotType",
		"EnvList",
		"ExecuteIn",
		"HookType",
		"JenkinsFileFromType",
		"JenkinsFilePath",
		"JenkinsFileStaticContent",
		"JobFromType",
		"Name",
		"ProjectId",
		"ScheduleList",
		"TriggerMethodList",
		"TriggerRemind",
		"UpdatedAt",
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

	varCodingCIJob := _CodingCIJob{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCodingCIJob)

	if err != nil {
		return err
	}

	*o = CodingCIJob(varCodingCIJob)

	return err
}

type NullableCodingCIJob struct {
	value *CodingCIJob
	isSet bool
}

func (v NullableCodingCIJob) Get() *CodingCIJob {
	return v.value
}

func (v *NullableCodingCIJob) Set(val *CodingCIJob) {
	v.value = val
	v.isSet = true
}

func (v NullableCodingCIJob) IsSet() bool {
	return v.isSet
}

func (v *NullableCodingCIJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodingCIJob(val *CodingCIJob) *NullableCodingCIJob {
	return &NullableCodingCIJob{value: val, isSet: true}
}

func (v NullableCodingCIJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodingCIJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


