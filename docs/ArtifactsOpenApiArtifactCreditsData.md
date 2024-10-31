# ArtifactsOpenApiArtifactCreditsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditKey** | **string** | 授信清单Key | [default to ""]
**Ranges** | [**[]ArtifactsOpenApiArtifactCreditsRangeData**](ArtifactsOpenApiArtifactCreditsRangeData.md) | 授信清单适用范围列表 | 
**Type** | **string** | 类型: NORMAL&#x3D;普通类型，SYNC&#x3D;来源其他系统同步 | [default to ""]
**Enabled** | **bool** | 是否启用 | [default to false]
**Id** | **int64** | 授信清单ID | [default to 0]
**Rules** | [**[]ArtifactsOpenApiArtifactCreditsRuleData**](ArtifactsOpenApiArtifactCreditsRuleData.md) | 授信清单适用范围列表 | 
**Name** | **string** | 授信清单名称 | [default to ""]

## Methods

### NewArtifactsOpenApiArtifactCreditsData

`func NewArtifactsOpenApiArtifactCreditsData(creditKey string, ranges []ArtifactsOpenApiArtifactCreditsRangeData, type_ string, enabled bool, id int64, rules []ArtifactsOpenApiArtifactCreditsRuleData, name string, ) *ArtifactsOpenApiArtifactCreditsData`

NewArtifactsOpenApiArtifactCreditsData instantiates a new ArtifactsOpenApiArtifactCreditsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiArtifactCreditsDataWithDefaults

`func NewArtifactsOpenApiArtifactCreditsDataWithDefaults() *ArtifactsOpenApiArtifactCreditsData`

NewArtifactsOpenApiArtifactCreditsDataWithDefaults instantiates a new ArtifactsOpenApiArtifactCreditsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditKey

`func (o *ArtifactsOpenApiArtifactCreditsData) GetCreditKey() string`

GetCreditKey returns the CreditKey field if non-nil, zero value otherwise.

### GetCreditKeyOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetCreditKeyOk() (*string, bool)`

GetCreditKeyOk returns a tuple with the CreditKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditKey

`func (o *ArtifactsOpenApiArtifactCreditsData) SetCreditKey(v string)`

SetCreditKey sets CreditKey field to given value.


### GetRanges

`func (o *ArtifactsOpenApiArtifactCreditsData) GetRanges() []ArtifactsOpenApiArtifactCreditsRangeData`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetRangesOk() (*[]ArtifactsOpenApiArtifactCreditsRangeData, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ArtifactsOpenApiArtifactCreditsData) SetRanges(v []ArtifactsOpenApiArtifactCreditsRangeData)`

SetRanges sets Ranges field to given value.


### GetType

`func (o *ArtifactsOpenApiArtifactCreditsData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ArtifactsOpenApiArtifactCreditsData) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *ArtifactsOpenApiArtifactCreditsData) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ArtifactsOpenApiArtifactCreditsData) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *ArtifactsOpenApiArtifactCreditsData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactsOpenApiArtifactCreditsData) SetId(v int64)`

SetId sets Id field to given value.


### GetRules

`func (o *ArtifactsOpenApiArtifactCreditsData) GetRules() []ArtifactsOpenApiArtifactCreditsRuleData`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetRulesOk() (*[]ArtifactsOpenApiArtifactCreditsRuleData, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ArtifactsOpenApiArtifactCreditsData) SetRules(v []ArtifactsOpenApiArtifactCreditsRuleData)`

SetRules sets Rules field to given value.


### GetName

`func (o *ArtifactsOpenApiArtifactCreditsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactsOpenApiArtifactCreditsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactsOpenApiArtifactCreditsData) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


