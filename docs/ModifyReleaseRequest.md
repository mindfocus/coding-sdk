# ModifyReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **int64** | 修改之后的处理人ID | [optional] 
**DelIterationCodes** | Pointer to **[]int64** | 需要删除的版本已关联的迭代的列表 | [optional] 
**DelLabelIds** | Pointer to **[]int64** | 需要删除的标签的ID列表 | [optional] 
**Description** | Pointer to **string** | 修改之后的版本描述 | [optional] 
**IterationCodes** | Pointer to **[]int64** | 版本关联的迭代code列表 | [optional] 
**LabelIds** | Pointer to **[]int64** | 需要添加的标签的ID列表 | [optional] 
**Name** | Pointer to **string** | 修改之后的版本名 | [optional] 
**ProjectName** | **string** | 项目名称  | 
**ReleaseCode** | **int64** | 页面上版本ID | 
**ReleaseDate** | Pointer to **string** | 修改之后的版本发布日期 | [optional] 
**StartDate** | Pointer to **string** | 修改之后的版本开始日期 | [optional] 
**StatusId** | Pointer to **int64** | 修改之后的状态ID | [optional] 

## Methods

### NewModifyReleaseRequest

`func NewModifyReleaseRequest(projectName string, releaseCode int64, ) *ModifyReleaseRequest`

NewModifyReleaseRequest instantiates a new ModifyReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyReleaseRequestWithDefaults

`func NewModifyReleaseRequestWithDefaults() *ModifyReleaseRequest`

NewModifyReleaseRequestWithDefaults instantiates a new ModifyReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ModifyReleaseRequest) GetAssignee() int64`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ModifyReleaseRequest) GetAssigneeOk() (*int64, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ModifyReleaseRequest) SetAssignee(v int64)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ModifyReleaseRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDelIterationCodes

`func (o *ModifyReleaseRequest) GetDelIterationCodes() []int64`

GetDelIterationCodes returns the DelIterationCodes field if non-nil, zero value otherwise.

### GetDelIterationCodesOk

`func (o *ModifyReleaseRequest) GetDelIterationCodesOk() (*[]int64, bool)`

GetDelIterationCodesOk returns a tuple with the DelIterationCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelIterationCodes

`func (o *ModifyReleaseRequest) SetDelIterationCodes(v []int64)`

SetDelIterationCodes sets DelIterationCodes field to given value.

### HasDelIterationCodes

`func (o *ModifyReleaseRequest) HasDelIterationCodes() bool`

HasDelIterationCodes returns a boolean if a field has been set.

### GetDelLabelIds

`func (o *ModifyReleaseRequest) GetDelLabelIds() []int64`

GetDelLabelIds returns the DelLabelIds field if non-nil, zero value otherwise.

### GetDelLabelIdsOk

`func (o *ModifyReleaseRequest) GetDelLabelIdsOk() (*[]int64, bool)`

GetDelLabelIdsOk returns a tuple with the DelLabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelLabelIds

`func (o *ModifyReleaseRequest) SetDelLabelIds(v []int64)`

SetDelLabelIds sets DelLabelIds field to given value.

### HasDelLabelIds

`func (o *ModifyReleaseRequest) HasDelLabelIds() bool`

HasDelLabelIds returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyReleaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyReleaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyReleaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyReleaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIterationCodes

`func (o *ModifyReleaseRequest) GetIterationCodes() []int64`

GetIterationCodes returns the IterationCodes field if non-nil, zero value otherwise.

### GetIterationCodesOk

`func (o *ModifyReleaseRequest) GetIterationCodesOk() (*[]int64, bool)`

GetIterationCodesOk returns a tuple with the IterationCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCodes

`func (o *ModifyReleaseRequest) SetIterationCodes(v []int64)`

SetIterationCodes sets IterationCodes field to given value.

### HasIterationCodes

`func (o *ModifyReleaseRequest) HasIterationCodes() bool`

HasIterationCodes returns a boolean if a field has been set.

### GetLabelIds

`func (o *ModifyReleaseRequest) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *ModifyReleaseRequest) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *ModifyReleaseRequest) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *ModifyReleaseRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetName

`func (o *ModifyReleaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyReleaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyReleaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyReleaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyReleaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyReleaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyReleaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseCode

`func (o *ModifyReleaseRequest) GetReleaseCode() int64`

GetReleaseCode returns the ReleaseCode field if non-nil, zero value otherwise.

### GetReleaseCodeOk

`func (o *ModifyReleaseRequest) GetReleaseCodeOk() (*int64, bool)`

GetReleaseCodeOk returns a tuple with the ReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCode

`func (o *ModifyReleaseRequest) SetReleaseCode(v int64)`

SetReleaseCode sets ReleaseCode field to given value.


### GetReleaseDate

`func (o *ModifyReleaseRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ModifyReleaseRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ModifyReleaseRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ModifyReleaseRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetStartDate

`func (o *ModifyReleaseRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModifyReleaseRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModifyReleaseRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModifyReleaseRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatusId

`func (o *ModifyReleaseRequest) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ModifyReleaseRequest) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ModifyReleaseRequest) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *ModifyReleaseRequest) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


