# DescribeArtifactCreditList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactsOpenApiArtifactCreditsData**](ArtifactsOpenApiArtifactCreditsData.md) | 授信清单列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeArtifactCreditList200ResponseResponse

`func NewDescribeArtifactCreditList200ResponseResponse() *DescribeArtifactCreditList200ResponseResponse`

NewDescribeArtifactCreditList200ResponseResponse instantiates a new DescribeArtifactCreditList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactCreditList200ResponseResponseWithDefaults

`func NewDescribeArtifactCreditList200ResponseResponseWithDefaults() *DescribeArtifactCreditList200ResponseResponse`

NewDescribeArtifactCreditList200ResponseResponseWithDefaults instantiates a new DescribeArtifactCreditList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *DescribeArtifactCreditList200ResponseResponse) GetInstanceSet() []ArtifactsOpenApiArtifactCreditsData`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *DescribeArtifactCreditList200ResponseResponse) GetInstanceSetOk() (*[]ArtifactsOpenApiArtifactCreditsData, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *DescribeArtifactCreditList200ResponseResponse) SetInstanceSet(v []ArtifactsOpenApiArtifactCreditsData)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *DescribeArtifactCreditList200ResponseResponse) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeArtifactCreditList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeArtifactCreditList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeArtifactCreditList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeArtifactCreditList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


