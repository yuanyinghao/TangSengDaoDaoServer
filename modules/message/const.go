package message

const (
	// 消息已删除
	CMDMessageDeleted = "messageDeleted"
	// CMDMessageErase 消息擦除
	CMDMessageErase       = "messageEerase"
	sensitiveWordsVersion = 1
)
const ReadedCount = "readedCount" // 消息已读数量

type ReminderType int

const (
	ReminderTypeMentionMe      = 1 // 有人@我
	ReminderTypeApplyJoinGroup = 2 // 申请加群
)

var sensitive_words = []string{
	"银行卡",
	"微信",
	"qq",
	"密码",
	"支付宝",
	"钱包",
	"转账",
	"彩票",
	"股票",
	"人民币",
	"RMB",
	"兼职",
	"网络",
	"招聘",
	"有意者",
	"到货",
	"本店",
	"代购",
	"扣扣",
	"微店",
	"兼值",
	"淘宝",
	"小姐",
	"妓女",
	"包夜",
	"3P",
	"LY",
	"JS",
	"狼友",
	"技师",
	"推油",
	"胸推",
	"BT",
	"毒龙",
	"口爆",
	"楼凤",
	"足交",
	"口暴",
	"口交",
	"全套",
	"SM",
	"桑拿",
	"吞精",
	"咪咪",
	"婊子",
	"乳方",
	"操逼",
	"全职",
	"性伴侣",
	"网购",
	"网络工作",
	"代理",
	"专业代理",
	"帮忙点一下",
	"帮忙点下",
	"请点击进入",
	"详情请进入",
	"私人侦探",
	"私家侦探",
	"针孔摄象",
	"调查婚外情",
	"信用卡提现",
	"无抵押贷款",
	"广告代理",
	"原音铃声",
	"借腹生子",
	"找个妈妈",
	"找个爸爸",
	"代孕妈妈",
	"代生孩子",
	"代开发票",
	"腾讯客服电话",
	"销售热线",
	"免费订购热线",
	"低价出售",
	"款到发货",
	"回复可见",
	"连锁加盟",
	"加盟连锁",
	"免费二级域名",
	"免费使用",
	"免费索取",
	"蚁力神",
	"婴儿汤",
	"售肾",
	"刻章办",
	"买小车",
	"套牌车",
	"玛雅网",
	"电脑传讯",
	"视频来源",
	"下载速度",
	"高清在线",
	"全集在线",
	"在线播放",
	"txt下载",
	"六位qq",
	"6位qq",
	"位的qq",
	"个qb",
	"送qb",
	"用刀横向切腹",
	"完全自杀手册",
	"四海帮",
	"足球投注",
	"地下钱庄",
	"中国复兴党",
	"阿波罗网",
	"曾道人",
	"六合彩",
	"改卷内幕",
	"替考试",
	"隐形耳机",
	"出售答案",
	"考中答案",
	"答an",
	"da案",
	"资金周转",
	"救市",
	"股市圈钱",
	"崩盘",
	"资金短缺",
	"证监会",
	"质押贷款",
	"小额贷款",
	"周小川",
	"刘明康",
	"尚福林",
	"孔丹",
}
