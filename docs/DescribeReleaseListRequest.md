# DescribeReleaseListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IssueCondition**](IssueCondition.md) | 筛选条件列表，每一个值都是一个筛选条件，条件取值可以参考页面上的对应的HTTP接口  | [optional] 
**Page** | Pointer to **int64** | 分页查询中的页数，page从1开始计数  | [optional] 
**PageSize** | Pointer to **int64** | 分页查询中每页的大小  | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SortKey** | Pointer to **string** | 排序KEY  | [optional] 
**SortValue** | Pointer to **string** | 排序VALUE  | [optional] 

## Methods

### NewDescribeReleaseListRequest

`func NewDescribeReleaseListRequest(projectName string, ) *DescribeReleaseListRequest`

NewDescribeReleaseListRequest instantiates a new DescribeReleaseListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeReleaseListRequestWithDefaults

`func NewDescribeReleaseListRequestWithDefaults() *DescribeReleaseListRequest`

NewDescribeReleaseListRequestWithDefaults instantiates a new DescribeReleaseListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DescribeReleaseListRequest) GetConditions() []IssueCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DescribeReleaseListRequest) GetConditionsOk() (*[]IssueCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DescribeReleaseListRequest) SetConditions(v []IssueCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DescribeReleaseListRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPage

`func (o *DescribeReleaseListRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DescribeReleaseListRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DescribeReleaseListRequest) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DescribeReleaseListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeReleaseListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeReleaseListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeReleaseListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeReleaseListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeReleaseListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeReleaseListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeReleaseListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSortKey

`func (o *DescribeReleaseListRequest) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *DescribeReleaseListRequest) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *DescribeReleaseListRequest) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *DescribeReleaseListRequest) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortValue

`func (o *DescribeReleaseListRequest) GetSortValue() string`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *DescribeReleaseListRequest) GetSortValueOk() (*string, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *DescribeReleaseListRequest) SetSortValue(v string)`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *DescribeReleaseListRequest) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


