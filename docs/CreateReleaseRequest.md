# CreateReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **int64** | 版本处理人ID | [optional] 
**Description** | Pointer to **string** | 版本描述 | [optional] 
**IterationCodes** | Pointer to **[]int64** | 关联迭代code列表 | [optional] 
**LabelIds** | Pointer to **[]int64** | 标签ID列表 | [optional] 
**Name** | **string** | 版本名称 | 
**ProjectName** | **string** | 项目名称 | 
**ReleaseDate** | Pointer to **string** | 版本发布日期 | [optional] 
**StartDate** | Pointer to **string** | 版本开始时间 | [optional] 

## Methods

### NewCreateReleaseRequest

`func NewCreateReleaseRequest(name string, projectName string, ) *CreateReleaseRequest`

NewCreateReleaseRequest instantiates a new CreateReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseRequestWithDefaults

`func NewCreateReleaseRequestWithDefaults() *CreateReleaseRequest`

NewCreateReleaseRequestWithDefaults instantiates a new CreateReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *CreateReleaseRequest) GetAssignee() int64`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *CreateReleaseRequest) GetAssigneeOk() (*int64, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *CreateReleaseRequest) SetAssignee(v int64)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *CreateReleaseRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDescription

`func (o *CreateReleaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReleaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReleaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateReleaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIterationCodes

`func (o *CreateReleaseRequest) GetIterationCodes() []int64`

GetIterationCodes returns the IterationCodes field if non-nil, zero value otherwise.

### GetIterationCodesOk

`func (o *CreateReleaseRequest) GetIterationCodesOk() (*[]int64, bool)`

GetIterationCodesOk returns a tuple with the IterationCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCodes

`func (o *CreateReleaseRequest) SetIterationCodes(v []int64)`

SetIterationCodes sets IterationCodes field to given value.

### HasIterationCodes

`func (o *CreateReleaseRequest) HasIterationCodes() bool`

HasIterationCodes returns a boolean if a field has been set.

### GetLabelIds

`func (o *CreateReleaseRequest) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreateReleaseRequest) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreateReleaseRequest) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreateReleaseRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetName

`func (o *CreateReleaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReleaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReleaseRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectName

`func (o *CreateReleaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateReleaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateReleaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseDate

`func (o *CreateReleaseRequest) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *CreateReleaseRequest) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *CreateReleaseRequest) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *CreateReleaseRequest) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateReleaseRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateReleaseRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateReleaseRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateReleaseRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


