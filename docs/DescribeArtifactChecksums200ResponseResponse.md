# DescribeArtifactChecksums200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactChecksum**](ArtifactChecksum.md) | 制品Checksum列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeArtifactChecksums200ResponseResponse

`func NewDescribeArtifactChecksums200ResponseResponse() *DescribeArtifactChecksums200ResponseResponse`

NewDescribeArtifactChecksums200ResponseResponse instantiates a new DescribeArtifactChecksums200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactChecksums200ResponseResponseWithDefaults

`func NewDescribeArtifactChecksums200ResponseResponseWithDefaults() *DescribeArtifactChecksums200ResponseResponse`

NewDescribeArtifactChecksums200ResponseResponseWithDefaults instantiates a new DescribeArtifactChecksums200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *DescribeArtifactChecksums200ResponseResponse) GetInstanceSet() []ArtifactChecksum`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *DescribeArtifactChecksums200ResponseResponse) GetInstanceSetOk() (*[]ArtifactChecksum, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *DescribeArtifactChecksums200ResponseResponse) SetInstanceSet(v []ArtifactChecksum)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *DescribeArtifactChecksums200ResponseResponse) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeArtifactChecksums200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeArtifactChecksums200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeArtifactChecksums200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeArtifactChecksums200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


