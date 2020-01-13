package v1

// Route apisix route object
// +k8s:deepcopy-gen=true
type Route struct {
	ID         *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Host       string    `json:"host,omitempty" yaml:"host,omitempty"`
	Path       *string   `json:"path,omitempty" yaml:"path,omitempty"`
	Name       *string   `json:"name,omitempty" yaml:"name,omitempty"`
	Methods    []*string `json:"methods,omitempty" yaml:"methods,omitempty"`
	ServiceId  *string   `json:"service_id,omitempty" yaml:"service_id,omitempty"`
	UpstreamId *string   `json:"upstream_id,omitempty" yaml:"upstream_id,omitempty"`
	Plugins    []*Plugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// Plugin customize plugin struct
// +k8s:deepcopy-gen=true
type Plugin struct {
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

// Service apisix service
// +k8s:deepcopy-gen=true
type Service struct {
	ID         *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Name       *string   `json:"name,omitempty" yaml:"name,omitempty"`
	UpstreamId *string   `json:"upstream_id,omitempty" yaml:"upstream_id,omitempty"`
	Plugins    []*Plugin `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// Upstream apisix upstream
// +k8s:deepcopy-gen=true
type Upstream struct {
	ID    *string `json:"id,omitempty" yaml:"id,omitempty"`
	Name  *string `json:"name,omitempty" yaml:"name,omitempty"`
	Type  *string `json:"type,omitempty" yaml:"type,omitempty"`
	Key   *string `json:"key,omitempty" yaml:"key,omitempty"`
	Nodes []*Node `json:"nodes,omitempty" yaml:"nodes,omitempty"`
}

// Node the node in upstream
// +k8s:deepcopy-gen=true
type Node struct {
	IP     *string `json:"ip,omitempty" yaml:"ip,omitempty"`
	Port   *int    `json:"port,omitempty" yaml:"port,omitempty"`
	Weight *int    `json:"weight,omitempty" yaml:"weight,omitempty"`
}
