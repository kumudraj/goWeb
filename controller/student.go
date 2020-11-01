package controller

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/kumudraj/goWeb/model"
	"github.com/kumudraj/goWeb/storage"
	"github.com/labstack/echo/v4"
)

var (
	student = map[int]*model.Students{}
	seq     = 1
)

// GetAllStudents <- e.GET("/all_students", controller.GetAllStudents)
func GetAllStudents(c echo.Context) error {
	students, _ := genAllStudents()
	return c.JSON(http.StatusOK, students)
}

// GenAllStudents select  all student
func genAllStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}
	db.Select(&students, "SELECT * FROM students ORDER BY id DESC")

	return students, nil
}

// GetStudent <- e.GET("/student/:id", GetStudent)
func GetStudent(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	x, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	students, _ := genStudent(x)
	return c.JSON(http.StatusOK, students)
}

// GenStudent select  student  by id
func genStudent(id int64) ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}
	db.Select(&students, "SELECT * FROM students where id=$1", id)
	return students, nil
}

// SaveStudent <- e.POST("/student", controller.SaveStudent)
func SaveStudent(c echo.Context) error {

	u := &model.Students{
		Id: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	student[u.Id] = u
	seq++
	err := insertStudentInDb(u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, u)
	}
	return c.JSON(http.StatusCreated, u)
}

func insertStudentInDb(data *model.Students) error {
	q := `INSERT INTO students (id, name) VALUES ($1, $2)`
	db := storage.GetDBInstance()
	result := db.MustExec(q, data.Id, data.Name)

	if result != nil {
		fmt.Println("#####", result)
	}
	return nil
}
