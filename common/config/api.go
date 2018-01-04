/*
Copyright IBM Corp. 2017 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/
package config

import (
	"github.com/hyperledger/fabric/common/resourcesconfig"
	cb "github.com/hyperledger/fabric/protos/common"
)

// Config encapsulates config (channel or resource) tree
type Config interface {
	// ConfigProto returns the current config
	ConfigProto() *cb.Config

	// ProposeConfigUpdate attempts to validate a new configtx against the current config state
	ProposeConfigUpdate(configtx *cb.Envelope) (*cb.ConfigEnvelope, error)
}

// Manager provides access to the resource config
type Manager interface {
	// GetChannelConfig defines methods that are related to channel configuration
	GetChannelConfig(channel string) Config

	// GetResourceConfig defines methods that are related to resource configuration
	GetResourceConfig(channel string) Config

	// GetPolicyMapper returns API to the policy mapper
	GetPolicyMapper(channel string) resourcesconfig.PolicyMapper
}
