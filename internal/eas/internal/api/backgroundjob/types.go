package backgroundjob

import "github.com/elevenode/terraform-provider-expo/internal/eas/internal/graphql"

type State string

const (
	StateQueued     State = "QUEUED"
	StateInProgress State = "IN_PROGRESS"
	StateSuccess    State = "SUCCESS"
	StateFailure    State = "FAILURE"
)

// Data is a receipt for an asynchronous job. EAS returns one from every
// schedule*Deletion mutation; poll it by id until State is terminal.
type Data struct {
	Id           string  `json:"id"`
	State        State   `json:"state"`
	Tries        int     `json:"tries"`
	WillRetry    bool    `json:"willRetry"`
	ResultId     *string `json:"resultId"`
	ResultType   string  `json:"resultType"`
	ErrorCode    *string `json:"errorCode"`
	ErrorMessage *string `json:"errorMessage"`
}

// Done reports whether the job reached a terminal state and will not retry.
func (data *Data) Done() bool {
	switch data.State {
	case StateSuccess:
		return true
	case StateFailure:
		return !data.WillRetry
	default:
		return false
	}
}

type Service interface {
	GetById(id string) (*Data, error)
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
