# CreateArtifactCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ranges** | [**[]ArtifactsOpenApiCreateArtifactCreditsRangeData**](ArtifactsOpenApiCreateArtifactCreditsRangeData.md) | 授信清单适用范围 | 
**Enabled** | **bool** | 是否启用 | [default to false]
**Rules** | [**[]ArtifactsOpenApiArtifactCreditsRuleData**](ArtifactsOpenApiArtifactCreditsRuleData.md) | 授信规则 | 
**Name** | **string** | 授信清单名称 | [default to ""]

## Methods

### NewCreateArtifactCreditRequest

`func NewCreateArtifactCreditRequest(ranges []ArtifactsOpenApiCreateArtifactCreditsRangeData, enabled bool, rules []ArtifactsOpenApiArtifactCreditsRuleData, name string, ) *CreateArtifactCreditRequest`

NewCreateArtifactCreditRequest instantiates a new CreateArtifactCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateArtifactCreditRequestWithDefaults

`func NewCreateArtifactCreditRequestWithDefaults() *CreateArtifactCreditRequest`

NewCreateArtifactCreditRequestWithDefaults instantiates a new CreateArtifactCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRanges

`func (o *CreateArtifactCreditRequest) GetRanges() []ArtifactsOpenApiCreateArtifactCreditsRangeData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *CreateArtifactCreditRequest) GetRangesOk() (*[]ArtifactsOpenApiCreateArtifactCreditsRangeData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *CreateArtifactCreditRequest) SetRanges(v []ArtifactsOpenApiCreateArtifactCreditsRangeData)`

SetRanges sets Ranges field to given value.


### GetEnabled

`func (o *CreateArtifactCreditRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateArtifactCreditRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateArtifactCreditRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetRules

`func (o *CreateArtifactCreditRequest) GetRules() []ArtifactsOpenApiArtifactCreditsRuleData`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateArtifactCreditRequest) GetRulesOk() (*[]ArtifactsOpenApiArtifactCreditsRuleData, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateArtifactCreditRequest) SetRules(v []ArtifactsOpenApiArtifactCreditsRuleData)`

SetRules sets Rules field to given value.


### GetName

`func (o *CreateArtifactCreditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateArtifactCreditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateArtifactCreditRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


