# DescribeProjectAnnouncements200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int64** | 公告总数 | [optional] [default to 0]
**PageSize** | Pointer to **int64** | 每页数量 | [optional] [default to 0]
**PageNumber** | Pointer to **int64** | 页数 | [optional] [default to 0]
**List** | Pointer to [**[]ProjectAnnouncementProjectAnnouncement**](ProjectAnnouncementProjectAnnouncement.md) | 公告列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProjectAnnouncements200ResponseResponse

`func NewDescribeProjectAnnouncements200ResponseResponse() *DescribeProjectAnnouncements200ResponseResponse`

NewDescribeProjectAnnouncements200ResponseResponse instantiates a new DescribeProjectAnnouncements200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectAnnouncements200ResponseResponseWithDefaults

`func NewDescribeProjectAnnouncements200ResponseResponseWithDefaults() *DescribeProjectAnnouncements200ResponseResponse`

NewDescribeProjectAnnouncements200ResponseResponseWithDefaults instantiates a new DescribeProjectAnnouncements200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeProjectAnnouncements200ResponseResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribeProjectAnnouncements200ResponseResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectAnnouncements200ResponseResponse) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeProjectAnnouncements200ResponseResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectAnnouncements200ResponseResponse) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeProjectAnnouncements200ResponseResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetList

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetList() []ProjectAnnouncementProjectAnnouncement`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetListOk() (*[]ProjectAnnouncementProjectAnnouncement, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DescribeProjectAnnouncements200ResponseResponse) SetList(v []ProjectAnnouncementProjectAnnouncement)`

SetList sets List field to given value.

### HasList

`func (o *DescribeProjectAnnouncements200ResponseResponse) HasList() bool`

HasList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProjectAnnouncements200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProjectAnnouncements200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProjectAnnouncements200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


