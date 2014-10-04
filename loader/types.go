package loader

type Conditions struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type Expectations struct {
	StatusCode int `json:"status_code"`
}

type Dependency struct {
	Service string `json:"service"`
}

type Case struct {
	When  *Conditions   `json:"when"`
	Then  *Expectations `json:"then"`
	While []*Dependency `json:"while"`
}

type Endpoint struct {
	Name  string  `json:"name"`
	Cases []*Case `json:"cases"`
}

type Spec struct {
	Endpoints []*Endpoint `json:"endpoints"`
}
