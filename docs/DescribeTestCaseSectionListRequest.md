# DescribeTestCaseSectionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **int32** | 父级 ID，默认 0 | [optional] 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeTestCaseSectionListRequest

`func NewDescribeTestCaseSectionListRequest(projectName string, ) *DescribeTestCaseSectionListRequest`

NewDescribeTestCaseSectionListRequest instantiates a new DescribeTestCaseSectionListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTestCaseSectionListRequestWithDefaults

`func NewDescribeTestCaseSectionListRequestWithDefaults() *DescribeTestCaseSectionListRequest`

NewDescribeTestCaseSectionListRequestWithDefaults instantiates a new DescribeTestCaseSectionListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *DescribeTestCaseSectionListRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DescribeTestCaseSectionListRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DescribeTestCaseSectionListRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *DescribeTestCaseSectionListRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeTestCaseSectionListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTestCaseSectionListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTestCaseSectionListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


