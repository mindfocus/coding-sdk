# IssueProjectLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **string** | 标签颜色，例如：#5A606B | [optional] [default to ""]
**Id** | Pointer to **int64** | 项目标签 Id | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]

## Methods

### NewIssueProjectLabel

`func NewIssueProjectLabel() *IssueProjectLabel`

NewIssueProjectLabel instantiates a new IssueProjectLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueProjectLabelWithDefaults

`func NewIssueProjectLabelWithDefaults() *IssueProjectLabel`

NewIssueProjectLabelWithDefaults instantiates a new IssueProjectLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *IssueProjectLabel) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *IssueProjectLabel) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *IssueProjectLabel) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *IssueProjectLabel) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetId

`func (o *IssueProjectLabel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueProjectLabel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueProjectLabel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueProjectLabel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IssueProjectLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueProjectLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueProjectLabel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueProjectLabel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


