# CreateDepartmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 部门名 | [default to ""]
**ParentId** | **int64** | 父部门ID | [default to 0]

## Methods

### NewCreateDepartmentRequest

`func NewCreateDepartmentRequest(name string, parentId int64, ) *CreateDepartmentRequest`

NewCreateDepartmentRequest instantiates a new CreateDepartmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDepartmentRequestWithDefaults

`func NewCreateDepartmentRequestWithDefaults() *CreateDepartmentRequest`

NewCreateDepartmentRequestWithDefaults instantiates a new CreateDepartmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDepartmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDepartmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDepartmentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CreateDepartmentRequest) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateDepartmentRequest) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateDepartmentRequest) SetParentId(v int64)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


