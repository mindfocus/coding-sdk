# DeleteTestCaseSectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**SectionId** | **int32** | 分组 ID | 

## Methods

### NewDeleteTestCaseSectionRequest

`func NewDeleteTestCaseSectionRequest(projectName string, sectionId int32, ) *DeleteTestCaseSectionRequest`

NewDeleteTestCaseSectionRequest instantiates a new DeleteTestCaseSectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTestCaseSectionRequestWithDefaults

`func NewDeleteTestCaseSectionRequestWithDefaults() *DeleteTestCaseSectionRequest`

NewDeleteTestCaseSectionRequestWithDefaults instantiates a new DeleteTestCaseSectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DeleteTestCaseSectionRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteTestCaseSectionRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteTestCaseSectionRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *DeleteTestCaseSectionRequest) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *DeleteTestCaseSectionRequest) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *DeleteTestCaseSectionRequest) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


