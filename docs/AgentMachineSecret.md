# AgentMachineSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | Pointer to **NullableString** | Secret 值 | [optional] [default to ""]

## Methods

### NewAgentMachineSecret

`func NewAgentMachineSecret() *AgentMachineSecret`

NewAgentMachineSecret instantiates a new AgentMachineSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentMachineSecretWithDefaults

`func NewAgentMachineSecretWithDefaults() *AgentMachineSecret`

NewAgentMachineSecretWithDefaults instantiates a new AgentMachineSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *AgentMachineSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AgentMachineSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AgentMachineSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AgentMachineSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *AgentMachineSecret) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *AgentMachineSecret) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


