# ServiceHookPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** | 分页页码 | [optional] 
**PageSize** | Pointer to **int64** | 分页大小 | [optional] 
**ServiceHook** | Pointer to [**[]ServiceHook**](ServiceHook.md) | Service Hook 列表 | [optional] 
**TotalCount** | Pointer to **int64** | 总记录数 | [optional] 

## Methods

### NewServiceHookPage

`func NewServiceHookPage() *ServiceHookPage`

NewServiceHookPage instantiates a new ServiceHookPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookPageWithDefaults

`func NewServiceHookPageWithDefaults() *ServiceHookPage`

NewServiceHookPageWithDefaults instantiates a new ServiceHookPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ServiceHookPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ServiceHookPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ServiceHookPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ServiceHookPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ServiceHookPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ServiceHookPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ServiceHookPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ServiceHookPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetServiceHook

`func (o *ServiceHookPage) GetServiceHook() []ServiceHook`

GetServiceHook returns the ServiceHook field if non-nil, zero value otherwise.

### GetServiceHookOk

`func (o *ServiceHookPage) GetServiceHookOk() (*[]ServiceHook, bool)`

GetServiceHookOk returns a tuple with the ServiceHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHook

`func (o *ServiceHookPage) SetServiceHook(v []ServiceHook)`

SetServiceHook sets ServiceHook field to given value.

### HasServiceHook

`func (o *ServiceHookPage) HasServiceHook() bool`

HasServiceHook returns a boolean if a field has been set.

### GetTotalCount

`func (o *ServiceHookPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceHookPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceHookPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ServiceHookPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


