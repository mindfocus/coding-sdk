# DeleteProjectAnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名 | [default to ""]
**Id** | **[]int64** | 要删除的公告ID列表 | 

## Methods

### NewDeleteProjectAnnouncementRequest

`func NewDeleteProjectAnnouncementRequest(projectName string, id []int64, ) *DeleteProjectAnnouncementRequest`

NewDeleteProjectAnnouncementRequest instantiates a new DeleteProjectAnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProjectAnnouncementRequestWithDefaults

`func NewDeleteProjectAnnouncementRequestWithDefaults() *DeleteProjectAnnouncementRequest`

NewDeleteProjectAnnouncementRequestWithDefaults instantiates a new DeleteProjectAnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DeleteProjectAnnouncementRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteProjectAnnouncementRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteProjectAnnouncementRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetId

`func (o *DeleteProjectAnnouncementRequest) GetId() []int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteProjectAnnouncementRequest) GetIdOk() (*[]int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteProjectAnnouncementRequest) SetId(v []int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


