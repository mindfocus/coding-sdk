# IssueFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentString** | Pointer to **string** | 筛选组合，JSON 字符串&lt;br /&gt;例如：&#x60;{\&quot;filterIssueType\&quot;:\&quot;ALL\&quot;,\&quot;sort\&quot;:{\&quot;key\&quot;:\&quot;PRIORITY\&quot;,\&quot;value\&quot;:\&quot;DESC\&quot;},\&quot;conditions\&quot;:[{\&quot;value\&quot;:[\&quot;TODO\&quot;,\&quot;PROCESSING\&quot;],\&quot;key\&quot;:\&quot;STATUS_TYPE\&quot;,\&quot;fixed\&quot;:true,\&quot;filterIssueType\&quot;:\&quot;ALL\&quot;,\&quot;projectId\&quot;:1},{\&quot;value\&quot;:[],\&quot;key\&quot;:\&quot;ASSIGNEE\&quot;,\&quot;fixed\&quot;:true,\&quot;constValue\&quot;:[\&quot;UNSPECIFIC\&quot;],\&quot;filterIssueType\&quot;:\&quot;ALL\&quot;,\&quot;projectId\&quot;:1}]}&#x60; | [optional] [default to ""]
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**CreatorId** | Pointer to **int64** | 创建人 ID | [optional] 
**Id** | Pointer to **int64** | 过滤器 ID | [optional] 
**IsDefault** | Pointer to **bool** | 默认筛选器 | [optional] [default to false]
**IsSystem** | Pointer to **bool** | 是否是系统自带 | [optional] [default to false]
**IssueType** | Pointer to **string** | 事项类型 | [optional] [default to ""]
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**SharedTeam** | Pointer to **bool** | 是否是团队筛选器 | [optional] [default to false]
**UpdatedAt** | Pointer to **int64** | 修改时间戳 | [optional] 

## Methods

### NewIssueFilter

`func NewIssueFilter() *IssueFilter`

NewIssueFilter instantiates a new IssueFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFilterWithDefaults

`func NewIssueFilterWithDefaults() *IssueFilter`

NewIssueFilterWithDefaults instantiates a new IssueFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentString

`func (o *IssueFilter) GetContentString() string`

GetContentString returns the ContentString field if non-nil, zero value otherwise.

### GetContentStringOk

`func (o *IssueFilter) GetContentStringOk() (*string, bool)`

GetContentStringOk returns a tuple with the ContentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentString

`func (o *IssueFilter) SetContentString(v string)`

SetContentString sets ContentString field to given value.

### HasContentString

`func (o *IssueFilter) HasContentString() bool`

HasContentString returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IssueFilter) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueFilter) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueFilter) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorId

`func (o *IssueFilter) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *IssueFilter) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *IssueFilter) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *IssueFilter) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetId

`func (o *IssueFilter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueFilter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueFilter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *IssueFilter) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IssueFilter) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IssueFilter) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IssueFilter) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsSystem

`func (o *IssueFilter) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *IssueFilter) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *IssueFilter) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *IssueFilter) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetIssueType

`func (o *IssueFilter) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *IssueFilter) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *IssueFilter) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *IssueFilter) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetName

`func (o *IssueFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSharedTeam

`func (o *IssueFilter) GetSharedTeam() bool`

GetSharedTeam returns the SharedTeam field if non-nil, zero value otherwise.

### GetSharedTeamOk

`func (o *IssueFilter) GetSharedTeamOk() (*bool, bool)`

GetSharedTeamOk returns a tuple with the SharedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedTeam

`func (o *IssueFilter) SetSharedTeam(v bool)`

SetSharedTeam sets SharedTeam field to given value.

### HasSharedTeam

`func (o *IssueFilter) HasSharedTeam() bool`

HasSharedTeam returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueFilter) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueFilter) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueFilter) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueFilter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


