# ModifyTestCaseSectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 分组名称 | 
**ParentId** | Pointer to **int32** | 父级 ID，默认 0 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SectionId** | **int32** | 分组 ID | 

## Methods

### NewModifyTestCaseSectionRequest

`func NewModifyTestCaseSectionRequest(name string, projectName string, sectionId int32, ) *ModifyTestCaseSectionRequest`

NewModifyTestCaseSectionRequest instantiates a new ModifyTestCaseSectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyTestCaseSectionRequestWithDefaults

`func NewModifyTestCaseSectionRequestWithDefaults() *ModifyTestCaseSectionRequest`

NewModifyTestCaseSectionRequestWithDefaults instantiates a new ModifyTestCaseSectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModifyTestCaseSectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyTestCaseSectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyTestCaseSectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *ModifyTestCaseSectionRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ModifyTestCaseSectionRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ModifyTestCaseSectionRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ModifyTestCaseSectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyTestCaseSectionRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyTestCaseSectionRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyTestCaseSectionRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *ModifyTestCaseSectionRequest) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ModifyTestCaseSectionRequest) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ModifyTestCaseSectionRequest) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


