# DescribeAllProjectsIssueWorkLogListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAt** | Pointer to **int64** | 结束时间 | [optional] 
**Limit** | Pointer to **int32** | 每页数量，默认 20，最大1000 | [optional] 
**Offset** | Pointer to **int32** | 偏移量，默认 0 | [optional] 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] 
**StartAt** | Pointer to **int64** | 开始时间 | [optional] 
**UserId** | Pointer to **int64** | 用户 ID | [optional] 

## Methods

### NewDescribeAllProjectsIssueWorkLogListRequest

`func NewDescribeAllProjectsIssueWorkLogListRequest() *DescribeAllProjectsIssueWorkLogListRequest`

NewDescribeAllProjectsIssueWorkLogListRequest instantiates a new DescribeAllProjectsIssueWorkLogListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAllProjectsIssueWorkLogListRequestWithDefaults

`func NewDescribeAllProjectsIssueWorkLogListRequestWithDefaults() *DescribeAllProjectsIssueWorkLogListRequest`

NewDescribeAllProjectsIssueWorkLogListRequestWithDefaults instantiates a new DescribeAllProjectsIssueWorkLogListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetEndAt() int64`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetEndAtOk() (*int64, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetEndAt(v int64)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### GetLimit

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetStartAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeAllProjectsIssueWorkLogListRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeAllProjectsIssueWorkLogListRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DescribeAllProjectsIssueWorkLogListRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


