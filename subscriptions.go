package adn

type Subscription struct {
	Counts struct {
		Messages int `json:"messages"`
	} `json:"counts"`
	HasUnread bool `json:"has_unread"`
	Id int `json:"id"`
	Owner User `json:"owner"`
	Readers struct {
		AnyUser bool `json:"any_user"`
		Immutable bool `json:"immutable"`
		Public bool `json:"public"`
		UserIds []string `json:"user_ids"`
		You bool `json:"you"`
	} `json:"readers"`
	Type string `json:"type"`
	Writers struct {
		AnyUser bool `json:"any_user"`
		Immutable bool `json:"immutable"`
		Public bool `json:"public"`
		UserIds []string `json:"user_ids"`
		You bool `json:"you"`
	} `json:"writers"`
	YouCanEdit bool `json:"you_can_edit"`
	YouSubscribed bool `json:"you_subscribed"`
}

type SubscriptionList struct {
	Data []Subscription `json:"data"`
	Meta struct {
		Code int `json:"code"`
		MaxId int `json:"max_id"`
		MinId int `json:"min_id"`
		More bool `json:"more"`
	} `json:"meta"`
}
