# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** | 内容 | [optional] [default to ""]
**CommitSha** | Pointer to **NullableString** | commit Sha 值 | [optional] [default to ""]
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**CreatorId** | Pointer to **NullableInt64** | 创建者 | [optional] 
**DepotId** | Pointer to **NullableInt64** | 仓库 Id | [optional] 
**Html** | Pointer to **NullableString** | html内容 | [optional] [default to ""]
**Id** | Pointer to **NullableInt64** | 版本 Id | [optional] 
**Pre** | Pointer to **NullableBool** | 是否预发布 | [optional] [default to false]
**ProjectId** | Pointer to **NullableInt64** | 项目 Id | [optional] 
**ReleaseId** | Pointer to **NullableInt64** | 版本序号Id | [optional] 
**TagName** | Pointer to **NullableString** | 标签名字 | [optional] [default to ""]
**TargetCommitBranch** | Pointer to **NullableString** | 目标 commit Sha 值 | [optional] [default to ""]
**Title** | Pointer to **NullableString** | 标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 

## Methods

### NewRelease

`func NewRelease() *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Release) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Release) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Release) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Release) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Release) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Release) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetCommitSha

`func (o *Release) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *Release) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *Release) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *Release) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### SetCommitShaNil

`func (o *Release) SetCommitShaNil(b bool)`

 SetCommitShaNil sets the value for CommitSha to be an explicit nil

### UnsetCommitSha
`func (o *Release) UnsetCommitSha()`

UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
### GetCreatedAt

`func (o *Release) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Release) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Release) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Release) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Release) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Release) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatorId

`func (o *Release) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Release) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Release) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Release) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *Release) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Release) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetDepotId

`func (o *Release) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *Release) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *Release) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *Release) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### SetDepotIdNil

`func (o *Release) SetDepotIdNil(b bool)`

 SetDepotIdNil sets the value for DepotId to be an explicit nil

### UnsetDepotId
`func (o *Release) UnsetDepotId()`

UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
### GetHtml

`func (o *Release) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *Release) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *Release) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *Release) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### SetHtmlNil

`func (o *Release) SetHtmlNil(b bool)`

 SetHtmlNil sets the value for Html to be an explicit nil

### UnsetHtml
`func (o *Release) UnsetHtml()`

UnsetHtml ensures that no value is present for Html, not even an explicit nil
### GetId

`func (o *Release) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Release) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Release) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Release) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Release) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Release) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPre

`func (o *Release) GetPre() bool`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *Release) GetPreOk() (*bool, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *Release) SetPre(v bool)`

SetPre sets Pre field to given value.

### HasPre

`func (o *Release) HasPre() bool`

HasPre returns a boolean if a field has been set.

### SetPreNil

`func (o *Release) SetPreNil(b bool)`

 SetPreNil sets the value for Pre to be an explicit nil

### UnsetPre
`func (o *Release) UnsetPre()`

UnsetPre ensures that no value is present for Pre, not even an explicit nil
### GetProjectId

`func (o *Release) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Release) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Release) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Release) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Release) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Release) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetReleaseId

`func (o *Release) GetReleaseId() int64`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *Release) GetReleaseIdOk() (*int64, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *Release) SetReleaseId(v int64)`

SetReleaseId sets ReleaseId field to given value.

### HasReleaseId

`func (o *Release) HasReleaseId() bool`

HasReleaseId returns a boolean if a field has been set.

### SetReleaseIdNil

`func (o *Release) SetReleaseIdNil(b bool)`

 SetReleaseIdNil sets the value for ReleaseId to be an explicit nil

### UnsetReleaseId
`func (o *Release) UnsetReleaseId()`

UnsetReleaseId ensures that no value is present for ReleaseId, not even an explicit nil
### GetTagName

`func (o *Release) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *Release) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *Release) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *Release) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### SetTagNameNil

`func (o *Release) SetTagNameNil(b bool)`

 SetTagNameNil sets the value for TagName to be an explicit nil

### UnsetTagName
`func (o *Release) UnsetTagName()`

UnsetTagName ensures that no value is present for TagName, not even an explicit nil
### GetTargetCommitBranch

`func (o *Release) GetTargetCommitBranch() string`

GetTargetCommitBranch returns the TargetCommitBranch field if non-nil, zero value otherwise.

### GetTargetCommitBranchOk

`func (o *Release) GetTargetCommitBranchOk() (*string, bool)`

GetTargetCommitBranchOk returns a tuple with the TargetCommitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitBranch

`func (o *Release) SetTargetCommitBranch(v string)`

SetTargetCommitBranch sets TargetCommitBranch field to given value.

### HasTargetCommitBranch

`func (o *Release) HasTargetCommitBranch() bool`

HasTargetCommitBranch returns a boolean if a field has been set.

### SetTargetCommitBranchNil

`func (o *Release) SetTargetCommitBranchNil(b bool)`

 SetTargetCommitBranchNil sets the value for TargetCommitBranch to be an explicit nil

### UnsetTargetCommitBranch
`func (o *Release) UnsetTargetCommitBranch()`

UnsetTargetCommitBranch ensures that no value is present for TargetCommitBranch, not even an explicit nil
### GetTitle

`func (o *Release) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Release) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Release) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Release) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Release) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Release) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUpdatedAt

`func (o *Release) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Release) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Release) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Release) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *Release) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Release) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


