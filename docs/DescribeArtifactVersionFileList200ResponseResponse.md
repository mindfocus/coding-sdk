# DescribeArtifactVersionFileList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactVersionFileBean**](ArtifactVersionFileBean.md) | 文件列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeArtifactVersionFileList200ResponseResponse

`func NewDescribeArtifactVersionFileList200ResponseResponse() *DescribeArtifactVersionFileList200ResponseResponse`

NewDescribeArtifactVersionFileList200ResponseResponse instantiates a new DescribeArtifactVersionFileList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactVersionFileList200ResponseResponseWithDefaults

`func NewDescribeArtifactVersionFileList200ResponseResponseWithDefaults() *DescribeArtifactVersionFileList200ResponseResponse`

NewDescribeArtifactVersionFileList200ResponseResponseWithDefaults instantiates a new DescribeArtifactVersionFileList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *DescribeArtifactVersionFileList200ResponseResponse) GetInstanceSet() []ArtifactVersionFileBean`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *DescribeArtifactVersionFileList200ResponseResponse) GetInstanceSetOk() (*[]ArtifactVersionFileBean, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *DescribeArtifactVersionFileList200ResponseResponse) SetInstanceSet(v []ArtifactVersionFileBean)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *DescribeArtifactVersionFileList200ResponseResponse) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeArtifactVersionFileList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeArtifactVersionFileList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeArtifactVersionFileList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeArtifactVersionFileList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


