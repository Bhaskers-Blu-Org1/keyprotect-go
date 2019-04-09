/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 2 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect

// The request for creating a transport encryption key.
type CreateLockerRequest struct {
	// The time in seconds from the creation of a transport key that determines how long the key remains valid.     The minimum value is `300` seconds (5 minutes), and the maximum value is `86400` (24 hours). The default value is `600` (10 minutes).
	Expiration float32 `json:"expiration,omitempty"`
	// The number of times that a transport key can be retrieved within its expiration time before it is no longer accessible. 
	MaxAllowedRetrievals float32 `json:"maxAllowedRetrievals,omitempty"`
}
