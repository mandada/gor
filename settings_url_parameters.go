// settings_url_parameters.go
package main

import (
	"errors"
	"fmt"
	"strings"
)

type UrlParameter struct {
	parakey   string
	paravalue string
}

type UrlParametes []UrlParameter

func (p *UrlParametes) String() string {
	return fmt.Sprint(*p)
}

func (p *UrlParametes) Set(value string) error {
	v := strings.SplitN(value, "=", 2)
	if len(v) != 2 {
		return errors.New("Expected `Key=Value`")
	}
	parameter := UrlParameter{
		strings.TrimSpace(v[0]),
		strings.TrimSpace(v[1]),
	}

	*p = append(*p, parameter)
	return nil
}

func (p *UrlParametes) SetParameter() string {
	var additionalpara string
	for _, para := range *p {
		additionalpara = fmt.Sprintf("%v%v=%v&", additionalpara, para.parakey, para.paravalue)
	}
	return additionalpara
}
