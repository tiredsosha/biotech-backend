package server

type Basic struct {
	Command string `json:"command" binding:"required"`
}

type Advansed struct {
	Command string `json:"command" binding:"required"`
	Zone    string `json:"zone" binding:"required"`
}
