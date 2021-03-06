/*
Copyright 2018 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfs

import (
	"github.com/rook/nfs/cmd/rook/rook"
	operator "github.com/rook/nfs/pkg/operator/nfs"
	"github.com/rook/nfs/pkg/util/flags"
	"github.com/spf13/cobra"
)

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "Runs the NFS operator to deploy and manage NFS server in kubernetes clusters",
	Long: `Runs the NFS operator to deploy and manage NFS server in kubernetes clusters.
https://github.com/rook/nfs`,
}

func init() {
	flags.SetFlagsFromEnv(operatorCmd.Flags(), rook.RookEnvVarPrefix)
	flags.SetLoggingFlags(operatorCmd.Flags())

	operatorCmd.RunE = startOperator
}

func startOperator(cmd *cobra.Command, args []string) error {
	rook.SetLogLevel()
	rook.LogStartupInfo(operatorCmd.Flags())

	logger.Infof("starting NFS operator")
	context := rook.NewContext()
	op := operator.New(context)
	err := op.Run()
	rook.TerminateOnError(err, "failed to run operator")

	return nil
}
