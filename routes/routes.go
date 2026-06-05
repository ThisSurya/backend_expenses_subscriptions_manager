package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(router *gin.RouterGroup, c *controllers.ExpenseController) {
	expense := router.Group("/expenses")

	{
		expense.GET("/", c.GetAllExpenses)
		expense.GET("/:id", c.GetExpenseDetail)
		expense.GET("/user/:user_id", c.GetExpenseByUserId)
		expense.POST("/", c.CreateExpense)
		expense.PUT("/:id", c.UpdateExpenses)
		expense.DELETE("/:id", c.DeleteExpense)
	}
}

func CategoryRoutes(router *gin.RouterGroup, c *controllers.CategoryController) {
	category := router.Group("/categories")
	{
		category.GET("/", c.GetByUserId)
		category.GET("/:id_category", c.GetById)
		category.POST("/", c.CreateCategory)
		category.PUT("/:id_category", c.UpdateCategory)
		category.DELETE("/:id_category", c.DeleteCategory)
	}

}

func SubscriptionRoutes(router *gin.RouterGroup, c *controllers.SubscriptionController) {
	subscription := router.Group("/subscriptions")
	{
		subscription.GET("/", c.GetSubscriptionByUserId)
		subscription.GET("/:id", c.GetSubscriptionDetail)
		subscription.POST("/", c.CreateSubscription)
		subscription.PUT("/:id", c.UpdateSubscription)
		subscription.DELETE("/:id", c.DeleteSubscription)
	}
}
