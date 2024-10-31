# Contributor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to **NullableInt64** | 提交次数 | [optional] 
**Email** | Pointer to **NullableString** | 邮箱 | [optional] [default to ""]
**Name** | Pointer to **NullableString** | 名字 | [optional] [default to ""]

## Methods

### NewContributor

`func NewContributor() *Contributor`

NewContributor instantiates a new Contributor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContributorWithDefaults

`func NewContributorWithDefaults() *Contributor`

NewContributorWithDefaults instantiates a new Contributor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *Contributor) GetCommits() int64`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *Contributor) GetCommitsOk() (*int64, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *Contributor) SetCommits(v int64)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *Contributor) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### SetCommitsNil

`func (o *Contributor) SetCommitsNil(b bool)`

 SetCommitsNil sets the value for Commits to be an explicit nil

### UnsetCommits
`func (o *Contributor) UnsetCommits()`

UnsetCommits ensures that no value is present for Commits, not even an explicit nil
### GetEmail

`func (o *Contributor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contributor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contributor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Contributor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *Contributor) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *Contributor) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *Contributor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Contributor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Contributor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Contributor) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Contributor) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Contributor) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


