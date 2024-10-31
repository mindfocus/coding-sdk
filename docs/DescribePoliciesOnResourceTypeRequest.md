# DescribePoliciesOnResourceTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**DescribePoliciesOnResourceTypeRequestFilter**](DescribePoliciesOnResourceTypeRequestFilter.md) |  | 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 

## Methods

### NewDescribePoliciesOnResourceTypeRequest

`func NewDescribePoliciesOnResourceTypeRequest(filter DescribePoliciesOnResourceTypeRequestFilter, pageNumber int32, pageSize int32, ) *DescribePoliciesOnResourceTypeRequest`

NewDescribePoliciesOnResourceTypeRequest instantiates a new DescribePoliciesOnResourceTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePoliciesOnResourceTypeRequestWithDefaults

`func NewDescribePoliciesOnResourceTypeRequestWithDefaults() *DescribePoliciesOnResourceTypeRequest`

NewDescribePoliciesOnResourceTypeRequestWithDefaults instantiates a new DescribePoliciesOnResourceTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *DescribePoliciesOnResourceTypeRequest) GetFilter() DescribePoliciesOnResourceTypeRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DescribePoliciesOnResourceTypeRequest) GetFilterOk() (*DescribePoliciesOnResourceTypeRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DescribePoliciesOnResourceTypeRequest) SetFilter(v DescribePoliciesOnResourceTypeRequestFilter)`

SetFilter sets Filter field to given value.


### GetPageNumber

`func (o *DescribePoliciesOnResourceTypeRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribePoliciesOnResourceTypeRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribePoliciesOnResourceTypeRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribePoliciesOnResourceTypeRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribePoliciesOnResourceTypeRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribePoliciesOnResourceTypeRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


