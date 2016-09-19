package main

import (
	"errors"
	"fmt"
	"github.com/appcelerator/amp/api/client"
	"github.com/appcelerator/amp/api/rpc/topic"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
)

var (
	deleteTopicCmd = &cobra.Command{
		Use:   "rm [OPTIONS] TOPIC",
		Short: "Delete topic",
		Long:  `Delete topic.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := deleteTopic(AMP, cmd, args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	TopicCmd.AddCommand(deleteTopicCmd)
}

func deleteTopic(amp *client.AMP, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("must specify topic id")
	}
	id := args[0]
	if id == "" {
		return errors.New("must specify topic id")
	}

	request := &topic.DeleteRequest{Id: id}

	client := topic.NewTopicClient(amp.Conn)
	_, err := client.Delete(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}
