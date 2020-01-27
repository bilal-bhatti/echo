package templates

var Post = `
func ContactsServiceCreateHandlerFunc(funk ContactsServiceCreateHandlerType) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		contactRequest := &models.ContactRequest{}

		err := json.NewDecoder(req.Body).Decode(contactRequest)
		if err != nil {
			errors.Wrap(err, "Failed to decode json request body")
		}

		if err != nil {
			// write error response
			// invalid request error
			panic(err)
		}

		res, err := funk(context.TODO(), contactRequest)
		if err != nil {
			// write error response
			// internal error
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			// write error response
			panic(err)
		}
	}
}
`
