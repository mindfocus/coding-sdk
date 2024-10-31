# DescribeProjectDepotInfoListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **string** | 页号 | [optional] 
**PageSize** | Pointer to **string** | 每页的个数 | [optional] 
**ProjectId** | **int64** | 项目id | 

## Methods

### NewDescribeProjectDepotInfoListRequest

`func NewDescribeProjectDepotInfoListRequest(projectId int64, ) *DescribeProjectDepotInfoListRequest`

NewDescribeProjectDepotInfoListRequest instantiates a new DescribeProjectDepotInfoListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectDepotInfoListRequestWithDefaults

`func NewDescribeProjectDepotInfoListRequestWithDefaults() *DescribeProjectDepotInfoListRequest`

NewDescribeProjectDepotInfoListRequestWithDefaults instantiates a new DescribeProjectDepotInfoListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeProjectDepotInfoListRequest) GetPageNumber() string`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectDepotInfoListRequest) GetPageNumberOk() (*string, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectDepotInfoListRequest) SetPageNumber(v string)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeProjectDepotInfoListRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeProjectDepotInfoListRequest) GetPageSize() string`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectDepotInfoListRequest) GetPageSizeOk() (*string, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectDepotInfoListRequest) SetPageSize(v string)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeProjectDepotInfoListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeProjectDepotInfoListRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectDepotInfoListRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectDepotInfoListRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


