package connectors

type ElasticSearchConnector struct {
}

func ProvideElasticSearchConnector() *ElasticSearchConnector {
	return &ElasticSearchConnector{}
}
