package terraform

import "github.com/hashicorp/hcl"

type Parser struct{}

func (s *Parser) Unmarshal(p []byte, v interface{}) error {
	return hcl.Unmarshal(p, v)
}
