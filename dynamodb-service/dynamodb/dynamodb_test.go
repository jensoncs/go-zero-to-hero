package dynamodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGetCorrectZipCodeRelatedToTheLeadCusCode(t *testing.T) {
	resp, _ := GetZipCodeItemsFromDynamoDb("0212xyz")
	assert.Equal(t, "560029", resp.ZipCode)
}
func TestShouldGetCorrectZoneRelatedToTheLeadCusCode(t *testing.T) {
	resp, _ := GetZipCodeItemsFromDynamoDb("0212xyz")
	assert.Equal(t, "majestic", resp.Zone)
}
