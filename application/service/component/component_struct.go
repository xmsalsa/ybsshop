/*
 * @Created By:  宅职社 -- unknown
 * @Description: 不怕有BUG, 就怕你不背锅--
 * @Author: Unknown
 * @Date: 2021-04-14 17:36:21
 */
package component

type ComponentBox struct {
	Title  string  `json:"title"`
	Action string  `json:"action"`
	Method string  `json:"method"`
	Info   string  `json:"info"`
	Status bool    `json:"status"`
	Rules  []Rules `json:"rules"`
}

type Rules struct {
	Type     string     `json:"type"`
	Field    string     `json:"field"`
	Value    string     `json:"value"`
	Title    string     `json:"title,omitempty"`
	Props    Props      `json:"props,omitempty"`
	Options  []Options  `json:"options,omitempty"`
	Validate []Validate `json:"validate,omitempty"`
}

type Props struct {
	Type        string `json:"type,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
}

type Options struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type Validate struct {
	Message  string `json:"message"`
	Required bool   `json:"required"`
	Type     string `json:"type"`
	Trigger  string `json:"trigger"`
}
