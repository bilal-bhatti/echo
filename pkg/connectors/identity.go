package connectors

type IdentityConnector struct {
}

func ProvideIdentityConnector() *IdentityConnector {
	return &IdentityConnector{}
}
