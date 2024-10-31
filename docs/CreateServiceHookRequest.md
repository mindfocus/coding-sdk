# CreateServiceHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionProperty** | **string** | 发送方式属性 | 
**Enabled** | **bool** | 事件开关 | 
**Event** | **[]string** | 事件名，取值范围：ITERATION_CREATED,ITERATION_DELETED,ITERATION_UPDATED,ISSUE_CREATED,ISSUE_DELETED,ISSUE_COMMENT_CREATED,ISSUE_STATUS_UPDATED,ISSUE_ASSIGNEE_CHANGED,ISSUE_ITERATION_CHANGED,ISSUE_RELATIONSHIP_CHANGED,ISSUE_UPDATED,GIT_PUSHED,GIT_MR_CREATED,GIT_MR_UPDATED,GIT_MR_MERGED,GIT_MR_CLOSED,ARTIFACTS_VERSION_CREATED,ARTIFACTS_VERSION_UPDATED,ARTIFACTS_VERSION_DOWNLOADED,ARTIFACTS_VERSION_DELETED,ARTIFACTS_VERSION_RELEASED,ARTIFACTS_VERSION_DOWNLOAD_FORBIDDEN,ARTIFACTS_VERSION_DOWNLOAD_ALLOWED,ARTIFACTS_VERSION_DOWNLOAD_BLOCKED,ARTIFACTS_REPO_CREATED,ARTIFACTS_REPO_UPDATED,ARTIFACTS_REPO_DELETED,CI_JOB_CREATED,CI_JOB_UPDATED,CI_JOB_DELETED,CI_JOB_STARTED,CI_JOB_FINISHED,FILE_CREATED,FILE_UPDATED,FILE_RENAMED,FILE_SHARE_UPDATED,FILE_MOVED,FILE_COPIED,FILE_MOVED_TO_RECYCLE_BIN,FILE_RESTORED_FROM_RECYCLE_BIN,FILE_DELETED,WIKI_CREATED,WIKI_UPDATED,WIKI_MOVED,WIKI_SHARE_UPDATED,WIKI_ACCESS_UPDATED,WIKI_COPIED,WIKI_MOVED_TO_RECYCLE_BIN,WIKI_RESTORED_FROM_RECYCLE_BIN,WIKI_DELETED,MEMBER_CREATED,MEMBER_DELETED,MEMBER_ROLE_UPDATED,TEST_PLAN_CREATED, TEST_PLAN_UPDATED, TEST_PLAN_FINISHED, TEST_TASK_ASSIGNED, TEST_REPORT_CREATED, FLEXIBLE_TESTX_REVIEW_CREATED, FLEXIBLE_TESTX_REVIEW_COMMENTED, FLEXIBLE_TESTX_REVIEW_UPDATED, FLEXIBLE_TESTX_REVIEW_COMPLETED, FLEXIBLE_TESTX_PLAN_CREATED, FLEXIBLE_TESTX_PLAN_TASK_ASSIGNED, FLEXIBLE_TESTX_PLAN_UPDATED, FLEXIBLE_TESTX_PLAN_FINISHED, FLEXIBLE_TESTX_REPORT_CREATED, CODE_DOG_CREATE_JOB, CODE_DOG_RESULT_NOTIFY, PLAN_CREATED, PLAN_DELETED, PLAN_COMMENT_CREATED, PLAN_STATUE_CHANGED, PLAN_ASSIGNEE_CHANGED, PLAN_UPDATED, RISK_CREATED, RISK_DELETED, RISK_COMMENT_CREATED, RISK_STATUS_CHANGED, RISK_ASSIGNEE_CHANGED, RISK_UPDATED | 
**FilterProperty** | Pointer to **string** | 过滤器属性 | [optional] 
**Name** | Pointer to **string** | 备注名 | [optional] 
**ProjectId** | **int64** | 项目编号或者研发空间编号 | 
**Service** | **string** | 服务名，取值范围：WebHook、WeCom、DingDing、Jenkins、FeiShu。 | 
**ServiceAction** | **string** | 发送方式，取值范围：dingding_group_chat_robot,wecom_group_chat_robot,jenkins_generic_build_job,feishu_group_chat_robot,webhook_http_post,webhook_http_get | 
**TargetType** | Pointer to **string** | 目标数据类型：PROJECT,SPACE_NODE,PROGRAM,默认PROJECT | [optional] 

## Methods

### NewCreateServiceHookRequest

`func NewCreateServiceHookRequest(actionProperty string, enabled bool, event []string, projectId int64, service string, serviceAction string, ) *CreateServiceHookRequest`

NewCreateServiceHookRequest instantiates a new CreateServiceHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceHookRequestWithDefaults

`func NewCreateServiceHookRequestWithDefaults() *CreateServiceHookRequest`

NewCreateServiceHookRequestWithDefaults instantiates a new CreateServiceHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionProperty

`func (o *CreateServiceHookRequest) GetActionProperty() string`

GetActionProperty returns the ActionProperty field if non-nil, zero value otherwise.

### GetActionPropertyOk

`func (o *CreateServiceHookRequest) GetActionPropertyOk() (*string, bool)`

GetActionPropertyOk returns a tuple with the ActionProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionProperty

`func (o *CreateServiceHookRequest) SetActionProperty(v string)`

SetActionProperty sets ActionProperty field to given value.


### GetEnabled

`func (o *CreateServiceHookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateServiceHookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateServiceHookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEvent

`func (o *CreateServiceHookRequest) GetEvent() []string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CreateServiceHookRequest) GetEventOk() (*[]string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CreateServiceHookRequest) SetEvent(v []string)`

SetEvent sets Event field to given value.


### GetFilterProperty

`func (o *CreateServiceHookRequest) GetFilterProperty() string`

GetFilterProperty returns the FilterProperty field if non-nil, zero value otherwise.

### GetFilterPropertyOk

`func (o *CreateServiceHookRequest) GetFilterPropertyOk() (*string, bool)`

GetFilterPropertyOk returns a tuple with the FilterProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProperty

`func (o *CreateServiceHookRequest) SetFilterProperty(v string)`

SetFilterProperty sets FilterProperty field to given value.

### HasFilterProperty

`func (o *CreateServiceHookRequest) HasFilterProperty() bool`

HasFilterProperty returns a boolean if a field has been set.

### GetName

`func (o *CreateServiceHookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceHookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceHookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateServiceHookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateServiceHookRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateServiceHookRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateServiceHookRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetService

`func (o *CreateServiceHookRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CreateServiceHookRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CreateServiceHookRequest) SetService(v string)`

SetService sets Service field to given value.


### GetServiceAction

`func (o *CreateServiceHookRequest) GetServiceAction() string`

GetServiceAction returns the ServiceAction field if non-nil, zero value otherwise.

### GetServiceActionOk

`func (o *CreateServiceHookRequest) GetServiceActionOk() (*string, bool)`

GetServiceActionOk returns a tuple with the ServiceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAction

`func (o *CreateServiceHookRequest) SetServiceAction(v string)`

SetServiceAction sets ServiceAction field to given value.


### GetTargetType

`func (o *CreateServiceHookRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CreateServiceHookRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CreateServiceHookRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CreateServiceHookRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


