package grpc

func isPublicMethod(method string) bool {
	switch method {
	case "/auth.v1.AuthService/Login":
		return true
	case "/auth.v1.AuthService/Register":
		return true
	case "/auth.v1.AuthService/Refresh":
		return true
	default:
		return false
	}
}
