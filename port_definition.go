package marathon

// PortDefinition is the definition for one application portDefinitions in Marathon
type PortDefinition struct {
	Port     int               `json:"port,omitempty"`
	Protocol string            `json:"protocol,omitempty"`
	Name     string            `json:"name,omitempty"`
	Labels   map[string]string `json:"labels,omitempty"`
}
