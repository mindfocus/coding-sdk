# RefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | 分支名称 | [optional] [default to ""]

## Methods

### NewRefInfo

`func NewRefInfo() *RefInfo`

NewRefInfo instantiates a new RefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefInfoWithDefaults

`func NewRefInfoWithDefaults() *RefInfo`

NewRefInfoWithDefaults instantiates a new RefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RefInfo) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RefInfo) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RefInfo) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RefInfo) HasRef() bool`

HasRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


