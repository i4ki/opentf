// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opentf

import (
	"github.com/placeholderplaceholderplaceholder/opentf/internal/addrs"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/configs/configschema"
	"github.com/placeholderplaceholderplaceholder/opentf/internal/providers"
)

func simpleTestSchemas() *Schemas {
	provider := simpleMockProvider()
	provisioner := simpleMockProvisioner()

	return &Schemas{
		Providers: map[addrs.Provider]providers.ProviderSchema{
			addrs.NewDefaultProvider("test"): provider.GetProviderSchema(),
		},
		Provisioners: map[string]*configschema.Block{
			"test": provisioner.GetSchemaResponse.Provisioner,
		},
	}
}

// schemaOnlyProvidersForTesting is a testing helper that constructs a
// plugin library that contains a set of providers that only know how to
// return schema, and will exhibit undefined behavior if used for any other
// purpose.
//
// The intended use for this is in testing components that use schemas to
// drive other behavior, such as reference analysis during graph construction,
// but that don't actually need to interact with providers otherwise.
func schemaOnlyProvidersForTesting(schemas map[addrs.Provider]providers.ProviderSchema) *contextPlugins {
	factories := make(map[addrs.Provider]providers.Factory, len(schemas))

	for providerAddr, schema := range schemas {
		schema := schema

		provider := &MockProvider{
			GetProviderSchemaResponse: &schema,
		}

		factories[providerAddr] = func() (providers.Interface, error) {
			return provider, nil
		}
	}

	return newContextPlugins(factories, nil)
}
