package recipe

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
)

// Request wraps request
type Request struct {
	*models.Recipe
}

// Response wraps response
type Response struct {
	*models.Recipe
}

// Bind binds additional parameters on IngredientRequest after decode
func (i *Request) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
