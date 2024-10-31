# DescribeGrantObjectsOnResourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **int32** | 页码 | 
**PageSize** | **int32** | 每页条数  | 
**Resource** | [**ResourceInfo**](ResourceInfo.md) |  | 

## Methods

### NewDescribeGrantObjectsOnResourceRequest

`func NewDescribeGrantObjectsOnResourceRequest(pageNumber int32, pageSize int32, resource ResourceInfo, ) *DescribeGrantObjectsOnResourceRequest`

NewDescribeGrantObjectsOnResourceRequest instantiates a new DescribeGrantObjectsOnResourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGrantObjectsOnResourceRequestWithDefaults

`func NewDescribeGrantObjectsOnResourceRequestWithDefaults() *DescribeGrantObjectsOnResourceRequest`

NewDescribeGrantObjectsOnResourceRequestWithDefaults instantiates a new DescribeGrantObjectsOnResourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeGrantObjectsOnResourceRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGrantObjectsOnResourceRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGrantObjectsOnResourceRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeGrantObjectsOnResourceRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGrantObjectsOnResourceRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGrantObjectsOnResourceRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetResource

`func (o *DescribeGrantObjectsOnResourceRequest) GetResource() ResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeGrantObjectsOnResourceRequest) GetResourceOk() (*ResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeGrantObjectsOnResourceRequest) SetResource(v ResourceInfo)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


