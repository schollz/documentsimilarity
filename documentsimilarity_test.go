package documentsimilarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJaccardSimilarity(t *testing.T) {
	documents := []string{"the cow jumped over the fence", "the rabbit ate a carrot", "the fox jumped over the moon"}
	ds, err := New(documents)
	assert.Nil(t, err)
	fmt.Printf("%+v", ds)

	similarities, err := ds.JaccardSimilarity("the rabbit jumped over the moon")
	assert.Nil(t, err)
	// the second sentence is the most similar
	assert.Equal(t, 2, similarities[0].Index)
	// the first sentence is the second-most similar
	assert.Equal(t, 0, similarities[1].Index)
}

func TestCosineSimilarity(t *testing.T) {
	documents := []string{"the cow jumped over the fence", "the rabbit ate a carrot", "the fox jumped over the moon"}
	ds, err := New(documents)
	assert.Nil(t, err)
	fmt.Printf("%+v", ds)

	similarities, err := ds.CosineSimilarity("the rabbit jumped over the moon")
	assert.Nil(t, err)
	// the second sentence is the most similar
	assert.Equal(t, 2, similarities[0].Index)
	// the first sentence is the second-most similar
	assert.Equal(t, 0, similarities[1].Index)
}
