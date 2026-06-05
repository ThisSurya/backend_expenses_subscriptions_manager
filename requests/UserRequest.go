package requests

type UserRole string

const (
	UserRoleBasic   UserRole = "basic"
	UserRolePremium UserRole = "premium"
	UserRoleAdmin   UserRole = "admin"
)

type UserRequest struct {
	GoogleId        *string   `json:"google_id"`
	Username        *string   `json:"username"`
	Email           string    `json:"email" binding:"required,email"`
	Password        string    `json:"password" binding:"required,min=6"`
	AvatarUrl       *string   `json:"avatar_url"`
	ReminderEnabled *bool     `json:"reminder_enabled"`
	ReminderDays    *int      `json:"reminder_days"`
	Role            *UserRole `json:"role"`
	Timezone        *string   `json:"timezone"`
}

type UserRegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
