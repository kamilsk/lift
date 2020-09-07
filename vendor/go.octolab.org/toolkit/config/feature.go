package config

import (
	"fmt"
	"strings"
)

// Feature describe a feature.
type Feature struct {
	ID      [16]byte
	Name    string
	Brief   string
	Docs    string
	RFC     string
	Enabled bool
}

// String returns a string representation of the feature.
func (feature Feature) String() string {
	return fmt.Sprintf("%s=%v", feature.Name, feature.Enabled)
}

// Features defines a list of features.
type Features []Feature

// String returns a string representation of the feature list.
func (features Features) String() string {
	if len(features) == 0 {
		return "-"
	}
	list := make([]string, 0, len(features))
	for _, feature := range features {
		list = append(list, feature.String())
	}
	return strings.Join(list, ", ")
}

// FindByID finds and returns Feature by passed ID.
// If nothing found it returns empty Feature.
func (features Features) FindByID(id [16]byte) Feature {
	for _, feature := range features {
		if feature.ID == id {
			return feature
		}
	}
	return Feature{}
}

// FindByName finds and returns Feature by passed name.
// If nothing found it returns empty Feature.
func (features Features) FindByName(name string) Feature {
	for _, feature := range features {
		if feature.Name == name {
			return feature
		}
	}
	return Feature{}
}
