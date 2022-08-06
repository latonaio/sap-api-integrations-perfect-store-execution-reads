package sap_api_output_formatter

type PerfectStoreExecution struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	APISchema     string `json:"api_schema"`
	StoreValuationCode  string `json:"store_valuation_code"`
	Deleted       bool   `json:"deleted"`
}

type StoreValuationCollection struct {
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
}