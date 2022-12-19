/*
 * Api Documentation
 *
 * Api Documentation
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VisitDtoReq struct {
	Comment string `json:"comment,omitempty"`
	Date *TimestampReq `json:"date,omitempty"`
	Id int64 `json:"id,omitempty"`
	ServiceInfoId int64 `json:"serviceInfoId,omitempty"`
	UserId int64 `json:"userId,omitempty"`
}