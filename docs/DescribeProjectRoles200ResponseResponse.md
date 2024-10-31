# DescribeProjectRoles200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]Role**](Role.md) | 用户组 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProjectRoles200ResponseResponse

`func NewDescribeProjectRoles200ResponseResponse() *DescribeProjectRoles200ResponseResponse`

NewDescribeProjectRoles200ResponseResponse instantiates a new DescribeProjectRoles200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectRoles200ResponseResponseWithDefaults

`func NewDescribeProjectRoles200ResponseResponseWithDefaults() *DescribeProjectRoles200ResponseResponse`

NewDescribeProjectRoles200ResponseResponseWithDefaults instantiates a new DescribeProjectRoles200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *DescribeProjectRoles200ResponseResponse) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *DescribeProjectRoles200ResponseResponse) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *DescribeProjectRoles200ResponseResponse) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *DescribeProjectRoles200ResponseResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProjectRoles200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProjectRoles200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProjectRoles200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProjectRoles200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


