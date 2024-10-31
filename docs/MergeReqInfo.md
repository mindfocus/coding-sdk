# MergeReqInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Describe** | Pointer to **string** | 描述,为 markdown 格式 | [optional] [default to ""]
**SourceBranch** | Pointer to **string** | 源分支 | [optional] [default to ""]
**Status** | Pointer to **string** | 合并状态,状态值如下:  CANMERGE:状态可自动合并;  ACCEPTED:状态已接受;  CANNOTMERGE:状态不可自动合并;  REFUSED:状态已拒绝(关闭);  CANCEL: 取消;  MERGING:正在合并中;  ABNORMAL:状态异常; | [optional] [default to ""]
**TargetBranch** | Pointer to **string** | 目标分支 | [optional] [default to ""]
**Title** | Pointer to **string** | 合并标题 | [optional] [default to ""]

## Methods

### NewMergeReqInfo

`func NewMergeReqInfo() *MergeReqInfo`

NewMergeReqInfo instantiates a new MergeReqInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeReqInfoWithDefaults

`func NewMergeReqInfoWithDefaults() *MergeReqInfo`

NewMergeReqInfoWithDefaults instantiates a new MergeReqInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescribe

`func (o *MergeReqInfo) GetDescribe() string`

GetDescribe returns the Describe field if non-nil, zero value otherwise.

### GetDescribeOk

`func (o *MergeReqInfo) GetDescribeOk() (*string, bool)`

GetDescribeOk returns a tuple with the Describe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribe

`func (o *MergeReqInfo) SetDescribe(v string)`

SetDescribe sets Describe field to given value.

### HasDescribe

`func (o *MergeReqInfo) HasDescribe() bool`

HasDescribe returns a boolean if a field has been set.

### GetSourceBranch

`func (o *MergeReqInfo) GetSourceBranch() string`

GetSourceBranch returns the SourceBranch field if non-nil, zero value otherwise.

### GetSourceBranchOk

`func (o *MergeReqInfo) GetSourceBranchOk() (*string, bool)`

GetSourceBranchOk returns a tuple with the SourceBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBranch

`func (o *MergeReqInfo) SetSourceBranch(v string)`

SetSourceBranch sets SourceBranch field to given value.

### HasSourceBranch

`func (o *MergeReqInfo) HasSourceBranch() bool`

HasSourceBranch returns a boolean if a field has been set.

### GetStatus

`func (o *MergeReqInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MergeReqInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MergeReqInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MergeReqInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetBranch

`func (o *MergeReqInfo) GetTargetBranch() string`

GetTargetBranch returns the TargetBranch field if non-nil, zero value otherwise.

### GetTargetBranchOk

`func (o *MergeReqInfo) GetTargetBranchOk() (*string, bool)`

GetTargetBranchOk returns a tuple with the TargetBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetBranch

`func (o *MergeReqInfo) SetTargetBranch(v string)`

SetTargetBranch sets TargetBranch field to given value.

### HasTargetBranch

`func (o *MergeReqInfo) HasTargetBranch() bool`

HasTargetBranch returns a boolean if a field has been set.

### GetTitle

`func (o *MergeReqInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MergeReqInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MergeReqInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MergeReqInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


