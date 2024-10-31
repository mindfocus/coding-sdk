# DeleteProjectLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int64** | 项目ID | [default to 0]
**LabelId** | **int64** | 标签Id | [default to 0]

## Methods

### NewDeleteProjectLabelRequest

`func NewDeleteProjectLabelRequest(projectId int64, labelId int64, ) *DeleteProjectLabelRequest`

NewDeleteProjectLabelRequest instantiates a new DeleteProjectLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProjectLabelRequestWithDefaults

`func NewDeleteProjectLabelRequestWithDefaults() *DeleteProjectLabelRequest`

NewDeleteProjectLabelRequestWithDefaults instantiates a new DeleteProjectLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *DeleteProjectLabelRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteProjectLabelRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteProjectLabelRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetLabelId

`func (o *DeleteProjectLabelRequest) GetLabelId() int64`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *DeleteProjectLabelRequest) GetLabelIdOk() (*int64, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *DeleteProjectLabelRequest) SetLabelId(v int64)`

SetLabelId sets LabelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


