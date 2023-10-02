package drivinglicense

import (
	"errors"

	"github.com/google/uuid"
)

type Applicant interface {
	IsAdult() bool
	HoldDoubleLicenses() bool
}

type NumberGenerator struct{}

func NewNumberGenerator() NumberGenerator {
	return NumberGenerator{}
}

func (n NumberGenerator) Generate(applicant Applicant) (string, error) {
	if applicant.IsAdult() && !applicant.HoldDoubleLicenses() {
		// normal case. IsAdult and Not holding double licenses
		return uuid.New().String(), nil
	}

	if applicant.IsAdult() && applicant.HoldDoubleLicenses() {
		// error case: Adult that holding double licenses
		return "", errors.New("holding double licenses, you must not holding a license before")
	}

	if !applicant.IsAdult() && !applicant.HoldDoubleLicenses() {
		// error case: Underaged that not holding license
		return "", errors.New("underaged applicant, you must be 17 for a license")
	}
	return "", nil
}
