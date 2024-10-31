# IssueFilterListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFilterList** | Pointer to [**[]IssueFilter**](IssueFilter.md) | 自定义过滤器列表 | [optional] 
**SystemFilterList** | Pointer to [**[]IssueFilter**](IssueFilter.md) | 系统过滤器列表 | [optional] 

## Methods

### NewIssueFilterListData

`func NewIssueFilterListData() *IssueFilterListData`

NewIssueFilterListData instantiates a new IssueFilterListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFilterListDataWithDefaults

`func NewIssueFilterListDataWithDefaults() *IssueFilterListData`

NewIssueFilterListDataWithDefaults instantiates a new IssueFilterListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFilterList

`func (o *IssueFilterListData) GetCustomFilterList() []IssueFilter`

GetCustomFilterList returns the CustomFilterList field if non-nil, zero value otherwise.

### GetCustomFilterListOk

`func (o *IssueFilterListData) GetCustomFilterListOk() (*[]IssueFilter, bool)`

GetCustomFilterListOk returns a tuple with the CustomFilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFilterList

`func (o *IssueFilterListData) SetCustomFilterList(v []IssueFilter)`

SetCustomFilterList sets CustomFilterList field to given value.

### HasCustomFilterList

`func (o *IssueFilterListData) HasCustomFilterList() bool`

HasCustomFilterList returns a boolean if a field has been set.

### GetSystemFilterList

`func (o *IssueFilterListData) GetSystemFilterList() []IssueFilter`

GetSystemFilterList returns the SystemFilterList field if non-nil, zero value otherwise.

### GetSystemFilterListOk

`func (o *IssueFilterListData) GetSystemFilterListOk() (*[]IssueFilter, bool)`

GetSystemFilterListOk returns a tuple with the SystemFilterList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFilterList

`func (o *IssueFilterListData) SetSystemFilterList(v []IssueFilter)`

SetSystemFilterList sets SystemFilterList field to given value.

### HasSystemFilterList

`func (o *IssueFilterListData) HasSystemFilterList() bool`

HasSystemFilterList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


