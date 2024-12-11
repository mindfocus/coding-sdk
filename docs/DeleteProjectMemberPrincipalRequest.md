# DeleteProjectMemberPrincipalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principals** | [**[]PrincipalDel**](PrincipalDel.md) | 成员主体信息 | 
**ProjectId** | **int32** | 项目Id | 

## Methods

### NewDeleteProjectMemberPrincipalRequest

`func NewDeleteProjectMemberPrincipalRequest(principals []PrincipalDel, projectId int32, ) *DeleteProjectMemberPrincipalRequest`

NewDeleteProjectMemberPrincipalRequest instantiates a new DeleteProjectMemberPrincipalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteProjectMemberPrincipalRequestWithDefaults

`func NewDeleteProjectMemberPrincipalRequestWithDefaults() *DeleteProjectMemberPrincipalRequest`

NewDeleteProjectMemberPrincipalRequestWithDefaults instantiates a new DeleteProjectMemberPrincipalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipals

`func (o *DeleteProjectMemberPrincipalRequest) GetPrincipals() []PrincipalDel`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *DeleteProjectMemberPrincipalRequest) GetPrincipalsOk() (*[]PrincipalDel, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *DeleteProjectMemberPrincipalRequest) SetPrincipals(v []PrincipalDel)`

SetPrincipals sets Principals field to given value.


### GetProjectId

`func (o *DeleteProjectMemberPrincipalRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteProjectMemberPrincipalRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteProjectMemberPrincipalRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


