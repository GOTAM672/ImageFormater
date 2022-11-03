// Copyright 2018 The NuxUI Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"nuxui.org/nuxui/nux"
	_ "nuxui.org/nuxui/ui"
)

func init(){

        nux.RegisterType( (*Home)(nil), func(manifest nux.Attr) any { return NewHome(manifest) } )

}

func main() {
	nux.App().Init(manifest)
	nux.App().Run()
}
