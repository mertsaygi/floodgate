/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ResponseEntity struct {
	StatusCode string `json:"statusCode,omitempty"`
	StatusCodeValue int32 `json:"statusCodeValue,omitempty"`
	Body *interface{} `json:"body,omitempty"`
}