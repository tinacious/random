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

	result, err := subject.RandomFirstNameFemale()
	expected := []string{"Samantha", "Sally"}

	assert.Nil(t, err)
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

	result, err := subject.RandomFirstNameMale()
	expected := []string{"John", "Bob", "Noah"}

	assert.Nil(t, err)
	assert.Contains(t, expected, result)
}

func TestNameService_GetFirstNamesUnisex(t *testing.T) {
	fileInput := `Billie
Charlie
Devin
Frankie
`
	subject := NewNameService(
		WithUnisexNamesList(fileInput),
	)

	result := subject.GetFirstNamesUnisex()
	expected := []string{"Billie", "Charlie", "Devin", "Frankie"}

	assert.Equal(t, 4, len(result))
	assert.Equal(t, expected, result)
}

func TestNameService_RandomFirstNameUnisex(t *testing.T) {
	fileInput := `Billie
Charlie
Devin
Frankie
`
	subject := NewNameService(
		WithUnisexNamesList(fileInput),
	)

	result, err := subject.RandomFirstNameUnisex()
	expected := []string{"Billie", "Charlie", "Devin", "Frankie"}

	assert.Nil(t, err)
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

	result, err := subject.RandomLastName()
	expected := []string{"Robinson", "Nguyen", "Mitchell", "Reyes", "Cook"}

	assert.Nil(t, err)
	assert.Contains(t, expected, result)
}

func TestNameGenderKey_NameGenderFemale(t *testing.T) {
	result, err := NameGenderFemale.Key()

	assert.Nil(t, err)
	assert.Equal(t, "f", result)
}

func TestNameGenderKey_NameGenderMale(t *testing.T) {
	result, err := NameGenderMale.Key()

	assert.Nil(t, err)
	assert.Equal(t, "m", result)
}

func TestNameGenderKey_NameGenderUnisex(t *testing.T) {
	result, err := NameGenderUnisex.Key()

	assert.Nil(t, err)
	assert.Equal(t, "u", result)
}

func TestNameGenderKey_NameGenderAny(t *testing.T) {
	result, err := NameGenderAny.Key()

	assert.Nil(t, err)
	assert.Equal(t, "a", result)
}

func TestNameGenderKey_InvalidNameGender(t *testing.T) {
	var input NameGender = 999
	result, err := input.Key()

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestNameGenderDescription_NameGenderFemale(t *testing.T) {
	result, err := NameGenderFemale.Description()

	assert.Nil(t, err)
	assert.Equal(t, "female", result)
}

func TestNameGenderDescription_NameGenderMale(t *testing.T) {
	result, err := NameGenderMale.Description()

	assert.Nil(t, err)
	assert.Equal(t, "male", result)
}

func TestNameGenderDescription_NameGenderUnisex(t *testing.T) {
	result, err := NameGenderUnisex.Description()

	assert.Nil(t, err)
	assert.Equal(t, "unisex", result)
}

func TestNameGenderDescription_NameGenderAny(t *testing.T) {
	result, err := NameGenderAny.Description()

	assert.Nil(t, err)
	assert.Equal(t, "any/all", result)
}

func TestNameGenderDescriptionWithKey_NameGenderFemale(t *testing.T) {
	result, err := NameGenderFemale.DescriptionWithKey()

	assert.Nil(t, err)
	assert.Equal(t, "f = female", result)
}

func TestNameGenderDescriptionWithKey_NameGenderMale(t *testing.T) {
	result, err := NameGenderMale.DescriptionWithKey()

	assert.Nil(t, err)
	assert.Equal(t, "m = male", result)
}

func TestNameGenderDescriptionWithKey_NameGenderUnisex(t *testing.T) {
	result, err := NameGenderUnisex.DescriptionWithKey()

	assert.Nil(t, err)
	assert.Equal(t, "u = unisex", result)
}

func TestNameGenderDescriptionWithKey_NameGenderAny(t *testing.T) {
	result, err := NameGenderAny.DescriptionWithKey()

	assert.Nil(t, err)
	assert.Equal(t, "a = any/all", result)
}

func TestNameGenderDescriptionWithKey_InvalidGender(t *testing.T) {
	var input NameGender = 999
	result, err := input.DescriptionWithKey()

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestNameGenderDescription_InvalidNameGender(t *testing.T) {
	var input NameGender = 999
	result, err := input.Description()

	assert.NotNil(t, err)
	assert.Empty(t, result)
}

func TestNameGenderFromString_WithF(t *testing.T) {
	input := "f"

	result := NameGenderFromString(input)

	assert.Equal(t, NameGenderFemale, result)
}

func TestNameGenderFromString_WithM(t *testing.T) {
	input := "m"

	result := NameGenderFromString(input)

	assert.Equal(t, NameGenderMale, result)
}

func TestNameGenderFromString_WithU(t *testing.T) {
	input := "u"

	result := NameGenderFromString(input)

	assert.Equal(t, NameGenderUnisex, result)
}

func TestNameGenderFromString_WithA(t *testing.T) {
	input := "a"

	result := NameGenderFromString(input)

	assert.Equal(t, NameGenderAny, result)
}

func TestNameGenderFromString_WithInvalidInput(t *testing.T) {
	input := "x"

	result := NameGenderFromString(input)

	assert.Equal(t, NameGenderAny, result)
}

func TestAllGenderKeys(t *testing.T) {
	expected := []string{"u", "f", "m", "a"}

	result, err := AllNameGenderKeys()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestAllGenderDescriptions(t *testing.T) {
	expected := []string{"unisex", "female", "male", "any/all"}

	result, err := AllNameGenderDescriptions()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestAllGenderDescriptionsWithKeys(t *testing.T) {
	expected := []string{"u = unisex", "f = female", "m = male", "a = any/all"}

	result, err := AllNameGenderDescriptionsWithKeys()

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestRandomFirstName_WithNameGenderFemale_ReturnsFemaleName(t *testing.T) {
	input := NameGenderFemale

	subject := NewNameService(
		WithFemaleNamesList("tina"),
		WithMaleNamesList("bob"),
		WithUnisexNamesList("devin"),
	)

	result, err := subject.RandomFirstName(input)

	assert.Nil(t, err)
	assert.Equal(t, "Tina", result)
}

func TestRandomFirstName_WithNameGenderMale_ReturnsMaleName(t *testing.T) {
	input := NameGenderMale

	subject := NewNameService(
		WithFemaleNamesList("tina"),
		WithMaleNamesList("bob"),
		WithUnisexNamesList("devin"),
	)

	result, err := subject.RandomFirstName(input)

	assert.Nil(t, err)
	assert.Equal(t, "Bob", result)
}

func TestRandomFirstName_WithNameGenderUnisex_ReturnsUnisexName(t *testing.T) {
	input := NameGenderUnisex

	subject := NewNameService(
		WithFemaleNamesList("tina"),
		WithMaleNamesList("bob"),
		WithUnisexNamesList("devin"),
	)

	result, err := subject.RandomFirstName(input)

	assert.Nil(t, err)
	assert.Equal(t, "Devin", result)
}

func TestRandomFirstName_WithNameGenderAny_ReturnsAnyName(t *testing.T) {
	input := NameGenderAny

	subject := NewNameService(
		WithFemaleNamesList("tina"),
		WithMaleNamesList("bob"),
		WithUnisexNamesList("devin"),
	)

	result, err := subject.RandomFirstName(input)

	assert.Nil(t, err)
	assert.Contains(t, []string{"Tina", "Bob", "Devin"}, result)
}

func TestRandomFirstName_WithUnsupportedGender_ReturnsAnyName(t *testing.T) {
	var input NameGender = 999 // invalid

	subject := NewNameService(
		WithFemaleNamesList("tina"),
		WithMaleNamesList("bob"),
		WithUnisexNamesList("devin"),
	)

	result, err := subject.RandomFirstName(input)

	assert.Nil(t, err)
	assert.Contains(t, []string{"Tina", "Bob", "Devin"}, result)
}

func TestNewPopulatedService_HasManyNames(t *testing.T) {
	subject := NewPopulatedService()

	assert.Equal(t, 1000, len(subject.GetFirstNamesFemale()))
	assert.Equal(t, 1000, len(subject.GetFirstNamesMale()))
	assert.Equal(t, 300, len(subject.GetFirstNamesUnisex()))
}
