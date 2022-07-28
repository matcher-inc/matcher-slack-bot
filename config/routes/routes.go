package routes

import (
	"go-bot-test/features/deploy"
	"go-bot-test/lib/feature"
)

type Route struct {
	Path    string
	Feature feature.Feature
}

var Rounting = []Route{
	{Path: "deploy", Feature: deploy.Feature},
}
