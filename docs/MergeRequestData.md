# MergeRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionAt** | Pointer to **NullableInt64** | 合并时间 | [optional] 
**ActionAuthor** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**Author** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**BaseSha** | Pointer to **NullableString** | 基础sha | [optional] [default to ""]
**CommentCount** | Pointer to **NullableInt64** | 评论数 | [optional] 
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**DepotId** | Pointer to **int64** | 仓库 Id | [optional] 
**Describe** | Pointer to **NullableString** | MR描述 | [optional] [default to ""]
**Granted** | Pointer to **NullableInt64** | 授权数 | [optional] 
**Id** | Pointer to **int64** | 合并请求 Id | [optional] 
**Labels** | Pointer to **[]string** | 标签列表 | [optional] 
**MergeCommitSha** | Pointer to **NullableString** | 合并Commit Sha | [optional] [default to ""]
**MergeId** | Pointer to **int64** | 合并请求 IId | [optional] 
**Path** | Pointer to **NullableString** | 合并请求 web 页面路径 | [optional] [default to ""]
**ProjectId** | Pointer to **int64** | 项目ID | [optional] 
**Reminded** | Pointer to **NullableBool** | 是否提醒 | [optional] [default to false]
**Reviewers** | Pointer to [**[]CodingUser**](CodingUser.md) | 评审者列表 | [optional] 
**SourceBranch** | Pointer to **NullableString** | 源分支 | [optional] [default to ""]
**SourceBranchSha** | Pointer to **NullableString** | 源分支Commit Sha | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 合并请求状态 | [optional] [default to ""]
**StickingPoint** | Pointer to **NullableString** | MR阻塞点 | [optional] [default to ""]
**TargetBranch** | Pointer to **NullableString** | 目标分支 | [optional] [default to ""]
**TargetBranchProtected** | Pointer to **NullableBool** | 目标分支是否保护分支 | [optional] [default to false]
**TargetBranchSha** | Pointer to **NullableString** | 目标分支Commit Sha | [optional] [default to ""]
**Title** | Pointer to **string** | 合并请求标题 | [optional] [default to ""]
**UpdateAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 

## Methods

### NewMergeRequestData

`func NewMergeRequestData() *MergeRequestData`

NewMergeRequestData instantiates a new MergeRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestDataWithDefaults

`func NewMergeRequestDataWithDefaults() *MergeRequestData`

NewMergeRequestDataWithDefaults instantiates a new MergeRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionAt

`func (o *MergeRequestData) GetActionAt() int64`

GetActionAt returns the ActionAt field if non-nil, zero value otherwise.

### GetActionAtOk

`func (o *MergeRequestData) GetActionAtOk() (*int64, bool)`

GetActionAtOk returns a tuple with the ActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAt

`func (o *MergeRequestData) SetActionAt(v int64)`

SetActionAt sets ActionAt field to given value.

### HasActionAt

`func (o *MergeRequestData) HasActionAt() bool`

HasActionAt returns a boolean if a field has been set.

### SetActionAtNil

`func (o *MergeRequestData) SetActionAtNil(b bool)`

 SetActionAtNil sets the value for ActionAt to be an explicit nil

### UnsetActionAt
`func (o *MergeRequestData) UnsetActionAt()`

UnsetActionAt ensures that no value is present for ActionAt, not even an explicit nil
### GetActionAuthor

`func (o *MergeRequestData) GetActionAuthor() CodingUser`

GetActionAuthor returns the ActionAuthor field if non-nil, zero value otherwise.

### GetActionAuthorOk

`func (o *MergeRequestData) GetActionAuthorOk() (*CodingUser, bool)`

GetActionAuthorOk returns a tuple with the ActionAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAuthor

`func (o *MergeRequestData) SetActionAuthor(v CodingUser)`

SetActionAuthor sets ActionAuthor field to given value.

### HasActionAuthor

`func (o *MergeRequestData) HasActionAuthor() bool`

HasActionAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *MergeRequestData) GetAuthor() CodingUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MergeRequestData) GetAuthorOk() (*CodingUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MergeRequestData) SetAuthor(v CodingUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MergeRequestData) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBaseSha

`func (o *MergeRequestData) GetBaseSha() string`

GetBaseSha returns the BaseSha field if non-nil, zero value otherwise.

### GetBaseShaOk

`func (o *MergeRequestData) GetBaseShaOk() (*string, bool)`

GetBaseShaOk returns a tuple with the BaseSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSha

`func (o *MergeRequestData) SetBaseSha(v string)`

SetBaseSha sets BaseSha field to given value.

### HasBaseSha

`func (o *MergeRequestData) HasBaseSha() bool`

HasBaseSha returns a boolean if a field has been set.

### SetBaseShaNil

`func (o *MergeRequestData) SetBaseShaNil(b bool)`

 SetBaseShaNil sets the value for BaseSha to be an explicit nil

### UnsetBaseSha
`func (o *MergeRequestData) UnsetBaseSha()`

UnsetBaseSha ensures that no value is present for BaseSha, not even an explicit nil
### GetCommentCount

`func (o *MergeRequestData) GetCommentCount() int64`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *MergeRequestData) GetCommentCountOk() (*int64, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *MergeRequestData) SetCommentCount(v int64)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *MergeRequestData) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### SetCommentCountNil

`func (o *MergeRequestData) SetCommentCountNil(b bool)`

 SetCommentCountNil sets the value for CommentCount to be an explicit nil

### UnsetCommentCount
`func (o *MergeRequestData) UnsetCommentCount()`

UnsetCommentCount ensures that no value is present for CommentCount, not even an explicit nil
### GetCreatedAt

`func (o *MergeRequestData) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MergeRequestData) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MergeRequestData) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MergeRequestData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *MergeRequestData) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *MergeRequestData) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDepotId

`func (o *MergeRequestData) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *MergeRequestData) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *MergeRequestData) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *MergeRequestData) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDescribe

`func (o *MergeRequestData) GetDescribe() string`

GetDescribe returns the Describe field if non-nil, zero value otherwise.

### GetDescribeOk

`func (o *MergeRequestData) GetDescribeOk() (*string, bool)`

GetDescribeOk returns a tuple with the Describe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribe

`func (o *MergeRequestData) SetDescribe(v string)`

SetDescribe sets Describe field to given value.

### HasDescribe

`func (o *MergeRequestData) HasDescribe() bool`

HasDescribe returns a boolean if a field has been set.

### SetDescribeNil

`func (o *MergeRequestData) SetDescribeNil(b bool)`

 SetDescribeNil sets the value for Describe to be an explicit nil

### UnsetDescribe
`func (o *MergeRequestData) UnsetDescribe()`

UnsetDescribe ensures that no value is present for Describe, not even an explicit nil
### GetGranted

`func (o *MergeRequestData) GetGranted() int64`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *MergeRequestData) GetGrantedOk() (*int64, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *MergeRequestData) SetGranted(v int64)`

SetGranted sets Granted field to given value.

### HasGranted

`func (o *MergeRequestData) HasGranted() bool`

HasGranted returns a boolean if a field has been set.

### SetGrantedNil

`func (o *MergeRequestData) SetGrantedNil(b bool)`

 SetGrantedNil sets the value for Granted to be an explicit nil

### UnsetGranted
`func (o *MergeRequestData) UnsetGranted()`

UnsetGranted ensures that no value is present for Granted, not even an explicit nil
### GetId

`func (o *MergeRequestData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergeRequestData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergeRequestData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MergeRequestData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *MergeRequestData) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MergeRequestData) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MergeRequestData) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MergeRequestData) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *MergeRequestData) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *MergeRequestData) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMergeCommitSha

`func (o *MergeRequestData) GetMergeCommitSha() string`

GetMergeCommitSha returns the MergeCommitSha field if non-nil, zero value otherwise.

### GetMergeCommitShaOk

`func (o *MergeRequestData) GetMergeCommitShaOk() (*string, bool)`

GetMergeCommitShaOk returns a tuple with the MergeCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitSha

`func (o *MergeRequestData) SetMergeCommitSha(v string)`

SetMergeCommitSha sets MergeCommitSha field to given value.

### HasMergeCommitSha

`func (o *MergeRequestData) HasMergeCommitSha() bool`

HasMergeCommitSha returns a boolean if a field has been set.

### SetMergeCommitShaNil

`func (o *MergeRequestData) SetMergeCommitShaNil(b bool)`

 SetMergeCommitShaNil sets the value for MergeCommitSha to be an explicit nil

### UnsetMergeCommitSha
`func (o *MergeRequestData) UnsetMergeCommitSha()`

UnsetMergeCommitSha ensures that no value is present for MergeCommitSha, not even an explicit nil
### GetMergeId

`func (o *MergeRequestData) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *MergeRequestData) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *MergeRequestData) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.

### HasMergeId

`func (o *MergeRequestData) HasMergeId() bool`

HasMergeId returns a boolean if a field has been set.

### GetPath

`func (o *MergeRequestData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MergeRequestData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MergeRequestData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MergeRequestData) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MergeRequestData) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MergeRequestData) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProjectId

`func (o *MergeRequestData) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MergeRequestData) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MergeRequestData) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MergeRequestData) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReminded

`func (o *MergeRequestData) GetReminded() bool`

GetReminded returns the Reminded field if non-nil, zero value otherwise.

### GetRemindedOk

`func (o *MergeRequestData) GetRemindedOk() (*bool, bool)`

GetRemindedOk returns a tuple with the Reminded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminded

`func (o *MergeRequestData) SetReminded(v bool)`

SetReminded sets Reminded field to given value.

### HasReminded

`func (o *MergeRequestData) HasReminded() bool`

HasReminded returns a boolean if a field has been set.

### SetRemindedNil

`func (o *MergeRequestData) SetRemindedNil(b bool)`

 SetRemindedNil sets the value for Reminded to be an explicit nil

### UnsetReminded
`func (o *MergeRequestData) UnsetReminded()`

UnsetReminded ensures that no value is present for Reminded, not even an explicit nil
### GetReviewers

`func (o *MergeRequestData) GetReviewers() []CodingUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *MergeRequestData) GetReviewersOk() (*[]CodingUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *MergeRequestData) SetReviewers(v []CodingUser)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *MergeRequestData) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### SetReviewersNil

`func (o *MergeRequestData) SetReviewersNil(b bool)`

 SetReviewersNil sets the value for Reviewers to be an explicit nil

### UnsetReviewers
`func (o *MergeRequestData) UnsetReviewers()`

UnsetReviewers ensures that no value is present for Reviewers, not even an explicit nil
### GetSourceBranch

`func (o *MergeRequestData) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *MergeRequestData) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *MergeRequestData) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *MergeRequestData) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### SetSourceBranchNil

`func (o *MergeRequestData) SetSourceBranchNil(b bool)`

 SetSourceBranchNil sets the value for SourceBranch to be an explicit nil

### UnsetSourceBranch
`func (o *MergeRequestData) UnsetSourceBranch()`

UnsetSourceBranch ensures that no value is present for SourceBranch, not even an explicit nil
### GetSourceBranchSha

`func (o *MergeRequestData) GetSourceBranchSha() string`

GetSourceBranchSha returns the SourceBranchSha field if non-nil, zero value otherwise.

### GetSourceBranchShaOk

`func (o *MergeRequestData) GetSourceBranchShaOk() (*string, bool)`

GetSourceBranchShaOk returns a tuple with the SourceBranchSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranchSha

`func (o *MergeRequestData) SetSourceBranchSha(v string)`

SetSourceBranchSha sets SourceBranchSha field to given value.

### HasSourceBranchSha

`func (o *MergeRequestData) HasSourceBranchSha() bool`

HasSourceBranchSha returns a boolean if a field has been set.

### SetSourceBranchShaNil

`func (o *MergeRequestData) SetSourceBranchShaNil(b bool)`

 SetSourceBranchShaNil sets the value for SourceBranchSha to be an explicit nil

### UnsetSourceBranchSha
`func (o *MergeRequestData) UnsetSourceBranchSha()`

UnsetSourceBranchSha ensures that no value is present for SourceBranchSha, not even an explicit nil
### GetStatus

`func (o *MergeRequestData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MergeRequestData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MergeRequestData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MergeRequestData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MergeRequestData) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MergeRequestData) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStickingPoint

`func (o *MergeRequestData) GetStickingPoint() string`

GetStickingPoint returns the StickingPoint field if non-nil, zero value otherwise.

### GetStickingPointOk

`func (o *MergeRequestData) GetStickingPointOk() (*string, bool)`

GetStickingPointOk returns a tuple with the StickingPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickingPoint

`func (o *MergeRequestData) SetStickingPoint(v string)`

SetStickingPoint sets StickingPoint field to given value.

### HasStickingPoint

`func (o *MergeRequestData) HasStickingPoint() bool`

HasStickingPoint returns a boolean if a field has been set.

### SetStickingPointNil

`func (o *MergeRequestData) SetStickingPointNil(b bool)`

 SetStickingPointNil sets the value for StickingPoint to be an explicit nil

### UnsetStickingPoint
`func (o *MergeRequestData) UnsetStickingPoint()`

UnsetStickingPoint ensures that no value is present for StickingPoint, not even an explicit nil
### GetTargetBranch

`func (o *MergeRequestData) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *MergeRequestData) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *MergeRequestData) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *MergeRequestData) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### SetTargetBranchNil

`func (o *MergeRequestData) SetTargetBranchNil(b bool)`

 SetTargetBranchNil sets the value for TargetBranch to be an explicit nil

### UnsetTargetBranch
`func (o *MergeRequestData) UnsetTargetBranch()`

UnsetTargetBranch ensures that no value is present for TargetBranch, not even an explicit nil
### GetTargetBranchProtected

`func (o *MergeRequestData) GetTargetBranchProtected() bool`

GetTargetBranchProtected returns the TargetBranchProtected field if non-nil, zero value otherwise.

### GetTargetBranchProtectedOk

`func (o *MergeRequestData) GetTargetBranchProtectedOk() (*bool, bool)`

GetTargetBranchProtectedOk returns a tuple with the TargetBranchProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchProtected

`func (o *MergeRequestData) SetTargetBranchProtected(v bool)`

SetTargetBranchProtected sets TargetBranchProtected field to given value.

### HasTargetBranchProtected

`func (o *MergeRequestData) HasTargetBranchProtected() bool`

HasTargetBranchProtected returns a boolean if a field has been set.

### SetTargetBranchProtectedNil

`func (o *MergeRequestData) SetTargetBranchProtectedNil(b bool)`

 SetTargetBranchProtectedNil sets the value for TargetBranchProtected to be an explicit nil

### UnsetTargetBranchProtected
`func (o *MergeRequestData) UnsetTargetBranchProtected()`

UnsetTargetBranchProtected ensures that no value is present for TargetBranchProtected, not even an explicit nil
### GetTargetBranchSha

`func (o *MergeRequestData) GetTargetBranchSha() string`

GetTargetBranchSha returns the TargetBranchSha field if non-nil, zero value otherwise.

### GetTargetBranchShaOk

`func (o *MergeRequestData) GetTargetBranchShaOk() (*string, bool)`

GetTargetBranchShaOk returns a tuple with the TargetBranchSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchSha

`func (o *MergeRequestData) SetTargetBranchSha(v string)`

SetTargetBranchSha sets TargetBranchSha field to given value.

### HasTargetBranchSha

`func (o *MergeRequestData) HasTargetBranchSha() bool`

HasTargetBranchSha returns a boolean if a field has been set.

### SetTargetBranchShaNil

`func (o *MergeRequestData) SetTargetBranchShaNil(b bool)`

 SetTargetBranchShaNil sets the value for TargetBranchSha to be an explicit nil

### UnsetTargetBranchSha
`func (o *MergeRequestData) UnsetTargetBranchSha()`

UnsetTargetBranchSha ensures that no value is present for TargetBranchSha, not even an explicit nil
### GetTitle

`func (o *MergeRequestData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MergeRequestData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MergeRequestData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MergeRequestData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdateAt

`func (o *MergeRequestData) GetUpdateAt() int64`

GetUpdateAt returns the UpdateAt field if non-nil, zero value otherwise.

### GetUpdateAtOk

`func (o *MergeRequestData) GetUpdateAtOk() (*int64, bool)`

GetUpdateAtOk returns a tuple with the UpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAt

`func (o *MergeRequestData) SetUpdateAt(v int64)`

SetUpdateAt sets UpdateAt field to given value.

### HasUpdateAt

`func (o *MergeRequestData) HasUpdateAt() bool`

HasUpdateAt returns a boolean if a field has been set.

### SetUpdateAtNil

`func (o *MergeRequestData) SetUpdateAtNil(b bool)`

 SetUpdateAtNil sets the value for UpdateAt to be an explicit nil

### UnsetUpdateAt
`func (o *MergeRequestData) UnsetUpdateAt()`

UnsetUpdateAt ensures that no value is present for UpdateAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


