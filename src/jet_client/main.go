package main

import (
	"context"
	"flag"
	"log"
	"time"
	"fmt"
//	"io"
//	"reflect"
	"encoding/json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	dds "labnet.com/proto/2/dds_service"
	auth "labnet.com/proto/2/auth_service"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var (
	serverAddr = flag.String("server_addr", "b1-re:32767", "The connection point")
	tls = flag.Bool("tls", false, "Connection uses TLS if true")
	caFile = flag.String("ca_file", "b1-re.pem", "The file containing the CA root cert file")
	serverHostOverride = flag.String("server_host_override", "b1-re", "The server name as specified in the cert reqeust")
	username = flag.String("user", "lab", "username")
	passwd = flag.String("passwordd", "lab123", "password")
	clientId = flag.String("client_id", "42", "Client ID for the session")
	subscribeTopic = flag.String("sub_topic", "/Root/Type/net::juniper::rtnh::Nexthop", "Topic string for subscription")
	kafkaBrokers = flag.String("kafka_brokers", "b1-re", "Kafka broker for output")
	kafkaTopic = flag.String("kafka_topic", "test_topic", "Kafka topic for output")
)

func publish_stream(topic_stream dds.DDS_TopicSubscribeClient) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": *kafkaBrokers})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := *kafkaTopic
	//for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	//	p.Produce(&kafka.Message{
	//		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	//		Value:          []byte(word),
	//	}, nil)
	//}
	for {
		msg, err :=  topic_stream.Recv()
		if err != nil {
			log.Fatalf("Something bad happened %v", err)
		}
		entry := msg.Objects
		for _, e := range entry {
			out_entry, _ := json.Marshal(e)
			//fmt.Printf("%+v\n", out_entry)
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value: out_entry,
			}, nil)
		}
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	if *tls {
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Get new login client using our grpc conn
	c := auth.NewAuthenticationClient(conn)

	// Perform login check against the JET auth API
	r, err := c.Login(context.Background(), &auth.LoginRequest{
		Username: *username,
		Password: *passwd,
		ClientId: *clientId,
	})

	if err != nil {
		log.Fatalf("Could not connect. Check IP address or domain name: %v\n", err)
	} else {
		if r.Status.Code == 0 {
			log.Printf("Connect to %s successful\n", *username)
		}
	}
	client := dds.NewDDSClient(conn)

	//request := dds.TopicListRequest{RequestId: uint64(42)}
	//req := &request
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	
	request := dds.TopicSubscribeRequest{ContinuousSync: true, SubscriptionTopics: []string{*subscribeTopic}}
	req := &request

	topicStream, err := client.TopicSubscribe(ctx, req)

	publish_stream(topicStream)	

	//topicList, err := client.TopicListGet(ctx,req)
	//if err != nil {
	//	log.Fatalf("%v.TopicSubscribe(_) = _, %v", client, err)
	//}

	//tl := &topicList
	//fmt.Printf("%+v\n", topicStream)
	//b, err := json.Marshal(tl)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(b))
//	vtl := reflect.ValueOf(topicList)
//	typeOfTl := vtl.Type()
//
//	for i := 0; i< vtl.NumField(); i++ {
//		fmt.Printf("Field: %s\tValue:%v\n", typeOfTl.Field(i).Name, vtl.Field(i).Interface())
//	}
	//topics := topicList.GetTopics()

	//for _,t := range topics {
	//	fmt.Printf(t.Topic,"\n")
	//}
}
