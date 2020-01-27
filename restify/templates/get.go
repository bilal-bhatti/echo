package templates

var get = `
func ContactsServiceGetOneHandler(w http.ResponseWriter, req *http.Request) {
	{{.Preamble}}

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// invalid request error
		panic(err)
	}

	contactResponse, err := contacts.GetOne(id)
	if err != nil {
		// write error response
		// internal error
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	err = json.NewEncoder(w).Encode(contactResponse)
	if err != nil {
		// write error response
		panic(err)
	}
}
`
