# ModifyArtifactCreditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ranges** | [**[]ArtifactsOpenApiCreateArtifactCreditsRangeData**](ArtifactsOpenApiCreateArtifactCreditsRangeData.md) | 授信清单范围 | 
**Enabled** | **bool** | 是否启用 | [default to false]
**Id** | **int64** | 授信清单ID | [default to 0]
**Rules** | [**[]ArtifactsOpenApiArtifactCreditsRuleData**](ArtifactsOpenApiArtifactCreditsRuleData.md) | 授信清单规则 | 
**Name** | **string** | 授信清单名称 | [default to ""]

## Methods

### NewModifyArtifactCreditRequest

`func NewModifyArtifactCreditRequest(ranges []ArtifactsOpenApiCreateArtifactCreditsRangeData, enabled bool, id int64, rules []ArtifactsOpenApiArtifactCreditsRuleData, name string, ) *ModifyArtifactCreditRequest`

NewModifyArtifactCreditRequest instantiates a new ModifyArtifactCreditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyArtifactCreditRequestWithDefaults

`func NewModifyArtifactCreditRequestWithDefaults() *ModifyArtifactCreditRequest`

NewModifyArtifactCreditRequestWithDefaults instantiates a new ModifyArtifactCreditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRanges

`func (o *ModifyArtifactCreditRequest) GetRanges() []ArtifactsOpenApiCreateArtifactCreditsRangeData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ModifyArtifactCreditRequest) GetRangesOk() (*[]ArtifactsOpenApiCreateArtifactCreditsRangeData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ModifyArtifactCreditRequest) SetRanges(v []ArtifactsOpenApiCreateArtifactCreditsRangeData)`

SetRanges sets Ranges field to given value.


### GetEnabled

`func (o *ModifyArtifactCreditRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModifyArtifactCreditRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModifyArtifactCreditRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ModifyArtifactCreditRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyArtifactCreditRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyArtifactCreditRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetRules

`func (o *ModifyArtifactCreditRequest) GetRules() []ArtifactsOpenApiArtifactCreditsRuleData`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ModifyArtifactCreditRequest) GetRulesOk() (*[]ArtifactsOpenApiArtifactCreditsRuleData, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ModifyArtifactCreditRequest) SetRules(v []ArtifactsOpenApiArtifactCreditsRuleData)`

SetRules sets Rules field to given value.


### GetName

`func (o *ModifyArtifactCreditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyArtifactCreditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyArtifactCreditRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


