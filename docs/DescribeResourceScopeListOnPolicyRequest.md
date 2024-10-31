# DescribeResourceScopeListOnPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**DescribeResourceScopeListOnPolicyRequestFilter**](DescribeResourceScopeListOnPolicyRequestFilter.md) |  | [optional] 
**PageNumber** | **int32** | 请求页码 | 
**PageSize** | **int32** | 请求条数 | 
**PolicyId** | **int64** | 权限组 ID | 

## Methods

### NewDescribeResourceScopeListOnPolicyRequest

`func NewDescribeResourceScopeListOnPolicyRequest(pageNumber int32, pageSize int32, policyId int64, ) *DescribeResourceScopeListOnPolicyRequest`

NewDescribeResourceScopeListOnPolicyRequest instantiates a new DescribeResourceScopeListOnPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceScopeListOnPolicyRequestWithDefaults

`func NewDescribeResourceScopeListOnPolicyRequestWithDefaults() *DescribeResourceScopeListOnPolicyRequest`

NewDescribeResourceScopeListOnPolicyRequestWithDefaults instantiates a new DescribeResourceScopeListOnPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *DescribeResourceScopeListOnPolicyRequest) GetFilter() DescribeResourceScopeListOnPolicyRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DescribeResourceScopeListOnPolicyRequest) GetFilterOk() (*DescribeResourceScopeListOnPolicyRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DescribeResourceScopeListOnPolicyRequest) SetFilter(v DescribeResourceScopeListOnPolicyRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DescribeResourceScopeListOnPolicyRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeResourceScopeListOnPolicyRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeResourceScopeListOnPolicyRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPolicyId

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DescribeResourceScopeListOnPolicyRequest) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DescribeResourceScopeListOnPolicyRequest) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


