# DescribeArtifactProperties200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactProperty**](ArtifactProperty.md) | 制品属性列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeArtifactProperties200ResponseResponse

`func NewDescribeArtifactProperties200ResponseResponse() *DescribeArtifactProperties200ResponseResponse`

NewDescribeArtifactProperties200ResponseResponse instantiates a new DescribeArtifactProperties200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactProperties200ResponseResponseWithDefaults

`func NewDescribeArtifactProperties200ResponseResponseWithDefaults() *DescribeArtifactProperties200ResponseResponse`

NewDescribeArtifactProperties200ResponseResponseWithDefaults instantiates a new DescribeArtifactProperties200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *DescribeArtifactProperties200ResponseResponse) GetInstanceSet() []ArtifactProperty`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *DescribeArtifactProperties200ResponseResponse) GetInstanceSetOk() (*[]ArtifactProperty, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *DescribeArtifactProperties200ResponseResponse) SetInstanceSet(v []ArtifactProperty)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *DescribeArtifactProperties200ResponseResponse) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeArtifactProperties200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeArtifactProperties200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeArtifactProperties200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeArtifactProperties200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


