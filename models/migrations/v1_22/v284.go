// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_22 //nolint
import (
	"xorm.io/xorm"
)

func AddIgnoreStaleApprovalsColumnToProtectedBranchTable(x *xorm.Engine) error {
	type ProtectedBranch struct {
		IgnoreStaleApprovals bool `xorm:"NOT NULL DEFAULT false"`
	}
	return x.Sync(new(ProtectedBranch))
}
