package api

type Err struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err,omitempty"`
}

var internal = Err{false, "internal error, try again later"}
