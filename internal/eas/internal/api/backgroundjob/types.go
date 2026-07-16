package backgroundjob

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

// State mirrors the BackgroundJobState enum of the EAS GraphQL schema.
type State string

const (
	StateQueued     State = "QUEUED"
	StateInProgress State = "IN_PROGRESS"
	StateSuccess    State = "SUCCESS"
	StateFailure    State = "FAILURE"
)

type Data struct {
	Id    string `json:"id"`
	State State  `json:"state"`
	Tries int    `json:"tries"`
	// WillRetry reports whether EAS will retry a FAILURE, which makes it non-terminal.
	WillRetry    bool   `json:"willRetry"`
	ResultId     string `json:"resultId"`
	ResultType   string `json:"resultType"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type Service interface {
	Get(id string) (*Data, error)
}

type service struct {
	graphql graphql.GraphQL
}

var _ Service = (*service)(nil)

func NewService(graphql graphql.GraphQL) Service {
	return &service{
		graphql: graphql,
	}
}
