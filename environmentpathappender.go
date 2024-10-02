// Package environmentpathappender a plugin to use environment variables in headers
package environmentpathappender

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	Env string `json:"env,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type environmentPathAppenderPlugin struct {
	EnvironmentVar string
	next           http.Handler
}

// New creates a new EnvironmentHeader plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	environmentVarName := config.Env
	if config.Env == "" {
		return nil, fmt.Errorf("missing env parameter %v", environmentVarName)
	}

	var environmentVar = os.Getenv(environmentVarName)
	if environmentVar == "" {
		slog.Warn("missing env variable " + environmentVarName)
	}

	if strings.Contains(environmentVar, "/") {
		return nil, fmt.Errorf("no / allowed in the path addition: %v", environmentVar)
	}

	return &environmentPathAppenderPlugin{
		EnvironmentVar: environmentVar,
		next:           next,
	}, nil
}

func (c *environmentPathAppenderPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if c.EnvironmentVar != "" {
		req.URL.Path = req.URL.Path + "/" + c.EnvironmentVar
	}
	c.next.ServeHTTP(rw, req)
}
