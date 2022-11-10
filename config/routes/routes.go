package routes

import (
	"errors"
	"go-bot-test/features/deploy"
	"go-bot-test/features/dialog"
	"go-bot-test/lib/feature"
)

type Route struct {
	Path    string
	Feature feature.Feature
}

var rounting = []Route{
	{Path: "deployy", Feature: deploy.Feature},
	{Path: "dialog", Feature: dialog.Feature},
}

func GetRoute(path string) (*Route, error) {
	for _, route := range rounting {
		if route.Path == path {
			return &route, nil
		}
	}
	return nil, errors.New("404 Not found")
}
