package repo

type TaskId struct {
	TaskId int `json:"taskId" db:"id"`
}

func (t TaskId) GetAllTask() error {

	return nil
}
