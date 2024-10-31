# DescribePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccurateActions** | **bool** | 针对模糊配置的 action 定义，是否转换成精确的 action 定义返回 | 
**PolicyId** | **int64** | 权限组 ID | 

## Methods

### NewDescribePolicyRequest

`func NewDescribePolicyRequest(accurateActions bool, policyId int64, ) *DescribePolicyRequest`

NewDescribePolicyRequest instantiates a new DescribePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePolicyRequestWithDefaults

`func NewDescribePolicyRequestWithDefaults() *DescribePolicyRequest`

NewDescribePolicyRequestWithDefaults instantiates a new DescribePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccurateActions

`func (o *DescribePolicyRequest) GetAccurateActions() bool`

GetAccurateActions returns the AccurateActions field if non-nil, zero value otherwise.

### GetAccurateActionsOk

`func (o *DescribePolicyRequest) GetAccurateActionsOk() (*bool, bool)`

GetAccurateActionsOk returns a tuple with the AccurateActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccurateActions

`func (o *DescribePolicyRequest) SetAccurateActions(v bool)`

SetAccurateActions sets AccurateActions field to given value.


### GetPolicyId

`func (o *DescribePolicyRequest) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DescribePolicyRequest) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DescribePolicyRequest) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


