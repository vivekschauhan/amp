package topic

import (
	"github.com/appcelerator/amp/data/storage"
	"golang.org/x/net/context"
	"path"
)

// Server is used to implement topic.TopicServer
type Server struct {
	Store storage.Interface
}

// Create implements topic.TopicServer
func (s *Server) Create(ctx context.Context, in *CreateRequest) (*CreateReply, error) {
	if err := s.Store.Create(ctx, path.Join("topics", in.Topic.Name), nil, nil, 0); err != nil {
		return nil, err
	}
	reply := &CreateReply{}
	return reply, nil
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
