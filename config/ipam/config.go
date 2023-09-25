package ipam

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("netbox_prefix", func(r *config.Resource) {
		r.ShortGroup = "ipam.netbox"
		r.UseAsync = false
	})

	p.AddResourceConfigurator("netbox_available_prefix", func(r *config.Resource) {
		r.ShortGroup = "ipam.netbox"
		r.Kind = "AvailablePrefix"
		r.UseAsync = false
	})
}
