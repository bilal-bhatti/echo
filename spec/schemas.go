package spec

// import (
// 	"https://github.com/go-openapi/spec"
// )

// type Resource interface {
// 	func() string
// }

// type ObjectSchema struct {
// }

// type StringSchema struct {
// 	Description string
// 	Required    bool
// }

// type ContactsService struct {
// 	data     *connectors.DataConnector
// 	identity *connectors.IdentityConnector
// }

// Constructor with whatever args
// func NewContactsService() {

// }

// func (cs ContactsService) Post(req ContactRequest) (ContactResponse, error) {

// 	return ContactResponse{}, nil
// }

// func (cs ContactsService) PostHandler(w http.ResponseWriter, req *http.Request) {
// 	// creq := &ContactRequest{}
// 	// json.NewDecoder(req.GetBody()).Decode(creq)

// }

// func reg() {
// 	mux := mux.NewRouter()

// 	// Register route handlers
// 	cs := services.ContactsService{}

// 	mux.HandleFunc("/", cs.PostHandler).Methods("POST")
// }

// type Bundle struct {
// 	Data []byte
// }

// func ParseHTTPRequest(req *http.Request) ([]byte, error) {
// 	bites, err := ioutil.ReadAll(req.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return bites, nil
// }

// type Error struct {
// 	err    string  //error description
// 	length float64 //length which caused the error
// 	width  float64 //width which caused the error
// }

// type Builder struct {
// 	path    string
// 	method  string
// 	handler func(w http.ResponseWriter, req *http.Request)
// }

// // func build(m, path string, handler func(req ContactRequest) (ContactResponse, error)) string {

func Post(path string, handler interface{}) {
	panic("Implementation not generated, run restify")
}

func Get(path string, handler interface{}) {
	panic("Implementation not generated, run restify")
}

func spec() {
	// panic(build("POST", "/contacts", func(req services.ContactRequest) (*services.ContactResponse, error) { panic("not implemented") }))
}
