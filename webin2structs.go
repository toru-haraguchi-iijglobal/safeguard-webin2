package main

type Definition struct {
	Asset          string   `json:"asset" yaml:"asset"`                     // -asset=[asset name of SPP]
	UseEdge        bool     `json:"use_edge" yaml:"use_edge"`               // Use MS Edge instead of Google Chrome
	Secret         bool     `json:"secret" yaml:"secret"`                   // Open with Secret Window
	CertValidation bool     `json:"cert_validation" yaml:"cert_validation"` // Do or Don't Cert Validation
	Actions        []Action `json:"actions" yaml:"actions"`                 // array of actions to perform
}

type Action struct {
	Type   string `json:"type" yaml:"type"`     // chromedp action type
	Target string `json:"target" yaml:"target"` // url or element selector
	Value  int    `json:"value" yaml:"value"`   // value to set or click
}
