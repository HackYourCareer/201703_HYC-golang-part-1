// 1. First approach
func GetDataFromRequest(r *http.Request) (interface{}, error)

//2. Second approach
func GetDataFromRequest(r *http.Request) (MyData, error)

// 3. Third approach
//Create an interafce
type Dto interface {
	Unmarshall(r *http.Request) error
}

func GetDataFromRequest(r *http.Request, d Dto) error {
	return d.Unmarshal(r)
}
// Implement interface
func (p *MyData) Unmarshal(r *http.Request) error {
// ...
}
//Consume
var d MyData
if err := GetDataFromRequest(req, &d); err != nil {
// ...
}


