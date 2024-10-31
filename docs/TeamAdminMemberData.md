# TeamAdminMemberData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]TeamAdminMember**](TeamAdminMember.md) | 团队管理员列表信息 | [optional] 
**PageNumber** | Pointer to **int32** | 第几页 | [optional] 
**PageSize** | Pointer to **int32** | 每页条数 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewTeamAdminMemberData

`func NewTeamAdminMemberData() *TeamAdminMemberData`

NewTeamAdminMemberData instantiates a new TeamAdminMemberData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamAdminMemberDataWithDefaults

`func NewTeamAdminMemberDataWithDefaults() *TeamAdminMemberData`

NewTeamAdminMemberDataWithDefaults instantiates a new TeamAdminMemberData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *TeamAdminMemberData) GetMembers() []TeamAdminMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TeamAdminMemberData) GetMembersOk() (*[]TeamAdminMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TeamAdminMemberData) SetMembers(v []TeamAdminMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TeamAdminMemberData) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPageNumber

`func (o *TeamAdminMemberData) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *TeamAdminMemberData) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *TeamAdminMemberData) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *TeamAdminMemberData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *TeamAdminMemberData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TeamAdminMemberData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TeamAdminMemberData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TeamAdminMemberData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *TeamAdminMemberData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamAdminMemberData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamAdminMemberData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamAdminMemberData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


