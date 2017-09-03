package commands

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/sh3rp/databox/config"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/util"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/yhat/scrape"
	"google.golang.org/grpc"
)

var linkId string
var linkName string
var linkUrl string
var linkBoxId string
var linkTags string

var searchTerm string
var searchCount int
var searchPage int
var searchLoadLinks bool

var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Manage your links",
}

var LinkAddCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a link",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		if linkName == "" {
			resp, err := http.Get(linkUrl)
			if err != nil {
				panic(err)
			}

			dom, err := html.Parse(resp.Body)

			if err != nil {
				panic(err)
			}

			if title, ok := scrape.Find(dom, scrape.ByTag(atom.Title)); ok {
				linkName = scrape.Text(title)
			} else {
				linkName = "Unknown"
			}
		}

		if linkBoxId == "" {
			cfg := &config.ClientConfig{}
			cfg.Read()
			linkBoxId = cfg.DefaultBoxId
		}

		if linkBoxId == "" {
			fmt.Printf("Error: unable to determine box id.\n")
			os.Exit(1)
		}

		link, err := client.NewLink(context.Background(), &msg.Link{Name: linkName, Url: linkUrl, BoxId: linkBoxId})

		if err != nil {
			fmt.Printf("Error creating link: %v\n", err)
		} else {
			util.PrettyPrint(link)
		}
	},
}

var LinkGetLinksCmd = &cobra.Command{
	Use:   "ls",
	Short: "Get a link(s)",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		if linkBoxId == "" {
			cfg := &config.ClientConfig{}
			cfg.Read()
			linkBoxId = cfg.DefaultBoxId
		}

		if linkBoxId == "" {
			fmt.Printf("Error: unable to determine box id.\n")
			os.Exit(1)
		}

		links, err := client.GetLinksByBoxId(context.Background(), &msg.Box{Id: linkBoxId})

		if err != nil {
			fmt.Printf("Error getting links: %v\n", err)
		} else {
			util.PrettyPrint(links)
		}
	},
}

var LinkLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load a link(s)",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)

		link, err := client.GetLinkById(context.Background(), &msg.Link{Id: linkId})

		if err != nil {
			fmt.Printf("Error getting link: %v\n", err)
		} else {
			open.Run(link.Url)
		}
	},
}

var LinkTagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Tag a link",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)
		link, err := client.GetLinkById(context.Background(), &msg.Link{Id: linkId, BoxId: linkBoxId})

		if err != nil {
			fmt.Printf("Error getting link: %v\n", err)
			os.Exit(1)
		}

		var tags []string
		if strings.Contains(linkTags, ",") {
			tags = strings.Split(linkTags, ",")
			var newTags []string
			for _, tag := range tags {
				newTags = append(newTags, strings.TrimSpace(tag))
			}
			tags = newTags
		} else {
			tags = []string{linkTags}
		}

		link.Tags = append(link.Tags, tags...)

		link, err = client.SaveLink(context.Background(), link)

		if err != nil {
			fmt.Printf("Error saving tags: %v\n", err)
		} else {
			util.PrettyPrint(link)
		}
	},
}

var LinkSearchCmd = &cobra.Command{
	Use:   "find",
	Short: "Find links based on tag search",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("127.0.0.1:5656", grpc.WithInsecure())

		if err != nil {
			panic(err)
		}

		defer conn.Close()

		client := msg.NewBoxServiceClient(conn)
		links, err := client.SearchLinks(context.Background(), &msg.Search{
			Term:  searchTerm,
			Count: int32(searchCount),
			Page:  int32(searchPage),
		})
		if err != nil {
			fmt.Printf("Error searching tags: %v\n", err)
		} else {
			util.PrettyPrint(links)
		}

		if searchLoadLinks {
			for _, link := range links.Links {
				open.Run(link.Url)
			}
		}
	},
}
