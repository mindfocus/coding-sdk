# DescribeResourceScopeListOnPolicyResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 | [optional] 
**PolicyResourceScopeList** | Pointer to [**[]PolicyResourceScopeInfo**](PolicyResourceScopeInfo.md) | 权限组可用资源范围列表 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewDescribeResourceScopeListOnPolicyResponseData

`func NewDescribeResourceScopeListOnPolicyResponseData() *DescribeResourceScopeListOnPolicyResponseData`

NewDescribeResourceScopeListOnPolicyResponseData instantiates a new DescribeResourceScopeListOnPolicyResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceScopeListOnPolicyResponseDataWithDefaults

`func NewDescribeResourceScopeListOnPolicyResponseDataWithDefaults() *DescribeResourceScopeListOnPolicyResponseData`

NewDescribeResourceScopeListOnPolicyResponseDataWithDefaults instantiates a new DescribeResourceScopeListOnPolicyResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeResourceScopeListOnPolicyResponseData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeResourceScopeListOnPolicyResponseData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeResourceScopeListOnPolicyResponseData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeResourceScopeListOnPolicyResponseData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPolicyResourceScopeList

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPolicyResourceScopeList() []PolicyResourceScopeInfo`

GetPolicyResourceScopeList returns the PolicyResourceScopeList field if non-nil, zero value otherwise.

### GetPolicyResourceScopeListOk

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetPolicyResourceScopeListOk() (*[]PolicyResourceScopeInfo, bool)`

GetPolicyResourceScopeListOk returns a tuple with the PolicyResourceScopeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyResourceScopeList

`func (o *DescribeResourceScopeListOnPolicyResponseData) SetPolicyResourceScopeList(v []PolicyResourceScopeInfo)`

SetPolicyResourceScopeList sets PolicyResourceScopeList field to given value.

### HasPolicyResourceScopeList

`func (o *DescribeResourceScopeListOnPolicyResponseData) HasPolicyResourceScopeList() bool`

HasPolicyResourceScopeList returns a boolean if a field has been set.

### SetPolicyResourceScopeListNil

`func (o *DescribeResourceScopeListOnPolicyResponseData) SetPolicyResourceScopeListNil(b bool)`

 SetPolicyResourceScopeListNil sets the value for PolicyResourceScopeList to be an explicit nil

### UnsetPolicyResourceScopeList
`func (o *DescribeResourceScopeListOnPolicyResponseData) UnsetPolicyResourceScopeList()`

UnsetPolicyResourceScopeList ensures that no value is present for PolicyResourceScopeList, not even an explicit nil
### GetTotalCount

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeResourceScopeListOnPolicyResponseData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeResourceScopeListOnPolicyResponseData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribeResourceScopeListOnPolicyResponseData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


