package main

import (
	"image"
	"sync"
)

var (
	loadIconOnce sync.Once
	iconsV2      map[string]image.Image
)

func IconV2(name string) image.Image {
	// sync.Once Do 方法只会执行一次
	loadIconOnce.Do(loadIconsV2)

	return iconsV2[name]
}

func loadIconsV2() {
	icons = map[string]image.Image{
		// TODO
	}
}
