package models

// Creamos una varible del tipo Struct con la letra inicial en mayuscula para
// poder exportar y tener acceso a ello desde cualquier archivo

type User struct {
	Record  int32  `json:"record"`
	User_Id string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Is_test bool   `json:"is_test"`
}

// var Count = 0
// var ResMap map[string]string
var StoreUsersId = make([]string, 3)
