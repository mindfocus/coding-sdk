# CreateReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentIds** | Pointer to **[]int32** | 附件 ID 数组：来自“生成附件预上传信息”接口 | [optional] 
**Name** | **string** | 测试报告标题 | 
**ProjectName** | **string** | 项目名称 | 
**RunIds** | **[]int32** | 测试计划 ID 数组 | 

## Methods

### NewCreateReportRequest

`func NewCreateReportRequest(name string, projectName string, runIds []int32, ) *CreateReportRequest`

NewCreateReportRequest instantiates a new CreateReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportRequestWithDefaults

`func NewCreateReportRequestWithDefaults() *CreateReportRequest`

NewCreateReportRequestWithDefaults instantiates a new CreateReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentIds

`func (o *CreateReportRequest) GetAttachmentIds() []int32`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *CreateReportRequest) GetAttachmentIdsOk() (*[]int32, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *CreateReportRequest) SetAttachmentIds(v []int32)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *CreateReportRequest) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### GetName

`func (o *CreateReportRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReportRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReportRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectName

`func (o *CreateReportRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateReportRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateReportRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunIds

`func (o *CreateReportRequest) GetRunIds() []int32`

GetRunIds returns the RunIds field if non-nil, zero value otherwise.

### GetRunIdsOk

`func (o *CreateReportRequest) GetRunIdsOk() (*[]int32, bool)`

GetRunIdsOk returns a tuple with the RunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunIds

`func (o *CreateReportRequest) SetRunIds(v []int32)`

SetRunIds sets RunIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


