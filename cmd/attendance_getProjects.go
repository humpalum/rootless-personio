// SPDX-FileCopyrightText: 2022 Risk.Ident GmbH <contact@riskident.com>
//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var attendanceGetProjects = &cobra.Command{
	Use:     "getProjects",
	Short:   "Prints available projects",
	Long:    `Prints available projects`,
	Example: `getProjects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("expected no argument, got %d", len(args))
		}

		client, err := newLoggedInClient()
		if err != nil {
			return err
		}
		projects, err := client.GetProjects()
		if err != nil {
			return fmt.Errorf("failed to get projects: %w", err)
		}
		if len(projects) == 0 {
			return errors.New("no projects found")
		}

		projectNames := make([]string, len(projects))
		for i, project := range projects {
			projectNames[i] = project.Attributes.Name
		}

		return printOutputJSONOrYAML(projectNames)
	},
}

func init() {
	attendanceCmd.AddCommand(attendanceGetProjects)

}
