# ArtifactFilterRuleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | 名称匹配规则（支持 “EQUAL” 和 “REGEX”） | [default to ""]
**Value** | **string** | 名称匹配值 | [default to ""]

## Methods

### NewArtifactFilterRuleDetail

`func NewArtifactFilterRuleDetail(algorithm string, value string, ) *ArtifactFilterRuleDetail`

NewArtifactFilterRuleDetail instantiates a new ArtifactFilterRuleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactFilterRuleDetailWithDefaults

`func NewArtifactFilterRuleDetailWithDefaults() *ArtifactFilterRuleDetail`

NewArtifactFilterRuleDetailWithDefaults instantiates a new ArtifactFilterRuleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *ArtifactFilterRuleDetail) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ArtifactFilterRuleDetail) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ArtifactFilterRuleDetail) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetValue

`func (o *ArtifactFilterRuleDetail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArtifactFilterRuleDetail) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArtifactFilterRuleDetail) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


