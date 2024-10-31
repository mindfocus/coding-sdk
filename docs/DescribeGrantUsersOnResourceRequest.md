# DescribeGrantUsersOnResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserIdScope** | Pointer to **[]int64** |  | [optional] 
**PolicyIdScope** | Pointer to **[]int64** |  | [optional] 
**PageSize** | **int64** | 请求条数 | [default to 0]
**PageNumber** | **int64** | 请求页数 | [default to 0]
**Resource** | [**RamGrantResourceInfoRequest**](RamGrantResourceInfoRequest.md) |  | 

## Methods

### NewDescribeGrantUsersOnResourceRequest

`func NewDescribeGrantUsersOnResourceRequest(pageSize int64, pageNumber int64, resource RamGrantResourceInfoRequest, ) *DescribeGrantUsersOnResourceRequest`

NewDescribeGrantUsersOnResourceRequest instantiates a new DescribeGrantUsersOnResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGrantUsersOnResourceRequestWithDefaults

`func NewDescribeGrantUsersOnResourceRequestWithDefaults() *DescribeGrantUsersOnResourceRequest`

NewDescribeGrantUsersOnResourceRequestWithDefaults instantiates a new DescribeGrantUsersOnResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserIdScope

`func (o *DescribeGrantUsersOnResourceRequest) GetUserIdScope() []int64`

GetUserIdScope returns the UserIdScope field if non-nil, zero value otherwise.

### GetUserIdScopeOk

`func (o *DescribeGrantUsersOnResourceRequest) GetUserIdScopeOk() (*[]int64, bool)`

GetUserIdScopeOk returns a tuple with the UserIdScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdScope

`func (o *DescribeGrantUsersOnResourceRequest) SetUserIdScope(v []int64)`

SetUserIdScope sets UserIdScope field to given value.

### HasUserIdScope

`func (o *DescribeGrantUsersOnResourceRequest) HasUserIdScope() bool`

HasUserIdScope returns a boolean if a field has been set.

### GetPolicyIdScope

`func (o *DescribeGrantUsersOnResourceRequest) GetPolicyIdScope() []int64`

GetPolicyIdScope returns the PolicyIdScope field if non-nil, zero value otherwise.

### GetPolicyIdScopeOk

`func (o *DescribeGrantUsersOnResourceRequest) GetPolicyIdScopeOk() (*[]int64, bool)`

GetPolicyIdScopeOk returns a tuple with the PolicyIdScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIdScope

`func (o *DescribeGrantUsersOnResourceRequest) SetPolicyIdScope(v []int64)`

SetPolicyIdScope sets PolicyIdScope field to given value.

### HasPolicyIdScope

`func (o *DescribeGrantUsersOnResourceRequest) HasPolicyIdScope() bool`

HasPolicyIdScope returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeGrantUsersOnResourceRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGrantUsersOnResourceRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGrantUsersOnResourceRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPageNumber

`func (o *DescribeGrantUsersOnResourceRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGrantUsersOnResourceRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGrantUsersOnResourceRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetResource

`func (o *DescribeGrantUsersOnResourceRequest) GetResource() RamGrantResourceInfoRequest`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeGrantUsersOnResourceRequest) GetResourceOk() (*RamGrantResourceInfoRequest, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeGrantUsersOnResourceRequest) SetResource(v RamGrantResourceInfoRequest)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


