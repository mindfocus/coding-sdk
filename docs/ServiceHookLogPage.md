# ServiceHookLogPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Log** | Pointer to [**[]ServiceHookLog**](ServiceHookLog.md) | Service Hook 发送记录列表 | [optional] 
**PageNumber** | Pointer to **int64** | 分页页码 | [optional] 
**PageSize** | Pointer to **int64** | 分页大小 | [optional] 
**TotalCount** | Pointer to **int64** | 总记录数 | [optional] 

## Methods

### NewServiceHookLogPage

`func NewServiceHookLogPage() *ServiceHookLogPage`

NewServiceHookLogPage instantiates a new ServiceHookLogPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookLogPageWithDefaults

`func NewServiceHookLogPageWithDefaults() *ServiceHookLogPage`

NewServiceHookLogPageWithDefaults instantiates a new ServiceHookLogPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLog

`func (o *ServiceHookLogPage) GetLog() []ServiceHookLog`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ServiceHookLogPage) GetLogOk() (*[]ServiceHookLog, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ServiceHookLogPage) SetLog(v []ServiceHookLog)`

SetLog sets Log field to given value.

### HasLog

`func (o *ServiceHookLogPage) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetPageNumber

`func (o *ServiceHookLogPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ServiceHookLogPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ServiceHookLogPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ServiceHookLogPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceHookLogPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceHookLogPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceHookLogPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceHookLogPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ServiceHookLogPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceHookLogPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceHookLogPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServiceHookLogPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


