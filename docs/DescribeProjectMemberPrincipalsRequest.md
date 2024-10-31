# DescribeProjectMemberPrincipalsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | 关键字搜索：成员名/邮箱 | [optional] 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 
**PolicyId** | Pointer to **int64** | 权限组Id | [optional] 
**ProjectId** | **int32** | 项目Id | 

## Methods

### NewDescribeProjectMemberPrincipalsRequest

`func NewDescribeProjectMemberPrincipalsRequest(pageNumber int32, pageSize int32, projectId int32, ) *DescribeProjectMemberPrincipalsRequest`

NewDescribeProjectMemberPrincipalsRequest instantiates a new DescribeProjectMemberPrincipalsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectMemberPrincipalsRequestWithDefaults

`func NewDescribeProjectMemberPrincipalsRequestWithDefaults() *DescribeProjectMemberPrincipalsRequest`

NewDescribeProjectMemberPrincipalsRequestWithDefaults instantiates a new DescribeProjectMemberPrincipalsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *DescribeProjectMemberPrincipalsRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeProjectMemberPrincipalsRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeProjectMemberPrincipalsRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeProjectMemberPrincipalsRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeProjectMemberPrincipalsRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectMemberPrincipalsRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectMemberPrincipalsRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeProjectMemberPrincipalsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectMemberPrincipalsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectMemberPrincipalsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPolicyId

`func (o *DescribeProjectMemberPrincipalsRequest) GetPolicyId() int64`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DescribeProjectMemberPrincipalsRequest) GetPolicyIdOk() (*int64, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DescribeProjectMemberPrincipalsRequest) SetPolicyId(v int64)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *DescribeProjectMemberPrincipalsRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeProjectMemberPrincipalsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectMemberPrincipalsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectMemberPrincipalsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


