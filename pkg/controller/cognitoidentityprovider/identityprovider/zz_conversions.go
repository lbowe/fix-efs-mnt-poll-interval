/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package identityprovider

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cognitoidentityprovider/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeIdentityProviderInput returns input for read
// operation.
func GenerateDescribeIdentityProviderInput(cr *svcapitypes.IdentityProvider) *svcsdk.DescribeIdentityProviderInput {
	res := &svcsdk.DescribeIdentityProviderInput{}

	if cr.Status.AtProvider.ProviderName != nil {
		res.SetProviderName(*cr.Status.AtProvider.ProviderName)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// GenerateIdentityProvider returns the current state in the form of *svcapitypes.IdentityProvider.
func GenerateIdentityProvider(resp *svcsdk.DescribeIdentityProviderOutput) *svcapitypes.IdentityProvider {
	cr := &svcapitypes.IdentityProvider{}

	if resp.IdentityProvider.AttributeMapping != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range resp.IdentityProvider.AttributeMapping {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		cr.Spec.ForProvider.AttributeMapping = f0
	} else {
		cr.Spec.ForProvider.AttributeMapping = nil
	}
	if resp.IdentityProvider.CreationDate != nil {
		cr.Status.AtProvider.CreationDate = &metav1.Time{*resp.IdentityProvider.CreationDate}
	} else {
		cr.Status.AtProvider.CreationDate = nil
	}
	if resp.IdentityProvider.IdpIdentifiers != nil {
		f2 := []*string{}
		for _, f2iter := range resp.IdentityProvider.IdpIdentifiers {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		cr.Spec.ForProvider.IDpIdentifiers = f2
	} else {
		cr.Spec.ForProvider.IDpIdentifiers = nil
	}
	if resp.IdentityProvider.LastModifiedDate != nil {
		cr.Status.AtProvider.LastModifiedDate = &metav1.Time{*resp.IdentityProvider.LastModifiedDate}
	} else {
		cr.Status.AtProvider.LastModifiedDate = nil
	}
	if resp.IdentityProvider.ProviderName != nil {
		cr.Status.AtProvider.ProviderName = resp.IdentityProvider.ProviderName
	} else {
		cr.Status.AtProvider.ProviderName = nil
	}
	if resp.IdentityProvider.ProviderType != nil {
		cr.Spec.ForProvider.ProviderType = resp.IdentityProvider.ProviderType
	} else {
		cr.Spec.ForProvider.ProviderType = nil
	}
	if resp.IdentityProvider.UserPoolId != nil {
		cr.Status.AtProvider.UserPoolID = resp.IdentityProvider.UserPoolId
	} else {
		cr.Status.AtProvider.UserPoolID = nil
	}

	return cr
}

// GenerateCreateIdentityProviderInput returns a create input.
func GenerateCreateIdentityProviderInput(cr *svcapitypes.IdentityProvider) *svcsdk.CreateIdentityProviderInput {
	res := &svcsdk.CreateIdentityProviderInput{}

	if cr.Spec.ForProvider.AttributeMapping != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range cr.Spec.ForProvider.AttributeMapping {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		res.SetAttributeMapping(f0)
	}
	if cr.Spec.ForProvider.IDpIdentifiers != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.IDpIdentifiers {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetIdpIdentifiers(f1)
	}
	if cr.Spec.ForProvider.ProviderType != nil {
		res.SetProviderType(*cr.Spec.ForProvider.ProviderType)
	}

	return res
}

// GenerateUpdateIdentityProviderInput returns an update input.
func GenerateUpdateIdentityProviderInput(cr *svcapitypes.IdentityProvider) *svcsdk.UpdateIdentityProviderInput {
	res := &svcsdk.UpdateIdentityProviderInput{}

	if cr.Spec.ForProvider.AttributeMapping != nil {
		f0 := map[string]*string{}
		for f0key, f0valiter := range cr.Spec.ForProvider.AttributeMapping {
			var f0val string
			f0val = *f0valiter
			f0[f0key] = &f0val
		}
		res.SetAttributeMapping(f0)
	}
	if cr.Spec.ForProvider.IDpIdentifiers != nil {
		f1 := []*string{}
		for _, f1iter := range cr.Spec.ForProvider.IDpIdentifiers {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetIdpIdentifiers(f1)
	}
	if cr.Status.AtProvider.ProviderName != nil {
		res.SetProviderName(*cr.Status.AtProvider.ProviderName)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// GenerateDeleteIdentityProviderInput returns a deletion input.
func GenerateDeleteIdentityProviderInput(cr *svcapitypes.IdentityProvider) *svcsdk.DeleteIdentityProviderInput {
	res := &svcsdk.DeleteIdentityProviderInput{}

	if cr.Status.AtProvider.ProviderName != nil {
		res.SetProviderName(*cr.Status.AtProvider.ProviderName)
	}
	if cr.Status.AtProvider.UserPoolID != nil {
		res.SetUserPoolId(*cr.Status.AtProvider.UserPoolID)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}