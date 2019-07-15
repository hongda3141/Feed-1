package model

type ClickLog struct {
	ClickId  string `json:"click_id,omitempty"`
	Uri      string `json:"uri,omitempty"`
	Ts       int64  `json:"ts,omitempty"`
	Ip       string `json:"ip,omitempty"`
	Ua       string `json:"ua,omitempty"`
	Country  string `json:"country,omitempty"`
	Platform string `json:"platform,omitempty"`
	OfferId  string `json:"offer_id,omitempty"`
	AdvId    string `json:"adv_id,omitempty"`
	AffId    string `json:"aff_id,omitempty"`
	Gaid     string `json:"gaid,omitempty"`
	Aid      string `json:"aid,omitempty"`
	Idfa     string `json:"idfa,omitempty"`
	SubId    string `json:"sub_id,omitempty"`
	Pn       string `json:"pn,omitempty"`
	Osv      string `json:"osv,omitempty"`
	AffSub1  string `json:"aff_sub1,omitempty"`
	AffSub2  string `json:"aff_sub2,omitempty"`
	AffSub3  string `json:"aff_sub3,omitempty"`
	AffSub4  string `json:"aff_sub4,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Status   bool   `json:"status,omitempty"`
}

type EsResponse struct {
	Took     int64 `json:"took"`
	TimedOut bool  `json:"timed_out"`
	_Shards  *Shards
	Hits     *Hits
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}
type Hits struct {
	Total    int64   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []*HitItems
}
type HitItems struct {
	Index  string      `json:"_index"`
	Type   string      `json:"_type"`
	Id     string      `json:"_id"`
	Score  float64     `json:"_score"`
	Source interface{} `json:"_source"`
}

type Query struct {
	Query *Match `json:"query"`
}
type Match struct {
	Match interface{} `json:"match"`
}
