# IssueThirdLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 第三方链接 Id | [optional] 
**Link** | Pointer to **string** | 链接地址 | [optional] [default to ""]
**ThirdType** | Pointer to **string** | 第三方链接类型，MODAO - 墨刀 | [optional] [default to ""]
**Title** | Pointer to **string** | 名称 | [optional] [default to ""]

## Methods

### NewIssueThirdLink

`func NewIssueThirdLink() *IssueThirdLink`

NewIssueThirdLink instantiates a new IssueThirdLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueThirdLinkWithDefaults

`func NewIssueThirdLinkWithDefaults() *IssueThirdLink`

NewIssueThirdLinkWithDefaults instantiates a new IssueThirdLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IssueThirdLink) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueThirdLink) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueThirdLink) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueThirdLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLink

`func (o *IssueThirdLink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IssueThirdLink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IssueThirdLink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *IssueThirdLink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetThirdType

`func (o *IssueThirdLink) GetThirdType() string`

GetThirdType returns the ThirdType field if non-nil, zero value otherwise.

### GetThirdTypeOk

`func (o *IssueThirdLink) GetThirdTypeOk() (*string, bool)`

GetThirdTypeOk returns a tuple with the ThirdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdType

`func (o *IssueThirdLink) SetThirdType(v string)`

SetThirdType sets ThirdType field to given value.

### HasThirdType

`func (o *IssueThirdLink) HasThirdType() bool`

HasThirdType returns a boolean if a field has been set.

### GetTitle

`func (o *IssueThirdLink) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueThirdLink) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueThirdLink) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssueThirdLink) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


