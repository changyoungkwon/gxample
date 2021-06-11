package service

import (
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type ctxKeyImageUpload int

// ImageUploadKey is the key holds the unique rueqest
const (
	imageHandleKey     ctxKeyImageUpload = 0
	maximumContentSize int64             = 32 << 20
)

var (
	validFileFormName = regexp.MustCompile(`^(file|step_[0-9]+)$`)
	validJSONFormName = regexp.MustCompile(`^json$`)
)

// ImageHandleMiddlware is a middlware that handles multipart formdata, insert key-value
func imageHandleMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// inserts key to the context
		imagePaths, err := saveMultipartFiles(r)
		// when not part of multipart-form
		if err != nil {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, imageHandleKey, imagePaths)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
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

// saveMultipartFiles save images, returns formdata key and saved path
func saveMultipartFiles(r *http.Request) (map[string]string, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	keyPathMap := make(map[string]string)
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}

		// validation
		defer part.Close()
		key := part.FormName()
		if validJSONFormName.MatchString(key) {
			continue
		} else if !validFileFormName.MatchString(key) {
			return nil, errors.New("invalid multipart key")
		}
		dirname, _ := uuid.NewUUID()
		filename := path.Join(StaticRootPath, dirname.String(), part.FileName())
		err = saveMultipart(part, filename, true)
		if err != nil {
			return nil, err
		}
		keyPathMap[key] = filename
	}
	return keyPathMap, nil
}
