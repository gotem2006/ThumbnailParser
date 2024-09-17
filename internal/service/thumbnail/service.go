package thumbnail_service


import(
	"fmt"
	"io"
	"net/http"

	"github.com/gotem2006/thumbnail/internal/cache"
)



type ThumbnailService interface{
	GetThumbnail(videoId string) (thumbnail *[]byte, err error) 
}

type thumbnailService struct{
	cache.Cache
}

func NewThumbnailService() ThumbnailService{
	return &thumbnailService{
		cache.NewCache(),
	}
}


func (t *thumbnailService) GetThumbnail(
	videoId string,
) (thumbnail *[]byte, err error) {
	thumbnail = t.GetThumbnailFromCache(videoId)
	if thumbnail == nil{
		thumbnail, err = getThumbnailFromYt(videoId)
		if err != nil{
			return nil, err
		}
		t.SetThumbnailToCache(videoId, thumbnail)
	}
	return thumbnail, nil 
}


func getThumbnailFromYt(videoId string) (*[]byte, error) {
	thumbnailUrl := fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", videoId)
	response, err := http.Get(thumbnailUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	thumbnail, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	return &thumbnail, nil
}
