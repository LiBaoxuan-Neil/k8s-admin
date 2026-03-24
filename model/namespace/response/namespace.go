package response

type NameSpace struct {
	Name              string `json:"name"`
	CreationTimestamp int64  `json:"creationTimestamp"`
	Status            string `json:"status"`
}
