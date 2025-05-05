package utils

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed res/names_first_f.txt
var namesF string

//go:embed res/names_first_m.txt
var namesM string

//go:embed res/names_first_u.txt
var namesU string

//go:embed res/names_last.txt
var namesLast string

type NameService interface {
	GetFirstNamesFemale() []string
	GetFirstNamesMale() []string
	GetFirstNamesNonBinary() []string
	GetLastNames() []string

	RandomFirstNameFemale() (string, error)
	RandomFirstNameMale() (string, error)
	RandomFirstNameNonBinary() (string, error)
	RandomFirstNameAny() (string, error)
	RandomFirstName(NameGender) (string, error)
	RandomLastName() (string, error)
}

type nameService struct {
	female    []string
	male      []string
	nonBinary []string
	last      []string
}

type NameGender int

const (
	NameGenderUnisex NameGender = iota
	NameGenderFemale
	NameGenderMale
	NameGenderAny
)

var SupportedGenders = []NameGender{
	NameGenderUnisex,
	NameGenderFemale,
	NameGenderMale,
	NameGenderAny,
}

func NameGenderFromString(value string) NameGender {
	if value == "f" {
		return NameGenderFemale
	}
	if value == "m" {
		return NameGenderMale
	}
	if value == "u" {
		return NameGenderUnisex
	}

	// Invalid option defaults to any
	return NameGenderAny
}

func (g NameGender) Key() (string, error) {
	if g == NameGenderFemale {
		return "f", nil
	}
	if g == NameGenderMale {
		return "m", nil
	}
	if g == NameGenderUnisex {
		return "u", nil
	}
	if g == NameGenderAny {
		return "a", nil
	}

	return "", fmt.Errorf("unsupported gender: %d", g)
}

func (g NameGender) Description() (string, error) {
	if g == NameGenderFemale {
		return "female", nil
	}
	if g == NameGenderMale {
		return "male", nil
	}
	if g == NameGenderUnisex {
		return "unisex", nil
	}
	if g == NameGenderAny {
		return "any/all", nil
	}

	return "", fmt.Errorf("unsupported gender: %d", g)
}

func AllNameGenderKeys() ([]string, error) {
	keys := make([]string, len(SupportedGenders))

	for i, g := range SupportedGenders {
		k, err := g.Key()
		if err != nil {
			return nil, err
		}
		keys[i] = k
	}

	return keys, nil
}

func AllNameGenderDescriptions() ([]string, error) {
	keys := make([]string, len(SupportedGenders))

	for i, g := range SupportedGenders {
		k, err := g.Description()
		if err != nil {
			return nil, err
		}
		keys[i] = k
	}

	return keys, nil
}

func AllNameGenderDescriptionsWithKeys() ([]string, error) {
	keys := make([]string, len(SupportedGenders))

	for i, g := range SupportedGenders {
		k, err := g.DescriptionWithKey()
		if err != nil {
			return nil, err
		}
		keys[i] = k
	}

	return keys, nil
}

func (g NameGender) DescriptionWithKey() (string, error) {
	key, err := g.Key()
	if err != nil {
		return "", fmt.Errorf("unsupported gender: %d", g)
	}

	description, err := g.Description()
	if err != nil {
		return "", fmt.Errorf("unsupported gender: %d", g)
	}

	return fmt.Sprintf("%s = %s", key, description), nil
}

func NewNameService(options ...func(*nameService)) NameService {
	s := &nameService{}
	for _, o := range options {
		o(s)
	}
	return s
}

func NewPopulatedService(options ...func(*nameService)) NameService {
	return NewNameService(
		WithFemaleNamesList(namesF),
		WithMaleNamesList(namesM),
		WithNonBinaryNamesList(namesU),
		WithLastNamesList(namesLast),
	)
}

func WithFemaleNamesList(namesList string) func(*nameService) {
	return func(s *nameService) {
		s.female = namesFromTextList(namesList)
	}
}

func WithMaleNamesList(namesList string) func(*nameService) {
	return func(s *nameService) {
		s.male = namesFromTextList(namesList)
	}
}

func WithNonBinaryNamesList(namesList string) func(*nameService) {
	return func(s *nameService) {
		s.nonBinary = namesFromTextList(namesList)
	}
}

func WithLastNamesList(namesList string) func(*nameService) {
	return func(s *nameService) {
		s.last = namesFromTextList(namesList)
	}
}

// GetFirstNamesFemale returns 1000 names that are considered female according to the source
func (s *nameService) GetFirstNamesFemale() []string {
	return s.female
}

// GetFirstNamesMale returns 1000 names that are considered male according to the source
func (s *nameService) GetFirstNamesMale() []string {
	return s.male
}

// GetFirstNamesNonBinary returns 300 names that are considered unisex according to the source
func (s *nameService) GetFirstNamesNonBinary() []string {
	return s.nonBinary
}

// GetLastNames returns 1000 last names
func (s *nameService) GetLastNames() []string {
	return s.last
}

// RandomFirstNameFemale
func (s *nameService) RandomFirstNameFemale() (string, error) {
	names := s.GetFirstNamesFemale()
	name, err := RandomItemFromList(names)

	if err != nil {
		return "", err
	}

	return *name, nil
}

// RandomFirstNameMale
func (s *nameService) RandomFirstNameMale() (string, error) {
	names := s.GetFirstNamesMale()
	name, err := RandomItemFromList(names)

	if err != nil {
		return "", err
	}

	return *name, nil
}

// RandomFirstNameNonBinary
func (s *nameService) RandomFirstNameNonBinary() (string, error) {
	names := s.GetFirstNamesNonBinary()
	name, err := RandomItemFromList(names)

	if err != nil {
		return "", err
	}

	return *name, nil
}

// RandomFirstNameAny returns a first name of all available names (female, male, unisex)
func (s *nameService) RandomFirstNameAny() (string, error) {
	allNames := append(append(s.GetFirstNamesFemale(), s.GetFirstNamesMale()...), s.GetFirstNamesNonBinary()...)
	name, err := RandomItemFromList(allNames)

	if err != nil {
		return "", err
	}

	return *name, nil
}

// RandomFirstName returns a random first name based on the gender
func (s *nameService) RandomFirstName(gender NameGender) (string, error) {
	if gender == NameGenderFemale {
		return s.RandomFirstNameFemale()
	}
	if gender == NameGenderMale {
		return s.RandomFirstNameMale()
	}
	if gender == NameGenderUnisex {
		return s.RandomFirstNameNonBinary()
	}

	return s.RandomFirstNameAny()
}

// RandomLastName
func (s *nameService) RandomLastName() (string, error) {
	names := s.GetLastNames()
	name, err := RandomItemFromList(names)

	if err != nil {
		return "", err
	}

	return *name, nil
}

// namesFromTextList Parses the name list files in the res directory. Strips out new lines and lines that begin with #
func namesFromTextList(contents string) []string {
	re := regexp.MustCompile(`(?m)^\s*#.*\n?`)
	noComments := re.ReplaceAllString(contents, "")

	return fixCase(strings.Split(TrimEmptyLines(noComments), "\n"))
}

func fixCase(items []string) []string {
	fixed := make([]string, len(items))

	for i, v := range items {
		fixed[i] = ProperNounCase(v)
	}

	return fixed
}
