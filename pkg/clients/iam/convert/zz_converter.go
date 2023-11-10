// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.

package convert

import (
	types "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"time"
)

type ConverterImpl struct{}

func (c *ConverterImpl) DeepCopyAWSRole(source *types.Role) *types.Role {
	var pTypesRole *types.Role
	if source != nil {
		var typesRole types.Role
		var pString *string
		if (*source).Arn != nil {
			xstring := *(*source).Arn
			pString = &xstring
		}
		typesRole.Arn = pString
		typesRole.CreateDate = c.pTimeTimeToPTimeTime((*source).CreateDate)
		var pString2 *string
		if (*source).Path != nil {
			xstring2 := *(*source).Path
			pString2 = &xstring2
		}
		typesRole.Path = pString2
		var pString3 *string
		if (*source).RoleId != nil {
			xstring3 := *(*source).RoleId
			pString3 = &xstring3
		}
		typesRole.RoleId = pString3
		var pString4 *string
		if (*source).RoleName != nil {
			xstring4 := *(*source).RoleName
			pString4 = &xstring4
		}
		typesRole.RoleName = pString4
		var pString5 *string
		if (*source).AssumeRolePolicyDocument != nil {
			xstring5 := *(*source).AssumeRolePolicyDocument
			pString5 = &xstring5
		}
		typesRole.AssumeRolePolicyDocument = pString5
		var pString6 *string
		if (*source).Description != nil {
			xstring6 := *(*source).Description
			pString6 = &xstring6
		}
		typesRole.Description = pString6
		var pInt32 *int32
		if (*source).MaxSessionDuration != nil {
			xint32 := *(*source).MaxSessionDuration
			pInt32 = &xint32
		}
		typesRole.MaxSessionDuration = pInt32
		typesRole.PermissionsBoundary = c.pTypesAttachedPermissionsBoundaryToPTypesAttachedPermissionsBoundary((*source).PermissionsBoundary)
		typesRole.RoleLastUsed = c.pTypesRoleLastUsedToPTypesRoleLastUsed((*source).RoleLastUsed)
		var typesTagList []types.Tag
		if (*source).Tags != nil {
			typesTagList = make([]types.Tag, len((*source).Tags))
			for i := 0; i < len((*source).Tags); i++ {
				typesTagList[i] = c.typesTagToTypesTag((*source).Tags[i])
			}
		}
		typesRole.Tags = typesTagList
		pTypesRole = &typesRole
	}
	return pTypesRole
}
func (c *ConverterImpl) pTimeTimeToPTimeTime(source *time.Time) *time.Time {
	var pTimeTime *time.Time
	if source != nil {
		var timeTime time.Time
		_ = (*source)
		pTimeTime = &timeTime
	}
	return pTimeTime
}
func (c *ConverterImpl) pTypesAttachedPermissionsBoundaryToPTypesAttachedPermissionsBoundary(source *types.AttachedPermissionsBoundary) *types.AttachedPermissionsBoundary {
	var pTypesAttachedPermissionsBoundary *types.AttachedPermissionsBoundary
	if source != nil {
		var typesAttachedPermissionsBoundary types.AttachedPermissionsBoundary
		var pString *string
		if (*source).PermissionsBoundaryArn != nil {
			xstring := *(*source).PermissionsBoundaryArn
			pString = &xstring
		}
		typesAttachedPermissionsBoundary.PermissionsBoundaryArn = pString
		typesAttachedPermissionsBoundary.PermissionsBoundaryType = types.PermissionsBoundaryAttachmentType((*source).PermissionsBoundaryType)
		pTypesAttachedPermissionsBoundary = &typesAttachedPermissionsBoundary
	}
	return pTypesAttachedPermissionsBoundary
}
func (c *ConverterImpl) pTypesRoleLastUsedToPTypesRoleLastUsed(source *types.RoleLastUsed) *types.RoleLastUsed {
	var pTypesRoleLastUsed *types.RoleLastUsed
	if source != nil {
		var typesRoleLastUsed types.RoleLastUsed
		typesRoleLastUsed.LastUsedDate = c.pTimeTimeToPTimeTime((*source).LastUsedDate)
		var pString *string
		if (*source).Region != nil {
			xstring := *(*source).Region
			pString = &xstring
		}
		typesRoleLastUsed.Region = pString
		pTypesRoleLastUsed = &typesRoleLastUsed
	}
	return pTypesRoleLastUsed
}
func (c *ConverterImpl) typesTagToTypesTag(source types.Tag) types.Tag {
	var typesTag types.Tag
	var pString *string
	if source.Key != nil {
		xstring := *source.Key
		pString = &xstring
	}
	typesTag.Key = pString
	var pString2 *string
	if source.Value != nil {
		xstring2 := *source.Value
		pString2 = &xstring2
	}
	typesTag.Value = pString2
	return typesTag
}