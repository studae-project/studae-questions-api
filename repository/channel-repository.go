package repository

func GetAllowedChannelsToSendQuestions() map[string]string {
	return map[string]string{
		"java-kotlin":             "1008357078043201627",
		"outros":                  "asdsa",
		"frameworks":              "asdsa",
		"html-css-javascript-php": "sadsads",
	}
}

func GetAllowedChannelToSendQuestionByName(channelName string) (string, bool) {
	if val, ok := GetAllowedChannelsToSendQuestions()[channelName]; ok {
		return val, true
	}

	return "", false
}
