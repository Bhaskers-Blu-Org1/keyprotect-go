# ImportToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | [**time.Time**](time.Time.md) | The date the import token was created. The date format follows RFC 3339. | [optional] [readonly] 
**Expiration** | **float32** | The time in seconds from the creation of an import token that determines how long its associated public key  remains valid.     The minimum value is &#x60;300&#x60; seconds (5 minutes), and the maximum value is &#x60;86400&#x60; (24 hours). The default value is &#x60;600&#x60; (10 minutes). | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date the import token expires. The date format follows RFC 3339. | [optional] [readonly] 
**MaxAllowedRetrievals** | **float32** | The number of times that an import token can be retrieved within its expiration time before it is no longer accessible.  | [optional] 
**RemainingRetrievals** | **float32** | The number of retrievals that are available for the import token before it is no longer accessible.   | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


