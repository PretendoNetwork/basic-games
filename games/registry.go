package games

import (
	"fmt"
	"sort"
)

var registry = map[string]*Game{
	tetrisAxis.Slug: &tetrisAxis,
}

// Get returns the game with the given slug
func Get(slug string) (*Game, error) {
	game, exists := registry[slug]
	if !exists {
		return nil, fmt.Errorf("unknown game %q, known games are %v", slug, Slugs())
	}

	return game, nil
}

// Slugs returns every known game slug
func Slugs() []string {
	slugs := make([]string, 0, len(registry))
	for slug := range registry {
		slugs = append(slugs, slug)
	}

	sort.Strings(slugs)

	return slugs
}
