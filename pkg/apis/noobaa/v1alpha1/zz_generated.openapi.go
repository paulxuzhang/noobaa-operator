// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaa":       schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec":   schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus": schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref),
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaa is the Schema for the NooBaas API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Description: "Standard object metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaSpec defines the desired state of System",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image (optional) overrides the default image for server pods",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaStatus defines the observed state of System",
				Properties: map[string]spec.Schema{
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration is the most recent generation observed for this noobaa system. It corresponds to the CR generation, which is updated on mutation by the API Server.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"actualImage": {
						SchemaProps: spec.SchemaProps{
							Description: "ActualImage is set to report which image the operator is using",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accounts": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus"),
						},
					},
					"services": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus"),
						},
					},
					"readme": {
						SchemaProps: spec.SchemaProps{
							Description: "Readme is a user readable string with explanations on the system",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"observedGeneration", "phase", "actualImage", "accounts", "services", "readme"},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus"},
	}
}
