/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

//  CPU runtime statistics
type ProcessorStats struct {
	TotalRuntime100ns int32 `json:"TotalRuntime100ns,omitempty"`

	RuntimeUser100ns int32 `json:"RuntimeUser100ns,omitempty"`

	RuntimeKernel100ns int32 `json:"RuntimeKernel100ns,omitempty"`
}
