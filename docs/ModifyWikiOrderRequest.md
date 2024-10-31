# ModifyWikiOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**After** | Pointer to **int64** | 在第几层级之后 | [optional] 
**Before** | Pointer to **int64** | 在第几层级之前 | [optional] 
**Forced** | Pointer to **bool** | 是否检查权限,默认false | [optional] 
**Iid** | **int64** | wiki Iid | 
**ParentIid** | **int64** | 父级 Iid | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewModifyWikiOrderRequest

`func NewModifyWikiOrderRequest(iid int64, parentIid int64, projectName string, ) *ModifyWikiOrderRequest`

NewModifyWikiOrderRequest instantiates a new ModifyWikiOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWikiOrderRequestWithDefaults

`func NewModifyWikiOrderRequestWithDefaults() *ModifyWikiOrderRequest`

NewModifyWikiOrderRequestWithDefaults instantiates a new ModifyWikiOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfter

`func (o *ModifyWikiOrderRequest) GetAfter() int64`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ModifyWikiOrderRequest) GetAfterOk() (*int64, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ModifyWikiOrderRequest) SetAfter(v int64)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ModifyWikiOrderRequest) HasAfter() bool`

HasAfter returns a boolean if a field has been set.

### GetBefore

`func (o *ModifyWikiOrderRequest) GetBefore() int64`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ModifyWikiOrderRequest) GetBeforeOk() (*int64, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ModifyWikiOrderRequest) SetBefore(v int64)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ModifyWikiOrderRequest) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetForced

`func (o *ModifyWikiOrderRequest) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *ModifyWikiOrderRequest) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *ModifyWikiOrderRequest) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *ModifyWikiOrderRequest) HasForced() bool`

HasForced returns a boolean if a field has been set.

### GetIid

`func (o *ModifyWikiOrderRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *ModifyWikiOrderRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *ModifyWikiOrderRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetParentIid

`func (o *ModifyWikiOrderRequest) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *ModifyWikiOrderRequest) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *ModifyWikiOrderRequest) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.


### GetProjectName

`func (o *ModifyWikiOrderRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyWikiOrderRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyWikiOrderRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


