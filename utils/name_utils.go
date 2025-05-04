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

type NameService struct {
	female    []string
	male      []string
	nonBinary []string
	last      []string
}

func NewNameService() NameService {
	return NameService{
		female:    namesFromNameFileContent(namesF),
		male:      namesFromNameFileContent(namesM),
		nonBinary: namesFromNameFileContent(namesU),
		last:      namesFromNameFileContent(namesLast),
	}
}

func (s *NameService) RandomName() (string, error) {
	return "", nil
}

// GetFirstNamesFemale returns 1000 names that are considered female according to the source
func (s *NameService) GetFirstNamesFemale() []string {
	return s.female
}

// GetFirstNamesMale returns 1000 names that are considered male according to the source
func (s *NameService) GetFirstNamesMale() []string {
	return s.male
}

// GetFirstNamesNonBinary returns 300 names that are considered unisex according to the source
func (s *NameService) GetFirstNamesNonBinary() []string {
	return s.nonBinary
}

// GetLastNames returns 1000 last names
func (s *NameService) GetLastNames() []string {
	return s.last
}

// namesFromNameFileContent Parses the name list files in the res directory. Strips out new lines and lines that begin with #
func namesFromNameFileContent(contents string) []string {
	re := regexp.MustCompile(`(?m)^\s*#.*\n?`)
	noComments := re.ReplaceAllString(contents, "")

	return strings.Split(TrimEmptyLines(noComments), "\n")
}
