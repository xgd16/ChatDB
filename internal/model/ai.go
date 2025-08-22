package model

type AIModelItem struct {
	ID         string                  `json:"id"`
	Object     string                  `json:"object"`
	Created    int64                   `json:"created"`
	OwnedBy    string                  `json:"owned_by"`
	Permission []AIModelItemPermission `json:"permission"`
	Root       string                  `json:"root"`
	Parent     interface{}             `json:"parent"`
}

type AIModelItemPermission struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Created            int64       `json:"created"`
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	Organization       string      `json:"organization"`
	Group              interface{} `json:"group"`
	IsBlocking         bool        `json:"is_blocking"`
}
