# DescribeCdDeployTimeByProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndAt** | **string** | 结束时间（格式：yyyy-MM-dd HH:mm:ss） | 
**ProjectName** | **string** | 项目名称 | 
**StartAt** | **string** | 开始时间（格式：yyyy-MM-dd HH:mm:ss） | 

## Methods

### NewDescribeCdDeployTimeByProjectRequest

`func NewDescribeCdDeployTimeByProjectRequest(endAt string, projectName string, startAt string, ) *DescribeCdDeployTimeByProjectRequest`

NewDescribeCdDeployTimeByProjectRequest instantiates a new DescribeCdDeployTimeByProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployTimeByProjectRequestWithDefaults

`func NewDescribeCdDeployTimeByProjectRequestWithDefaults() *DescribeCdDeployTimeByProjectRequest`

NewDescribeCdDeployTimeByProjectRequestWithDefaults instantiates a new DescribeCdDeployTimeByProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndAt

`func (o *DescribeCdDeployTimeByProjectRequest) GetEndAt() string`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *DescribeCdDeployTimeByProjectRequest) GetEndAtOk() (*string, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *DescribeCdDeployTimeByProjectRequest) SetEndAt(v string)`

SetEndAt sets EndAt field to given value.


### GetProjectName

`func (o *DescribeCdDeployTimeByProjectRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeCdDeployTimeByProjectRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeCdDeployTimeByProjectRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStartAt

`func (o *DescribeCdDeployTimeByProjectRequest) GetStartAt() string`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *DescribeCdDeployTimeByProjectRequest) GetStartAtOk() (*string, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *DescribeCdDeployTimeByProjectRequest) SetStartAt(v string)`

SetStartAt sets StartAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


