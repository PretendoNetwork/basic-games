package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// temp global storage for tracking any issues when loading the environment variables
type env struct {
	prefix   string
	problems []string
}

// lookup returns an environment variables value based on the selected games prefix
func (e *env) lookup(name string) string {
	if value, ok := os.LookupEnv(fmt.Sprintf("%s_%s", e.prefix, name)); ok {
		return strings.TrimSpace(value)
	}

	return ""
}

// name returns the fully qualified environment variable name
func (e *env) name(name string) string {
	return fmt.Sprintf("%s_%s", e.prefix, name)
}

// requiredString reads a string from an environment variable that MUST be set
func (e *env) requiredString(name string) string {
	value := e.lookup(name)
	if value == "" {
		e.problems = append(e.problems, fmt.Sprintf("%s is required but not set", e.name(name)))
	}

	return value
}

// requiredPort reads a port from an environment variable that MUST be set
func (e *env) requiredPort(name string) int {
	value := e.lookup(name)
	if value == "" {
		e.problems = append(e.problems, fmt.Sprintf("%s is required but not set", e.name(name)))
		return 0
	}

	return e.port(name, value)
}

// optionalPort reads a port from an environment variable that may be unset
func (e *env) optionalPort(name string) int {
	value := e.lookup(name)
	if value == "" {
		return 0
	}

	return e.port(name, value)
}

// port converts a string to a port number
func (e *env) port(name, value string) int {
	port, err := strconv.Atoi(value)
	if err != nil || port < 1 || port > 65535 {
		e.problems = append(e.problems, fmt.Sprintf("%s is not a valid port, expected 1-65535, got %q", e.name(name), value))
		return 0
	}

	return port
}

// err returns every problem found, or nil if all good
func (e *env) err() error {
	if len(e.problems) == 0 {
		return nil
	}

	return fmt.Errorf("invalid environment:\n  - %s", strings.Join(e.problems, "\n  - "))
}
