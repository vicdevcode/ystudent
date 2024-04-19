package dto

type Group struct {
	Name      string `json:"name"       binding:"required"`
	FacultyID uint   `json:"faculty_id" binding:"required"`
}

type CreateGroup Group

type UpdateGroupCurator struct {
	ID        uint `json:"id"         binding:"required"`
	CuratorID uint `json:"curator_id" binding:"required"`
}
