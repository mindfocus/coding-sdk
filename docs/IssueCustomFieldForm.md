# IssueCustomFieldForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 自定义属性值，多选值使用 \&quot;,\&quot; 分隔 | [default to ""]
**Id** | **int64** | 事项自定义属性 Id | 

## Methods

### NewIssueCustomFieldForm

`func NewIssueCustomFieldForm(content string, id int64, ) *IssueCustomFieldForm`

NewIssueCustomFieldForm instantiates a new IssueCustomFieldForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueCustomFieldFormWithDefaults

`func NewIssueCustomFieldFormWithDefaults() *IssueCustomFieldForm`

NewIssueCustomFieldFormWithDefaults instantiates a new IssueCustomFieldForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *IssueCustomFieldForm) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *IssueCustomFieldForm) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *IssueCustomFieldForm) SetContent(v string)`

SetContent sets Content field to given value.


### GetId

`func (o *IssueCustomFieldForm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueCustomFieldForm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueCustomFieldForm) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


