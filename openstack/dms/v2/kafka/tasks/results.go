package tasks

type Task struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	UserName  string `json:"user_name"`
	UserId    string `json:"user_id"`
	Params    string `json:"params"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
