package main

import (


"nuxui.org/nuxui/nux"

)

type Home interface{


nux.Component

}

type home struct{


*nux.ComponentBase

}


func NewHome(manifest nux.Attr) Home{


 me := &home{}
 me.ComponentBase = nux.NewComponentBase(me, manifest)
 nux.InflateLayout(me, me.layout(), nux.InflateStyle(me.style()))
 return me


}

func (me *home)style() string {


    return `
{
        import: {

               ui: nuxui.org/nuxui/ui,

        }
        style: {

              radio: {

                          type: ui.Radio,
                          textColor: #000000,
                          font: {size: 14},
              }


        }


}

    `



}

func (me *home)layout()string{

    return `
    {
          import: {
              ui: nuxui.org/nuxui/ui,
          }

          layout: {
              type: ui.Column,
              width: 100%,
              height: 100%,
              background: #00ff0033,
              children: [
                  {
                      type: ui.Column,
                      width: 100%,
                      height: 50%,
                      background: #ff000033,
                  },{
                      type: ui.Column,
                      width: 100%,
                      height: 50%,
                      background: #0000ff33,
                  }
              ],
          }

    }

    `

}