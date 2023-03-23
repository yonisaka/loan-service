package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yonisaka/loan-service/consts"
	"github.com/yonisaka/loan-service/domain/entity"
	"github.com/yonisaka/loan-service/rest/dto"
	pbBook "github.com/yonisaka/protobank/book"
	pbUser "github.com/yonisaka/protobank/user"
)

type LoanHandler struct {
	*Handler
}

func NewLoanHandler(h *Handler) *LoanHandler {
	return &LoanHandler{h}
}

func (r *LoanHandler) GetLoanList(c *gin.Context) {
	rsp, err := r.repo.Loan.Get(c)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithData(rsp),
	)
	return
}

func (r *LoanHandler) CreateLoan(c *gin.Context) {
	var req entity.Loan
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			*dto.NewResponse().WithCode(http.StatusBadRequest).WithMessage(err.Error()),
		)
		return
	}
	
	_, err := r.client.GetUser(c, &pbUser.UserByIDRequest{Id: int64(req.UserID)})
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			*dto.NewResponse().WithCode(http.StatusNotFound).WithMessage(err.Error()),
		)
		return
	}

	_, err = r.client.GetBook(c, &pbBook.BookByIDRequest{Id: int64(req.BookID)})
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			*dto.NewResponse().WithCode(http.StatusNotFound).WithMessage(err.Error()),
		)
		return
	}
	
	err = r.repo.Loan.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			*dto.NewResponse().WithCode(http.StatusInternalServerError).WithMessage(err.Error()),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		*dto.NewResponse().WithCode(http.StatusOK).WithMessage(consts.MessageSuccessCreate).WithData(req),
	)
	return
}