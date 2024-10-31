# DescribeWorkbenchIssueListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | 关键字 | [optional] 
**PageNumber** | Pointer to **int64** | 分页查询的分页数，不填的话从第一页开始查询 | [optional] 
**PageSize** | Pointer to **int64** | 每页展示数，默认 20 ，最大值 500 | [optional] 
**ProjectId** | Pointer to **int64** | 不填或者填0则查询团队内参与的所有项目 | [optional] 
**Type** | Pointer to **string** | 事项类型，取值如“REQUIREMENT”、“DEFECT”等 | [optional] 

## Methods

### NewDescribeWorkbenchIssueListRequest

`func NewDescribeWorkbenchIssueListRequest() *DescribeWorkbenchIssueListRequest`

NewDescribeWorkbenchIssueListRequest instantiates a new DescribeWorkbenchIssueListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkbenchIssueListRequestWithDefaults

`func NewDescribeWorkbenchIssueListRequestWithDefaults() *DescribeWorkbenchIssueListRequest`

NewDescribeWorkbenchIssueListRequestWithDefaults instantiates a new DescribeWorkbenchIssueListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *DescribeWorkbenchIssueListRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeWorkbenchIssueListRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeWorkbenchIssueListRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeWorkbenchIssueListRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeWorkbenchIssueListRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeWorkbenchIssueListRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeWorkbenchIssueListRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeWorkbenchIssueListRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeWorkbenchIssueListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeWorkbenchIssueListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeWorkbenchIssueListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeWorkbenchIssueListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeWorkbenchIssueListRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeWorkbenchIssueListRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeWorkbenchIssueListRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DescribeWorkbenchIssueListRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetType

`func (o *DescribeWorkbenchIssueListRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeWorkbenchIssueListRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeWorkbenchIssueListRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DescribeWorkbenchIssueListRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


