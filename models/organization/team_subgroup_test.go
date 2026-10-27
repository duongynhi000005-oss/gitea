// Copyright 2026 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package organization

import (
	"errors"
	"testing"

	"gitea.dev/modules/util"
	"github.com/stretchr/testify/assert"
)

func TestSubgroupHierarchyValidation(t *testing.T) {
	err := ValidateParentTeam(t.Context(), 10, 5, 0)
	assert.NoError(t, err)

	err = ValidateParentTeam(t.Context(), 10, 5, -1)
	assert.Error(t, err)
	assert.True(t, IsErrSubgroupHierarchy(err))
	assert.True(t, errors.Is(err, util.ErrInvalidArgument))

	err = ValidateParentTeam(t.Context(), 10, 5, 5)
	assert.Error(t, err)
	assert.True(t, IsErrSubgroupHierarchy(err))
}
