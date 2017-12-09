package banjo

// Config struct
//
// Allows you to create configuration to your banjo application
//
type Config struct {
	port  string
	host  string
	debug bool
}

// DefaultHost is default application host value
const DefaultHost = "127.0.0.1"

// DefaultPort is default application port value
const DefaultPort = "4321"

// DefaultConfig function
//
// Returns default configurations for
// Banjo application
//
// Params:
// - None
//
// Response:
// - config {Config} Config struct
//
func DefaultConfig() Config {
	return Config{
		port:  DefaultPort,
		host:  DefaultHost,
		debug: false,
	}
}
