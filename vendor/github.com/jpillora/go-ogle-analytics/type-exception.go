package ga

import "net/url"

//WARNING: This file was generated. Do not edit.

//Exception Hit Type
type Exception struct {
	description         string
	descriptionSet      bool
	isExceptionFatal    bool
	isExceptionFatalSet bool
}

// NewException creates a new Exception Hit Type.

func NewException() *Exception {
	h := &Exception{}
	return h
}

func (h *Exception) addFields(v url.Values) error {
	if h.descriptionSet {
		v.Add("exd", h.description)
	}
	if h.isExceptionFatalSet {
		v.Add("exf", bool2str(h.isExceptionFatal))
	}
	return nil
}

// Specifies the description of an exception.
func (h *Exception) Description(description string) *Exception {
	h.description = description
	h.descriptionSet = true
	return h
}

// Specifies whether the exception was fatal.
func (h *Exception) IsExceptionFatal(isExceptionFatal bool) *Exception {
	h.isExceptionFatal = isExceptionFatal
	h.isExceptionFatalSet = true
	return h
}

func (h *Exception) Copy() *Exception {
	c := *h
	return &c
}
