package drivinglicense_test

import (
	"testing"

	"github.com/google/uuid"
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

func (u *UnderAgeApplicant) GetInitials() string {
	return ""
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

func (s *SecondLicenseHolderApplicant) GetInitials() string {
	return ""
}

// hardcode the return value
// dummy test stub for Normal license holder applicant
type HolderApplicant struct {
	initials string
}

func (h *HolderApplicant) IsAdult() bool {
	return true
}

func (h *HolderApplicant) HoldDoubleLicenses() bool {
	return false
}

func (h *HolderApplicant) GetInitials() string {
	return h.initials
}

// Fake logger to spy the logger contents
type SpyLogger struct {
	callCount   int
	lastMessage string
}

func (spy *SpyLogger) LogStuffs(v string) {
	spy.callCount++
	spy.lastMessage = v
}

// Fake RandomNumbersGenerator
type FakeRandomGenerator struct{}

func (fake *FakeRandomGenerator) GetRandomNumbers() string {
	return uuid.New().String()
}

// TestDrivingLicenseSuite
func (s *DrivingLicenseSuite) TestUnderAgeApplicant() {
	underAgeApplicant := &UnderAgeApplicant{}
	logger := &SpyLogger{}
	generator := &FakeRandomGenerator{}

	lg := drivinglicense.NewNumberGenerator(logger, generator)
	_, err := lg.Generate(underAgeApplicant)
	s.Error(err)
	s.Contains(err.Error(), "underaged applicant")
	// Spying technique on logger
	s.Equal(1, logger.callCount)
	s.Contains(logger.lastMessage, "underaged applicant")
}

// TestDrivingLicenseSuite
func (s *DrivingLicenseSuite) TestDoubleLicensesApplicant() {
	holdingDoubleLicensesApplicant := &SecondLicenseHolderApplicant{}
	logger := &SpyLogger{}
	generator := &FakeRandomGenerator{}

	lg := drivinglicense.NewNumberGenerator(logger, generator)
	_, err := lg.Generate(holdingDoubleLicensesApplicant)
	s.Error(err)
	s.Contains(err.Error(), "holding double licenses")
	// Spying on logger
	s.Equal(1, logger.callCount)
	s.Contains(logger.lastMessage, "holding double licenses")
}

// TestDrivingLicenseSuite
func (s *DrivingLicenseSuite) TestNormalLicenseApplicant() {
	normalLicensesApplicant := &HolderApplicant{}
	logger := &SpyLogger{}
	generator := &FakeRandomGenerator{}

	lg := drivinglicense.NewNumberGenerator(logger, generator)
	license, err := lg.Generate(normalLicensesApplicant)
	// yield no error
	s.NoError(err)
	// license is not empty
	s.NotZero(len(license))
	// Spying the logger
	s.Equal(1, logger.callCount)
	s.Contains(logger.lastMessage, "normal applicant")
}

// TestDrivingLicenseSuite
func (s *DrivingLicenseSuite) TestLicenseGeneration() {
	normalLicensesApplicant := &HolderApplicant{initials: "MDB"}
	logger := &SpyLogger{}
	generator := &FakeRandomGenerator{}

	lg := drivinglicense.NewNumberGenerator(logger, generator)
	license, err := lg.Generate(normalLicensesApplicant)
	s.NoError(err)
	s.NotZero(len(license))
	s.Contains(license, normalLicensesApplicant.GetInitials())
	// Spying on logger
	s.Contains(logger.lastMessage, "normal applicant")
}
