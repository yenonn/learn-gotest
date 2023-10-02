package drivinglicense_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yenonn/learn-gotest/drivinglicense"
)

type DrivingLicenseSuite struct {
	suite.Suite
}

func TestDrivingLicenseSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenseSuite))
}

// dummy test stub for UnderAgeApplicant
// hardcode the return value
type UnderAgeApplicant struct{}

func (u *UnderAgeApplicant) IsAdult() bool {
	return false
}

func (u *UnderAgeApplicant) HoldDoubleLicenses() bool {
	return false
}

// hardcode the return value
// dummy test stub for SecondLicenseHolderApplicant
type SecondLicenseHolderApplicant struct{}

func (s *SecondLicenseHolderApplicant) IsAdult() bool {
	return true
}

func (s *SecondLicenseHolderApplicant) HoldDoubleLicenses() bool {
	return true
}

// hardcode the return value
// dummy test stub for Normal license holder applicant
type HolderApplicant struct{}

func (h *HolderApplicant) IsAdult() bool {
	return true
}

func (h *HolderApplicant) HoldDoubleLicenses() bool {
	return false
}

func (s *DrivingLicenseSuite) TestUnderAgeApplicant() {
	underAgeApplicant := &UnderAgeApplicant{}
	lg := drivinglicense.NewNumberGenerator()
	_, err := lg.Generate(underAgeApplicant)
	s.Error(err)
	s.Contains(err.Error(), "underaged applicant")
}

func (s *DrivingLicenseSuite) TestDoubleLicensesApplicant() {
	holdingDoubleLicensesApplicant := &SecondLicenseHolderApplicant{}
	lg := drivinglicense.NewNumberGenerator()
	_, err := lg.Generate(holdingDoubleLicensesApplicant)
	s.Error(err)
	s.Contains(err.Error(), "holding double licenses")
}

func (s *DrivingLicenseSuite) TestNormalLicenseApplicant() {
	normalLicensesApplicant := &HolderApplicant{}
	lg := drivinglicense.NewNumberGenerator()
	license, err := lg.Generate(normalLicensesApplicant)
	// yield no error
	s.NoError(err)
	// license is not empty
	s.NotZero(len(license))
}
