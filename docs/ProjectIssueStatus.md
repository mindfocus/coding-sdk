# ProjectIssueStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**IsDefault** | Pointer to **bool** | 是否是默认值 | [optional] [default to false]
**IssueStatus** | Pointer to [**IssueStatus**](IssueStatus.md) |  | [optional] 
**IssueStatusId** | Pointer to **int64** | 事项状态 ID | [optional] 
**IssueType** | Pointer to **string** | 事项类型 | [optional] [default to ""]
**Sort** | Pointer to **int64** | 排序 | [optional] 
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 

## Methods

### NewProjectIssueStatus

`func NewProjectIssueStatus() *ProjectIssueStatus`

NewProjectIssueStatus instantiates a new ProjectIssueStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIssueStatusWithDefaults

`func NewProjectIssueStatusWithDefaults() *ProjectIssueStatus`

NewProjectIssueStatusWithDefaults instantiates a new ProjectIssueStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ProjectIssueStatus) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectIssueStatus) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectIssueStatus) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProjectIssueStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetIsDefault

`func (o *ProjectIssueStatus) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ProjectIssueStatus) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ProjectIssueStatus) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *ProjectIssueStatus) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIssueStatus

`func (o *ProjectIssueStatus) GetIssueStatus() IssueStatus`

GetIssueStatus returns the IssueStatus field if non-nil, zero value otherwise.

### GetIssueStatusOk

`func (o *ProjectIssueStatus) GetIssueStatusOk() (*IssueStatus, bool)`

GetIssueStatusOk returns a tuple with the IssueStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatus

`func (o *ProjectIssueStatus) SetIssueStatus(v IssueStatus)`

SetIssueStatus sets IssueStatus field to given value.

### HasIssueStatus

`func (o *ProjectIssueStatus) HasIssueStatus() bool`

HasIssueStatus returns a boolean if a field has been set.

### GetIssueStatusId

`func (o *ProjectIssueStatus) GetIssueStatusId() int64`

GetIssueStatusId returns the IssueStatusId field if non-nil, zero value otherwise.

### GetIssueStatusIdOk

`func (o *ProjectIssueStatus) GetIssueStatusIdOk() (*int64, bool)`

GetIssueStatusIdOk returns a tuple with the IssueStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatusId

`func (o *ProjectIssueStatus) SetIssueStatusId(v int64)`

SetIssueStatusId sets IssueStatusId field to given value.

### HasIssueStatusId

`func (o *ProjectIssueStatus) HasIssueStatusId() bool`

HasIssueStatusId returns a boolean if a field has been set.

### GetIssueType

`func (o *ProjectIssueStatus) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *ProjectIssueStatus) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *ProjectIssueStatus) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *ProjectIssueStatus) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetSort

`func (o *ProjectIssueStatus) GetSort() int64`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ProjectIssueStatus) GetSortOk() (*int64, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ProjectIssueStatus) SetSort(v int64)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ProjectIssueStatus) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProjectIssueStatus) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectIssueStatus) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectIssueStatus) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProjectIssueStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


