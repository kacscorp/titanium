package employees

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/kacscorp/titanium/lib/config/handlers"
	"github.com/kacscorp/titanium/lib/report"
	"github.com/kacscorp/titanium/lib/sources/employees/v1/query"
)

//Source ...
type Source struct {
	db *gorm.DB
}

// NewSource ...
func NewSource(db *gorm.DB) *Source {
	return &Source{
		db: db,
	}
}

// SearchEmployeeByID takes an id object and search it on the database.
// If an error occurs then it is returned with a nil Response.
func (src Source) SearchEmployeeByID(id int64) (*report.Employee, error) {
	qr, err := query.NewSelectEmployeesByID(id)
	if err != nil {
		return nil, errors.New("error when building query")
	}

	output := query.Employees{}
	src.db.Raw(qr.Query, qr.Args...).Scan(&output)

	return &report.Employee{
		ID:                   int64(output.ID),
		IdentificationNumber: output.IdentificationNumber.String,
		Age:                  int(output.Age.Int64),
		FirstName:            output.FirstName.String,
		LastName:             output.LastName.String,
	}, nil
}

// GetHandler receives a *handlers.AppContext, a http.ResponseWriter and a *http.Request and
// searches for values in titanium.
//
// The result is written to the writer as a JSON.
//
// In case of error, it is returned.
func GetHandler(ctx *handlers.AppContext, w http.ResponseWriter, r *http.Request) (int, error) {
	if ctx == nil {
		return http.StatusInternalServerError, errors.New("nil context")
	}

	strID := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to parse ID query parameter into an int64")
	}

	src := NewSource(ctx.DB)
	response, err := src.SearchEmployeeByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	_, err = json.Marshal(response)
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to marshal response")
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		return http.StatusInternalServerError, errors.New("failed to encode employee into the http.ResponseWriter")
	}

	return http.StatusOK, nil
}
