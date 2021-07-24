/*


1736. 替换隐藏数字得到的最晚时间
给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。

有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。

替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。


示例 1：

输入：time = "2?:?0"
输出："23:50"
解释：以数字 '2' 开头的最晚一小时是 23 ，以 '0' 结尾的最晚一分钟是 50 。


*/

func maximumTime(time string) string {

	timeStr := []byte(time)

	if timeStr[0] == '?' {
		if timeStr[1] == '?' || timeStr < '4' {
			timeStr[0] = '2'
		} else {
			timeStr[0] = '1'
		}
	}

	if timeStr[1] == '?' {
		if timeStr[0] == '2' {
			timeStr[1] = '3'
		} else {
			timeStr[1] = '9'
		}
	}

	if timeStr[3] == '?' {
		timeStr[3] = '5'
	}

	if timeStr[4] == '?' {
		timeStr[4] = '9'
	}
	return string(timeStr)

}

