// Package classification Petstore API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v2
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: John Doe<john.doe@example.com> http://john.doe.com
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: KEY
//          in: header
//     oauth2:
//         type: oauth2
//         authorizationUrl: /oauth2/auth
//         tokenUrl: /oauth2/token
//         in: header
//         scopes:
//           bar: foo
//         flow: accessCode
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	routes := httprouter.New()

	routes.HandlerFunc(http.MethodGet, "/foo", fooHandler)
	routes.HandlerFunc(http.MethodGet, "/bar", barHandler)

	srv := http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	log.Fatal(srv.ListenAndServe())
}

// @Summary A dummy foo handler
// @Description foo handler
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string "FOO is called"
// @Router /foo [get]
func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("FOO is called"))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BAR is called"))
}
