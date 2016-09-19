package main

import (
	"fmt"
	"github.com/appcelerator/amp/api/client"
	"github.com/appcelerator/amp/api/rpc/topic"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	listTopicCmd = &cobra.Command{
		Use:   "ls [OPTIONS]",
		Short: "List topics",
		Long:  `List topics.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := listTopic(AMP, cmd, args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	TopicCmd.AddCommand(listTopicCmd)
}

func listTopic(amp *client.AMP, cmd *cobra.Command, args []string) error {
	request := &topic.ListRequest{}

	client := topic.NewTopicClient(amp.Conn)
	reply, err := client.List(context.Background(), request)
	if err != nil {
		return err
	}
	for _, topic := range reply.Topics {
		fmt.Printf("ID: %s, Name: %s, partitions: %d, replicationFactor: %d", topic.Id, topic.Name, topic.Partitions, topic.ReplicationFactor)
	}
	return nil
}
