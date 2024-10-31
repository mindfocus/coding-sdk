# DeleteReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**ReleaseCode** | **int64** | 版本code | 

## Methods

### NewDeleteReleaseRequest

`func NewDeleteReleaseRequest(projectName string, releaseCode int64, ) *DeleteReleaseRequest`

NewDeleteReleaseRequest instantiates a new DeleteReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteReleaseRequestWithDefaults

`func NewDeleteReleaseRequestWithDefaults() *DeleteReleaseRequest`

NewDeleteReleaseRequestWithDefaults instantiates a new DeleteReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DeleteReleaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteReleaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteReleaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseCode

`func (o *DeleteReleaseRequest) GetReleaseCode() int64`

GetReleaseCode returns the ReleaseCode field if non-nil, zero value otherwise.

### GetReleaseCodeOk

`func (o *DeleteReleaseRequest) GetReleaseCodeOk() (*int64, bool)`

GetReleaseCodeOk returns a tuple with the ReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCode

`func (o *DeleteReleaseRequest) SetReleaseCode(v int64)`

SetReleaseCode sets ReleaseCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


