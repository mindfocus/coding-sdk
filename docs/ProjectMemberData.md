# ProjectMemberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | 第几页 | [optional] 
**PageSize** | Pointer to **int32** | 每页条数 | [optional] 
**ProjectMembers** | Pointer to [**[]ProjectMemberUserData**](ProjectMemberUserData.md) | 项目成员列表信息 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewProjectMemberData

`func NewProjectMemberData() *ProjectMemberData`

NewProjectMemberData instantiates a new ProjectMemberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMemberDataWithDefaults

`func NewProjectMemberDataWithDefaults() *ProjectMemberData`

NewProjectMemberDataWithDefaults instantiates a new ProjectMemberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *ProjectMemberData) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ProjectMemberData) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ProjectMemberData) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ProjectMemberData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ProjectMemberData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ProjectMemberData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ProjectMemberData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ProjectMemberData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectMembers

`func (o *ProjectMemberData) GetProjectMembers() []ProjectMemberUserData`

GetProjectMembers returns the ProjectMembers field if non-nil, zero value otherwise.

### GetProjectMembersOk

`func (o *ProjectMemberData) GetProjectMembersOk() (*[]ProjectMemberUserData, bool)`

GetProjectMembersOk returns a tuple with the ProjectMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMembers

`func (o *ProjectMemberData) SetProjectMembers(v []ProjectMemberUserData)`

SetProjectMembers sets ProjectMembers field to given value.

### HasProjectMembers

`func (o *ProjectMemberData) HasProjectMembers() bool`

HasProjectMembers returns a boolean if a field has been set.

### GetTotalCount

`func (o *ProjectMemberData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectMemberData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectMemberData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProjectMemberData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


