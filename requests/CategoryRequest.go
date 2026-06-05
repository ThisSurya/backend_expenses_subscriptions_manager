package requests

type CategoryRequest struct {
	Name    string  `json:"name" binding:"required"`
	Color   string  `json:"color" binding:"required"`
	IconUrl *string `json:"icon_url"`
}
