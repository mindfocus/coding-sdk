# DescribePoliciesOnResourceTypeResponsePageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 | [optional] 
**PolicyList** | Pointer to [**[]PolicyInfo**](PolicyInfo.md) | 权限组列表 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewDescribePoliciesOnResourceTypeResponsePageData

`func NewDescribePoliciesOnResourceTypeResponsePageData() *DescribePoliciesOnResourceTypeResponsePageData`

NewDescribePoliciesOnResourceTypeResponsePageData instantiates a new DescribePoliciesOnResourceTypeResponsePageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePoliciesOnResourceTypeResponsePageDataWithDefaults

`func NewDescribePoliciesOnResourceTypeResponsePageDataWithDefaults() *DescribePoliciesOnResourceTypeResponsePageData`

NewDescribePoliciesOnResourceTypeResponsePageDataWithDefaults instantiates a new DescribePoliciesOnResourceTypeResponsePageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribePoliciesOnResourceTypeResponsePageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribePoliciesOnResourceTypeResponsePageData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribePoliciesOnResourceTypeResponsePageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribePoliciesOnResourceTypeResponsePageData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPolicyList

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPolicyList() []PolicyInfo`

GetPolicyList returns the PolicyList field if non-nil, zero value otherwise.

### GetPolicyListOk

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetPolicyListOk() (*[]PolicyInfo, bool)`

GetPolicyListOk returns a tuple with the PolicyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyList

`func (o *DescribePoliciesOnResourceTypeResponsePageData) SetPolicyList(v []PolicyInfo)`

SetPolicyList sets PolicyList field to given value.

### HasPolicyList

`func (o *DescribePoliciesOnResourceTypeResponsePageData) HasPolicyList() bool`

HasPolicyList returns a boolean if a field has been set.

### SetPolicyListNil

`func (o *DescribePoliciesOnResourceTypeResponsePageData) SetPolicyListNil(b bool)`

 SetPolicyListNil sets the value for PolicyList to be an explicit nil

### UnsetPolicyList
`func (o *DescribePoliciesOnResourceTypeResponsePageData) UnsetPolicyList()`

UnsetPolicyList ensures that no value is present for PolicyList, not even an explicit nil
### GetTotalCount

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribePoliciesOnResourceTypeResponsePageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribePoliciesOnResourceTypeResponsePageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribePoliciesOnResourceTypeResponsePageData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


