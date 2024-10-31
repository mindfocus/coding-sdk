# DescribeAvailablePoliciesOnResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**DescribeAvailablePoliciesOnResourceRequestFilter**](DescribeAvailablePoliciesOnResourceRequestFilter.md) |  | 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 
**Resource** | [**ResourceInfoOfPolicyScope**](ResourceInfoOfPolicyScope.md) |  | 

## Methods

### NewDescribeAvailablePoliciesOnResourceRequest

`func NewDescribeAvailablePoliciesOnResourceRequest(filter DescribeAvailablePoliciesOnResourceRequestFilter, pageNumber int32, pageSize int32, resource ResourceInfoOfPolicyScope, ) *DescribeAvailablePoliciesOnResourceRequest`

NewDescribeAvailablePoliciesOnResourceRequest instantiates a new DescribeAvailablePoliciesOnResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAvailablePoliciesOnResourceRequestWithDefaults

`func NewDescribeAvailablePoliciesOnResourceRequestWithDefaults() *DescribeAvailablePoliciesOnResourceRequest`

NewDescribeAvailablePoliciesOnResourceRequestWithDefaults instantiates a new DescribeAvailablePoliciesOnResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetFilter() DescribeAvailablePoliciesOnResourceRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetFilterOk() (*DescribeAvailablePoliciesOnResourceRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DescribeAvailablePoliciesOnResourceRequest) SetFilter(v DescribeAvailablePoliciesOnResourceRequestFilter)`

SetFilter sets Filter field to given value.


### GetPageNumber

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeAvailablePoliciesOnResourceRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeAvailablePoliciesOnResourceRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetResource

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetResource() ResourceInfoOfPolicyScope`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeAvailablePoliciesOnResourceRequest) GetResourceOk() (*ResourceInfoOfPolicyScope, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeAvailablePoliciesOnResourceRequest) SetResource(v ResourceInfoOfPolicyScope)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


