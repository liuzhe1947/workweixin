package callbacks

import "encoding/xml"

// 自动生成的回调结构，按需修改, 生成方式: make callback doc=微信文档地址url
// 文档: https://open.work.weixin.qq.com/api/doc/90001/90143/92655#更新部门事件

func init() {
    // 添加可解析的回调事件
    supportCallback(ThirdChangeContactUpdateParty{})
}

// XML was generated 2021-10-30 09:21:09 by insomnia on Insomnia.lan.
type ThirdChangeContactUpdateParty struct {
    XMLName xml.Name `xml:"xml"`
    Text    string   `xml:",chardata"`
    SuiteId struct {
        Text string `xml:",chardata"`
    } `xml:"SuiteId"`
    AuthCorpId struct {
        Text string `xml:",chardata"`
    } `xml:"AuthCorpId"`
    InfoType struct {
        Text string `xml:",chardata"`
    } `xml:"InfoType"`
    TimeStamp struct {
        Text string `xml:",chardata"`
    } `xml:"TimeStamp"`
    ChangeType struct {
        Text string `xml:",chardata"`
    } `xml:"ChangeType"`
    ID struct {
        Text string `xml:",chardata"`
    } `xml:"Id"`
    Name struct {
        Text string `xml:",chardata"`
    } `xml:"Name"`
    ParentId struct {
        Text string `xml:",chardata"`
    } `xml:"ParentId"`
}

func (ThirdChangeContactUpdateParty) GetMessageType() string {
    return "third"
}

func (ThirdChangeContactUpdateParty) GetEventType() string {
    return "change_contact"
}

func (ThirdChangeContactUpdateParty) GetChangeType() string {
    return "update_party"
}

func (m ThirdChangeContactUpdateParty) GetTypeKey() string {
    return m.GetMessageType() + ":" + m.GetEventType() + ":" + m.GetChangeType()
}

func (ThirdChangeContactUpdateParty) ParseFromXml(data []byte) (CallBackExtraInfoInterface, error) {
    var temp ThirdChangeContactUpdateParty
    err := xml.Unmarshal(data, &temp)
    return temp, err
}
