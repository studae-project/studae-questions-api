package repository

func GetAllowedChannelsToSendQuestions() map[string]string {
	return map[string]string{
		"java-kotlin":             "655390389498281984",
		"outros":                  "769362287261122610",
		"frameworks":              "766039785646325782",
		"html-css-javascript-php": "766039622400344085",
		"testing":                 "1011731564750721084",
	}
}

func GetAllowedChannelToSendQuestionByName(channelName string) (string, bool) {
	if val, ok := GetAllowedChannelsToSendQuestions()[channelName]; ok {
		return val, true
	}

	return "", false
}
