# CodingCIDepotDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Commit 内容 | [optional] [default to ""]
**Sha** | Pointer to **string** | CommitId | [optional] [default to ""]

## Methods

### NewCodingCIDepotDetail

`func NewCodingCIDepotDetail() *CodingCIDepotDetail`

NewCodingCIDepotDetail instantiates a new CodingCIDepotDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIDepotDetailWithDefaults

`func NewCodingCIDepotDetailWithDefaults() *CodingCIDepotDetail`

NewCodingCIDepotDetailWithDefaults instantiates a new CodingCIDepotDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodingCIDepotDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCIDepotDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCIDepotDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodingCIDepotDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSha

`func (o *CodingCIDepotDetail) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CodingCIDepotDetail) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CodingCIDepotDetail) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CodingCIDepotDetail) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


