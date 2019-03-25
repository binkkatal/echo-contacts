package api

import (
	"database/sql"
	"log"
	"net/http"

	errors "github.com/binkkatal/echo-contacts/errors"

	"github.com/binkkatal/echo-contacts/model"
	"github.com/labstack/echo"
)

const (
	INSERT_CONTACT_STATEMENT = `INSERT INTO contacts (first_name, last_name, email, organization, phone_number, website) 
	VALUES (?, ?, ?, ?, ?, ?);`

	LIST_CONTACTS_QUERY = `SELECT * FROM contacts`

	SELECT_CONTACT_BY_ID = `SELECT * FROM contacts WHERE id = ?`

	UPDATE_CONTACT_QUERY = `UPDATE contacts SET first_name = ?, last_name = ?, email =?, organization = ?, phone_number =?, website= ?
	WHERE id = ?;`

	DELETE_CONTACT_QUERY = `DELETE FROM contacts
	WHERE id = ?;`
)

type Api struct {
	*sql.DB
}

func (api *Api) Index(c echo.Context) (err error) {
	rows, err := api.DB.Query(LIST_CONTACTS_QUERY)

	if err != nil {
		log.Printf("Error getting contacts %s", err)
		return handleError(c, err)
	}

	var contacts []*model.Contact

	for rows.Next() {
		contact := &model.Contact{}

		err = rows.Scan(&contact.ID,
			&contact.FirstName,
			&contact.LastName,
			&contact.Email,
			&contact.Organization,
			&contact.PhoneNumber,
			&contact.Website,
		)

		if err != nil {
			log.Printf("Error Scanning rows %s", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		contacts = append(contacts, contact)
	}

	c.JSON(http.StatusOK, contacts)
	return
}

func (api *Api) Create(c echo.Context) (err error) {
	contact := new(model.Contact)
	if err = c.Bind(contact); err != nil {
		log.Printf("Create: Error binding contact (%+v)", err)
		return handleError(c, err)
	}
	log.Printf("Create: Received contact params", contact)
	result, err := api.DB.Exec(INSERT_CONTACT_STATEMENT,
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Organization,
		contact.PhoneNumber,
		contact.Website,
	)

	if err != nil {
		log.Printf("Create: Error Creating contact (%+v)", err)
		return handleError(c, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error retreiveing id of inserted row %+v", err)
		return handleError(c, err)
	}

	log.Printf("Added row to ID %d", id)
	contact.ID = id
	return c.JSON(http.StatusOK, contact)
}

func (api *Api) Show(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		log.Printf("Error: No contact 'id' passed ")
		return handleError(c, errors.MISSING_CONTACT_ID)
	}
	contact := new(model.Contact)
	err := api.QueryRow(SELECT_CONTACT_BY_ID, id).Scan(&contact.ID,
		&contact.FirstName,
		&contact.LastName,
		&contact.Email,
		&contact.Organization,
		&contact.PhoneNumber,
		&contact.Website,
	)

	if err != nil {
		log.Printf("Error Scanning rows %s", err)
		return handleError(c, err)
	}

	return c.JSON(http.StatusOK, contact)
}

func (api *Api) Update(c echo.Context) error {
	contact := new(model.Contact)

	id := c.Param("id")

	if id == "" {
		log.Printf("Error: No contact 'id' passed ")
		return handleError(c, errors.MISSING_CONTACT_ID)
	}

	if err := c.Bind(contact); err != nil {
		log.Printf("Create: Error binding contact (%+v)", err)
		return handleError(c, err)
	}

	log.Printf("Create: Received contact params", contact)

	result, err := api.DB.Exec(UPDATE_CONTACT_QUERY,
		contact.FirstName,
		contact.LastName,
		contact.Email,
		contact.Organization,
		contact.PhoneNumber,
		contact.Website,
		id,
	)

	if err != nil {
		log.Printf("Update: Error Updating contact (%+v)", err)
		return handleError(c, err)
	}
	count, err := result.RowsAffected()

	if err != nil {
		log.Printf("Update: Error Updating contact with id: %s (%+v)", id, err)
		return handleError(c, err)
	}

	if count == 0 {
		log.Printf("Update: Record was not Updated: %s (%+v)", id, err)
		return handleError(c, errors.Error("Record was not Updated"))
	}

	return c.JSON(http.StatusOK, "Record Successfully Updated")
}

func (api *Api) Delete(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		log.Printf("Error: No contact 'id' passed ")
		return handleError(c, errors.MISSING_CONTACT_ID)
	}

	result, err := api.DB.Exec(DELETE_CONTACT_QUERY, id)
	if err != nil {
		log.Printf("Create: Error Executing delete (%+v)", err)
		return handleError(c, err)
	}

	count, err := result.RowsAffected()

	if err != nil {
		log.Printf("DELETE: Error Deleting contact with id: %s (%+v)", id, err)
		return handleError(c, err)
	}

	if count == 0 {
		log.Printf("DELETE: Record was not deleted: %s (%+v)", id, err)
		return handleError(c, errors.Error("Record was not deleted"))
	}

	return c.JSON(http.StatusOK, "Record Deleted Successfully")
}

func handleError(c echo.Context, err error) error {
	var errMsg = ""
	if err != nil {
		errMsg = err.Error()
	}
	return c.JSON(http.StatusInternalServerError, errMsg)
}
