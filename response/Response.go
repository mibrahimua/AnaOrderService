package response

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type ResponsePaginate struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Meta    Meta        `json:"meta,omitempty"`
	Links   Links       `json:"links,omitempty"`
}

// Meta contains pagination information
type Meta struct {
	Page         int `json:"page"`
	PerPage      int `json:"perPage"`
	TotalPages   int `json:"totalPages"`
	TotalRecords int `json:"totalRecords"`
}

type Links struct {
	Prev string `json:"prev,omitempty"`
	Self string `json:"self"`
	Next string `json:"next,omitempty"`
}
