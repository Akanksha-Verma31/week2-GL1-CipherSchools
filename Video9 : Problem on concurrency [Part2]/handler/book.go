package handler

import (
	"log"
	"net/http"

	"go-application/database"
	"go-application/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBooks(in *gin.Context) {
	books, err := database.GetBooks(h.DB) //h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, books)
}
func (h *Handler) SaveBook(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SaveBook(h.DB, &book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
func (h *Handler) GetBookByID(in *gin.Context) {
	id := in.Params.ByName("id")
	book, err := database.GetBookByID(h.DB, id) //h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}
func (h *Handler) DeleteBookByID(in *gin.Context) {
	id := in.Params.ByName("id")
	err := database.DeleteBookByID(h.DB, id) //h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, id)
}
func (h *Handler) UpdateBookByID(in *gin.Context) {
	book := models.Book{}
	err := in.BindJSON(&book)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.UpdateBookByID(h.DB, &book) //h.DB is fully configured and can give the access of book table
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, book)
}

/*
h:=Handler{}
h.DB=GetDB()
h.GetBooks()
*/
