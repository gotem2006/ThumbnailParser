package main

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/gotem2006/thumbnail/pkg/thumbnail"
)

var initFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "async",
		Aliases: []string{"q"},
		Usage:   "async thumbnails donwload",
	},
}

func fetchThumbnail(client pb.ThumbnailApiServiceClient, videoUrl string) error {
	resp, err := client.GetThumbnail(
		context.Background(),
		&pb.GetThumbnailRequset{
			Url: videoUrl,
		},
	)
	if err != nil {
		return err
	}
	thumbnail, err := resp.Recv()
	if err != nil {
		return err
	}
	if err := os.WriteFile(thumbnail.Filename, thumbnail.GetThumbnail(), 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	var async bool
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "async",
				Value:       false,
				Usage:       "async thumbnails donwload",
				Destination: &async,
			},
		},
		Name:  "thumbnail",
		Usage: "donwload thumbnails",
		Action: func(ctx *cli.Context) error {
			conn, err := grpc.NewClient(
				"0.0.0.0:8082",
				grpc.WithTransportCredentials(
					insecure.NewCredentials(),
				),
			)
			if err != nil {
				return err
			}
			defer conn.Close()

			client := pb.NewThumbnailApiServiceClient(conn)
			urls := ctx.Args().Slice()

			if !async {
				for _, url := range urls {
					if err := fetchThumbnail(client, url); err != nil{
						log.Printf(err.Error())
					}
				}
				return nil
			} else {

				ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				var wg sync.WaitGroup

				for _, url := range urls{
					wg.Add(1)
					go func(videoURL string) {
						defer wg.Done()
						select{
						case <-ctx.Done():
							log.Printf("Cancelled fetching thumbnail for %s due to timeout", videoURL)
						default:
							if err := fetchThumbnail(client, videoURL); err != nil{
								log.Printf(err.Error())
							}
						}
					}(url)
				}
				wg.Wait()
				return nil
			}
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}