package core

import (
	"strings"

	"github.com/ShuBo6/cloudquery/pkg/config"
	"github.com/ShuBo6/cloudquery/pkg/plugin/registry"
)

func ParseProviderSource(requestedProvider *config.RequiredProvider) (org string, provider string, err error) {
	var requestedSource string
	if requestedProvider.Source == nil || *requestedProvider.Source == "" {
		requestedSource = requestedProvider.Name
	} else {
		requestedSource = *requestedProvider.Source
		if !strings.Contains(requestedSource, "/") {
			requestedSource = strings.Join([]string{requestedSource, requestedProvider.Name}, "/")
		}
	}
	return registry.ParseProviderName(requestedSource)
}
