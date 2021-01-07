package queries

type EsQuery struct {
	Equals []struct{
		Field string      `json:"field"`
		Value interface{} `json:"value"`
	}
}


