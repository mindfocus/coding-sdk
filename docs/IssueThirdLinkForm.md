# IssueThirdLinkForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 链接创建时间 | [optional] 
**Link** | **string** | 链接地址 | [default to ""]
**ThirdType** | **string** | 第三方链接类型  MODAO：墨刀 | [default to ""]
**Title** | **string** | 链接标题 | [default to ""]

## Methods

### NewIssueThirdLinkForm

`func NewIssueThirdLinkForm(link string, thirdType string, title string, ) *IssueThirdLinkForm`

NewIssueThirdLinkForm instantiates a new IssueThirdLinkForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueThirdLinkFormWithDefaults

`func NewIssueThirdLinkFormWithDefaults() *IssueThirdLinkForm`

NewIssueThirdLinkFormWithDefaults instantiates a new IssueThirdLinkForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IssueThirdLinkForm) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueThirdLinkForm) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueThirdLinkForm) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueThirdLinkForm) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLink

`func (o *IssueThirdLinkForm) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IssueThirdLinkForm) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IssueThirdLinkForm) SetLink(v string)`

SetLink sets Link field to given value.


### GetThirdType

`func (o *IssueThirdLinkForm) GetThirdType() string`

GetThirdType returns the ThirdType field if non-nil, zero value otherwise.

### GetThirdTypeOk

`func (o *IssueThirdLinkForm) GetThirdTypeOk() (*string, bool)`

GetThirdTypeOk returns a tuple with the ThirdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdType

`func (o *IssueThirdLinkForm) SetThirdType(v string)`

SetThirdType sets ThirdType field to given value.


### GetTitle

`func (o *IssueThirdLinkForm) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueThirdLinkForm) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueThirdLinkForm) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


