# ModifyMergeRequestBasicSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelMrGrantAfterPushSrc** | **bool** | 合并请求源分支有新提交时是否自动取消合并授权 | 
**DefaultFastForwardMerge** | **bool** | 是否默认以 Fast-Forward 模式合并 | 
**DefaultTargetBranch** | Pointer to **string** | 合并请求默认分支 | [optional] 
**DeleteSrcBranchAfterMerged** | **bool** | 是否默认删除源分支 | 
**DepotPath** | **string** | 仓库路径，格式：&lt;team&gt;/&lt;project&gt;/&lt;depot&gt; | 
**DepotStatusCheckRequired** | **bool** | 是否开启状态检查 | 
**MrCheckAllReviewersAllow** | **bool** | 合并前是否必须获得所有评审者的允许合并 | 
**SquashOnMerge** | Pointer to **string** | 合并请求选择方式no_squash:默认直接合并default_squash:默认Squash合并 force_squash:只能Squash合并 | [optional] 

## Methods

### NewModifyMergeRequestBasicSettingsRequest

`func NewModifyMergeRequestBasicSettingsRequest(cancelMrGrantAfterPushSrc bool, defaultFastForwardMerge bool, deleteSrcBranchAfterMerged bool, depotPath string, depotStatusCheckRequired bool, mrCheckAllReviewersAllow bool, ) *ModifyMergeRequestBasicSettingsRequest`

NewModifyMergeRequestBasicSettingsRequest instantiates a new ModifyMergeRequestBasicSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMergeRequestBasicSettingsRequestWithDefaults

`func NewModifyMergeRequestBasicSettingsRequestWithDefaults() *ModifyMergeRequestBasicSettingsRequest`

NewModifyMergeRequestBasicSettingsRequestWithDefaults instantiates a new ModifyMergeRequestBasicSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelMrGrantAfterPushSrc

`func (o *ModifyMergeRequestBasicSettingsRequest) GetCancelMrGrantAfterPushSrc() bool`

GetCancelMrGrantAfterPushSrc returns the CancelMrGrantAfterPushSrc field if non-nil, zero value otherwise.

### GetCancelMrGrantAfterPushSrcOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetCancelMrGrantAfterPushSrcOk() (*bool, bool)`

GetCancelMrGrantAfterPushSrcOk returns a tuple with the CancelMrGrantAfterPushSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelMrGrantAfterPushSrc

`func (o *ModifyMergeRequestBasicSettingsRequest) SetCancelMrGrantAfterPushSrc(v bool)`

SetCancelMrGrantAfterPushSrc sets CancelMrGrantAfterPushSrc field to given value.


### GetDefaultFastForwardMerge

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultFastForwardMerge() bool`

GetDefaultFastForwardMerge returns the DefaultFastForwardMerge field if non-nil, zero value otherwise.

### GetDefaultFastForwardMergeOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultFastForwardMergeOk() (*bool, bool)`

GetDefaultFastForwardMergeOk returns a tuple with the DefaultFastForwardMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFastForwardMerge

`func (o *ModifyMergeRequestBasicSettingsRequest) SetDefaultFastForwardMerge(v bool)`

SetDefaultFastForwardMerge sets DefaultFastForwardMerge field to given value.


### GetDefaultTargetBranch

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultTargetBranch() string`

GetDefaultTargetBranch returns the DefaultTargetBranch field if non-nil, zero value otherwise.

### GetDefaultTargetBranchOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDefaultTargetBranchOk() (*string, bool)`

GetDefaultTargetBranchOk returns a tuple with the DefaultTargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTargetBranch

`func (o *ModifyMergeRequestBasicSettingsRequest) SetDefaultTargetBranch(v string)`

SetDefaultTargetBranch sets DefaultTargetBranch field to given value.

### HasDefaultTargetBranch

`func (o *ModifyMergeRequestBasicSettingsRequest) HasDefaultTargetBranch() bool`

HasDefaultTargetBranch returns a boolean if a field has been set.

### GetDeleteSrcBranchAfterMerged

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDeleteSrcBranchAfterMerged() bool`

GetDeleteSrcBranchAfterMerged returns the DeleteSrcBranchAfterMerged field if non-nil, zero value otherwise.

### GetDeleteSrcBranchAfterMergedOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDeleteSrcBranchAfterMergedOk() (*bool, bool)`

GetDeleteSrcBranchAfterMergedOk returns a tuple with the DeleteSrcBranchAfterMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSrcBranchAfterMerged

`func (o *ModifyMergeRequestBasicSettingsRequest) SetDeleteSrcBranchAfterMerged(v bool)`

SetDeleteSrcBranchAfterMerged sets DeleteSrcBranchAfterMerged field to given value.


### GetDepotPath

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyMergeRequestBasicSettingsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetDepotStatusCheckRequired

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotStatusCheckRequired() bool`

GetDepotStatusCheckRequired returns the DepotStatusCheckRequired field if non-nil, zero value otherwise.

### GetDepotStatusCheckRequiredOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetDepotStatusCheckRequiredOk() (*bool, bool)`

GetDepotStatusCheckRequiredOk returns a tuple with the DepotStatusCheckRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotStatusCheckRequired

`func (o *ModifyMergeRequestBasicSettingsRequest) SetDepotStatusCheckRequired(v bool)`

SetDepotStatusCheckRequired sets DepotStatusCheckRequired field to given value.


### GetMrCheckAllReviewersAllow

`func (o *ModifyMergeRequestBasicSettingsRequest) GetMrCheckAllReviewersAllow() bool`

GetMrCheckAllReviewersAllow returns the MrCheckAllReviewersAllow field if non-nil, zero value otherwise.

### GetMrCheckAllReviewersAllowOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetMrCheckAllReviewersAllowOk() (*bool, bool)`

GetMrCheckAllReviewersAllowOk returns a tuple with the MrCheckAllReviewersAllow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrCheckAllReviewersAllow

`func (o *ModifyMergeRequestBasicSettingsRequest) SetMrCheckAllReviewersAllow(v bool)`

SetMrCheckAllReviewersAllow sets MrCheckAllReviewersAllow field to given value.


### GetSquashOnMerge

`func (o *ModifyMergeRequestBasicSettingsRequest) GetSquashOnMerge() string`

GetSquashOnMerge returns the SquashOnMerge field if non-nil, zero value otherwise.

### GetSquashOnMergeOk

`func (o *ModifyMergeRequestBasicSettingsRequest) GetSquashOnMergeOk() (*string, bool)`

GetSquashOnMergeOk returns a tuple with the SquashOnMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquashOnMerge

`func (o *ModifyMergeRequestBasicSettingsRequest) SetSquashOnMerge(v string)`

SetSquashOnMerge sets SquashOnMerge field to given value.

### HasSquashOnMerge

`func (o *ModifyMergeRequestBasicSettingsRequest) HasSquashOnMerge() bool`

HasSquashOnMerge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


