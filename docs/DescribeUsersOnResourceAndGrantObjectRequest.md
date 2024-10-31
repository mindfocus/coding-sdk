# DescribeUsersOnResourceAndGrantObjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grant** | [**DescribeUsersOnResourceAndGrantObjectGrantInfo**](DescribeUsersOnResourceAndGrantObjectGrantInfo.md) |  | 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 
**Resource** | [**ResourceInfo**](ResourceInfo.md) |  | 

## Methods

### NewDescribeUsersOnResourceAndGrantObjectRequest

`func NewDescribeUsersOnResourceAndGrantObjectRequest(grant DescribeUsersOnResourceAndGrantObjectGrantInfo, pageNumber int32, pageSize int32, resource ResourceInfo, ) *DescribeUsersOnResourceAndGrantObjectRequest`

NewDescribeUsersOnResourceAndGrantObjectRequest instantiates a new DescribeUsersOnResourceAndGrantObjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersOnResourceAndGrantObjectRequestWithDefaults

`func NewDescribeUsersOnResourceAndGrantObjectRequestWithDefaults() *DescribeUsersOnResourceAndGrantObjectRequest`

NewDescribeUsersOnResourceAndGrantObjectRequestWithDefaults instantiates a new DescribeUsersOnResourceAndGrantObjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrant

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetGrant() DescribeUsersOnResourceAndGrantObjectGrantInfo`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetGrantOk() (*DescribeUsersOnResourceAndGrantObjectGrantInfo, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) SetGrant(v DescribeUsersOnResourceAndGrantObjectGrantInfo)`

SetGrant sets Grant field to given value.


### GetPageNumber

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetResource

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetResource() ResourceInfo`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) GetResourceOk() (*ResourceInfo, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DescribeUsersOnResourceAndGrantObjectRequest) SetResource(v ResourceInfo)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


