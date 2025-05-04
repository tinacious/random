package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameService_GetFirstNamesFemale(t *testing.T) {
	fileInput := `Samantha
Sally`
	subject := NewNameService(
		WithFemaleNamesList(fileInput),
	)

	result := subject.GetFirstNamesFemale()
	expected := []string{"Samantha", "Sally"}

	assert.Equal(t, 2, len(result))
	assert.Equal(t, expected, result)
}

func TestNameService_RandomFirstNameFemale(t *testing.T) {
	fileInput := `Samantha
Sally`
	subject := NewNameService(
		WithFemaleNamesList(fileInput),
	)

	result := subject.RandomFirstNameFemale()
	expected := []string{"Samantha", "Sally"}

	assert.Contains(t, expected, result)
}

func TestNameService_GetFirstNamesMale(t *testing.T) {
	fileInput := `John
Bob
Noah
`
	subject := NewNameService(
		WithMaleNamesList(fileInput),
	)

	result := subject.GetFirstNamesMale()
	expected := []string{"John", "Bob", "Noah"}

	assert.Equal(t, 3, len(result))
	assert.Equal(t, expected, result)
}

func TestNameService_RandomFirstNameMale(t *testing.T) {
	fileInput := `John
Bob
Noah
`
	subject := NewNameService(
		WithMaleNamesList(fileInput),
	)

	result := subject.RandomFirstNameMale()
	expected := []string{"John", "Bob", "Noah"}

	assert.Contains(t, expected, result)
}

func TestNameService_GetFirstNamesNonBinary(t *testing.T) {
	fileInput := `Billie
Charlie
Devin
Frankie
`
	subject := NewNameService(
		WithNonBinaryNamesList(fileInput),
	)

	result := subject.GetFirstNamesNonBinary()
	expected := []string{"Billie", "Charlie", "Devin", "Frankie"}

	assert.Equal(t, 4, len(result))
	assert.Equal(t, expected, result)
}

func TestNameService_RandomFirstNameNonBinary(t *testing.T) {
	fileInput := `Billie
Charlie
Devin
Frankie
`
	subject := NewNameService(
		WithNonBinaryNamesList(fileInput),
	)

	result := subject.RandomFirstNameNonBinary()
	expected := []string{"Billie", "Charlie", "Devin", "Frankie"}

	assert.Contains(t, expected, result)
}

func TestNameService_GetLastNames(t *testing.T) {
	fileInput := `ROBINSON
NGUYEN
MITCHELL
REYES
COOK
`
	subject := NewNameService(
		WithLastNamesList(fileInput),
	)

	result := subject.GetLastNames()
	expected := []string{"Robinson", "Nguyen", "Mitchell", "Reyes", "Cook"}

	assert.Equal(t, 5, len(result))
	assert.Equal(t, expected, result)
}

func TestNameService_RandomLastName(t *testing.T) {
	fileInput := `ROBINSON
NGUYEN
MITCHELL
REYES
COOK
`
	subject := NewNameService(
		WithLastNamesList(fileInput),
	)

	result := subject.RandomLastName()
	expected := []string{"Robinson", "Nguyen", "Mitchell", "Reyes", "Cook"}

	assert.Contains(t, expected, result)
}
