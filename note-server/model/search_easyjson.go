// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel(in *jlexer.Lexer, out *SearchTagResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "took":
			out.Took = int64(in.Int64())
		case "hits":
			easyjsonD4176298Decode(in, &out.Hits)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel(out *jwriter.Writer, in SearchTagResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Took))
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		easyjsonD4176298Encode(out, in.Hits)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchTagResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchTagResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchTagResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchTagResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel(l, v)
}
func easyjsonD4176298Decode(in *jlexer.Lexer, out *struct {
	Total struct {
		Value int64 `json:"value"`
	} `json:"total"`
	Hits []*SearchTagHit `json:"hits"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "total":
			easyjsonD4176298Decode1(in, &out.Total)
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.Hits = nil
			} else {
				in.Delim('[')
				if out.Hits == nil {
					if !in.IsDelim(']') {
						out.Hits = make([]*SearchTagHit, 0, 8)
					} else {
						out.Hits = []*SearchTagHit{}
					}
				} else {
					out.Hits = (out.Hits)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *SearchTagHit
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(SearchTagHit)
						}
						if data := in.Raw(); in.Ok() {
							in.AddError((*v1).UnmarshalJSON(data))
						}
					}
					out.Hits = append(out.Hits, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298Encode(out *jwriter.Writer, in struct {
	Total struct {
		Value int64 `json:"value"`
	} `json:"total"`
	Hits []*SearchTagHit `json:"hits"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix[1:])
		easyjsonD4176298Encode1(out, in.Total)
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		if in.Hits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Hits {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					out.Raw((*v3).MarshalJSON())
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonD4176298Decode1(in *jlexer.Lexer, out *struct {
	Value int64 `json:"value"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "value":
			out.Value = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298Encode1(out *jwriter.Writer, in struct {
	Value int64 `json:"value"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"value\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Value))
	}
	out.RawByte('}')
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel1(in *jlexer.Lexer, out *SearchTagHit) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_score":
			out.Score = float64(in.Float64())
		case "_index":
			out.Index = string(in.String())
		case "_type":
			out.Type = string(in.String())
		case "_source":
			easyjsonD4176298DecodeECodingNetLogonodNoteServerModel2(in, &out.Source)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel1(out *jwriter.Writer, in SearchTagHit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_score\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Score))
	}
	{
		const prefix string = ",\"_index\":"
		out.RawString(prefix)
		out.String(string(in.Index))
	}
	{
		const prefix string = ",\"_type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix)
		easyjsonD4176298EncodeECodingNetLogonodNoteServerModel2(out, in.Source)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchTagHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchTagHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchTagHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchTagHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel1(l, v)
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel2(in *jlexer.Lexer, out *UserIdTagSearch) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "user_id":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(primitive.ObjectID)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.User).UnmarshalJSON(data))
				}
			}
		case "tag_name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(string)
				}
				*out.Name = string(in.String())
			}
		case "id":
			if in.IsNull() {
				in.Skip()
				out.ID = nil
			} else {
				if out.ID == nil {
					out.ID = new(primitive.ObjectID)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ID).UnmarshalJSON(data))
				}
			}
		case "created_at":
			if in.IsNull() {
				in.Skip()
				out.CreatedAt = nil
			} else {
				if out.CreatedAt == nil {
					out.CreatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreatedAt).UnmarshalJSON(data))
				}
			}
		case "updated_at":
			if in.IsNull() {
				in.Skip()
				out.UpdatedAt = nil
			} else {
				if out.UpdatedAt == nil {
					out.UpdatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.UpdatedAt).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel2(out *jwriter.Writer, in UserIdTagSearch) {
	out.RawByte('{')
	first := true
	_ = first
	if in.User != nil {
		const prefix string = ",\"user_id\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.User).MarshalJSON())
	}
	if in.Name != nil {
		const prefix string = ",\"tag_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Name))
	}
	if in.ID != nil {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.ID).MarshalJSON())
	}
	if in.CreatedAt != nil {
		const prefix string = ",\"created_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.CreatedAt).MarshalJSON())
	}
	if in.UpdatedAt != nil {
		const prefix string = ",\"updated_at\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel3(in *jlexer.Lexer, out *SearchCollectResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "took":
			out.Took = int64(in.Int64())
		case "hits":
			easyjsonD4176298Decode2(in, &out.Hits)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel3(out *jwriter.Writer, in SearchCollectResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"took\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Took))
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		easyjsonD4176298Encode2(out, in.Hits)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchCollectResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchCollectResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchCollectResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchCollectResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel3(l, v)
}
func easyjsonD4176298Decode2(in *jlexer.Lexer, out *struct {
	Total struct {
		Value int64 `json:"value"`
	} `json:"total"`
	Hits []*SearchCollectHit `json:"hits"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "total":
			easyjsonD4176298Decode1(in, &out.Total)
		case "hits":
			if in.IsNull() {
				in.Skip()
				out.Hits = nil
			} else {
				in.Delim('[')
				if out.Hits == nil {
					if !in.IsDelim(']') {
						out.Hits = make([]*SearchCollectHit, 0, 8)
					} else {
						out.Hits = []*SearchCollectHit{}
					}
				} else {
					out.Hits = (out.Hits)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *SearchCollectHit
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(SearchCollectHit)
						}
						if data := in.Raw(); in.Ok() {
							in.AddError((*v4).UnmarshalJSON(data))
						}
					}
					out.Hits = append(out.Hits, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298Encode2(out *jwriter.Writer, in struct {
	Total struct {
		Value int64 `json:"value"`
	} `json:"total"`
	Hits []*SearchCollectHit `json:"hits"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"total\":"
		out.RawString(prefix[1:])
		easyjsonD4176298Encode1(out, in.Total)
	}
	{
		const prefix string = ",\"hits\":"
		out.RawString(prefix)
		if in.Hits == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Hits {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					out.Raw((*v6).MarshalJSON())
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel4(in *jlexer.Lexer, out *SearchCollectHit) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "_score":
			out.Score = float64(in.Float64())
		case "_index":
			out.Index = string(in.String())
		case "_type":
			out.Type = string(in.String())
		case "_source":
			easyjsonD4176298DecodeECodingNetLogonodNoteServerModel5(in, &out.Source)
		case "highlight":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Highlight).UnmarshalJSON(data))
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel4(out *jwriter.Writer, in SearchCollectHit) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"_score\":"
		out.RawString(prefix[1:])
		out.Float64(float64(in.Score))
	}
	{
		const prefix string = ",\"_index\":"
		out.RawString(prefix)
		out.String(string(in.Index))
	}
	{
		const prefix string = ",\"_type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"_source\":"
		out.RawString(prefix)
		easyjsonD4176298EncodeECodingNetLogonodNoteServerModel5(out, in.Source)
	}
	if true {
		const prefix string = ",\"highlight\":"
		out.RawString(prefix)
		out.Raw((in.Highlight).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SearchCollectHit) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SearchCollectHit) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SearchCollectHit) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SearchCollectHit) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel4(l, v)
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel5(in *jlexer.Lexer, out *UserIdTagIdsCollect) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(string)
				}
				*out.Type = string(in.String())
			}
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(string)
				}
				*out.Title = string(in.String())
			}
		case "cover":
			if in.IsNull() {
				in.Skip()
				out.Cover = nil
			} else {
				if out.Cover == nil {
					out.Cover = new(string)
				}
				*out.Cover = string(in.String())
			}
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
			}
		case "url":
			if in.IsNull() {
				in.Skip()
				out.Url = nil
			} else {
				if out.Url == nil {
					out.Url = new(string)
				}
				*out.Url = string(in.String())
			}
		case "url_hash":
			if in.IsNull() {
				in.Skip()
				out.UrlHash = nil
			} else {
				if out.UrlHash == nil {
					out.UrlHash = new(string)
				}
				*out.UrlHash = string(in.String())
			}
		case "site_domain":
			if in.IsNull() {
				in.Skip()
				out.SiteDomain = nil
			} else {
				if out.SiteDomain == nil {
					out.SiteDomain = new(string)
				}
				*out.SiteDomain = string(in.String())
			}
		case "user_collected":
			if in.IsNull() {
				in.Skip()
				out.UserCollected = nil
			} else {
				if out.UserCollected == nil {
					out.UserCollected = new(int)
				}
				*out.UserCollected = int(in.Int())
			}
		case "id":
			if in.IsNull() {
				in.Skip()
				out.ID = nil
			} else {
				if out.ID == nil {
					out.ID = new(primitive.ObjectID)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.ID).UnmarshalJSON(data))
				}
			}
		case "created_at":
			if in.IsNull() {
				in.Skip()
				out.CreatedAt = nil
			} else {
				if out.CreatedAt == nil {
					out.CreatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.CreatedAt).UnmarshalJSON(data))
				}
			}
		case "updated_at":
			if in.IsNull() {
				in.Skip()
				out.UpdatedAt = nil
			} else {
				if out.UpdatedAt == nil {
					out.UpdatedAt = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.UpdatedAt).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel5(out *jwriter.Writer, in UserIdTagIdsCollect) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != nil {
		const prefix string = ",\"type\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(*in.Type))
	}
	if in.Title != nil {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Title))
	}
	if in.Cover != nil {
		const prefix string = ",\"cover\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Cover))
	}
	if in.Description != nil {
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Description))
	}
	if in.Url != nil {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.Url))
	}
	if in.UrlHash != nil {
		const prefix string = ",\"url_hash\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(*in.UrlHash))
	}
	{
		const prefix string = ",\"site_domain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.SiteDomain == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.SiteDomain))
		}
	}
	if in.UserCollected != nil {
		const prefix string = ",\"user_collected\":"
		out.RawString(prefix)
		out.Int(int(*in.UserCollected))
	}
	if in.ID != nil {
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.Raw((*in.ID).MarshalJSON())
	}
	if in.CreatedAt != nil {
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((*in.CreatedAt).MarshalJSON())
	}
	if in.UpdatedAt != nil {
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((*in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel6(in *jlexer.Lexer, out *Highlight) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				in.Delim('[')
				if out.Title == nil {
					if !in.IsDelim(']') {
						out.Title = make([]string, 0, 4)
					} else {
						out.Title = []string{}
					}
				} else {
					out.Title = (out.Title)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.Title = append(out.Title, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "full_text":
			if in.IsNull() {
				in.Skip()
				out.Fulltext = nil
			} else {
				in.Delim('[')
				if out.Fulltext == nil {
					if !in.IsDelim(']') {
						out.Fulltext = make([]string, 0, 4)
					} else {
						out.Fulltext = []string{}
					}
				} else {
					out.Fulltext = (out.Fulltext)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.Fulltext = append(out.Fulltext, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel6(out *jwriter.Writer, in Highlight) {
	out.RawByte('{')
	first := true
	_ = first
	if len(in.Title) != 0 {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		{
			out.RawByte('[')
			for v9, v10 := range in.Title {
				if v9 > 0 {
					out.RawByte(',')
				}
				out.String(string(v10))
			}
			out.RawByte(']')
		}
	}
	if len(in.Fulltext) != 0 {
		const prefix string = ",\"full_text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Fulltext {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Highlight) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Highlight) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Highlight) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Highlight) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel6(l, v)
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel7(in *jlexer.Lexer, out *ErrorResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "error":
			if in.IsNull() {
				in.Skip()
				out.Info = nil
			} else {
				if out.Info == nil {
					out.Info = new(ErrorInfo)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Info).UnmarshalJSON(data))
				}
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel7(out *jwriter.Writer, in ErrorResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Info != nil {
		const prefix string = ",\"error\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.Info).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel7(l, v)
}
func easyjsonD4176298DecodeECodingNetLogonodNoteServerModel8(in *jlexer.Lexer, out *ErrorInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "RootCause":
			if in.IsNull() {
				in.Skip()
				out.RootCause = nil
			} else {
				in.Delim('[')
				if out.RootCause == nil {
					if !in.IsDelim(']') {
						out.RootCause = make([]*ErrorInfo, 0, 8)
					} else {
						out.RootCause = []*ErrorInfo{}
					}
				} else {
					out.RootCause = (out.RootCause)[:0]
				}
				for !in.IsDelim(']') {
					var v13 *ErrorInfo
					if in.IsNull() {
						in.Skip()
						v13 = nil
					} else {
						if v13 == nil {
							v13 = new(ErrorInfo)
						}
						if data := in.Raw(); in.Ok() {
							in.AddError((*v13).UnmarshalJSON(data))
						}
					}
					out.RootCause = append(out.RootCause, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Type":
			out.Type = string(in.String())
		case "Reason":
			out.Reason = string(in.String())
		case "Phase":
			out.Phase = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD4176298EncodeECodingNetLogonodNoteServerModel8(out *jwriter.Writer, in ErrorInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"RootCause\":"
		out.RawString(prefix[1:])
		if in.RootCause == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.RootCause {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					out.Raw((*v15).MarshalJSON())
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"Reason\":"
		out.RawString(prefix)
		out.String(string(in.Reason))
	}
	{
		const prefix string = ",\"Phase\":"
		out.RawString(prefix)
		out.String(string(in.Phase))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ErrorInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ErrorInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD4176298EncodeECodingNetLogonodNoteServerModel8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ErrorInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ErrorInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD4176298DecodeECodingNetLogonodNoteServerModel8(l, v)
}
