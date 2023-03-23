package contract

// ProtectedMethods is a function to hold grpc service methods
// false value indicates that the method is not protected (no authorization needed)
func ProtectedMethods() map[string]bool {
	return map[string]bool{
		"/log.LogService/SaveHttpLog": true,
		"/user.UserService/GetUser": true,
		"/book.BookService/GetBook": true,	
	}
}
