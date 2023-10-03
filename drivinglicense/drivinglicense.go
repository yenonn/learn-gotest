package drivinglicense

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Applicant interface {
	IsAdult() bool
	HoldDoubleLicenses() bool
	GetInitials() string
}

type Logger interface {
	LogStuffs(v string)
}

type NumberGenerator struct {
	logger Logger
}

func NewNumberGenerator(logger Logger) NumberGenerator {
	return NumberGenerator{
		logger: logger,
	}
}

func (n NumberGenerator) Generate(applicant Applicant) (string, error) {
	if applicant.IsAdult() && !applicant.HoldDoubleLicenses() {
		// normal case. IsAdult and Not holding double licenses
		n.logger.LogStuffs("normal applicant")
		license := fmt.Sprintf("%s:%s", applicant.GetInitials(), uuid.New().String())
		return license, nil
	}

	if applicant.IsAdult() && applicant.HoldDoubleLicenses() {
		// error case: Adult that holding double licenses
		n.logger.LogStuffs("holding double licenses, you must not holding a licence before")
		return "", errors.New("holding double licenses, you must not holding a license before")
	}

	if !applicant.IsAdult() && !applicant.HoldDoubleLicenses() {
		// error case: Underaged that not holding license
		n.logger.LogStuffs("underaged applicant, you must be 17 for a license")
		return "", errors.New("underaged applicant, you must be 17 for a license")
	}
	return "", nil
}
