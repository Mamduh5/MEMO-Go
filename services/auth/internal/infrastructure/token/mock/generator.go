package mock

type Generator struct{}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateAccessToken(userID string) (string, error) {
	return "access-token-mock", nil
}

func (g *Generator) GenerateRefreshToken() (string, error) {
	return "refresh-token-mock", nil
}
