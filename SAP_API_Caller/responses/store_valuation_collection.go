package responses

type StoreValuationCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID             string `json:"ObjectID"`
			ETag                 string `json:"ETag"`
			TotalScoreValue      string `json:"TotalScoreValue"`
			ReferenceVisitUUID   string `json:"ReferenceVisitUUID"`
			StoreLayoutName      string `json:"StoreLayoutName"`
			ValidFromDate        string `json:"ValidFromDate"`
			ValidToDate          string `json:"ValidToDate"`
			ReferenceVisitID     string `json:"ReferenceVisitID"`
			StoreLayoutID        string `json:"StoreLayoutID"`
			StoreLayoutVersion   string `json:"StoreLayoutVersion"`
			EntityLastChangedOn  string `json:"EntityLastChangedOn"`
		} `json:"results"`
	} `json:"d"`
}
