/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package datastore

import (
	"errors"
	"flag"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/cli"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/govc/flags"
	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/vmware/govmomi/object"
	"github.com/apcera/libretto/Godeps/_workspace/src/golang.org/x/net/context"
)

type cp struct {
	*flags.OutputFlag
	*flags.DatastoreFlag

	force bool
}

func init() {
	cli.Register("datastore.cp", &cp{})
}

func (cmd *cp) Register(f *flag.FlagSet) {
	f.BoolVar(&cmd.force, "f", false, "If true, overwrite any identically named file at the destination")
}

func (cmd *cp) Process() error { return nil }

func (cmd *cp) Usage() string {
	return "SRC DST"
}

func (cmd *cp) Run(f *flag.FlagSet) error {
	args := f.Args()
	if len(args) != 2 {
		return errors.New("SRC and DST arguments are required")
	}

	c, err := cmd.Client()
	if err != nil {
		return err
	}

	dc, err := cmd.Datacenter()
	if err != nil {
		return err
	}

	// TODO: support cross-datacenter copy

	src, err := cmd.DatastorePath(args[0])
	if err != nil {
		return err
	}

	dst, err := cmd.DatastorePath(args[1])
	if err != nil {
		return err
	}

	m := object.NewFileManager(c)
	task, err := m.CopyDatastoreFile(context.TODO(), src, dc, dst, dc, cmd.force)
	if err != nil {
		return err
	}

	return task.Wait(context.TODO())
}
