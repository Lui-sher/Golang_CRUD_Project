package variables

// Creamos una varible del tipo Struct con la inicial en mayuscula para
// poder exportar y tener acceso a ello desde cualquier archivo

type User struct {
	Record  int32
	User_Id string
	Name    string `json:"name"`
	Email   string `json:"email"`
	Is_test bool
}

var Count = 0
var ResMap map[string]string
var StoreUsersId = make([]string, 3)
