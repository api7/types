package v1

import "encoding/json"

// Route apisix route object
// +k8s:deepcopy-gen=true
type Route struct {
	ID           *string   `json:"id,omitempty" yaml:"id,omitempty"`
	Host         *string   `json:"host,omitempty" yaml:"host,omitempty"`
	Path         *string   `json:"path,omitempty" yaml:"path,omitempty"`
	Name         *string   `json:"name,omitempty" yaml:"name,omitempty"`
	Methods      []*string `json:"methods,omitempty" yaml:"methods,omitempty"`
	ServiceId    *string   `json:"service_id,omitempty" yaml:"service_id,omitempty"`
	ServiceName  *string   `json:"service_name,omitempty" yaml:"service_name,omitempty"`
	UpstreamId   *string   `json:"upstream_id,omitempty" yaml:"upstream_id,omitempty"`
	UpstreamName *string   `json:"upstream_name,omitempty" yaml:"upstream_name,omitempty"`
	Plugins      *Plugins  `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

type Plugins map[string]interface{}

func (p *Plugins) DeepCopyInto(out *Plugins) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *Plugins) DeepCopy() *Plugins {
	if p == nil {
		return nil
	}
	out := new(Plugins)
	p.DeepCopyInto(out)
	return out
}

// Service apisix service
// +k8s:deepcopy-gen=true
type Service struct {
	ID           *string  `json:"id,omitempty" yaml:"id,omitempty"`
	Name         *string  `json:"name,omitempty" yaml:"name,omitempty"`
	UpstreamId   *string  `json:"upstream_id,omitempty" yaml:"upstream_id,omitempty"`
	UpstreamName *string  `json:"upstream_name,omitempty" yaml:"upstream_name,omitempty"`
	Plugins      *Plugins `json:"plugins,omitempty" yaml:"plugins,omitempty"`
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
