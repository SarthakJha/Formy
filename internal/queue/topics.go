package queue

type QueueTopic string

const(
	SHEETS QueueTopic = "formy-google-sheets"
	EMAIL_NOTIF QueueTopic = "formy-email-notif"
)