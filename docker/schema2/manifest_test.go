package schema2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	vtesting "github.com/Shopify/voucher/testing"
)

func TestToManifest(t *testing.T) {
	newManifest := vtesting.NewTestManifest()

	manifest := ToManifest(newManifest)
	assert.NotNil(t, manifest)
}
