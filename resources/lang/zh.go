package lang

import (
	"gitee.com/pangxianfei/frame/helpers/locale"

	"gitee.com/pangxianfei/frame/resources/lang"
)

func init() {
	validationTranslation := lang.ValidationTranslation{
		Required: "{0}为必填字段",
		Len: lang.EmbeddedRule{
			String:  "{0}长度必须是{1}",
			Numeric: "{0}必须等于{1}",
			Array:   "{0}必须包含{1}",
		},
		Min: lang.EmbeddedRule{
			String:  "{0}长度必须至少为{1}",
			Numeric: "{0}最小只能为{1}",
			Array:   "{0}必须至少包含{1}",
		},
		Max: lang.EmbeddedRule{
			String:  "{0}长度不能超过{1}",
			Numeric: "{0}必须小于或等于{1}",
			Array:   "{0}最多只能包含{1}",
		},
		Eq: "{0}不等于{1}",
		Ne: "{0}不能等于{1}",
		Lt: lang.EmbeddedRule{
			String:   "{0}长度必须小于{1}",
			Numeric:  "{0}必须小于{1}",
			Array:    "{0}必须包含少于{1}",
			Datetime: "{0}必须小于当前日期和时间",
		},
		Lte: lang.EmbeddedRule{
			String:   "{0}长度不能超过{1}",
			Numeric:  "{0}必须小于或等于{1}",
			Array:    "{0}最多只能包含{1}",
			Datetime: "{0}必须小于或等于当前日期和时间",
		},
		Gt: lang.EmbeddedRule{
			String:   "{0}长度必须大于{1}",
			Numeric:  "{0}必须大于{1}",
			Array:    "{0}必须大于{1}",
			Datetime: "{0}必须大于当前日期和时间",
		},
		Gte: lang.EmbeddedRule{
			String:   "{0}长度必须至少为{1}",
			Numeric:  "{0}必须大于或等于{1}",
			Array:    "{0}必须至少包含{1}",
			Datetime: "{0}必须大于或等于当前日期和时间",
		},
		Eqfield:       "{0}必须等于{1}",
		Eqcsfield:     "{0}必须等于{1}",
		Necsfield:     "{0}不能等于{1}",
		Gtcsfield:     "{0}必须大于{1}",
		Gtecsfield:    "{0}必须大于或等于{1}",
		Ltcsfield:     "{0}必须小于{1}",
		Ltecsfield:    "{0}必须小于或等于{1}",
		Nefield:       "{0}不能等于{1}",
		Gtfield:       "{0}必须大于{1}",
		Gtefield:      "{0}必须大于或等于{1}",
		Ltfield:       "{0}必须小于{1}",
		Ltefield:      "{0}必须小于或等于{1}",
		Alpha:         "{0}只能包含字母",
		Alphanum:      "{0}只能包含字母和数字",
		Numeric:       "{0}必须是一个有效的数值",
		Number:        "{0}必须是一个有效的数字",
		Hexadecimal:   "{0}必须是一个有效的十六进制",
		Hexcolor:      "{0}必须是一个有效的十六进制颜色",
		Rgb:           "{0}必须是一个有效的RGB颜色",
		Rgba:          "{0}必须是一个有效的RGBA颜色",
		Hsl:           "{0}必须是一个有效的HSL颜色",
		Hsla:          "{0}必须是一个有效的HSLA颜色",
		Email:         "{0}必须是一个有效的邮箱",
		Url:           "{0}必须是一个有效的URL",
		Uri:           "{0}必须是一个有效的URI",
		Base64:        "{0}必须是一个有效的Base64字符串",
		Base64url:     "{0}必须是一个有效的Base64的URL",
		UrnRfc2141:    "{0}必须是一个有效的URN字符串",
		Contains:      "{0}必须包含文本'{1}'",
		Containsany:   "{0}必须包含至少一个以下字符'{1}'",
		Excludes:      "{0}不能包含文本'{1}'",
		Excludesall:   "{0}不能包含以下任何字符'{1}'",
		Excludesrune:  "{0}不能包含'{1}'",
		Isbn:          "{0}必须是一个有效的ISBN编号",
		Isbn10:        "{0}必须是一个有效的ISBN-10编号",
		Isbn13:        "{0}必须是一个有效的ISBN-13编号",
		Uuid:          "{0}必须是一个有效的UUID",
		Uuid3:         "{0}必须是一个有效的V3 UUID",
		Uuid4:         "{0}必须是一个有效的V4 UUID",
		Uuid5:         "{0}必须是一个有效的V5 UUID",
		Ascii:         "{0}必须只包含ascii字符",
		Printascii:    "{0}必须只包含可打印的ascii字符",
		Multibyte:     "{0}必须包含多字节字符",
		Datauri:       "{0}必须包含有效的数据URI",
		Latitude:      "{0}必须包含有效的纬度坐标",
		Longitude:     "{0}必须包含有效的经度坐标",
		Ssn:           "{0}必须是一个有效的社会安全号码(SSN)",
		Ipv4:          "{0}必须是一个有效的IPv4地址",
		Ipv6:          "{0}必须是一个有效的IPv6地址",
		Ip:            "{0}必须是一个有效的IP地址",
		Cidr:          "{0}必须是一个有效的无类别域间路由(CIDR)",
		Cidrv4:        "{0}必须是一个包含IPv4地址的有效无类别域间路由(CIDR)",
		Cidrv6:        "{0}必须是一个包含IPv6地址的有效无类别域间路由(CIDR)",
		TcpAddr:       "{0}必须是一个有效的TCP地址",
		Tcp4Addr:      "{0}必须是一个有效的IPv4 TCP地址",
		Tcp6Addr:      "{0}必须是一个有效的IPv6 TCP地址",
		UdpAddr:       "{0}必须是一个有效的UDP地址",
		Udp4Addr:      "{0}必须是一个有效的IPv4 UDP地址",
		Udp6Addr:      "{0}必须是一个有效的IPv6 UDP地址",
		IpAddr:        "{0}必须是一个有效的IP地址",
		Ip4Addr:       "{0}必须是一个有效的IPv4地址",
		Ip6Addr:       "{0}必须是一个有效的IPv6地址",
		UnixAddr:      "{0}必须是一个有效的UNIX地址",
		Mac:           "{0}必须是一个有效的MAC地址",
		Unique:        "{0}必须是一个唯一值",
		Iscolor:       "{0}必须是一个有效的颜色",
		Oneof:         "{0}必须是[{1}]中的一个",
		BtcAddr:       "{0}必须是一个有效的BTC地址",
		BtcAddrBech32: "{0}必须是一个有效的BTC Bech32地址",
		EthAddr:       "{0}必须是一个有效的ETH地址",

		PluralRuleMap: map[string]lang.PluralRule{
			"character": {
				One:   "字符",
				Other: "字符",
			},
			"item": {
				One:   "集合",
				Other: "集合",
			},
		},
	}
	customTranslation := lang.CustomTranslation{
		"auth.register.failed_existed":              "user register failed, user has been registered before",
		"auth.register.failed_system_error":         "user register failed, system error",
		"auth.register.failed_token_generate_error": "user register failed, token generate failed",

		"auth.login.failed_not_exist":            "user login failed, user doesn't exist",
		"auth.login.failed_wrong_password":       "user login failed, user password incorrect",
		"auth.login.failed_token_generate_error": "user login failed, token generate failed",
	}
	locale.AddLocale("zh", &customTranslation, &validationTranslation)
}
