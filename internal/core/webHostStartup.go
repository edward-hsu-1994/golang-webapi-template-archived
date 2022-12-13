package core

// WebHostStartup requires `ConfigureServices` and `Configure`
type WebHostStartup struct {
	ConfigureServicesFunc any
	ConfigureFunc         any
}
