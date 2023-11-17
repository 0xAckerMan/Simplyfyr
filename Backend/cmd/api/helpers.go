package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type envelope map[string]interface{}

func (app *Application) readIDparam (r *http.Request) (int64, error){
    param := chi.URLParam(r, "id")
    idInt, err := strconv.ParseInt(param, 10, 64)
    if err != nil || idInt < 1 {
        return 0, errors.New("invalid url parameter")
    }
    return idInt,nil
}

func (app *Application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
    js, err := json.MarshalIndent(data, "", "\t")
    if err != nil {
        return err
    }
    js = append(js, '\n')
    for key, value := range(headers){
        w.Header()[key]=value
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(js)
    return nil
}

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Use http.MaxBytesReader() to limit the size of the request body to 1MB.
	maxByte := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxByte))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(dst)
	if err != nil {
        
		var syntaxError *json.SyntaxError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formatted json (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly formatted json")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect json type for  field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect json type(at character %d)", unmarshalTypeError.Offset)
			// If the request body exceeds 1MB in size the decode will now fail with the
			// error "http: request body too large".
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxByte)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		// If the JSON contains a field which cannot be mapped to the target destination
		case strings.HasPrefix(err.Error(), "json: unknown field"):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}
    err = dec.Decode(&struct{}{})
    if err != io.EOF{
        return errors.New("body must only contain one json value")
    }
	return nil
}

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

