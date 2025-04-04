// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterencoder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
)

type _highlight struct {
	v *types.Highlight
}

func NewHighlight() *_highlight {

	return &_highlight{v: types.NewHighlight()}

}

// A string that contains each boundary character.
func (s *_highlight) BoundaryChars(boundarychars string) *_highlight {

	s.v.BoundaryChars = &boundarychars

	return s
}

// How far to scan for boundary characters.
func (s *_highlight) BoundaryMaxScan(boundarymaxscan int) *_highlight {

	s.v.BoundaryMaxScan = &boundarymaxscan

	return s
}

// Specifies how to break the highlighted fragments: chars, sentence, or word.
// Only valid for the unified and fvh highlighters.
// Defaults to `sentence` for the `unified` highlighter. Defaults to `chars` for
// the `fvh` highlighter.
func (s *_highlight) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *_highlight {

	s.v.BoundaryScanner = &boundaryscanner
	return s
}

// Controls which locale is used to search for sentence and word boundaries.
// This parameter takes a form of a language tag, for example: `"en-US"`,
// `"fr-FR"`, `"ja-JP"`.
func (s *_highlight) BoundaryScannerLocale(boundaryscannerlocale string) *_highlight {

	s.v.BoundaryScannerLocale = &boundaryscannerlocale

	return s
}

func (s *_highlight) Encoder(encoder highlighterencoder.HighlighterEncoder) *_highlight {

	s.v.Encoder = &encoder
	return s
}

func (s *_highlight) Fields(fields map[string]types.HighlightField) *_highlight {

	s.v.Fields = fields
	return s
}

func (s *_highlight) AddField(key string, value types.HighlightFieldVariant) *_highlight {

	var tmp map[string]types.HighlightField
	if s.v.Fields == nil {
		s.v.Fields = make(map[string]types.HighlightField)
	} else {
		tmp = s.v.Fields
	}

	tmp[key] = *value.HighlightFieldCaster()

	s.v.Fields = tmp
	return s
}

func (s *_highlight) ForceSource(forcesource bool) *_highlight {

	s.v.ForceSource = &forcesource

	return s
}

// The size of the highlighted fragment in characters.
func (s *_highlight) FragmentSize(fragmentsize int) *_highlight {

	s.v.FragmentSize = &fragmentsize

	return s
}

// Specifies how text should be broken up in highlight snippets: `simple` or
// `span`.
// Only valid for the `plain` highlighter.
func (s *_highlight) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *_highlight {

	s.v.Fragmenter = &fragmenter
	return s
}

func (s *_highlight) HighlightFilter(highlightfilter bool) *_highlight {

	s.v.HighlightFilter = &highlightfilter

	return s
}

// Highlight matches for a query other than the search query.
// This is especially useful if you use a rescore query because those are not
// taken into account by highlighting by default.
func (s *_highlight) HighlightQuery(highlightquery types.QueryVariant) *_highlight {

	s.v.HighlightQuery = highlightquery.QueryCaster()

	return s
}

// If set to a non-negative value, highlighting stops at this defined maximum
// limit.
// The rest of the text is not processed, thus not highlighted and no error is
// returned
// The `max_analyzed_offset` query setting does not override the
// `index.highlight.max_analyzed_offset` setting, which prevails when it’s set
// to lower value than the query setting.
func (s *_highlight) MaxAnalyzedOffset(maxanalyzedoffset int) *_highlight {

	s.v.MaxAnalyzedOffset = &maxanalyzedoffset

	return s
}

func (s *_highlight) MaxFragmentLength(maxfragmentlength int) *_highlight {

	s.v.MaxFragmentLength = &maxfragmentlength

	return s
}

// The amount of text you want to return from the beginning of the field if
// there are no matching fragments to highlight.
func (s *_highlight) NoMatchSize(nomatchsize int) *_highlight {

	s.v.NoMatchSize = &nomatchsize

	return s
}

// The maximum number of fragments to return.
// If the number of fragments is set to `0`, no fragments are returned.
// Instead, the entire field contents are highlighted and returned.
// This can be handy when you need to highlight short texts such as a title or
// address, but fragmentation is not required.
// If `number_of_fragments` is `0`, `fragment_size` is ignored.
func (s *_highlight) NumberOfFragments(numberoffragments int) *_highlight {

	s.v.NumberOfFragments = &numberoffragments

	return s
}

func (s *_highlight) Options(options map[string]json.RawMessage) *_highlight {

	s.v.Options = options
	return s
}

func (s *_highlight) AddOption(key string, value json.RawMessage) *_highlight {

	var tmp map[string]json.RawMessage
	if s.v.Options == nil {
		s.v.Options = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Options
	}

	tmp[key] = value

	s.v.Options = tmp
	return s
}

// Sorts highlighted fragments by score when set to `score`.
// By default, fragments will be output in the order they appear in the field
// (order: `none`).
// Setting this option to `score` will output the most relevant fragments first.
// Each highlighter applies its own logic to compute relevancy scores.
func (s *_highlight) Order(order highlighterorder.HighlighterOrder) *_highlight {

	s.v.Order = &order
	return s
}

// Controls the number of matching phrases in a document that are considered.
// Prevents the `fvh` highlighter from analyzing too many phrases and consuming
// too much memory.
// When using `matched_fields`, `phrase_limit` phrases per matched field are
// considered. Raising the limit increases query time and consumes more memory.
// Only supported by the `fvh` highlighter.
func (s *_highlight) PhraseLimit(phraselimit int) *_highlight {

	s.v.PhraseLimit = &phraselimit

	return s
}

// Use in conjunction with `pre_tags` to define the HTML tags to use for the
// highlighted text.
// By default, highlighted text is wrapped in `<em>` and `</em>` tags.
func (s *_highlight) PostTags(posttags ...string) *_highlight {

	for _, v := range posttags {

		s.v.PostTags = append(s.v.PostTags, v)

	}
	return s
}

// Use in conjunction with `post_tags` to define the HTML tags to use for the
// highlighted text.
// By default, highlighted text is wrapped in `<em>` and `</em>` tags.
func (s *_highlight) PreTags(pretags ...string) *_highlight {

	for _, v := range pretags {

		s.v.PreTags = append(s.v.PreTags, v)

	}
	return s
}

// By default, only fields that contains a query match are highlighted.
// Set to `false` to highlight all fields.
func (s *_highlight) RequireFieldMatch(requirefieldmatch bool) *_highlight {

	s.v.RequireFieldMatch = &requirefieldmatch

	return s
}

// Set to `styled` to use the built-in tag schema.
func (s *_highlight) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *_highlight {

	s.v.TagsSchema = &tagsschema
	return s
}

func (s *_highlight) Type(type_ highlightertype.HighlighterType) *_highlight {

	s.v.Type = &type_
	return s
}

func (s *_highlight) HighlightCaster() *types.Highlight {
	return s.v
}
