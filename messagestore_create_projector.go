package gomessagestore

import (
	"github.com/blackhatbrigade/gomessagestore/projector"
)

func (ms *msgStore) CreateProjector() projector.Projector {
	return projector.CreateProjector(ms.repo)
}
