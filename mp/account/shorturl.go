package account

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 将一条长链接转成短链接.
func ShortURL(clt *core.Client, longURL string) (shortURL string, err error) {
	var request = struct {
		Action  string `json:"action"`
		LongURL string `json:"long_url"`
	}{
		Action:  "long2short",
		LongURL: longURL,
	}

	var result struct {
		core.Error
		ShortURL string `json:"short_url"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/shorturl?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	shortURL = result.ShortURL
	return
}