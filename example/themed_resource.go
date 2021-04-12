// auto-generated
// Code generated by 'fynematic'. DO NOT EDIT.

// Copyright (C) 2018 Fyne.io developers (see AUTHORS)
// All rights reserved.
//
// Use of the source code in this file is governed by a BSD 3-Clause License
// license that can be found at
// https://github.com/fyne-io/fyne/blob/v2.0.2/LICENSE
//
// Original: https://github.com/fyne-io/fyne/blob/v2.0.2/cmd/fyne_demo/data/icons.go

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// ThemedResource is a resource wrapper that will return an appropriate resource
// for the currently selected theme.
type ThemedResource struct {
	dark, light fyne.Resource
}

func isLight() bool {
	r, g, b, _ := theme.ForegroundColor().RGBA()
	return r < 0xaaaa && g < 0xaaaa && b < 0xaaaa
}

// Name returns the underlying resource name (used for caching)
func (res *ThemedResource) Name() string {
	if isLight() {
		return res.light.Name()
	}
	return res.dark.Name()
}

// Content returns the underlying content of the correct resource for the current theme
func (res *ThemedResource) Content() []byte {
	if isLight() {
		return res.light.Content()
	}
	return res.dark.Content()
}

// NewThemedResource creates a resource that adapts to the current theme setting.
func NewThemedResource(dark, light fyne.Resource) *ThemedResource {
	return &ThemedResource{dark, light}
}
