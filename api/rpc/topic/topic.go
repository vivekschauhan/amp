package topic

import (
	"github.com/appcelerator/amp-kafka-pilot/pilot/api/admin"
	"github.com/appcelerator/amp/data/kafka"
	"github.com/appcelerator/amp/data/storage"
	"github.com/docker/docker/pkg/stringid"
	"golang.org/x/net/context"
	"path"
)

// Server is used to implement topic.TopicServer
type Server struct {
	Store storage.Interface
	Kafka kafka.Kafka
}

// Create implements topic.TopicServer
func (s *Server) Create(ctx context.Context, in *CreateRequest) (*CreateReply, error) {
	request := &admin.CreateTopicRequest{
		Topic: &admin.TopicEntry{
			Name:              in.Topic.Name,
			Partitions:        in.Topic.Partitions,
			ReplicationFactor: in.Topic.ReplicationFactor,
		},
	}
	created, err := s.Kafka.Admin().CreateTopic(ctx, request)
	if err != nil {
		return nil, err
	}
	topic := &TopicEntry{
		Id:                stringid.GenerateNonCryptoID(),
		Name:              created.Topic.Name,
		Partitions:        created.Topic.Partitions,
		ReplicationFactor: created.Topic.ReplicationFactor,
	}
	if err := s.Store.Create(ctx, path.Join("topics", topic.Id), topic, nil, 0); err != nil {
		return nil, err
	}
	return &CreateReply{Topic: topic}, nil
}

// List implements topic.TopicServer
func (s *Server) List(ctx context.Context, in *ListRequest) (*ListReply, error) {
	//kvs, err := s.Store.ListRaw(ctx, "topics", storage.Everything)
	//if err != nil {
	//	return nil, err
	//}
	//for _, kv := range kvs {
	//	log.Println("topic: ", strings.TrimPrefix(string(kv.Key), "amp/topics/"))
	//}
	reply := &ListReply{}
	return reply, nil
}

// Delete implements topic.TopicServer
func (s *Server) Delete(ctx context.Context, in *DeleteRequest) (*DeleteReply, error) {
	//if err := s.Store.Delete(ctx, path.Join("topics", in.Name), false, nil); err != nil {
	//	return nil, err
	//}
	reply := &DeleteReply{}
	return reply, nil
}
