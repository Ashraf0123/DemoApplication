package v1

import (
	"database/sql"
	"fmt"
	"net/http"
	d1 "tempApi/Database"
	"tempApi/dto"
	"tempApi/helper"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

type CacheControllerInf interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

var Db *sql.DB
var Err error

type CacheControllerInstance struct {
}

func CacheController() CacheControllerInf {
	return new(CacheControllerInstance)
}
func (CacheControllerInstance) Create(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	err = helper.ValidateInput(value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	insertstmnt := `insert into Student(id,name) values($1,$2)`
	_, e := d1.GetPostgresMaster().Exec(insertstmnt, value.Id, value.Name)
	fmt.Println(e)
	c.JSON(http.StatusOK, "value set")

	return nil
}

func (CacheControllerInstance) Update(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	err = helper.ValidateInput(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	update := `update Student set "name"=$1 where "id"=$2`
	_, e := d1.GetPostgresMaster().Exec(update, value.Name, value.Id)
	fmt.Println(e)
	c.JSON(http.StatusOK, "Updated")

	return nil
}

func (CacheControllerInstance) Delete(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	err = helper.ValidateInput(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	delete := `delete from Student where "id"=$1`
	_, e := d1.GetPostgresMaster().Exec(delete, value.Id)
	fmt.Println(e)
	c.JSON(http.StatusOK, "Deleted")

	return nil
}
func (CacheControllerInstance) Get(c echo.Context) error {
	var value dto.Info
	err := c.Bind(&value)
	if err != nil {
		return err
	}
	statement := `select * from Student where "id"=$1`
	rows, err := d1.GetPostgresMaster().Query(statement, value.Id)
	for rows.Next() {
		var id string
		var name string
		err2 := rows.Scan(&id, &name)
		if err2 != nil {
			fmt.Println(err2.Error())
		} else {
			value = dto.Info{id, name}
		}
	}
	fmt.Println(err)

	c.JSON(http.StatusOK, value)

	return nil
}
