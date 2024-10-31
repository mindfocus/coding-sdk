# DescribeProjectAnnouncementsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | **int64** | 每页数量 | [default to 0]
**PageNumber** | **int64** | 页数 | [default to 0]
**ProjectName** | **string** | 项目名 | [default to ""]

## Methods

### NewDescribeProjectAnnouncementsRequest

`func NewDescribeProjectAnnouncementsRequest(pageSize int64, pageNumber int64, projectName string, ) *DescribeProjectAnnouncementsRequest`

NewDescribeProjectAnnouncementsRequest instantiates a new DescribeProjectAnnouncementsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectAnnouncementsRequestWithDefaults

`func NewDescribeProjectAnnouncementsRequestWithDefaults() *DescribeProjectAnnouncementsRequest`

NewDescribeProjectAnnouncementsRequestWithDefaults instantiates a new DescribeProjectAnnouncementsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *DescribeProjectAnnouncementsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectAnnouncementsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectAnnouncementsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPageNumber

`func (o *DescribeProjectAnnouncementsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectAnnouncementsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectAnnouncementsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetProjectName

`func (o *DescribeProjectAnnouncementsRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeProjectAnnouncementsRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeProjectAnnouncementsRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


