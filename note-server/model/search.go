package model

type ErrorResponse struct {
	Info *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	RootCause []*ErrorInfo
	Type      string
	Reason    string
	Phase     string
}

type SearchCollectResponse struct {
	Took int64 `json:"took"`
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*SearchCollectHit `json:"hits"`
	} `json:"hits"`
}

type SearchTagResponse struct {
	Took int64 `json:"took"`
	Hits struct {
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
		Hits []*SearchTagHit `json:"hits"`
	} `json:"hits"`
}

type Highlight struct {
	Title    []string `json:"title,omitempty"`
	Fulltext []string `json:"full_text,omitempty"`
}

type SearchCollectHit struct {
	Score float64 `json:"_score"`
	Index string  `json:"_index"`
	Type  string  `json:"_type"`

	Source    UserIdTagIdsCollect `json:"_source"`
	Highlight Highlight           `json:"highlight,omitempty"`
}

type SearchTagHit struct {
	Score float64 `json:"_score"`
	Index string  `json:"_index"`
	Type  string  `json:"_type"`

	Source UserIdTagSearch `json:"_source"`
}
