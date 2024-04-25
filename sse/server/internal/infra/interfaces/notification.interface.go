package interfaces

type NotificationInterface interface {
	// GetNotification() (interface{}, error)
	GetAll() ([]map[string]interface{}, error)
}