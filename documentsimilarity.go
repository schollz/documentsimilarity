package documentsimilarity

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

var removePuncuation *regexp.Regexp

func init() {
	var err error
	removePuncuation, err = regexp.Compile("[^a-z0-9]+")
	if err != nil {
		panic(err)
	}
}

// DocumentSimilarity used to create similar
type DocumentSimilarity struct {
	Documents []string
	Bags      []map[string]int
}

// New returns new DocumentSimilarity
func New(documents []string) (ds *DocumentSimilarity, err error) {
	ds = new(DocumentSimilarity)
	ds.Documents = documents
	ds.Bags = make([]map[string]int, len(documents))

	// create bags of words
	for i, doc := range ds.Documents {
		ds.Bags[i] = getBag(doc)
	}
	return
}

func getBag(document string) map[string]int {
	m := make(map[string]int)
	for _, word := range strings.Fields(removePuncuation.ReplaceAllString(strings.ToLower(document), " ")) {
		if _, ok := m[word]; !ok {
			m[word] = 0
		}
		m[word]++
	}
	return m
}

type IndexSimilarity struct {
	Index      int
	Similarity float64
}

// JaccardSimilarity analyzes the documents using Jaccard similarity
// https://stats.stackexchange.com/a/290740
func (ds *DocumentSimilarity) JaccardSimilarity(document string) (similarities []IndexSimilarity, err error) {
	indexBag := getBag(document)
	similarities = make([]IndexSimilarity, len(ds.Bags))
	for i, otherBag := range ds.Bags {
		intersectionLength, unionLength := getStats(indexBag, otherBag)
		similarities[i].Index = i
		similarities[i].Similarity = math.Round(float64(intersectionLength)/float64(unionLength)*10000) / 10000
	}

	sort.Slice(similarities, func(i, j int) bool {
		return similarities[i].Similarity > similarities[j].Similarity
	})
	return
}

func getStats(bag1, bag2 map[string]int) (intersectionLength int, unionLength int) {
	union := make(map[string]struct{})
	intersection := make(map[string]struct{})
	for k := range bag1 {
		union[k] = struct{}{}
		if _, ok := bag2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	for k := range bag2 {
		union[k] = struct{}{}
		if _, ok := bag1[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return len(intersection), len(union)
}
