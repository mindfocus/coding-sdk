# DescribeReportListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAt** | Pointer to **string** | 创建时间 | [optional] 
**Keyword** | Pointer to **string** | 报告名称关键词 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**StartAt** | Pointer to **string** | 创建时间 | [optional] 
**Status** | Pointer to **string** | 报告状态：CREATING 创建中，AVAILABLE 可用，UNAVAILABLE 不可用 | [optional] 

## Methods

### NewDescribeReportListRequest

`func NewDescribeReportListRequest(projectName string, ) *DescribeReportListRequest`

NewDescribeReportListRequest instantiates a new DescribeReportListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeReportListRequestWithDefaults

`func NewDescribeReportListRequestWithDefaults() *DescribeReportListRequest`

NewDescribeReportListRequestWithDefaults instantiates a new DescribeReportListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAt

`func (o *DescribeReportListRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *DescribeReportListRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *DescribeReportListRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *DescribeReportListRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetKeyword

`func (o *DescribeReportListRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeReportListRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeReportListRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeReportListRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeReportListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeReportListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeReportListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStartAt

`func (o *DescribeReportListRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *DescribeReportListRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *DescribeReportListRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *DescribeReportListRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeReportListRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeReportListRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeReportListRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeReportListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


