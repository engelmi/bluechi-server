/* SPDX-License-Identifier: LGPL-2.1-or-later */

package main

import (
	"os"

	"github.com/engelmi/bluechi-server/pkg/cmd/server"
	genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/component-base/cli"
)

func main() {
	stopCh := genericapiserver.SetupSignalHandler()
	options := server.NewBlueChiServerOptions(os.Stdout, os.Stderr)
	cmd := server.NewCommandStartBlueChiServer(options, stopCh)
	code := cli.Run(cmd)
	os.Exit(code)
}
