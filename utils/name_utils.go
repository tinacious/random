package utils

import (
	_ "embed"
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
	RandomLastName() string
	RandomFirstNameFemale() string
	RandomFirstNameMale() string
	RandomFirstNameNonBinary() string
}

type nameService struct {
	female    []string
	male      []string
	nonBinary []string
	last      []string
}

func NewNameService(options ...func(*nameService)) NameService {
	s := &nameService{}
	for _, o := range options {
		o(s)
	}
	return s
}

func NewPopulatedService(options ...func(*NameService)) NameService {
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
func (s *nameService) RandomFirstNameFemale() string {
	names := s.GetFirstNamesFemale()
	name, err := RandomItemFromList(names)

	if err != nil {
		panic(err)
	}

	return *name
}

// RandomFirstNameMale
func (s *nameService) RandomFirstNameMale() string {
	names := s.GetFirstNamesMale()
	name, err := RandomItemFromList(names)

	if err != nil {
		panic(err)
	}

	return *name
}

// RandomFirstNameNonBinary
func (s *nameService) RandomFirstNameNonBinary() string {
	names := s.GetFirstNamesNonBinary()
	name, err := RandomItemFromList(names)

	if err != nil {
		panic(err)
	}

	return *name
}

// RandomLastName
func (s *nameService) RandomLastName() string {
	names := s.GetLastNames()
	name, err := RandomItemFromList(names)

	if err != nil {
		panic(err)
	}

	return *name
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
