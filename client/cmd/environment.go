package cmd

type environmentSource struct {
	Key string `json:"key"`
}

type environment struct {
	ID        string            `json:"_id"`
	Key       string            `json:"key"`
	Name      string            `json:"name"`
	Color     string            `json:"color"`
	Source    environmentSource `json:"source"`
	ApiKey    string            `json:"apiKey"`
	MobileKey string            `json:"mobileKey"`
}
