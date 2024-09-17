package api

import (
	"fmt"
	"net/url"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	thumbnail_service "github.com/gotem2006/thumbnail/internal/service/thumbnail"
	pb "github.com/gotem2006/thumbnail/pkg/thumbnail"
)



type thumbnailAPI struct {
	pb.UnimplementedThumbnailApiServiceServer
	service thumbnail_service.ThumbnailService
}

func NewThumbnailAPI() pb.ThumbnailApiServiceServer {
	return &thumbnailAPI{
		service: thumbnail_service.NewThumbnailService(),
	}
}

func (t *thumbnailAPI) GetThumbnail(
	req *pb.GetThumbnailRequset,
	stream pb.ThumbnailApiService_GetThumbnailServer,
) (error) {
	videoId, err := getVideoId(req.Url)
	if err != nil{
		log.Error().Err(err).Msg("GetThumbnail - invaild argument")
		return status.Error(codes.InvalidArgument, err.Error())
	}
	thumbnail, err := t.service.GetThumbnail(videoId)
	if err != nil{
		log.Error().Err(err).Msg("GetThumbnail - failed get thumbnail from yt")
		return status.Error(codes.NotFound, err.Error())
	}
	if err := stream.Send(&pb.GetThumbnailResponse{Thumbnail: *thumbnail, Filename: videoId + ".jpg"});err != nil{
		log.Error().Err(err).Msg("GetThumbnail - failed send thumbnail")
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func getVideoId(Url string) (string, error) {
	videoUrl, err := url.Parse(Url)
	if err != nil{
		return "", fmt.Errorf("Failed parse url - %s", videoUrl)
	}
	videoId, ok := videoUrl.Query()["v"]
	if !ok{
		return "", fmt.Errorf("Failed get video id from url - %s", videoUrl)
	}
	return videoId[0], nil
}

