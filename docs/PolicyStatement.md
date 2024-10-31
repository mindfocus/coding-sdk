# PolicyStatement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **[]string** | 权限 action 的二段式描述 | 
**Effect** | **string** | 效能：ALLOW | DENY | [default to ""]
**Resource** | **[]string** | 授权资源的三段式描述，一般填 * 即可 | 

## Methods

### NewPolicyStatement

`func NewPolicyStatement(action []string, effect string, resource []string, ) *PolicyStatement`

NewPolicyStatement instantiates a new PolicyStatement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyStatementWithDefaults

`func NewPolicyStatementWithDefaults() *PolicyStatement`

NewPolicyStatementWithDefaults instantiates a new PolicyStatement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PolicyStatement) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PolicyStatement) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PolicyStatement) SetAction(v []string)`

SetAction sets Action field to given value.


### GetEffect

`func (o *PolicyStatement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PolicyStatement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PolicyStatement) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetResource

`func (o *PolicyStatement) GetResource() []string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyStatement) GetResourceOk() (*[]string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyStatement) SetResource(v []string)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


