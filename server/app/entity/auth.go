package entity

// TokenDetails jwtTokenをredisに保存する構造体
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails 取得したtokenIDを返す為の構造体
type AccessDetails struct {
	AccessUUID string
	UserID     uint64
}
