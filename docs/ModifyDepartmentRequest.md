# ModifyDepartmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | 部门ID | [default to 0]
**Name** | **string** | 部门名 | [default to ""]
**ParentId** | **int64** | 父级部门ID | [default to 0]

## Methods

### NewModifyDepartmentRequest

`func NewModifyDepartmentRequest(id int64, name string, parentId int64, ) *ModifyDepartmentRequest`

NewModifyDepartmentRequest instantiates a new ModifyDepartmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepartmentRequestWithDefaults

`func NewModifyDepartmentRequestWithDefaults() *ModifyDepartmentRequest`

NewModifyDepartmentRequestWithDefaults instantiates a new ModifyDepartmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModifyDepartmentRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyDepartmentRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyDepartmentRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ModifyDepartmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyDepartmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyDepartmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *ModifyDepartmentRequest) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModifyDepartmentRequest) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModifyDepartmentRequest) SetParentId(v int64)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


