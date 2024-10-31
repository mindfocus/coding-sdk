# DescribeUsersByGroupIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | 查询条件：GroupId | 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 

## Methods

### NewDescribeUsersByGroupIdRequest

`func NewDescribeUsersByGroupIdRequest(groupId int32, pageNumber int32, pageSize int32, ) *DescribeUsersByGroupIdRequest`

NewDescribeUsersByGroupIdRequest instantiates a new DescribeUsersByGroupIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersByGroupIdRequestWithDefaults

`func NewDescribeUsersByGroupIdRequestWithDefaults() *DescribeUsersByGroupIdRequest`

NewDescribeUsersByGroupIdRequestWithDefaults instantiates a new DescribeUsersByGroupIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *DescribeUsersByGroupIdRequest) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DescribeUsersByGroupIdRequest) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DescribeUsersByGroupIdRequest) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetPageNumber

`func (o *DescribeUsersByGroupIdRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeUsersByGroupIdRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeUsersByGroupIdRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeUsersByGroupIdRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeUsersByGroupIdRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeUsersByGroupIdRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


