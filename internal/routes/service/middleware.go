package service

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"regexp"

	"github.com/changyoungkwon/gxample/internal/logging"
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
			logging.Errorf("error during parsing multipart, %v", err)
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

// parseMultipartRequest save images, returns formdata key and saved path
func parseMultipartRequest(r *http.Request) (map[string]string, []byte, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, nil, err
	}

	keyPathMap := make(map[string]string)
	var receivedJSON []byte
	imageIndex := 0

	// generate directory to save resource
	dirname, _ := uuid.NewUUID()
	dirpath := path.Join(StaticRootPath, dirname.String())
	err = os.MkdirAll(dirpath, os.FileMode(0755))
	if err != nil {
		logging.Errorf("error creating directory, %v", err)
		return nil, nil, err
	}

	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		defer part.Close()

		// actions for each keyname
		key := part.FormName()
		if validJSONFormName.MatchString(key) {
			receivedJSON, err = ioutil.ReadAll(part)
			if err != nil {
				return nil, nil, err
			}
		} else if validFileFormName.MatchString(key) {
			// read all bytes
			bytes, err := ioutil.ReadAll(part)
			if err != nil {
				logging.Errorf("error during reading from multiparts, %v", err)
				return nil, nil, err
			}
			// check if valid image, and determines extension
			ext, err := getImageExt(bytes[:512])
			if err != nil {
				logging.Errorf("error due to invalid mime-type sent, %v", err)
				return nil, nil, err
			}
			// save file with image
			filename := path.Join(dirpath, fmt.Sprintf("image_%d.%s", imageIndex, ext))
			ioutil.WriteFile(filename, bytes, os.FileMode(0755))
			if err != nil {
				return nil, nil, err
			}
			// save
			keyPathMap[key] = "/" + filename
			imageIndex++
		} else {
			return nil, nil, errors.New("invalid multipart key")
		}
	}
	return keyPathMap, receivedJSON, nil
}

// getImageExt get image type based on first 512 bytes
func getImageExt(buf []byte) (string, error) {
	imagetype, _ := regexp.Compile("^image/[a-z]+$")
	mimetype := http.DetectContentType(buf)
	mediatype, _, err := mime.ParseMediaType(mimetype)
	if err != nil {
		logging.Errorf("error determining image ext, %v", err)
		return "", err
	}
	if !imagetype.MatchString(mediatype) {
		logging.Errorf("unexpected mediatype received, %s", mediatype)
		return "", fmt.Errorf("unexpected mimetype %s received", mimetype)
	}
	return mediatype[6:], nil
}
