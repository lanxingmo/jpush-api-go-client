package jpushclient

type Intent struct {
	URL string `json:"url,omitempty"`
}
type Notice struct {
	Alert    string          `json:"alert,omitempty"`
	Android  *AndroidNotice  `json:"android,omitempty"`
	IOS      *IOSNotice      `json:"ios,omitempty"`
	WINPhone *WinPhoneNotice `json:"winphone,omitempty"`
	Hmos     *Hmos           `json:"hmos"`
	Voip     *Voip           `json:"voip"`
	QuickApp *QuickApp       `json:"quickapp"`
}

type AndroidNotice struct {
	Alert     string `json:"alert"`
	Title     string `json:"title,omitempty"`
	BuilderId int    `json:"builder_id,omitempty"`
	LargeIcon string `json:"large_icon,omitempty"`

	Intent map[string]interface{} `json:"intent,omitempty"`
	Extras map[string]interface{} `json:"extras,omitempty"`
}

type IOSNotice struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            string                 `json:"badge,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
}

type WinPhoneNotice struct {
	Alert    string                 `json:"alert"`
	Title    string                 `json:"title,omitempty"`
	OpenPage string                 `json:"_open_page,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

// 定义 Hmos 结构体
type Hmos struct {
	Alert       string                 `json:"alert"`
	Title       string                 `json:"title"`
	Intent      map[string]interface{} `json:"intent,omitempty"`
	BadgeAddNum int                    `json:"badge_add_num"`
	BadgeSetNum int                    `json:"badge_set_num"`
	Extras      map[string]interface{} `json:"extras"`
	Category    string                 `json:"category"`
	TestMessage bool                   `json:"test_message"`
	ReceiptID   string                 `json:"receipt_id"`
	LargeIcon   string                 `json:"large_icon"`
	Style       int                    `json:"style"`
	PushType    int                    `json:"push_type"`
	// 如果有其他可选字段，可以在这里添加
}

// 定义 Voip 结构体
type Voip map[string]interface{} // 使用 map 来处理任意自定义 key/value 对

// 定义 QuickApp 结构体
type QuickApp struct {
	Alert string `json:"alert"`
	Title string `json:"title"`
	Page  string `json:"page"`
}

func (this *Notice) SetAlert(alert string) {
	this.Alert = alert
}

func (this *Notice) SetAndroidNotice(n *AndroidNotice) {
	this.Android = n
}

func (this *Notice) SetIOSNotice(n *IOSNotice) {
	this.IOS = n
}

func (this *Notice) SetWinPhoneNotice(n *WinPhoneNotice) {
	this.WINPhone = n
}
