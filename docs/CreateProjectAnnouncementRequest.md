# CreateProjectAnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名 | [default to ""]
**Content** | **string** | 公告内容 | [default to ""]

## Methods

### NewCreateProjectAnnouncementRequest

`func NewCreateProjectAnnouncementRequest(projectName string, content string, ) *CreateProjectAnnouncementRequest`

NewCreateProjectAnnouncementRequest instantiates a new CreateProjectAnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectAnnouncementRequestWithDefaults

`func NewCreateProjectAnnouncementRequestWithDefaults() *CreateProjectAnnouncementRequest`

NewCreateProjectAnnouncementRequestWithDefaults instantiates a new CreateProjectAnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *CreateProjectAnnouncementRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateProjectAnnouncementRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateProjectAnnouncementRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetContent

`func (o *CreateProjectAnnouncementRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateProjectAnnouncementRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateProjectAnnouncementRequest) SetContent(v string)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


