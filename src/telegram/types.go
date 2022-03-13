package telegram

// SetWebhookParams represents parameters that can be passed to setWebhook method
type SetWebhookParams struct {
	// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40.
	MaxConnections int
	// A JSON-serialized list of the update types you want your bot to receive. For example, specify
	// ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	AllowedUpdates []string
	// PublicCertificateURL Upload your public key certificate so that the root certificate in use can be checked
	// Should contain URL to public certificate file or be empty string to ignore
	PublicCertificateURL string
}

type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`
	// Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address"`
}

// This object represents a Telegram user or bot.
type user struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// inlineQuery object represents an incoming inline query. When the user sends an empty query,
// your bot could return some default or trending results.
type inlineQuery struct {
	// ID is a unique identifier for this query
	ID string
	// From is a user-sender
	From user
	// Query is the text of the query (up to 256 characters)
	Query string
	// Offset of the results to be returned, can be controlled by the bot
	Offset string
}

// inputMessageContent Represents the content of a text message to be sent as the result of an inline query.
type inputMessageContent struct {
	// MessageText Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`
}

type inlineQueryResultArticle struct {
	// Type of the result, must be "article"
	Type string `json:"type"`
	// ID is a unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`
	// InputMessageContent Content of the message to be sent
	InputMessageContent inputMessageContent `json:"input_message_content"`
	Title               string              `json:"title"`
	Description         string              `json:"description"`
}

// answerInlineQueryParams object for sending answers to inline queries
type answerInlineQueryParams struct {
	InlineQueryID string                     `json:"inline_query_id"`
	Results       []inlineQueryResultArticle `json:"results"`
	CacheTime     int                        `json:"cache_time"`
	IsPersonal    bool                       `json:"is_personal"`
}

// This object represents an incoming update. At most one of the optional parameters can be present in any given update.
type update struct {
	// UpdateID is The update's unique identifier
	UpdateID int `json:"update_id"`
	// InlineQuery is New incoming inline query
	InlineQuery inlineQuery `json:"inline_query"`
}
