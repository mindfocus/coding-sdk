# DescribeUserProjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | Pointer to **string** | 项目名称 | [optional] 
**UserId** | **int64** | 用户编号 | 

## Methods

### NewDescribeUserProjectsRequest

`func NewDescribeUserProjectsRequest(userId int64, ) *DescribeUserProjectsRequest`

NewDescribeUserProjectsRequest instantiates a new DescribeUserProjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserProjectsRequestWithDefaults

`func NewDescribeUserProjectsRequestWithDefaults() *DescribeUserProjectsRequest`

NewDescribeUserProjectsRequestWithDefaults instantiates a new DescribeUserProjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DescribeUserProjectsRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeUserProjectsRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeUserProjectsRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DescribeUserProjectsRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetUserId

`func (o *DescribeUserProjectsRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeUserProjectsRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeUserProjectsRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


