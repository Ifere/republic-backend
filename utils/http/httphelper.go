package httphelper


import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	responses "republic-backend/config/responses"
)

//C serves as a place holder for the http request and response
type C struct {
	W http.ResponseWriter
	R *http.Request
	S http.Server
}

//H defines a json type formate
type H map[string]interface{}

//BindJSON decodes http request body to a given object
func (c *C) BindJSON(data interface{}) {
	err := json.NewDecoder(c.R.Body).Decode(data)

	if err != nil {
		return
	}
}

//JSON returns a http response encoded in application/json format to the response writer
func responseJSON(res http.ResponseWriter, status int, object interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	err := json.NewEncoder(res).Encode(object)

	if err != nil {
		return
	}
}

//Params maps routes params to mux and returns the value of the key
func (c *C) Params(key string) string {
	return mux.Vars(c.R)[key]
}

//Response returns a json response to the requester
func (c C) Response(resp responses.GeneralResponse) {

	responseJSON(c.W, http.StatusOK, resp)

}

//Response400 returns a json response to the requester
func Response400(res http.ResponseWriter, resp responses.GeneralResponse) {

	responseJSON(res, http.StatusBadRequest, H{"error": resp.Error, "status": false, "message": resp.Message})
}

//Response401 returns a json response to the requester
func Response401(res http.ResponseWriter, resp responses.GeneralResponse) {
	responseJSON(res, http.StatusUnauthorized, H{"error": resp.Error, "status": false, "message": resp.Message})
}

