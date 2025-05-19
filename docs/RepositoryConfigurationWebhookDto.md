# RepositoryConfigurationWebhookDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Url** | **string** |  | 
**Events** | Pointer to **[]string** | Events the webhook should be triggered with. | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewRepositoryConfigurationWebhookDto

`func NewRepositoryConfigurationWebhookDto(name string, url string, ) *RepositoryConfigurationWebhookDto`

NewRepositoryConfigurationWebhookDto instantiates a new RepositoryConfigurationWebhookDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigurationWebhookDtoWithDefaults

`func NewRepositoryConfigurationWebhookDtoWithDefaults() *RepositoryConfigurationWebhookDto`

NewRepositoryConfigurationWebhookDtoWithDefaults instantiates a new RepositoryConfigurationWebhookDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryConfigurationWebhookDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryConfigurationWebhookDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryConfigurationWebhookDto) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *RepositoryConfigurationWebhookDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryConfigurationWebhookDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryConfigurationWebhookDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *RepositoryConfigurationWebhookDto) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *RepositoryConfigurationWebhookDto) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *RepositoryConfigurationWebhookDto) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *RepositoryConfigurationWebhookDto) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetConfiguration

`func (o *RepositoryConfigurationWebhookDto) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *RepositoryConfigurationWebhookDto) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *RepositoryConfigurationWebhookDto) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *RepositoryConfigurationWebhookDto) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


