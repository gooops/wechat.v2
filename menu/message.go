package menu

// 菜单的按钮.
// NOTE: (MenuButton.Type, MenuButton.Key) 不能和 MenuButton.SubButton 不能同时设置,
// 我们不会自动为你检查这个, 你要自己来检查.
type MenuButton struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Key  string `json:"key,omitempty"`
	// SubButton 的个数不能超过 SubMenuButtonCountLimit
	SubButton []*MenuButton `json:"sub_button,omitempty"`
}

// 如果总的子按钮数超过限制, 则截除多余的.
func (mbtn *MenuButton) AppendButton(btn ...*MenuButton) {
	if len(btn) <= 0 {
		return
	}

	switch n := SubMenuButtonCountLimit - len(mbtn.SubButton); {
	case n > 0:
		if len(btn) > n {
			btn = btn[:n]
		}
		mbtn.SubButton = append(mbtn.SubButton, btn...)
	case n == 0:
	default: // n < 0
		mbtn.SubButton = mbtn.SubButton[:SubMenuButtonCountLimit]
	}
}

type Menu struct {
	Button []*MenuButton `json:"button"`
}

// 如果总的按钮数超过限制, 则截除多余的.
func (m *Menu) AppendButton(btn ...*MenuButton) {
	if len(btn) <= 0 {
		return
	}

	switch n := MenuButtonCountLimit - len(m.Button); {
	case n > 0:
		if len(btn) > n {
			btn = btn[:n]
		}
		m.Button = append(m.Button, btn...)
	case n == 0:
	default: // n < 0
		m.Button = m.Button[:MenuButtonCountLimit]
	}
}

// 获取自定义菜单的回复
type GetMenuResponse struct {
	Menu Menu `json:"menu"`
}
