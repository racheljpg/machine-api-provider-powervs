// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/client/authentication"
	"github.com/IBM-Cloud/power-go-client/power/client/bluemix_service_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/catalog"
	"github.com/IBM-Cloud/power-go-client/power/client/hardware_platforms"
	"github.com/IBM-Cloud/power-go-client/power/client/iaas_service_broker"
	"github.com/IBM-Cloud/power-go-client/power/client/open_stacks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_cloud_connections"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_events"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_images"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_networks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_placement_groups"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_s_a_p"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_snapshots"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_storage_capacity"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_system_pools"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tasks"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_tenants_ssh_keys"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_volumes"
	"github.com/IBM-Cloud/power-go-client/power/client/service_bindings"
	"github.com/IBM-Cloud/power-go-client/power/client/service_instances"
	"github.com/IBM-Cloud/power-go-client/power/client/storage_types"
	"github.com/IBM-Cloud/power-go-client/power/client/swagger_spec"
)

// Default power iaas HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new power iaas HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PowerIaas {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new power iaas HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PowerIaas {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new power iaas client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PowerIaas {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PowerIaas)
	cli.Transport = transport

	cli.Authentication = authentication.New(transport, formats)

	cli.BluemixServiceInstances = bluemix_service_instances.New(transport, formats)

	cli.Catalog = catalog.New(transport, formats)

	cli.HardwarePlatforms = hardware_platforms.New(transport, formats)

	cli.IaasServiceBroker = iaas_service_broker.New(transport, formats)

	cli.OpenStacks = open_stacks.New(transport, formats)

	cli.PCloudCloudConnections = p_cloud_cloud_connections.New(transport, formats)

	cli.PCloudEvents = p_cloud_events.New(transport, formats)

	cli.PCloudImages = p_cloud_images.New(transport, formats)

	cli.PCloudInstances = p_cloud_instances.New(transport, formats)

	cli.PCloudNetworks = p_cloud_networks.New(transport, formats)

	cli.PCloudPVMInstances = p_cloud_p_vm_instances.New(transport, formats)

	cli.PCloudPlacementGroups = p_cloud_placement_groups.New(transport, formats)

	cli.PCloudSAP = p_cloud_s_a_p.New(transport, formats)

	cli.PCloudSnapshots = p_cloud_snapshots.New(transport, formats)

	cli.PCloudStorageCapacity = p_cloud_storage_capacity.New(transport, formats)

	cli.PCloudSystemPools = p_cloud_system_pools.New(transport, formats)

	cli.PCloudTasks = p_cloud_tasks.New(transport, formats)

	cli.PCloudTenants = p_cloud_tenants.New(transport, formats)

	cli.PCloudTenantsSSHKeys = p_cloud_tenants_ssh_keys.New(transport, formats)

	cli.PCloudVolumes = p_cloud_volumes.New(transport, formats)

	cli.ServiceBindings = service_bindings.New(transport, formats)

	cli.ServiceInstances = service_instances.New(transport, formats)

	cli.StorageTypes = storage_types.New(transport, formats)

	cli.SwaggerSpec = swagger_spec.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PowerIaas is a client for power iaas
type PowerIaas struct {
	Authentication *authentication.Client

	BluemixServiceInstances *bluemix_service_instances.Client

	Catalog *catalog.Client

	HardwarePlatforms *hardware_platforms.Client

	IaasServiceBroker *iaas_service_broker.Client

	OpenStacks *open_stacks.Client

	PCloudCloudConnections *p_cloud_cloud_connections.Client

	PCloudEvents *p_cloud_events.Client

	PCloudImages *p_cloud_images.Client

	PCloudInstances *p_cloud_instances.Client

	PCloudNetworks *p_cloud_networks.Client

	PCloudPVMInstances *p_cloud_p_vm_instances.Client

	PCloudPlacementGroups *p_cloud_placement_groups.Client

	PCloudSAP *p_cloud_s_a_p.Client

	PCloudSnapshots *p_cloud_snapshots.Client

	PCloudStorageCapacity *p_cloud_storage_capacity.Client

	PCloudSystemPools *p_cloud_system_pools.Client

	PCloudTasks *p_cloud_tasks.Client

	PCloudTenants *p_cloud_tenants.Client

	PCloudTenantsSSHKeys *p_cloud_tenants_ssh_keys.Client

	PCloudVolumes *p_cloud_volumes.Client

	ServiceBindings *service_bindings.Client

	ServiceInstances *service_instances.Client

	StorageTypes *storage_types.Client

	SwaggerSpec *swagger_spec.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PowerIaas) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authentication.SetTransport(transport)

	c.BluemixServiceInstances.SetTransport(transport)

	c.Catalog.SetTransport(transport)

	c.HardwarePlatforms.SetTransport(transport)

	c.IaasServiceBroker.SetTransport(transport)

	c.OpenStacks.SetTransport(transport)

	c.PCloudCloudConnections.SetTransport(transport)

	c.PCloudEvents.SetTransport(transport)

	c.PCloudImages.SetTransport(transport)

	c.PCloudInstances.SetTransport(transport)

	c.PCloudNetworks.SetTransport(transport)

	c.PCloudPVMInstances.SetTransport(transport)

	c.PCloudPlacementGroups.SetTransport(transport)

	c.PCloudSAP.SetTransport(transport)

	c.PCloudSnapshots.SetTransport(transport)

	c.PCloudStorageCapacity.SetTransport(transport)

	c.PCloudSystemPools.SetTransport(transport)

	c.PCloudTasks.SetTransport(transport)

	c.PCloudTenants.SetTransport(transport)

	c.PCloudTenantsSSHKeys.SetTransport(transport)

	c.PCloudVolumes.SetTransport(transport)

	c.ServiceBindings.SetTransport(transport)

	c.ServiceInstances.SetTransport(transport)

	c.StorageTypes.SetTransport(transport)

	c.SwaggerSpec.SetTransport(transport)

}
