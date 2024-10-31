# DescribeIterationListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]Iteration**](Iteration.md) | 迭代列表 | [optional] 
**TotalRow** | Pointer to **int64** | 结果总条数 | [optional] 

## Methods

### NewDescribeIterationListData

`func NewDescribeIterationListData() *DescribeIterationListData`

NewDescribeIterationListData instantiates a new DescribeIterationListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIterationListDataWithDefaults

`func NewDescribeIterationListDataWithDefaults() *DescribeIterationListData`

NewDescribeIterationListDataWithDefaults instantiates a new DescribeIterationListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DescribeIterationListData) GetList() []Iteration`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DescribeIterationListData) GetListOk() (*[]Iteration, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DescribeIterationListData) SetList(v []Iteration)`

SetList sets List field to given value.

### HasList

`func (o *DescribeIterationListData) HasList() bool`

HasList returns a boolean if a field has been set.

### GetTotalRow

`func (o *DescribeIterationListData) GetTotalRow() int64`

GetTotalRow returns the TotalRow field if non-nil, zero value otherwise.

### GetTotalRowOk

`func (o *DescribeIterationListData) GetTotalRowOk() (*int64, bool)`

GetTotalRowOk returns a tuple with the TotalRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRow

`func (o *DescribeIterationListData) SetTotalRow(v int64)`

SetTotalRow sets TotalRow field to given value.

### HasTotalRow

`func (o *DescribeIterationListData) HasTotalRow() bool`

HasTotalRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


