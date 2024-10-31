# ModifyProjectLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **int64** | 项目ID | [default to 0]
**LabelId** | **int64** | 标签Id | [default to 0]
**Name** | **string** | 标签名 | [default to ""]
**Color** | **string** | 颜色 | [default to ""]

## Methods

### NewModifyProjectLabelRequest

`func NewModifyProjectLabelRequest(projectId int64, labelId int64, name string, color string, ) *ModifyProjectLabelRequest`

NewModifyProjectLabelRequest instantiates a new ModifyProjectLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyProjectLabelRequestWithDefaults

`func NewModifyProjectLabelRequestWithDefaults() *ModifyProjectLabelRequest`

NewModifyProjectLabelRequestWithDefaults instantiates a new ModifyProjectLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ModifyProjectLabelRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModifyProjectLabelRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModifyProjectLabelRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetLabelId

`func (o *ModifyProjectLabelRequest) GetLabelId() int64`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *ModifyProjectLabelRequest) GetLabelIdOk() (*int64, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *ModifyProjectLabelRequest) SetLabelId(v int64)`

SetLabelId sets LabelId field to given value.


### GetName

`func (o *ModifyProjectLabelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyProjectLabelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyProjectLabelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *ModifyProjectLabelRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ModifyProjectLabelRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ModifyProjectLabelRequest) SetColor(v string)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


