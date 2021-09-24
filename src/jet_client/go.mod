module jet_client

go 1.16

require (
	github.com/confluentinc/confluent-kafka-go v1.7.0 // indirect
	google.golang.org/grpc v1.39.0
	gopkg.in/confluentinc/confluent-kafka-go.v1 v1.7.0
	labnet.com/proto v0.0.0-00010101000000-000000000000
)

replace labnet.com/proto => ../proto
