package models

type CheckList struct {
	Collection []Check `json:"results"`
}

type Check struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Md5        string `json:"md5"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
}

// has the page been updated?
func (self *Check) Changed(md5 string) bool {
	if self.Md5 != md5 {
		return true
	}
	return false
}
