# RelatedCase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 用例ID | [optional] 
**Name** | Pointer to **string** | 用例名称 | [optional] [default to ""]
**Priority** | Pointer to **int64** | 优先级, 0 &#x3D;&gt; &#39;紧急&#39;, 1 &#x3D;&gt; &#39;高&#39;, 2 &#x3D;&gt; &#39;中&#39;, 3 &#x3D;&gt; &#39;低 | [optional] 

## Methods

### NewRelatedCase

`func NewRelatedCase() *RelatedCase`

NewRelatedCase instantiates a new RelatedCase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedCaseWithDefaults

`func NewRelatedCaseWithDefaults() *RelatedCase`

NewRelatedCaseWithDefaults instantiates a new RelatedCase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RelatedCase) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedCase) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedCase) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedCase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelatedCase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedCase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedCase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelatedCase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *RelatedCase) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RelatedCase) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RelatedCase) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RelatedCase) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


