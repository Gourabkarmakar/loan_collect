package routers

import (
	"github.com/gin-gonic/gin"
	"text.com/controller"
)

func LoanRouters(r *gin.Engine) {

	// Loan routes
	loanRoutes := r.Group("/loan")
	{
		loanRoutes.POST("", controller.RequestLoan)
		loanRoutes.GET("", controller.GetAllLoans)
		loanRoutes.GET("/:id", controller.GetLoanInfo)
		loanRoutes.PUT("/:id", controller.PayLoan)
		loanRoutes.DELETE("/:id", controller.DeleteLoan)
	}

}
