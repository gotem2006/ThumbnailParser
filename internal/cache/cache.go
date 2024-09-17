package cache

import (
	"sync"
)

type Cache interface{
	GetThumbnailFromCache(videoId string) *[]byte
	SetThumbnailToCache(videoId string, thumbnail *[]byte)
}


type cache struct {
	thumbnails map[string]*[]byte
	mu         *sync.Mutex
}

func NewCache() *cache {
	return &cache{
		thumbnails: make(map[string]*[]byte),
		mu:         &sync.Mutex{},
	}
}

func (c *cache) GetThumbnailFromCache(videoId string) *[]byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	thumbnail, found := c.thumbnails[videoId]
	if !found {
		return nil
	}
	return thumbnail
}

func (c *cache) SetThumbnailToCache(videoId string, thumbnail *[]byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.thumbnails[videoId] = thumbnail
}
