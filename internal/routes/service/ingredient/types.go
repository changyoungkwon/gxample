package ingredient

import (
	"net/http"

	"github.com/changyoungkwon/gxample/internal/models"
)

// Request wraps request
type Request struct {
	*models.Ingredient
}

// Response wraps response
type Response struct {
	*models.Ingredient
}

// Bind binds additional parameters on Request after decode
func (i *Request) Bind(r *http.Request) error {
	return nil
}

// Render renders additional paramters before encode and response
func (i *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
