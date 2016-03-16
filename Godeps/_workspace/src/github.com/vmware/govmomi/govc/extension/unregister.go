/*
Copyright (c) 2015 VMware, Inc. All Rights Reserved.

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

package extension

import (
	"flag"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/cli"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/flags"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/object"
	"github.com/apcera/libretto/Godeps/_workspace/src/golang.org/x/net/context"
)

type unregister struct {
	*flags.ClientFlag
}

func init() {
	cli.Register("extension.unregister", &unregister{})
}

func (cmd *unregister) Register(f *flag.FlagSet) {}

func (cmd *unregister) Process() error { return nil }

func (cmd *unregister) Run(f *flag.FlagSet) error {
	ctx := context.TODO()

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	m, err := object.GetExtensionManager(c)
	if err != nil {
		return err
	}

	for _, key := range f.Args() {
		if err = m.Unregister(ctx, key); err != nil {
			return err
		}
	}

	return nil
}
