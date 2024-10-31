# DetachFromResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | [**[]GrantInfo**](GrantInfo.md) | 授权信息 | 
**Resource** | [**ResourceInfo**](ResourceInfo.md) |  | 

## Methods

### NewDetachFromResourceRequest

`func NewDetachFromResourceRequest(grants []GrantInfo, resource ResourceInfo, ) *DetachFromResourceRequest`

NewDetachFromResourceRequest instantiates a new DetachFromResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachFromResourceRequestWithDefaults

`func NewDetachFromResourceRequestWithDefaults() *DetachFromResourceRequest`

NewDetachFromResourceRequestWithDefaults instantiates a new DetachFromResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrants

`func (o *DetachFromResourceRequest) GetGrants() []GrantInfo`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### GetGrantsOk

`func (o *DetachFromResourceRequest) GetGrantsOk() (*[]GrantInfo, bool)`

GetGrantsOk returns a tuple with the Grants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrants

`func (o *DetachFromResourceRequest) SetGrants(v []GrantInfo)`

SetGrants sets Grants field to given value.


### GetResource

`func (o *DetachFromResourceRequest) GetResource() ResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DetachFromResourceRequest) GetResourceOk() (*ResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DetachFromResourceRequest) SetResource(v ResourceInfo)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


