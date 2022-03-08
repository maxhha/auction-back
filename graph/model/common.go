package model

import (
	"fmt"
	"io"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const dateTimeLayout = "2006-01-02T15:04:05.000Z"

func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(t.Format(dateTimeLayout)))
	})
}

func UnmarshalDateTime(v interface{}) (time.Time, error) {
	str, ok := v.(string)

	if !ok {
		return time.Time{}, fmt.Errorf("convert to string")
	}

	t, err := time.Parse(dateTimeLayout, str)

	if err != nil {
		return time.Time{}, fmt.Errorf("time parse: %w", err)
	}

	return t, nil
}