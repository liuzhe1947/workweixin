package callbacks

import "encoding/xml"

// 文档: https://work.weixin.qq.com/api/doc/90001/90143/90376#成员关注及取消关注事件

func init() {
    // 添加可解析的回调事件
    supportCallback(EventSubscribe{})
}

// XML was generated 2021-10-09 14:46:10 by insomnia on Insomnia.lan.
type EventSubscribe struct {
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
    AgentID struct {
        Text string `xml:",chardata"`
    } `xml:"AgentID"`
}

func (EventSubscribe) GetMessageType() string {
    return "event"
}

func (EventSubscribe) GetEventType() string {
    return "subscribe"
}

func (EventSubscribe) GetChangeType() string {
    return ""
}

func (m EventSubscribe) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (EventSubscribe) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp EventSubscribe
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
