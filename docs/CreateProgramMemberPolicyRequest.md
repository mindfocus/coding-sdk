# CreateProgramMemberPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policies** | **[]string** | 权限策略，默认策略名或者策略 ID | 
**PrincipalId** | **string** | 身份 ID | 
**PrincipalType** | **string** | 身份类型 | 
**ProgramId** | **int64** | 项目集 ID | [default to 0]

## Methods

### NewCreateProgramMemberPolicyRequest

`func NewCreateProgramMemberPolicyRequest(policies []string, principalId string, principalType string, programId int64, ) *CreateProgramMemberPolicyRequest`

NewCreateProgramMemberPolicyRequest instantiates a new CreateProgramMemberPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProgramMemberPolicyRequestWithDefaults

`func NewCreateProgramMemberPolicyRequestWithDefaults() *CreateProgramMemberPolicyRequest`

NewCreateProgramMemberPolicyRequestWithDefaults instantiates a new CreateProgramMemberPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicies

`func (o *CreateProgramMemberPolicyRequest) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *CreateProgramMemberPolicyRequest) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *CreateProgramMemberPolicyRequest) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.


### GetPrincipalId

`func (o *CreateProgramMemberPolicyRequest) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *CreateProgramMemberPolicyRequest) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *CreateProgramMemberPolicyRequest) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *CreateProgramMemberPolicyRequest) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *CreateProgramMemberPolicyRequest) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *CreateProgramMemberPolicyRequest) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.


### GetProgramId

`func (o *CreateProgramMemberPolicyRequest) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CreateProgramMemberPolicyRequest) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *CreateProgramMemberPolicyRequest) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


