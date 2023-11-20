package enum

type BizType struct {
	Index       int    `json:"index"`
	Description string `json:"description"`
	NameEnUs    string `json:"nameEnUs"`
	NameArSa    string `json:"nameArSa"`
}

var (
	MsgTemplate = newBizType(1, "MSG - Message Template",
		"Message Template", "قالب الرسالة")
	PdpPositionInfo = newBizType(2, "PDP - Ad Position Info",
		"Ad Position Info", "معلومات موضع الإعلان")
	CpsChannelMerchant = newBizType(3, "CPS - Channel Merchant",
		"Channel Merchant", "تاجر القناة")
	TipPositionConfig = newBizType(4, "Tip Position Config",
		"Tip Position Config", "تكوين موضع النصيحة")
)

func newBizType(index int, description string, nameEnUs string, nameArSa string) BizType {
	biztype := new(BizType)
	biztype.Index = index
	biztype.Description = description
	biztype.NameEnUs = nameEnUs
	biztype.NameArSa = nameArSa
	return *biztype
}

// GetAllBizType 获取业务类型列表
func GetAllBizType() []BizType {
	var list []BizType
	list = append(list, MsgTemplate, PdpPositionInfo, CpsChannelMerchant, TipPositionConfig)
	return list
}
