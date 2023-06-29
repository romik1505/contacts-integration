package errors

type AmoCRMError struct {
	Hint    string `json:"hint"`
	Title   string `json:"title"`
	Type    string `json:"type"`
	Status  int    `json:"status"`
	Details string `json:"details"`
}

func (c AmoCRMError) Error() string {
	return c.Title
}
