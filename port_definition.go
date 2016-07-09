package marathon

// PortDefinition is the definition for one application portDefinitions in Marathon
type PortDefinition struct {
	Port     int
	Protocol string
	Name     string
	Labels   map[string]string
}
