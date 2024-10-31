# AttachResourceScopeToPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | **int64** | 权限组 ID | 
**ResourceInfos** | [**[]ResourceInfoOfPolicyScope**](ResourceInfoOfPolicyScope.md) | 添加的资源 | 

## Methods

### NewAttachResourceScopeToPolicyRequest

`func NewAttachResourceScopeToPolicyRequest(policyId int64, resourceInfos []ResourceInfoOfPolicyScope, ) *AttachResourceScopeToPolicyRequest`

NewAttachResourceScopeToPolicyRequest instantiates a new AttachResourceScopeToPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachResourceScopeToPolicyRequestWithDefaults

`func NewAttachResourceScopeToPolicyRequestWithDefaults() *AttachResourceScopeToPolicyRequest`

NewAttachResourceScopeToPolicyRequestWithDefaults instantiates a new AttachResourceScopeToPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *AttachResourceScopeToPolicyRequest) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AttachResourceScopeToPolicyRequest) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AttachResourceScopeToPolicyRequest) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.


### GetResourceInfos

`func (o *AttachResourceScopeToPolicyRequest) GetResourceInfos() []ResourceInfoOfPolicyScope`

GetResourceInfos returns the ResourceInfos field if non-nil, zero value otherwise.

### GetResourceInfosOk

`func (o *AttachResourceScopeToPolicyRequest) GetResourceInfosOk() (*[]ResourceInfoOfPolicyScope, bool)`

GetResourceInfosOk returns a tuple with the ResourceInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfos

`func (o *AttachResourceScopeToPolicyRequest) SetResourceInfos(v []ResourceInfoOfPolicyScope)`

SetResourceInfos sets ResourceInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


