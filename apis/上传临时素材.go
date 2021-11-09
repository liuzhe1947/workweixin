package apis

import (
    "encoding/json"
    "errors"
    "fmt"
    "net/url"
)

// ReqUploadMedia 上传临时素材请求
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90389#上传临时素材

type ReqUploadMedia struct {
    // 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件（file）
    Type string `json:"type"`
    Media *Media `json:"media"`
}

var _ mediaUploader = ReqUploadMedia{}
var _ urlValuer = ReqUploadMedia{}

func (x ReqUploadMedia) intoURLValues() url.Values {
    var ret url.Values = make(map[string][]string)

    var vals map[string]interface{}
    jsonBytes, _ := json.Marshal(x)
    _ = json.Unmarshal(jsonBytes, &vals)

    for k, v := range vals {
        ret.Add(k, fmt.Sprintf("%v", v))
    }
    return ret
}

func (x ReqUploadMedia) getMedia() *Media {
    return x.Media
}

// RespUploadMedia 上传临时素材响应
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90389#上传临时素材
type RespUploadMedia struct {
    CommonResp
    Type      string `json:"type"`
    MediaId   string `json:"media_id"`
    CreatedAt string `json:"created_at"`
}

var _ bodyer = RespUploadMedia{}

func (x RespUploadMedia) intoBody() ([]byte, error) {
    result, err := json.Marshal(x)
    if err != nil {
        return nil, err
    }
    return result, nil
}

// execUploadMedia 上传临时素材
// 文档：https://work.weixin.qq.com/api/doc/90001/90143/90389#上传临时素材
func (c *ApiClient) ExecUploadMedia(req ReqUploadMedia) (RespUploadMedia, error) {
    var resp RespUploadMedia
    err := c.executeWXApiMediaUpload("/cgi-bin/media/upload", req, &resp, true)
    if err != nil {
        return RespUploadMedia{}, err
    }
    if bizErr := resp.TryIntoErr(); bizErr != nil {
        apiError, _ := bizErr.(*ClientError)
        if apiError.Code == ErrCode40011 {
            return RespUploadMedia{}, errors.New("发送素材过大，请使用链接发送")
        }
        return RespUploadMedia{}, bizErr
    }

    return resp, nil
}
