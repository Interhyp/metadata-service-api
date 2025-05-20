# RepositoryConfigurationWebhooksDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predefined** | Pointer to **[]string** | List of predefined webhooks | [optional] 
**Additional** | Pointer to [**[]RepositoryConfigurationWebhookDto**](RepositoryConfigurationWebhookDto.md) | Additional webhooks to be configured. | [optional] 

## Methods

### NewRepositoryConfigurationWebhooksDto

`func NewRepositoryConfigurationWebhooksDto() *RepositoryConfigurationWebhooksDto`

NewRepositoryConfigurationWebhooksDto instantiates a new RepositoryConfigurationWebhooksDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigurationWebhooksDtoWithDefaults

`func NewRepositoryConfigurationWebhooksDtoWithDefaults() *RepositoryConfigurationWebhooksDto`

NewRepositoryConfigurationWebhooksDtoWithDefaults instantiates a new RepositoryConfigurationWebhooksDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredefined

`func (o *RepositoryConfigurationWebhooksDto) GetPredefined() []string`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *RepositoryConfigurationWebhooksDto) GetPredefinedOk() (*[]string, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *RepositoryConfigurationWebhooksDto) SetPredefined(v []string)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *RepositoryConfigurationWebhooksDto) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.

### GetAdditional

`func (o *RepositoryConfigurationWebhooksDto) GetAdditional() []RepositoryConfigurationWebhookDto`

GetAdditional returns the Additional field if non-nil, zero value otherwise.

### GetAdditionalOk

`func (o *RepositoryConfigurationWebhooksDto) GetAdditionalOk() (*[]RepositoryConfigurationWebhookDto, bool)`

GetAdditionalOk returns a tuple with the Additional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditional

`func (o *RepositoryConfigurationWebhooksDto) SetAdditional(v []RepositoryConfigurationWebhookDto)`

SetAdditional sets Additional field to given value.

### HasAdditional

`func (o *RepositoryConfigurationWebhooksDto) HasAdditional() bool`

HasAdditional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


