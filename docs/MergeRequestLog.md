# MergeRequestLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | 操作方式 | [optional] [default to ""]
**Id** | Pointer to **int64** | 请求id | [optional] 
**Name** | Pointer to **string** | 操作人姓名 | [optional] [default to ""]

## Methods

### NewMergeRequestLog

`func NewMergeRequestLog() *MergeRequestLog`

NewMergeRequestLog instantiates a new MergeRequestLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMergeRequestLogWithDefaults

`func NewMergeRequestLogWithDefaults() *MergeRequestLog`

NewMergeRequestLogWithDefaults instantiates a new MergeRequestLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *MergeRequestLog) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *MergeRequestLog) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *MergeRequestLog) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *MergeRequestLog) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetId

`func (o *MergeRequestLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MergeRequestLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MergeRequestLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *MergeRequestLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MergeRequestLog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MergeRequestLog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MergeRequestLog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MergeRequestLog) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


