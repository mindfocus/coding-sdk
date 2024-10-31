# CreateProjectMemberPrincipalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principals** | [**[]PrincipalAdd**](PrincipalAdd.md) | 成员主体信息 | 
**ProjectId** | **int32** | 项目Id | 

## Methods

### NewCreateProjectMemberPrincipalRequest

`func NewCreateProjectMemberPrincipalRequest(principals []PrincipalAdd, projectId int32, ) *CreateProjectMemberPrincipalRequest`

NewCreateProjectMemberPrincipalRequest instantiates a new CreateProjectMemberPrincipalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProjectMemberPrincipalRequestWithDefaults

`func NewCreateProjectMemberPrincipalRequestWithDefaults() *CreateProjectMemberPrincipalRequest`

NewCreateProjectMemberPrincipalRequestWithDefaults instantiates a new CreateProjectMemberPrincipalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipals

`func (o *CreateProjectMemberPrincipalRequest) GetPrincipals() []PrincipalAdd`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *CreateProjectMemberPrincipalRequest) GetPrincipalsOk() (*[]PrincipalAdd, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *CreateProjectMemberPrincipalRequest) SetPrincipals(v []PrincipalAdd)`

SetPrincipals sets Principals field to given value.


### GetProjectId

`func (o *CreateProjectMemberPrincipalRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateProjectMemberPrincipalRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateProjectMemberPrincipalRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


