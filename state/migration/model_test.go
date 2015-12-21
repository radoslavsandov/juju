// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package migration_test

import (
	gc "gopkg.in/check.v1"

	"github.com/juju/juju/state/migration"
	"github.com/juju/juju/testing"
)

type ModelSuite struct {
	testing.BaseSuite
}

var _ = gc.Suite(&ModelSuite{})

func (*ModelSuite) TestMissingVersion(c *gc.C) {
	_, err := migration.NewModel(nil)
	c.Check(err, gc.ErrorMatches, "missing 'version'")
	_, err = migration.NewModel(map[string]interface{}{})
	c.Check(err, gc.ErrorMatches, "missing 'version'")
}
