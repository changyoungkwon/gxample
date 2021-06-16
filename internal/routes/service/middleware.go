package service

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type ctxKeyMultipart int

// ImageUploadKey is the key holds the unique rueqest
const (
	imageHandleKey     ctxKeyMultipart = 0
	jsonHandleKey      ctxKeyMultipart = 1
	maximumContentSize int64           = 32 << 20
)

var (
	validFileFormName = regexp.MustCompile(`^(file|step_[0-9]+)$`)
	validJSONFormName = regexp.MustCompile(`^json$`)
)

// multipartJSONHandler is a middlware that handles multipart formdata, and insert key-value
func multipartJSONHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// inserts key to the context
		imagePaths, receivedJSON, err := parseMultipartRequest(r)
		// when not part of multipart-form
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, imageHandleKey, imagePaths)
		ctx = context.WithValue(ctx, jsonHandleKey, receivedJSON)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

// getMultipartJSON get data saved by middleware
func getMultipartJSON(c context.Context) ([]byte, error) {
	data, ok := c.Value(jsonHandleKey).([]byte)
	if !ok {
		return nil, errors.New("unable to fetch JSON from context")
	}
	return data, nil
}

// saveMultipart save a file into specified name
// create directory recursively if force is true
func saveMultipart(part *multipart.Part, filename string, force bool) error {
	dirname := filepath.Dir(filename)
	if force {
		err := os.MkdirAll(dirname, os.FileMode(0755))
		if err != nil {
			return err
		}
	}
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, part)
	if err != nil {
		return err
	}
	return nil
}

// parseMultipartRequest save images, returns formdata key and saved path
func parseMultipartRequest(r *http.Request) (map[string]string, []byte, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, nil, err
	}

	keyPathMap := make(map[string]string)
	var receivedJSON []byte
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}

		// validation
		defer part.Close()
		key := part.FormName()
		if validJSONFormName.MatchString(key) {
			receivedJSON, err = ioutil.ReadAll(part)
			if err != nil {
				return nil, nil, err
			}
		} else if !validFileFormName.MatchString(key) {
			return nil, nil, errors.New("invalid multipart key")
		}
		dirname, _ := uuid.NewUUID()
		filename := path.Join(StaticRootPath, dirname.String(), part.FileName())
		err = saveMultipart(part, filename, true)
		if err != nil {
			return nil, nil, err
		}
		keyPathMap[key] = filename
	}
	return keyPathMap, receivedJSON, nil
}
