// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"xorm.io/xorm"
)

func addEstimatedTimeOnIssue(x *xorm.Engine) error {
	type Issue struct {
		EstimatedTime int64 `xorm:"INTEGER default NULL"`
	}

	return x.Sync2(new(Issue))
}
