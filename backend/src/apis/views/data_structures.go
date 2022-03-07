package views


type UserRow struct {
	Column_id          int64       `json:"id" binding:"required"`
	Column_uuid        string      `json:"uuid" binding:"required"`
	Column_email       string      `json:"email" binding:"required"`
	Column_name        string      `json:"name" binding:"required"`
	Column_description string      `json:"description" binding:"required"`
	Column_createdAt   interface{} `json:"createdAt"`
	Column_updatedAt   interface{} `json:"updatedAt"`
}


type NewUserRow struct {
	Column_id          int64       `json:"id"`
	Column_uuid        string      `json:"uuid"`
	Column_email       string      `json:"email" binding:"required"`
	Column_name        string      `json:"name" binding:"required"`
	Column_description string      `json:"description" binding:"required"`
	Column_createdAt   interface{} `json:"createdAt"`
	Column_updatedAt   interface{} `json:"updatedAt"`
}