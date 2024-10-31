# CustomStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 步骤 | [default to ""]
**Expected** | **string** | 预期 | [default to ""]

## Methods

### NewCustomStep

`func NewCustomStep(content string, expected string, ) *CustomStep`

NewCustomStep instantiates a new CustomStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomStepWithDefaults

`func NewCustomStepWithDefaults() *CustomStep`

NewCustomStepWithDefaults instantiates a new CustomStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CustomStep) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CustomStep) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CustomStep) SetContent(v string)`

SetContent sets Content field to given value.


### GetExpected

`func (o *CustomStep) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *CustomStep) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *CustomStep) SetExpected(v string)`

SetExpected sets Expected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


