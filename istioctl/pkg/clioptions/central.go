// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package clioptions

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// CentralControlPlaneOptions holds options common to all subcommands
// that invoke Istiod via xDS REST endpoint
type CentralControlPlaneOptions struct {
	// Xds is XDS endpoint, e.g. localhost:15010.
	Xds string

	// XdsPodLabel is a Kubernetes label on the Istiod pods
	XdsPodLabel string

	// XdsPodPort is a port exposing XDS (typically 15010 or 15012)
	XdsPodPort int

	// CertDir is the local directory containing certificates
	CertDir string

	// Timeout is how long to wait before giving up on XDS
	Timeout time.Duration
}

// AttachControlPlaneFlags attaches control-plane flags to a Cobra command.
// (Currently just --endpoint)
func (o *CentralControlPlaneOptions) AttachControlPlaneFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&o.Xds, "xds-address", "",
		"XDS Endpoint")
	cmd.PersistentFlags().StringVar(&o.CertDir, "cert-dir", "",
		"XDS Endpoint certificate directory")
	cmd.PersistentFlags().StringVar(&o.XdsPodLabel, "xds-label", "",
		"Istiod pod label selector")
	cmd.PersistentFlags().IntVar(&o.XdsPodPort, "xds-port", 15012,
		"Istiod pod port")
	cmd.PersistentFlags().DurationVar(&o.Timeout, "timeout", time.Second*30,
		"the duration to wait before failing")
}

// ValidateControlPlaneFlags checks arguments for valid values and combinations
func (o *CentralControlPlaneOptions) ValidateControlPlaneFlags() error {
	if o.Xds != "" && o.XdsPodLabel != "" {
		return fmt.Errorf("either --xds-address or --xds-label, not both")
	}
	return nil
}
