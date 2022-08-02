// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/nux"
	"nuxui.org/nuxui/theme"
)

func init() {
	nux.RegisterType((*Home)(nil), func(attr nux.Attr) any { return NewHome(attr) })
}

func main() {
	nux.ApplyTheme(nux.ThemeLight, theme.BootstrapLight)
	nux.Run(func(){
		nux.NewWindow(nux.Attr{
			"width":  "1:1",
			"height": "500px",
			"title":  "Image Fomater",
			"content": nux.Attr{
				"type": "main.Home",
			},
		})
	})
}
