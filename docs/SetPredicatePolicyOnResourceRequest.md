# SetPredicatePolicyOnResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ResourceInfo**](ResourceInfo.md) |  | 
**ResourcePredicatePolicy** | **string** |   SELF_PARENT // 同时使用父级资源+当前资源   SELF_NONE  // 只使用当前消息   NONE_PARENT  // 只使用父级资源 | 

## Methods

### NewSetPredicatePolicyOnResourceRequest

`func NewSetPredicatePolicyOnResourceRequest(resource ResourceInfo, resourcePredicatePolicy string, ) *SetPredicatePolicyOnResourceRequest`

NewSetPredicatePolicyOnResourceRequest instantiates a new SetPredicatePolicyOnResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPredicatePolicyOnResourceRequestWithDefaults

`func NewSetPredicatePolicyOnResourceRequestWithDefaults() *SetPredicatePolicyOnResourceRequest`

NewSetPredicatePolicyOnResourceRequestWithDefaults instantiates a new SetPredicatePolicyOnResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *SetPredicatePolicyOnResourceRequest) GetResource() ResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SetPredicatePolicyOnResourceRequest) GetResourceOk() (*ResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SetPredicatePolicyOnResourceRequest) SetResource(v ResourceInfo)`

SetResource sets Resource field to given value.


### GetResourcePredicatePolicy

`func (o *SetPredicatePolicyOnResourceRequest) GetResourcePredicatePolicy() string`

GetResourcePredicatePolicy returns the ResourcePredicatePolicy field if non-nil, zero value otherwise.

### GetResourcePredicatePolicyOk

`func (o *SetPredicatePolicyOnResourceRequest) GetResourcePredicatePolicyOk() (*string, bool)`

GetResourcePredicatePolicyOk returns a tuple with the ResourcePredicatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePredicatePolicy

`func (o *SetPredicatePolicyOnResourceRequest) SetResourcePredicatePolicy(v string)`

SetResourcePredicatePolicy sets ResourcePredicatePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


