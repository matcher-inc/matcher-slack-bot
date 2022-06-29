package routes

import "go-bot-test/lib/feature"

type Route struct {
	Path    string
	Feature feature.Feature
}

var Rounting = []Route{}
