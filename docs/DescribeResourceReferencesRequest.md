# DescribeResourceReferencesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**ResourceCode** | **int64** | 资源ID | 

## Methods

### NewDescribeResourceReferencesRequest

`func NewDescribeResourceReferencesRequest(projectName string, resourceCode int64, ) *DescribeResourceReferencesRequest`

NewDescribeResourceReferencesRequest instantiates a new DescribeResourceReferencesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeResourceReferencesRequestWithDefaults

`func NewDescribeResourceReferencesRequestWithDefaults() *DescribeResourceReferencesRequest`

NewDescribeResourceReferencesRequestWithDefaults instantiates a new DescribeResourceReferencesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DescribeResourceReferencesRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeResourceReferencesRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeResourceReferencesRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetResourceCode

`func (o *DescribeResourceReferencesRequest) GetResourceCode() int64`

GetResourceCode returns the ResourceCode field if non-nil, zero value otherwise.

### GetResourceCodeOk

`func (o *DescribeResourceReferencesRequest) GetResourceCodeOk() (*int64, bool)`

GetResourceCodeOk returns a tuple with the ResourceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceCode

`func (o *DescribeResourceReferencesRequest) SetResourceCode(v int64)`

SetResourceCode sets ResourceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


