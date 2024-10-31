# DescribeCodingCurrentUser200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**CurrentUser**](CurrentUser.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCodingCurrentUser200ResponseResponse

`func NewDescribeCodingCurrentUser200ResponseResponse() *DescribeCodingCurrentUser200ResponseResponse`

NewDescribeCodingCurrentUser200ResponseResponse instantiates a new DescribeCodingCurrentUser200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCurrentUser200ResponseResponseWithDefaults

`func NewDescribeCodingCurrentUser200ResponseResponseWithDefaults() *DescribeCodingCurrentUser200ResponseResponse`

NewDescribeCodingCurrentUser200ResponseResponseWithDefaults instantiates a new DescribeCodingCurrentUser200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *DescribeCodingCurrentUser200ResponseResponse) GetUser() CurrentUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DescribeCodingCurrentUser200ResponseResponse) GetUserOk() (*CurrentUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DescribeCodingCurrentUser200ResponseResponse) SetUser(v CurrentUser)`

SetUser sets User field to given value.

### HasUser

`func (o *DescribeCodingCurrentUser200ResponseResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeCodingCurrentUser200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCodingCurrentUser200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCodingCurrentUser200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCodingCurrentUser200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


