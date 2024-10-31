# ProjectsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | 第几页 | [optional] 
**PageSize** | Pointer to **int32** | 每页条数 | [optional] 
**ProjectList** | Pointer to [**[]Project**](Project.md) | 项目集合 | [optional] 
**TotalCount** | Pointer to **int32** | 总条数 | [optional] 

## Methods

### NewProjectsData

`func NewProjectsData() *ProjectsData`

NewProjectsData instantiates a new ProjectsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsDataWithDefaults

`func NewProjectsDataWithDefaults() *ProjectsData`

NewProjectsDataWithDefaults instantiates a new ProjectsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ProjectsData) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ProjectsData) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ProjectsData) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ProjectsData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ProjectsData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProjectsData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProjectsData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProjectsData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectList

`func (o *ProjectsData) GetProjectList() []Project`

GetProjectList returns the ProjectList field if non-nil, zero value otherwise.

### GetProjectListOk

`func (o *ProjectsData) GetProjectListOk() (*[]Project, bool)`

GetProjectListOk returns a tuple with the ProjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectList

`func (o *ProjectsData) SetProjectList(v []Project)`

SetProjectList sets ProjectList field to given value.

### HasProjectList

`func (o *ProjectsData) HasProjectList() bool`

HasProjectList returns a boolean if a field has been set.

### SetProjectListNil

`func (o *ProjectsData) SetProjectListNil(b bool)`

 SetProjectListNil sets the value for ProjectList to be an explicit nil

### UnsetProjectList
`func (o *ProjectsData) UnsetProjectList()`

UnsetProjectList ensures that no value is present for ProjectList, not even an explicit nil
### GetTotalCount

`func (o *ProjectsData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectsData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectsData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProjectsData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


