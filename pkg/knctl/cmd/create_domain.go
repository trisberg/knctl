/*
Copyright 2018 The Knative Authors

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

package cmd

import (
	"fmt"

	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/spf13/cobra"
)

type CreateDomainOptions struct {
	ui          ui.UI
	depsFactory DepsFactory

	Domain  string
	Default bool
}

func NewCreateDomainOptions(ui ui.UI, depsFactory DepsFactory) *CreateDomainOptions {
	return &CreateDomainOptions{ui: ui, depsFactory: depsFactory}
}

func NewCreateDomainCmd(o *CreateDomainOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "domain",
		Aliases: []string{"dom"},
		Short:   "Create domain",
		Example: `
  # Create domain 'example.com' and set it as default
  knctl create domain -d example.com --default`,
		RunE: func(_ *cobra.Command, _ []string) error { return o.Run() },
	}

	cmd.Flags().StringVarP(&o.Domain, "domain", "d", "", "Specified domain (example: domain.com)")
	cmd.MarkFlagRequired("domain")

	cmd.Flags().BoolVar(&o.Default, "default", false, "Set domain as default (currently required to be provided)")
	cmd.MarkFlagRequired("default")

	return cmd
}

func (o *CreateDomainOptions) Run() error {
	if !o.Default {
		return fmt.Errorf("Currently --default flag is required")
	}

	coreClient, err := o.depsFactory.CoreClient()
	if err != nil {
		return err
	}

	domains := NewDomains(coreClient)
	if err != nil {
		return err
	}

	return domains.Create(Domain{Name: o.Domain, Default: o.Default})
}
