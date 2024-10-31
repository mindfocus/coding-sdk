# DeletePoliciesById200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedPolicyId** | Pointer to **[]int64** | 成功删除的权限组 ID | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDeletePoliciesById200ResponseResponse

`func NewDeletePoliciesById200ResponseResponse() *DeletePoliciesById200ResponseResponse`

NewDeletePoliciesById200ResponseResponse instantiates a new DeletePoliciesById200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePoliciesById200ResponseResponseWithDefaults

`func NewDeletePoliciesById200ResponseResponseWithDefaults() *DeletePoliciesById200ResponseResponse`

NewDeletePoliciesById200ResponseResponseWithDefaults instantiates a new DeletePoliciesById200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedPolicyId

`func (o *DeletePoliciesById200ResponseResponse) GetDeletedPolicyId() []int64`

GetDeletedPolicyId returns the DeletedPolicyId field if non-nil, zero value otherwise.

### GetDeletedPolicyIdOk

`func (o *DeletePoliciesById200ResponseResponse) GetDeletedPolicyIdOk() (*[]int64, bool)`

GetDeletedPolicyIdOk returns a tuple with the DeletedPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedPolicyId

`func (o *DeletePoliciesById200ResponseResponse) SetDeletedPolicyId(v []int64)`

SetDeletedPolicyId sets DeletedPolicyId field to given value.

### HasDeletedPolicyId

`func (o *DeletePoliciesById200ResponseResponse) HasDeletedPolicyId() bool`

HasDeletedPolicyId returns a boolean if a field has been set.

### SetDeletedPolicyIdNil

`func (o *DeletePoliciesById200ResponseResponse) SetDeletedPolicyIdNil(b bool)`

 SetDeletedPolicyIdNil sets the value for DeletedPolicyId to be an explicit nil

### UnsetDeletedPolicyId
`func (o *DeletePoliciesById200ResponseResponse) UnsetDeletedPolicyId()`

UnsetDeletedPolicyId ensures that no value is present for DeletedPolicyId, not even an explicit nil
### GetRequestId

`func (o *DeletePoliciesById200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DeletePoliciesById200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DeletePoliciesById200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DeletePoliciesById200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


