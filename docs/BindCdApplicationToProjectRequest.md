# BindCdApplicationToProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** | CD 应用名 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewBindCdApplicationToProjectRequest

`func NewBindCdApplicationToProjectRequest(application string, projectName string, ) *BindCdApplicationToProjectRequest`

NewBindCdApplicationToProjectRequest instantiates a new BindCdApplicationToProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindCdApplicationToProjectRequestWithDefaults

`func NewBindCdApplicationToProjectRequestWithDefaults() *BindCdApplicationToProjectRequest`

NewBindCdApplicationToProjectRequestWithDefaults instantiates a new BindCdApplicationToProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *BindCdApplicationToProjectRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *BindCdApplicationToProjectRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *BindCdApplicationToProjectRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetProjectName

`func (o *BindCdApplicationToProjectRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BindCdApplicationToProjectRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BindCdApplicationToProjectRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


