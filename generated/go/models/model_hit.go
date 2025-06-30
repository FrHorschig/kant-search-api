package models

type Hit struct {

	FmtText string `json:"fmtText"`

	HighlightText string `json:"highlightText"`

	// LineByIndex is a map of fmtText string indices (rune, not byte indices) of the start of ks-meta-line tags to the line number inside the tag. This fields is used to determine the line where a search hit starts.
	LineByIndex []IndexNumberPair `json:"lineByIndex"`

	Ordinal int32 `json:"ordinal"`

	// PageByIndex is a map of fmtText string indices (rune, not byte indices) of the start of ks-meta-page tags to the page number inside the tag. This field is used to determine the page where a search hit starts.
	PageByIndex []IndexNumberPair `json:"pageByIndex"`

	Pages []int32 `json:"pages"`

	// wordIndexMap is a map of searchText string indices of the words of the text to fmtText string indices (both rune, not byte indices) of the same words. For example, the [k, v] pair [28, 847] would mean that the word at index 28 of SearchText is the same word as the one at index 847 in fmtText. This field is used to map ES search hit highlights (the highlightText field is the searchText plus ES highlighting) to fmtText.
	WordIndexMap map[string]int32 `json:"wordIndexMap"`
}
