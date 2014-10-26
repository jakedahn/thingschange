package models

import (
	"fmt"
	"sync"
)

type Check struct {
	id         string
	url        string
	md5        string
	content    string
	created_at int64
	updated_at int64
}

// has the page been updated?
func (self *Check) Changed(md5 string) bool {
	if self.md5 != md5 {
		return true
	}
	return false
}
