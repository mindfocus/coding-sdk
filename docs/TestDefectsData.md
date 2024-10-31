# TestDefectsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Defects** | Pointer to [**[]TestDefect**](TestDefect.md) | 测试任务的缺陷信息 | [optional] 

## Methods

### NewTestDefectsData

`func NewTestDefectsData() *TestDefectsData`

NewTestDefectsData instantiates a new TestDefectsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestDefectsDataWithDefaults

`func NewTestDefectsDataWithDefaults() *TestDefectsData`

NewTestDefectsDataWithDefaults instantiates a new TestDefectsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefects

`func (o *TestDefectsData) GetDefects() []TestDefect`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *TestDefectsData) GetDefectsOk() (*[]TestDefect, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *TestDefectsData) SetDefects(v []TestDefect)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *TestDefectsData) HasDefects() bool`

HasDefects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


