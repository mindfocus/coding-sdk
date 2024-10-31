# Committer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | 提交者的e-mail地址 | [optional] [default to ""]
**Name** | Pointer to **string** | 提交者的姓名 | [optional] [default to ""]

## Methods

### NewCommitter

`func NewCommitter() *Committer`

NewCommitter instantiates a new Committer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitterWithDefaults

`func NewCommitterWithDefaults() *Committer`

NewCommitterWithDefaults instantiates a new Committer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Committer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Committer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Committer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Committer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *Committer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Committer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Committer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Committer) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


