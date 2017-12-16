package util

var charList = []string{
	`1`, `2`, `3`, `4`, `5`, `6`, `7`,
	`8`, `9`, `a`, `b`, `c`, `d`, `e`, `f`,
	`g`, `h`, `i`, `j`, `k`, `l`, `m`, `n`,
	`p`, `q`, `r`, `s`, `t`, `u`, `v`,
	`w`, `x`, `y`, `z`,
}

//GetInviteCodeByUserID userId转8位34进制inviteCode
func GetInviteCodeByUserID(userID int64) string {
	if userID < 34 {
		return charList[userID]
	}
	var re int64
	var n string
	re = userID
	for re > 0 {
		n = n + charList[re%34]
		re = re / 34
	}
	return n
}

//GetUserIDByInviteCode 8位34进制inviteCode转userId
func GetUserIDByInviteCode(inviteCode string) int64 {
	var n, m int64
	i := 0
	for ; i < len(inviteCode); i++ {
		if m > 0 {
			m = m * 34
		} else {
			m = 1
		}
		for j, v := range charList {
			if string(inviteCode[i]) == v {
				n = n + int64(j)*m
			}
		}
	}
	return n
}
