# ModifyDepartment200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Department** | Pointer to [**DepartmentDepartmentData**](DepartmentDepartmentData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewModifyDepartment200ResponseResponse

`func NewModifyDepartment200ResponseResponse() *ModifyDepartment200ResponseResponse`

NewModifyDepartment200ResponseResponse instantiates a new ModifyDepartment200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepartment200ResponseResponseWithDefaults

`func NewModifyDepartment200ResponseResponseWithDefaults() *ModifyDepartment200ResponseResponse`

NewModifyDepartment200ResponseResponseWithDefaults instantiates a new ModifyDepartment200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartment

`func (o *ModifyDepartment200ResponseResponse) GetDepartment() DepartmentDepartmentData`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ModifyDepartment200ResponseResponse) GetDepartmentOk() (*DepartmentDepartmentData, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ModifyDepartment200ResponseResponse) SetDepartment(v DepartmentDepartmentData)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ModifyDepartment200ResponseResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetRequestId

`func (o *ModifyDepartment200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ModifyDepartment200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ModifyDepartment200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ModifyDepartment200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


