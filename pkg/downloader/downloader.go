package downloader

type Downloader struct {
	url string `json:"url"`
}

// 初始化链接
func (d *Downloader) NewUrl(url string) *Downloader {
	d.url = url
	return d
}

func (d *Downloader) Start(url string) *Downloader {
	d.url = url
	return d
}

func (d *Downloader) Merge() error {
	return nil
}
