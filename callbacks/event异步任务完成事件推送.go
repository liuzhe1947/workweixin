package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90376#异步任务完成事件推送

func init() {
    // 添加可解析的回调事件
    supportCallback(EventBatchJobResult{})
}

// XML was generated 2021-10-09 14:46:10 by insomnia on Insomnia.lan.
type EventBatchJobResult struct {
    XMLName    xml.Name `xml:"xml"`
    Text       string   `xml:",chardata"`
    ToUserName struct {
        Text string `xml:",chardata"`
    } `xml:"ToUserName"`
    FromUserName struct {
        Text string `xml:",chardata"`
    } `xml:"FromUserName"`
    CreateTime struct {
        Text string `xml:",chardata"`
    } `xml:"CreateTime"`
    MsgType struct {
        Text string `xml:",chardata"`
    } `xml:"MsgType"`
    Event struct {
        Text string `xml:",chardata"`
    } `xml:"Event"`
    BatchJob struct {
        Text  string `xml:",chardata"`
        JobId struct {
            Text string `xml:",chardata"`
        } `xml:"JobId"`
        JobType struct {
            Text string `xml:",chardata"`
        } `xml:"JobType"`
        ErrCode struct {
            Text string `xml:",chardata"`
        } `xml:"ErrCode"`
        ErrMsg struct {
            Text string `xml:",chardata"`
        } `xml:"ErrMsg"`
    } `xml:"BatchJob"`
}

func (EventBatchJobResult) GetMessageType() string {
    return "event"
}

func (EventBatchJobResult) GetEventType() string {
    return "batch_job_result"
}

func (EventBatchJobResult) GetChangeType() string {
    return ""
}

func (m EventBatchJobResult) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventBatchJobResult) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp EventBatchJobResult
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
