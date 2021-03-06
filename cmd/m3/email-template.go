// This file is part of MinIO Kubernetes Cloud
// Copyright (c) 2019 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"github.com/minio/cli"
)

// commands to interact with cluster admins
var emailTemplateCmd = cli.Command{
	Name:    "email-template",
	Aliases: []string{"et"},
	Usage:   "Email templates commands",
	Action:  emailTemplate,
	Subcommands: []cli.Command{
		emailSetCmd,
	},
}

func emailTemplate(ctx *cli.Context) error {
	return cli.ShowAppHelp(ctx)
}
