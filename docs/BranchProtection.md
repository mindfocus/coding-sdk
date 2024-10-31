# BranchProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoAddReviewer** | Pointer to **bool** | 自动添加评审着 | [optional] [default to false]
**BranchProtectionId** | Pointer to **int64** | 保护分支规则id | [optional] 
**DenyForcePush** | Pointer to **bool** | 是否禁止强制推送 | [optional] [default to false]
**ForceSquash** | Pointer to **bool** | 是否使用squash推送 | [optional] [default to false]
**MatcherCount** | Pointer to **int64** | 匹配到的分支数量 | [optional] 
**RequiredCodeOwnerGrant** | Pointer to **bool** | 是否需要文件 owner 授权 | [optional] [default to false]
**RequiredDiscussionResolved** | Pointer to **bool** | 已解决必需的讨论 | [optional] [default to false]
**RequiredGrantCount** | Pointer to **int64** | 合并请求允许合并授权数量 | [optional] 
**RequiredStatusCheck** | Pointer to **bool** | 是否开启状态检查 | [optional] [default to false]
**Rule** | Pointer to **string** | 保护分支规则 | [optional] [default to ""]
**SrcMustMergedTo** | Pointer to **string** | 源分支必须已经合并到的某个分支,默认位空串 | [optional] [default to ""]

## Methods

### NewBranchProtection

`func NewBranchProtection() *BranchProtection`

NewBranchProtection instantiates a new BranchProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchProtectionWithDefaults

`func NewBranchProtectionWithDefaults() *BranchProtection`

NewBranchProtectionWithDefaults instantiates a new BranchProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoAddReviewer

`func (o *BranchProtection) GetAutoAddReviewer() bool`

GetAutoAddReviewer returns the AutoAddReviewer field if non-nil, zero value otherwise.

### GetAutoAddReviewerOk

`func (o *BranchProtection) GetAutoAddReviewerOk() (*bool, bool)`

GetAutoAddReviewerOk returns a tuple with the AutoAddReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddReviewer

`func (o *BranchProtection) SetAutoAddReviewer(v bool)`

SetAutoAddReviewer sets AutoAddReviewer field to given value.

### HasAutoAddReviewer

`func (o *BranchProtection) HasAutoAddReviewer() bool`

HasAutoAddReviewer returns a boolean if a field has been set.

### GetBranchProtectionId

`func (o *BranchProtection) GetBranchProtectionId() int64`

GetBranchProtectionId returns the BranchProtectionId field if non-nil, zero value otherwise.

### GetBranchProtectionIdOk

`func (o *BranchProtection) GetBranchProtectionIdOk() (*int64, bool)`

GetBranchProtectionIdOk returns a tuple with the BranchProtectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtectionId

`func (o *BranchProtection) SetBranchProtectionId(v int64)`

SetBranchProtectionId sets BranchProtectionId field to given value.

### HasBranchProtectionId

`func (o *BranchProtection) HasBranchProtectionId() bool`

HasBranchProtectionId returns a boolean if a field has been set.

### GetDenyForcePush

`func (o *BranchProtection) GetDenyForcePush() bool`

GetDenyForcePush returns the DenyForcePush field if non-nil, zero value otherwise.

### GetDenyForcePushOk

`func (o *BranchProtection) GetDenyForcePushOk() (*bool, bool)`

GetDenyForcePushOk returns a tuple with the DenyForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyForcePush

`func (o *BranchProtection) SetDenyForcePush(v bool)`

SetDenyForcePush sets DenyForcePush field to given value.

### HasDenyForcePush

`func (o *BranchProtection) HasDenyForcePush() bool`

HasDenyForcePush returns a boolean if a field has been set.

### GetForceSquash

`func (o *BranchProtection) GetForceSquash() bool`

GetForceSquash returns the ForceSquash field if non-nil, zero value otherwise.

### GetForceSquashOk

`func (o *BranchProtection) GetForceSquashOk() (*bool, bool)`

GetForceSquashOk returns a tuple with the ForceSquash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSquash

`func (o *BranchProtection) SetForceSquash(v bool)`

SetForceSquash sets ForceSquash field to given value.

### HasForceSquash

`func (o *BranchProtection) HasForceSquash() bool`

HasForceSquash returns a boolean if a field has been set.

### GetMatcherCount

`func (o *BranchProtection) GetMatcherCount() int64`

GetMatcherCount returns the MatcherCount field if non-nil, zero value otherwise.

### GetMatcherCountOk

`func (o *BranchProtection) GetMatcherCountOk() (*int64, bool)`

GetMatcherCountOk returns a tuple with the MatcherCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherCount

`func (o *BranchProtection) SetMatcherCount(v int64)`

SetMatcherCount sets MatcherCount field to given value.

### HasMatcherCount

`func (o *BranchProtection) HasMatcherCount() bool`

HasMatcherCount returns a boolean if a field has been set.

### GetRequiredCodeOwnerGrant

`func (o *BranchProtection) GetRequiredCodeOwnerGrant() bool`

GetRequiredCodeOwnerGrant returns the RequiredCodeOwnerGrant field if non-nil, zero value otherwise.

### GetRequiredCodeOwnerGrantOk

`func (o *BranchProtection) GetRequiredCodeOwnerGrantOk() (*bool, bool)`

GetRequiredCodeOwnerGrantOk returns a tuple with the RequiredCodeOwnerGrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredCodeOwnerGrant

`func (o *BranchProtection) SetRequiredCodeOwnerGrant(v bool)`

SetRequiredCodeOwnerGrant sets RequiredCodeOwnerGrant field to given value.

### HasRequiredCodeOwnerGrant

`func (o *BranchProtection) HasRequiredCodeOwnerGrant() bool`

HasRequiredCodeOwnerGrant returns a boolean if a field has been set.

### GetRequiredDiscussionResolved

`func (o *BranchProtection) GetRequiredDiscussionResolved() bool`

GetRequiredDiscussionResolved returns the RequiredDiscussionResolved field if non-nil, zero value otherwise.

### GetRequiredDiscussionResolvedOk

`func (o *BranchProtection) GetRequiredDiscussionResolvedOk() (*bool, bool)`

GetRequiredDiscussionResolvedOk returns a tuple with the RequiredDiscussionResolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDiscussionResolved

`func (o *BranchProtection) SetRequiredDiscussionResolved(v bool)`

SetRequiredDiscussionResolved sets RequiredDiscussionResolved field to given value.

### HasRequiredDiscussionResolved

`func (o *BranchProtection) HasRequiredDiscussionResolved() bool`

HasRequiredDiscussionResolved returns a boolean if a field has been set.

### GetRequiredGrantCount

`func (o *BranchProtection) GetRequiredGrantCount() int64`

GetRequiredGrantCount returns the RequiredGrantCount field if non-nil, zero value otherwise.

### GetRequiredGrantCountOk

`func (o *BranchProtection) GetRequiredGrantCountOk() (*int64, bool)`

GetRequiredGrantCountOk returns a tuple with the RequiredGrantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGrantCount

`func (o *BranchProtection) SetRequiredGrantCount(v int64)`

SetRequiredGrantCount sets RequiredGrantCount field to given value.

### HasRequiredGrantCount

`func (o *BranchProtection) HasRequiredGrantCount() bool`

HasRequiredGrantCount returns a boolean if a field has been set.

### GetRequiredStatusCheck

`func (o *BranchProtection) GetRequiredStatusCheck() bool`

GetRequiredStatusCheck returns the RequiredStatusCheck field if non-nil, zero value otherwise.

### GetRequiredStatusCheckOk

`func (o *BranchProtection) GetRequiredStatusCheckOk() (*bool, bool)`

GetRequiredStatusCheckOk returns a tuple with the RequiredStatusCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatusCheck

`func (o *BranchProtection) SetRequiredStatusCheck(v bool)`

SetRequiredStatusCheck sets RequiredStatusCheck field to given value.

### HasRequiredStatusCheck

`func (o *BranchProtection) HasRequiredStatusCheck() bool`

HasRequiredStatusCheck returns a boolean if a field has been set.

### GetRule

`func (o *BranchProtection) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *BranchProtection) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *BranchProtection) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *BranchProtection) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSrcMustMergedTo

`func (o *BranchProtection) GetSrcMustMergedTo() string`

GetSrcMustMergedTo returns the SrcMustMergedTo field if non-nil, zero value otherwise.

### GetSrcMustMergedToOk

`func (o *BranchProtection) GetSrcMustMergedToOk() (*string, bool)`

GetSrcMustMergedToOk returns a tuple with the SrcMustMergedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMustMergedTo

`func (o *BranchProtection) SetSrcMustMergedTo(v string)`

SetSrcMustMergedTo sets SrcMustMergedTo field to given value.

### HasSrcMustMergedTo

`func (o *BranchProtection) HasSrcMustMergedTo() bool`

HasSrcMustMergedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


