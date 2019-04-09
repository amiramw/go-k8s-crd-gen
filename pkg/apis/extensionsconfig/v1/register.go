package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/amiramw/go-k8s-crd-gen/pkg/apis/extensionsconfig"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   extensionsconfig.GroupName,
	Version: "v1",
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// addKnownTypes adds our types to the API scheme by registering
// ExtensionConfig type and its list
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&ExtensionsConfig{},
		&ExtensionsConfigList{},
	)

	// register the type in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
