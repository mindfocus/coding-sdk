# TeamMemberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | 第几页 | [optional] 
**PageSize** | Pointer to **int32** | 每页条数 | [optional] 
**TeamMembers** | Pointer to [**[]UserData**](UserData.md) | 成员列表信息 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewTeamMemberData

`func NewTeamMemberData() *TeamMemberData`

NewTeamMemberData instantiates a new TeamMemberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMemberDataWithDefaults

`func NewTeamMemberDataWithDefaults() *TeamMemberData`

NewTeamMemberDataWithDefaults instantiates a new TeamMemberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *TeamMemberData) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *TeamMemberData) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *TeamMemberData) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *TeamMemberData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *TeamMemberData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TeamMemberData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TeamMemberData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TeamMemberData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTeamMembers

`func (o *TeamMemberData) GetTeamMembers() []UserData`

GetTeamMembers returns the TeamMembers field if non-nil, zero value otherwise.

### GetTeamMembersOk

`func (o *TeamMemberData) GetTeamMembersOk() (*[]UserData, bool)`

GetTeamMembersOk returns a tuple with the TeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMembers

`func (o *TeamMemberData) SetTeamMembers(v []UserData)`

SetTeamMembers sets TeamMembers field to given value.

### HasTeamMembers

`func (o *TeamMemberData) HasTeamMembers() bool`

HasTeamMembers returns a boolean if a field has been set.

### GetTotalCount

`func (o *TeamMemberData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamMemberData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamMemberData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamMemberData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


