# DescribeServiceHooksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatorBy** | Pointer to **[]int64** | 创建人编号 | [optional] 
**Enabled** | Pointer to **string** | 事件开关，取值范围：true、false | [optional] 
**Event** | Pointer to **[]string** | 事件名，取值范围：ITERATION_CREATED,ITERATION_DELETED,ITERATION_UPDATED,ISSUE_CREATED,ISSUE_DELETED,ISSUE_COMMENT_CREATED,ISSUE_STATUS_UPDATED,ISSUE_ASSIGNEE_CHANGED,ISSUE_ITERATION_CHANGED,ISSUE_RELATIONSHIP_CHANGED,ISSUE_UPDATED,GIT_PUSHED,GIT_MR_CREATED,GIT_MR_UPDATED,GIT_MR_MERGED,GIT_MR_CLOSED,ARTIFACTS_VERSION_CREATED,ARTIFACTS_VERSION_UPDATED,ARTIFACTS_VERSION_DOWNLOADED,ARTIFACTS_VERSION_DELETED,ARTIFACTS_VERSION_RELEASED,ARTIFACTS_VERSION_DOWNLOAD_FORBIDDEN,ARTIFACTS_VERSION_DOWNLOAD_ALLOWED,ARTIFACTS_VERSION_DOWNLOAD_BLOCKED,ARTIFACTS_REPO_CREATED,ARTIFACTS_REPO_UPDATED,ARTIFACTS_REPO_DELETED,CI_JOB_CREATED,CI_JOB_UPDATED,CI_JOB_DELETED,CI_JOB_STARTED,CI_JOB_FINISHED,FILE_CREATED,FILE_UPDATED,FILE_RENAMED,FILE_SHARE_UPDATED,FILE_MOVED,FILE_COPIED,FILE_MOVED_TO_RECYCLE_BIN,FILE_RESTORED_FROM_RECYCLE_BIN,FILE_DELETED,WIKI_CREATED,WIKI_UPDATED,WIKI_MOVED,WIKI_SHARE_UPDATED,WIKI_ACCESS_UPDATED,WIKI_COPIED,WIKI_MOVED_TO_RECYCLE_BIN,WIKI_RESTORED_FROM_RECYCLE_BIN,WIKI_DELETED,MEMBER_CREATED,MEMBER_DELETED,MEMBER_ROLE_UPDATED,TEST_PLAN_CREATED, TEST_PLAN_UPDATED, TEST_PLAN_FINISHED, TEST_TASK_ASSIGNED, TEST_REPORT_CREATED, FLEXIBLE_TESTX_REVIEW_CREATED, FLEXIBLE_TESTX_REVIEW_COMMENTED, FLEXIBLE_TESTX_REVIEW_UPDATED, FLEXIBLE_TESTX_REVIEW_COMPLETED, FLEXIBLE_TESTX_PLAN_CREATED, FLEXIBLE_TESTX_PLAN_TASK_ASSIGNED, FLEXIBLE_TESTX_PLAN_UPDATED, FLEXIBLE_TESTX_PLAN_FINISHED, FLEXIBLE_TESTX_REPORT_CREATED, CODE_DOG_CREATE_JOB, CODE_DOG_RESULT_NOTIFY, PLAN_CREATED, PLAN_DELETED, PLAN_COMMENT_CREATED, PLAN_STATUE_CHANGED, PLAN_ASSIGNEE_CHANGED, PLAN_UPDATED, RISK_CREATED, RISK_DELETED, RISK_COMMENT_CREATED, RISK_STATUS_CHANGED, RISK_ASSIGNEE_CHANGED, RISK_UPDATED | [optional] 
**Name** | Pointer to **string** | 备注名，支持模糊匹配 | [optional] 
**PageNumber** | **int64** | 分页页码 | 
**PageSize** | **int64** | 分页大小 | 
**ProjectId** | **int64** | 项目编号 | 
**Service** | Pointer to **[]string** | 服务名，取值范围：WebHook、WeCom、DingDing、Jenkins、FeiShu。 | [optional] 
**Status** | Pointer to **string** | 状态，取值范围：SUCCESS、FAILURE | [optional] 
**TargetType** | **string** | 目标数据类型：PROJECT,SPACE_NODE,PROGRAM,默认PROJECT | 

## Methods

### NewDescribeServiceHooksRequest

`func NewDescribeServiceHooksRequest(pageNumber int64, pageSize int64, projectId int64, targetType string, ) *DescribeServiceHooksRequest`

NewDescribeServiceHooksRequest instantiates a new DescribeServiceHooksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceHooksRequestWithDefaults

`func NewDescribeServiceHooksRequestWithDefaults() *DescribeServiceHooksRequest`

NewDescribeServiceHooksRequestWithDefaults instantiates a new DescribeServiceHooksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatorBy

`func (o *DescribeServiceHooksRequest) GetCreatorBy() []int64`

GetCreatorBy returns the CreatorBy field if non-nil, zero value otherwise.

### GetCreatorByOk

`func (o *DescribeServiceHooksRequest) GetCreatorByOk() (*[]int64, bool)`

GetCreatorByOk returns a tuple with the CreatorBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorBy

`func (o *DescribeServiceHooksRequest) SetCreatorBy(v []int64)`

SetCreatorBy sets CreatorBy field to given value.

### HasCreatorBy

`func (o *DescribeServiceHooksRequest) HasCreatorBy() bool`

HasCreatorBy returns a boolean if a field has been set.

### GetEnabled

`func (o *DescribeServiceHooksRequest) GetEnabled() string`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DescribeServiceHooksRequest) GetEnabledOk() (*string, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DescribeServiceHooksRequest) SetEnabled(v string)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DescribeServiceHooksRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEvent

`func (o *DescribeServiceHooksRequest) GetEvent() []string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DescribeServiceHooksRequest) GetEventOk() (*[]string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DescribeServiceHooksRequest) SetEvent(v []string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *DescribeServiceHooksRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetName

`func (o *DescribeServiceHooksRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DescribeServiceHooksRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DescribeServiceHooksRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DescribeServiceHooksRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeServiceHooksRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeServiceHooksRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeServiceHooksRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeServiceHooksRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeServiceHooksRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeServiceHooksRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetProjectId

`func (o *DescribeServiceHooksRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeServiceHooksRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeServiceHooksRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetService

`func (o *DescribeServiceHooksRequest) GetService() []string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DescribeServiceHooksRequest) GetServiceOk() (*[]string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DescribeServiceHooksRequest) SetService(v []string)`

SetService sets Service field to given value.

### HasService

`func (o *DescribeServiceHooksRequest) HasService() bool`

HasService returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeServiceHooksRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeServiceHooksRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeServiceHooksRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeServiceHooksRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetType

`func (o *DescribeServiceHooksRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *DescribeServiceHooksRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *DescribeServiceHooksRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


