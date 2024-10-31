# ApiFileFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | 文件夹ID | [default to 0]
**Name** | **string** | 文件夹名 | [default to ""]

## Methods

### NewApiFileFolder

`func NewApiFileFolder(id int64, name string, ) *ApiFileFolder`

NewApiFileFolder instantiates a new ApiFileFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFileFolderWithDefaults

`func NewApiFileFolderWithDefaults() *ApiFileFolder`

NewApiFileFolderWithDefaults instantiates a new ApiFileFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiFileFolder) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiFileFolder) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiFileFolder) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ApiFileFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiFileFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiFileFolder) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


