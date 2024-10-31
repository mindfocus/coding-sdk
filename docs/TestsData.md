# TestsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tests** | Pointer to [**[]Test**](Test.md) | 测试任务信息 | [optional] 

## Methods

### NewTestsData

`func NewTestsData() *TestsData`

NewTestsData instantiates a new TestsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestsDataWithDefaults

`func NewTestsDataWithDefaults() *TestsData`

NewTestsDataWithDefaults instantiates a new TestsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTests

`func (o *TestsData) GetTests() []Test`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *TestsData) GetTestsOk() (*[]Test, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *TestsData) SetTests(v []Test)`

SetTests sets Tests field to given value.

### HasTests

`func (o *TestsData) HasTests() bool`

HasTests returns a boolean if a field has been set.

### SetTestsNil

`func (o *TestsData) SetTestsNil(b bool)`

 SetTestsNil sets the value for Tests to be an explicit nil

### UnsetTests
`func (o *TestsData) UnsetTests()`

UnsetTests ensures that no value is present for Tests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


