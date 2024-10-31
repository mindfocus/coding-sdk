# MergeRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionAt** | Pointer to **NullableInt64** | 是否允许合并时间 | [optional] 
**ActionAuthor** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**Author** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**Content** | Pointer to **NullableString** | 合并请求描述 | [optional] [default to ""]
**ContentHtml** | Pointer to **NullableString** | 合并请求描述（html 格式） | [optional] [default to ""]
**CreateAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**DepotId** | Pointer to **NullableInt64** | 仓库 Id | [optional] 
**Id** | Pointer to **NullableInt64** | 合并请求 Id | [optional] 
**MergeId** | Pointer to **NullableInt64** | 合并请求 Iid | [optional] 
**MergeStatus** | Pointer to **NullableString** | 合并请求状态 | [optional] [default to ""]
**MergedSha** | Pointer to **NullableString** | 合并 Sha | [optional] [default to ""]
**SourceBranch** | Pointer to **NullableString** | 源分支名 | [optional] [default to ""]
**SourceSha** | Pointer to **NullableString** | 源分支 Sha | [optional] [default to ""]
**TargetBranch** | Pointer to **NullableString** | 目标分支名 | [optional] [default to ""]
**TargetSha** | Pointer to **NullableString** | 目标分支 Sha | [optional] [default to ""]
**Title** | Pointer to **NullableString** | 合并请求标题 | [optional] [default to ""]
**UpdateAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 

## Methods

### NewMergeRequestDetail

`func NewMergeRequestDetail() *MergeRequestDetail`

NewMergeRequestDetail instantiates a new MergeRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestDetailWithDefaults

`func NewMergeRequestDetailWithDefaults() *MergeRequestDetail`

NewMergeRequestDetailWithDefaults instantiates a new MergeRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionAt

`func (o *MergeRequestDetail) GetActionAt() int64`

GetActionAt returns the ActionAt field if non-nil, zero value otherwise.

### GetActionAtOk

`func (o *MergeRequestDetail) GetActionAtOk() (*int64, bool)`

GetActionAtOk returns a tuple with the ActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAt

`func (o *MergeRequestDetail) SetActionAt(v int64)`

SetActionAt sets ActionAt field to given value.

### HasActionAt

`func (o *MergeRequestDetail) HasActionAt() bool`

HasActionAt returns a boolean if a field has been set.

### SetActionAtNil

`func (o *MergeRequestDetail) SetActionAtNil(b bool)`

 SetActionAtNil sets the value for ActionAt to be an explicit nil

### UnsetActionAt
`func (o *MergeRequestDetail) UnsetActionAt()`

UnsetActionAt ensures that no value is present for ActionAt, not even an explicit nil
### GetActionAuthor

`func (o *MergeRequestDetail) GetActionAuthor() CodingUser`

GetActionAuthor returns the ActionAuthor field if non-nil, zero value otherwise.

### GetActionAuthorOk

`func (o *MergeRequestDetail) GetActionAuthorOk() (*CodingUser, bool)`

GetActionAuthorOk returns a tuple with the ActionAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAuthor

`func (o *MergeRequestDetail) SetActionAuthor(v CodingUser)`

SetActionAuthor sets ActionAuthor field to given value.

### HasActionAuthor

`func (o *MergeRequestDetail) HasActionAuthor() bool`

HasActionAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *MergeRequestDetail) GetAuthor() CodingUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MergeRequestDetail) GetAuthorOk() (*CodingUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MergeRequestDetail) SetAuthor(v CodingUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MergeRequestDetail) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetContent

`func (o *MergeRequestDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MergeRequestDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MergeRequestDetail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MergeRequestDetail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MergeRequestDetail) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MergeRequestDetail) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentHtml

`func (o *MergeRequestDetail) GetContentHtml() string`

GetContentHtml returns the ContentHtml field if non-nil, zero value otherwise.

### GetContentHtmlOk

`func (o *MergeRequestDetail) GetContentHtmlOk() (*string, bool)`

GetContentHtmlOk returns a tuple with the ContentHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHtml

`func (o *MergeRequestDetail) SetContentHtml(v string)`

SetContentHtml sets ContentHtml field to given value.

### HasContentHtml

`func (o *MergeRequestDetail) HasContentHtml() bool`

HasContentHtml returns a boolean if a field has been set.

### SetContentHtmlNil

`func (o *MergeRequestDetail) SetContentHtmlNil(b bool)`

 SetContentHtmlNil sets the value for ContentHtml to be an explicit nil

### UnsetContentHtml
`func (o *MergeRequestDetail) UnsetContentHtml()`

UnsetContentHtml ensures that no value is present for ContentHtml, not even an explicit nil
### GetCreateAt

`func (o *MergeRequestDetail) GetCreateAt() int64`

GetCreateAt returns the CreateAt field if non-nil, zero value otherwise.

### GetCreateAtOk

`func (o *MergeRequestDetail) GetCreateAtOk() (*int64, bool)`

GetCreateAtOk returns a tuple with the CreateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAt

`func (o *MergeRequestDetail) SetCreateAt(v int64)`

SetCreateAt sets CreateAt field to given value.

### HasCreateAt

`func (o *MergeRequestDetail) HasCreateAt() bool`

HasCreateAt returns a boolean if a field has been set.

### SetCreateAtNil

`func (o *MergeRequestDetail) SetCreateAtNil(b bool)`

 SetCreateAtNil sets the value for CreateAt to be an explicit nil

### UnsetCreateAt
`func (o *MergeRequestDetail) UnsetCreateAt()`

UnsetCreateAt ensures that no value is present for CreateAt, not even an explicit nil
### GetDepotId

`func (o *MergeRequestDetail) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *MergeRequestDetail) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *MergeRequestDetail) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *MergeRequestDetail) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### SetDepotIdNil

`func (o *MergeRequestDetail) SetDepotIdNil(b bool)`

 SetDepotIdNil sets the value for DepotId to be an explicit nil

### UnsetDepotId
`func (o *MergeRequestDetail) UnsetDepotId()`

UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
### GetId

`func (o *MergeRequestDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergeRequestDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergeRequestDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MergeRequestDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MergeRequestDetail) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MergeRequestDetail) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMergeId

`func (o *MergeRequestDetail) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *MergeRequestDetail) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *MergeRequestDetail) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.

### HasMergeId

`func (o *MergeRequestDetail) HasMergeId() bool`

HasMergeId returns a boolean if a field has been set.

### SetMergeIdNil

`func (o *MergeRequestDetail) SetMergeIdNil(b bool)`

 SetMergeIdNil sets the value for MergeId to be an explicit nil

### UnsetMergeId
`func (o *MergeRequestDetail) UnsetMergeId()`

UnsetMergeId ensures that no value is present for MergeId, not even an explicit nil
### GetMergeStatus

`func (o *MergeRequestDetail) GetMergeStatus() string`

GetMergeStatus returns the MergeStatus field if non-nil, zero value otherwise.

### GetMergeStatusOk

`func (o *MergeRequestDetail) GetMergeStatusOk() (*string, bool)`

GetMergeStatusOk returns a tuple with the MergeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeStatus

`func (o *MergeRequestDetail) SetMergeStatus(v string)`

SetMergeStatus sets MergeStatus field to given value.

### HasMergeStatus

`func (o *MergeRequestDetail) HasMergeStatus() bool`

HasMergeStatus returns a boolean if a field has been set.

### SetMergeStatusNil

`func (o *MergeRequestDetail) SetMergeStatusNil(b bool)`

 SetMergeStatusNil sets the value for MergeStatus to be an explicit nil

### UnsetMergeStatus
`func (o *MergeRequestDetail) UnsetMergeStatus()`

UnsetMergeStatus ensures that no value is present for MergeStatus, not even an explicit nil
### GetMergedSha

`func (o *MergeRequestDetail) GetMergedSha() string`

GetMergedSha returns the MergedSha field if non-nil, zero value otherwise.

### GetMergedShaOk

`func (o *MergeRequestDetail) GetMergedShaOk() (*string, bool)`

GetMergedShaOk returns a tuple with the MergedSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedSha

`func (o *MergeRequestDetail) SetMergedSha(v string)`

SetMergedSha sets MergedSha field to given value.

### HasMergedSha

`func (o *MergeRequestDetail) HasMergedSha() bool`

HasMergedSha returns a boolean if a field has been set.

### SetMergedShaNil

`func (o *MergeRequestDetail) SetMergedShaNil(b bool)`

 SetMergedShaNil sets the value for MergedSha to be an explicit nil

### UnsetMergedSha
`func (o *MergeRequestDetail) UnsetMergedSha()`

UnsetMergedSha ensures that no value is present for MergedSha, not even an explicit nil
### GetSourceBranch

`func (o *MergeRequestDetail) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *MergeRequestDetail) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *MergeRequestDetail) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *MergeRequestDetail) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### SetSourceBranchNil

`func (o *MergeRequestDetail) SetSourceBranchNil(b bool)`

 SetSourceBranchNil sets the value for SourceBranch to be an explicit nil

### UnsetSourceBranch
`func (o *MergeRequestDetail) UnsetSourceBranch()`

UnsetSourceBranch ensures that no value is present for SourceBranch, not even an explicit nil
### GetSourceSha

`func (o *MergeRequestDetail) GetSourceSha() string`

GetSourceSha returns the SourceSha field if non-nil, zero value otherwise.

### GetSourceShaOk

`func (o *MergeRequestDetail) GetSourceShaOk() (*string, bool)`

GetSourceShaOk returns a tuple with the SourceSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSha

`func (o *MergeRequestDetail) SetSourceSha(v string)`

SetSourceSha sets SourceSha field to given value.

### HasSourceSha

`func (o *MergeRequestDetail) HasSourceSha() bool`

HasSourceSha returns a boolean if a field has been set.

### SetSourceShaNil

`func (o *MergeRequestDetail) SetSourceShaNil(b bool)`

 SetSourceShaNil sets the value for SourceSha to be an explicit nil

### UnsetSourceSha
`func (o *MergeRequestDetail) UnsetSourceSha()`

UnsetSourceSha ensures that no value is present for SourceSha, not even an explicit nil
### GetTargetBranch

`func (o *MergeRequestDetail) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *MergeRequestDetail) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *MergeRequestDetail) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *MergeRequestDetail) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### SetTargetBranchNil

`func (o *MergeRequestDetail) SetTargetBranchNil(b bool)`

 SetTargetBranchNil sets the value for TargetBranch to be an explicit nil

### UnsetTargetBranch
`func (o *MergeRequestDetail) UnsetTargetBranch()`

UnsetTargetBranch ensures that no value is present for TargetBranch, not even an explicit nil
### GetTargetSha

`func (o *MergeRequestDetail) GetTargetSha() string`

GetTargetSha returns the TargetSha field if non-nil, zero value otherwise.

### GetTargetShaOk

`func (o *MergeRequestDetail) GetTargetShaOk() (*string, bool)`

GetTargetShaOk returns a tuple with the TargetSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSha

`func (o *MergeRequestDetail) SetTargetSha(v string)`

SetTargetSha sets TargetSha field to given value.

### HasTargetSha

`func (o *MergeRequestDetail) HasTargetSha() bool`

HasTargetSha returns a boolean if a field has been set.

### SetTargetShaNil

`func (o *MergeRequestDetail) SetTargetShaNil(b bool)`

 SetTargetShaNil sets the value for TargetSha to be an explicit nil

### UnsetTargetSha
`func (o *MergeRequestDetail) UnsetTargetSha()`

UnsetTargetSha ensures that no value is present for TargetSha, not even an explicit nil
### GetTitle

`func (o *MergeRequestDetail) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MergeRequestDetail) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MergeRequestDetail) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MergeRequestDetail) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MergeRequestDetail) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MergeRequestDetail) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdateAt

`func (o *MergeRequestDetail) GetUpdateAt() int64`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *MergeRequestDetail) GetUpdateAtOk() (*int64, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *MergeRequestDetail) SetUpdateAt(v int64)`

SetUpdateAt sets UpdateAt field to given value.

### HasUpdateAt

`func (o *MergeRequestDetail) HasUpdateAt() bool`

HasUpdateAt returns a boolean if a field has been set.

### SetUpdateAtNil

`func (o *MergeRequestDetail) SetUpdateAtNil(b bool)`

 SetUpdateAtNil sets the value for UpdateAt to be an explicit nil

### UnsetUpdateAt
`func (o *MergeRequestDetail) UnsetUpdateAt()`

UnsetUpdateAt ensures that no value is present for UpdateAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


