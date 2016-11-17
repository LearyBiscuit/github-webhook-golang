package webhook

import (
	"encoding/json"
)

// User represents a Github user.
type User struct {
	ID					string
	Login				string
}

// Repository represents a Github's repository.
type Repository struct {
	Name        string
	FullName    string
	Private     bool
	HTMLURL     string `json:"html_url"`
	Description string
	Fork        bool
	URL         string
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	PushedAt    int64  `json:"pushed_at"`
}

// Commit represents a Github's commit.
type Commit struct {
	ID        string
	Distinct  bool
	Message   string
	Timestamp string
	URL       string
	Added     []string
	Removed   []string
	Modified  []string
}

// PullRequest represents a Github's pull request.
type PullRequest struct {
	ID         string
	URL        string
	State      string
	Title      string
	User       []User
	Body       string
	Repository Repository
}
// PushEvent represents a Github's webhook push event.
type PushEvent struct {
	Ref        string
	Before     string
	After      string
	Created    bool
	Deleted    bool
	Forced     bool
	Compare    string
	Commits    []Commit
	Repository Repository
}

// PullRequestEvent represents a Github's pull request event.
type PullRequestEvent struct {
	Action      string
	Number      string
	PullRequest PullRequest
}

func (r *Event) GetEvent *string {
	return r.Header.EventType;
}

// PushEvent returns a PushEvent struct.
func (r *Event) PushEvent() *PushEvent {
	if r.Header.EventType != "push" {
		return nil
	}
	event := new(PushEvent)
	if err := json.Unmarshal(r.Body, event); err != nil {
		logErr(err)
		return nil
	}
	return event
}

// PUllRequestEvent returns a PullRequestEvent struct.
func (r *Event) PullRequestEvent() *PullRequestEvent {
	if r.Header.EventType != "pull_request" {
		return nil
	}
	event := new(PullRequestEvent)
	if err := json.Unmarshal(r.Body, event); err != nil {
		logErr(err)
		return nil
	}
	return event
}
