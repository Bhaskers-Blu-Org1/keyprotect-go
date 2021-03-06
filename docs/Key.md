# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **[]string** | Up to 30 tags can be created. Tags can be between 2-30 characters, including spaces. Special characters not permitted include the angled bracket, comma, colon, ampersand, and vertical pipe character (|).    To protect your privacy, do not use personal data, such as your name or location, as a tag for your key.  | [optional] 
**Crn** | **string** | The Cloud Resource Name (CRN) that uniquely identifies your cloud network resources. | [optional] [readonly] 
**Description** | **string** | A text field used to provide a more detailed description of the key. The maximum length is 240 characters.    To protect your privacy, do not use personal data, such as your name or location, as a description for your  key. | [optional] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date the key material expires. The date format follows RFC 3339. You can set an expiration date on any  key on its creation. If you create a key without specifying an expiration date, the key does not expire. | [optional] 
**Extractable** | **bool** | A boolean value that determines whether the key material can leave the service.       If set to &#x60;false&#x60;, Key Protect designates the key as a nonextractable root key used for &#x60;wrap&#x60; and &#x60;unwrap&#x60; actions. If set to &#x60;true&#x60;, Key Protect designates the key as a standard key that you can store in your apps and services. Once set to &#x60;false&#x60; it cannot be changed to &#x60;true&#x60;. | [optional] [default to true]
**Id** | **string** | The v4 UUID used to uniquely identify the resource, as specified by RFC 4122. | [optional] [readonly] 
**Imported** | **bool** | A boolean value that shows whether your key was originally imported or generated in Key Protect. The value is set by Key Protect based on how the key material is initially added to the service.    A value of &#x60;true&#x60; indicates that you must provide new key material when it&#39;s time to rotate the key. A value  of &#x60;false&#x60; indicates that Key Protect will generate the new key material on a &#x60;rotate&#x60; operation, as it did in key creation. | [optional] [readonly] [default to false]
**Name** | **string** | A unique, human-readable alias to assign to your key.    To protect your privacy, do not use personal data, such as your name or location, as an alias for your key. | 
**State** | **int32** | The key state based on NIST SP 800-57. States are integers and correspond to the Pre-activation &#x3D; 0, Active &#x3D; 1, Deactivated &#x3D; 3, and Destroyed &#x3D; 5 values. | [optional] [readonly] 
**Type** | **string** | Specifies the MIME type that represents the key resource. Currently, only the default is supported. | [default to application/vnd.ibm.kms.key+json]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


