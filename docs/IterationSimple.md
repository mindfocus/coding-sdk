# IterationSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | 迭代代码 | [optional] 
**Id** | Pointer to **int64** | 迭代Id | [optional] 
**Name** | Pointer to **string** | 迭代名称 | [optional] [default to ""]
**Status** | Pointer to **string** | 迭代状态：WAIT_PROCESS,PROCESSING,COMPLETED | [optional] [default to ""]

## Methods

### NewIterationSimple

`func NewIterationSimple() *IterationSimple`

NewIterationSimple instantiates a new IterationSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIterationSimpleWithDefaults

`func NewIterationSimpleWithDefaults() *IterationSimple`

NewIterationSimpleWithDefaults instantiates a new IterationSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *IterationSimple) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IterationSimple) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IterationSimple) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *IterationSimple) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetId

`func (o *IterationSimple) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IterationSimple) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IterationSimple) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IterationSimple) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IterationSimple) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IterationSimple) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IterationSimple) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IterationSimple) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IterationSimple) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IterationSimple) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IterationSimple) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IterationSimple) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


