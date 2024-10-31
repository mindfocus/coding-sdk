# MergeRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionAt** | Pointer to **int64** | 操作时间 | [optional] 
**ActionAuthor** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**Author** | Pointer to [**CodingUser**](CodingUser.md) |  | [optional] 
**BaseSha** | Pointer to **string** | 基础Sha | [optional] [default to ""]
**CommentCount** | Pointer to **int64** | 评论数 | [optional] 
**Conflicts** | Pointer to **[]string** | 冲突文件列表 | [optional] 
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**DepotId** | Pointer to **int64** | 仓库ID | [optional] 
**Describe** | Pointer to **string** | 描述,为 markdown 格式 | [optional] [default to ""]
**Granted** | Pointer to **int64** | 是否授权 | [optional] 
**Id** | Pointer to **int64** | 合并请求ID | [optional] 
**Labels** | Pointer to **[]string** | 合并请求标签列表 | [optional] 
**MergeCommitSha** | Pointer to **string** | 合并Commit Sha | [optional] [default to ""]
**MergeId** | Pointer to **int64** | 合并请求iid | [optional] 
**Mission** | Pointer to [**Mission**](Mission.md) |  | [optional] 
**Path** | Pointer to **string** | 合并请求 web 页面路径 | [optional] [default to ""]
**ProjectId** | Pointer to **int64** | 项目ID | [optional] 
**Reminded** | Pointer to **NullableBool** | 是否提醒 | [optional] [default to false]
**Reviewers** | Pointer to [**[]CodingUser**](CodingUser.md) | MR评审者 | [optional] 
**SourceBranch** | Pointer to **string** | 源分支 | [optional] [default to ""]
**SourceBranchSha** | Pointer to **string** | 源分支Commit Sha | [optional] [default to ""]
**Status** | Pointer to **string** | 合并状态,CANMERGE(状态可自动合并),ACCEPTED(状态已接受), CANNOTMERGE(状态不可自动合并), REFUSED(状态已拒绝(关闭)), CANCEL(取消), MERGING(正在合并中), ABNORMAL(状态异常) | [optional] [default to ""]
**StickingPoint** | Pointer to **string** | MR阻塞点 | [optional] [default to ""]
**TargetBranch** | Pointer to **string** | 目标分支 | [optional] [default to ""]
**TargetBranchProtected** | Pointer to **bool** | 目标分支是否为保护分支 | [optional] [default to false]
**TargetBranchSha** | Pointer to **string** | 目标分支Commit Sha | [optional] [default to ""]
**Title** | Pointer to **string** | 合并标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int64** | 更新时间 | [optional] 

## Methods

### NewMergeRequestInfo

`func NewMergeRequestInfo() *MergeRequestInfo`

NewMergeRequestInfo instantiates a new MergeRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestInfoWithDefaults

`func NewMergeRequestInfoWithDefaults() *MergeRequestInfo`

NewMergeRequestInfoWithDefaults instantiates a new MergeRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionAt

`func (o *MergeRequestInfo) GetActionAt() int64`

GetActionAt returns the ActionAt field if non-nil, zero value otherwise.

### GetActionAtOk

`func (o *MergeRequestInfo) GetActionAtOk() (*int64, bool)`

GetActionAtOk returns a tuple with the ActionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAt

`func (o *MergeRequestInfo) SetActionAt(v int64)`

SetActionAt sets ActionAt field to given value.

### HasActionAt

`func (o *MergeRequestInfo) HasActionAt() bool`

HasActionAt returns a boolean if a field has been set.

### GetActionAuthor

`func (o *MergeRequestInfo) GetActionAuthor() CodingUser`

GetActionAuthor returns the ActionAuthor field if non-nil, zero value otherwise.

### GetActionAuthorOk

`func (o *MergeRequestInfo) GetActionAuthorOk() (*CodingUser, bool)`

GetActionAuthorOk returns a tuple with the ActionAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionAuthor

`func (o *MergeRequestInfo) SetActionAuthor(v CodingUser)`

SetActionAuthor sets ActionAuthor field to given value.

### HasActionAuthor

`func (o *MergeRequestInfo) HasActionAuthor() bool`

HasActionAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *MergeRequestInfo) GetAuthor() CodingUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MergeRequestInfo) GetAuthorOk() (*CodingUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MergeRequestInfo) SetAuthor(v CodingUser)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MergeRequestInfo) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBaseSha

`func (o *MergeRequestInfo) GetBaseSha() string`

GetBaseSha returns the BaseSha field if non-nil, zero value otherwise.

### GetBaseShaOk

`func (o *MergeRequestInfo) GetBaseShaOk() (*string, bool)`

GetBaseShaOk returns a tuple with the BaseSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSha

`func (o *MergeRequestInfo) SetBaseSha(v string)`

SetBaseSha sets BaseSha field to given value.

### HasBaseSha

`func (o *MergeRequestInfo) HasBaseSha() bool`

HasBaseSha returns a boolean if a field has been set.

### GetCommentCount

`func (o *MergeRequestInfo) GetCommentCount() int64`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *MergeRequestInfo) GetCommentCountOk() (*int64, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *MergeRequestInfo) SetCommentCount(v int64)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *MergeRequestInfo) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetConflicts

`func (o *MergeRequestInfo) GetConflicts() []string`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *MergeRequestInfo) GetConflictsOk() (*[]string, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *MergeRequestInfo) SetConflicts(v []string)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *MergeRequestInfo) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MergeRequestInfo) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MergeRequestInfo) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MergeRequestInfo) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MergeRequestInfo) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepotId

`func (o *MergeRequestInfo) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *MergeRequestInfo) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *MergeRequestInfo) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *MergeRequestInfo) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDescribe

`func (o *MergeRequestInfo) GetDescribe() string`

GetDescribe returns the Describe field if non-nil, zero value otherwise.

### GetDescribeOk

`func (o *MergeRequestInfo) GetDescribeOk() (*string, bool)`

GetDescribeOk returns a tuple with the Describe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribe

`func (o *MergeRequestInfo) SetDescribe(v string)`

SetDescribe sets Describe field to given value.

### HasDescribe

`func (o *MergeRequestInfo) HasDescribe() bool`

HasDescribe returns a boolean if a field has been set.

### GetGranted

`func (o *MergeRequestInfo) GetGranted() int64`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *MergeRequestInfo) GetGrantedOk() (*int64, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *MergeRequestInfo) SetGranted(v int64)`

SetGranted sets Granted field to given value.

### HasGranted

`func (o *MergeRequestInfo) HasGranted() bool`

HasGranted returns a boolean if a field has been set.

### GetId

`func (o *MergeRequestInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergeRequestInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergeRequestInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MergeRequestInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *MergeRequestInfo) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *MergeRequestInfo) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *MergeRequestInfo) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *MergeRequestInfo) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *MergeRequestInfo) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *MergeRequestInfo) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMergeCommitSha

`func (o *MergeRequestInfo) GetMergeCommitSha() string`

GetMergeCommitSha returns the MergeCommitSha field if non-nil, zero value otherwise.

### GetMergeCommitShaOk

`func (o *MergeRequestInfo) GetMergeCommitShaOk() (*string, bool)`

GetMergeCommitShaOk returns a tuple with the MergeCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeCommitSha

`func (o *MergeRequestInfo) SetMergeCommitSha(v string)`

SetMergeCommitSha sets MergeCommitSha field to given value.

### HasMergeCommitSha

`func (o *MergeRequestInfo) HasMergeCommitSha() bool`

HasMergeCommitSha returns a boolean if a field has been set.

### GetMergeId

`func (o *MergeRequestInfo) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *MergeRequestInfo) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *MergeRequestInfo) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.

### HasMergeId

`func (o *MergeRequestInfo) HasMergeId() bool`

HasMergeId returns a boolean if a field has been set.

### GetMission

`func (o *MergeRequestInfo) GetMission() Mission`

GetMission returns the Mission field if non-nil, zero value otherwise.

### GetMissionOk

`func (o *MergeRequestInfo) GetMissionOk() (*Mission, bool)`

GetMissionOk returns a tuple with the Mission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMission

`func (o *MergeRequestInfo) SetMission(v Mission)`

SetMission sets Mission field to given value.

### HasMission

`func (o *MergeRequestInfo) HasMission() bool`

HasMission returns a boolean if a field has been set.

### GetPath

`func (o *MergeRequestInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MergeRequestInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MergeRequestInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MergeRequestInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProjectId

`func (o *MergeRequestInfo) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MergeRequestInfo) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MergeRequestInfo) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *MergeRequestInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetReminded

`func (o *MergeRequestInfo) GetReminded() bool`

GetReminded returns the Reminded field if non-nil, zero value otherwise.

### GetRemindedOk

`func (o *MergeRequestInfo) GetRemindedOk() (*bool, bool)`

GetRemindedOk returns a tuple with the Reminded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminded

`func (o *MergeRequestInfo) SetReminded(v bool)`

SetReminded sets Reminded field to given value.

### HasReminded

`func (o *MergeRequestInfo) HasReminded() bool`

HasReminded returns a boolean if a field has been set.

### SetRemindedNil

`func (o *MergeRequestInfo) SetRemindedNil(b bool)`

 SetRemindedNil sets the value for Reminded to be an explicit nil

### UnsetReminded
`func (o *MergeRequestInfo) UnsetReminded()`

UnsetReminded ensures that no value is present for Reminded, not even an explicit nil
### GetReviewers

`func (o *MergeRequestInfo) GetReviewers() []CodingUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *MergeRequestInfo) GetReviewersOk() (*[]CodingUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *MergeRequestInfo) SetReviewers(v []CodingUser)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *MergeRequestInfo) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### SetReviewersNil

`func (o *MergeRequestInfo) SetReviewersNil(b bool)`

 SetReviewersNil sets the value for Reviewers to be an explicit nil

### UnsetReviewers
`func (o *MergeRequestInfo) UnsetReviewers()`

UnsetReviewers ensures that no value is present for Reviewers, not even an explicit nil
### GetSourceBranch

`func (o *MergeRequestInfo) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *MergeRequestInfo) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *MergeRequestInfo) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *MergeRequestInfo) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### GetSourceBranchSha

`func (o *MergeRequestInfo) GetSourceBranchSha() string`

GetSourceBranchSha returns the SourceBranchSha field if non-nil, zero value otherwise.

### GetSourceBranchShaOk

`func (o *MergeRequestInfo) GetSourceBranchShaOk() (*string, bool)`

GetSourceBranchShaOk returns a tuple with the SourceBranchSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranchSha

`func (o *MergeRequestInfo) SetSourceBranchSha(v string)`

SetSourceBranchSha sets SourceBranchSha field to given value.

### HasSourceBranchSha

`func (o *MergeRequestInfo) HasSourceBranchSha() bool`

HasSourceBranchSha returns a boolean if a field has been set.

### GetStatus

`func (o *MergeRequestInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MergeRequestInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MergeRequestInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MergeRequestInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStickingPoint

`func (o *MergeRequestInfo) GetStickingPoint() string`

GetStickingPoint returns the StickingPoint field if non-nil, zero value otherwise.

### GetStickingPointOk

`func (o *MergeRequestInfo) GetStickingPointOk() (*string, bool)`

GetStickingPointOk returns a tuple with the StickingPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickingPoint

`func (o *MergeRequestInfo) SetStickingPoint(v string)`

SetStickingPoint sets StickingPoint field to given value.

### HasStickingPoint

`func (o *MergeRequestInfo) HasStickingPoint() bool`

HasStickingPoint returns a boolean if a field has been set.

### GetTargetBranch

`func (o *MergeRequestInfo) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *MergeRequestInfo) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *MergeRequestInfo) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *MergeRequestInfo) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### GetTargetBranchProtected

`func (o *MergeRequestInfo) GetTargetBranchProtected() bool`

GetTargetBranchProtected returns the TargetBranchProtected field if non-nil, zero value otherwise.

### GetTargetBranchProtectedOk

`func (o *MergeRequestInfo) GetTargetBranchProtectedOk() (*bool, bool)`

GetTargetBranchProtectedOk returns a tuple with the TargetBranchProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchProtected

`func (o *MergeRequestInfo) SetTargetBranchProtected(v bool)`

SetTargetBranchProtected sets TargetBranchProtected field to given value.

### HasTargetBranchProtected

`func (o *MergeRequestInfo) HasTargetBranchProtected() bool`

HasTargetBranchProtected returns a boolean if a field has been set.

### GetTargetBranchSha

`func (o *MergeRequestInfo) GetTargetBranchSha() string`

GetTargetBranchSha returns the TargetBranchSha field if non-nil, zero value otherwise.

### GetTargetBranchShaOk

`func (o *MergeRequestInfo) GetTargetBranchShaOk() (*string, bool)`

GetTargetBranchShaOk returns a tuple with the TargetBranchSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranchSha

`func (o *MergeRequestInfo) SetTargetBranchSha(v string)`

SetTargetBranchSha sets TargetBranchSha field to given value.

### HasTargetBranchSha

`func (o *MergeRequestInfo) HasTargetBranchSha() bool`

HasTargetBranchSha returns a boolean if a field has been set.

### GetTitle

`func (o *MergeRequestInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MergeRequestInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MergeRequestInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MergeRequestInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MergeRequestInfo) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MergeRequestInfo) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MergeRequestInfo) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MergeRequestInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


