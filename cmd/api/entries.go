//Filename: cmd/api/entries.go

package main
import (
	"fmt"
	"net/http"
	"strconv"
	

	"github.com/julienschmidt/httprouter"
)

//createSchoolHandler 
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "create an entry")
}

func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the id parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	//display the school id
	fmt.Fprintf(w, "Show the Entry details %d\n", id)
}