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
                      children: [
                          {
                               type: ui.Row,
                               width: 100%,
                               align: {vertical: center},
                               children: [
                                   {
                                       type: ui.Text,
                                       text: "Convert To:",
                                   },{
                                       type: ui.Options,
                                       single: true,
                                       content: {
                                           type: ui.Column,
                                           children: [
                                               {
                                                   type: ui.Row,
                                                   children: [
                                                   {style: [radio], text: bmp},
                                                   {style: [radio], text: ico},
                                                   {style: [radio], text: jpg},
                                                   {style: [radio], text: png},
                                                   {style: [radio], text: tif},
                                                   ],

                                               },{
                                                   type: ui.Row,
                                                   children: [
                                                   {style: [radio], text: gif},
                                                   {style: [radio], text: webp},
                                                   {style: [radio], text: pcx},
                                                   {style: [radio], text: tga},
                                                   ],
                                               }
                                           ],
                                       },
                                   }
                               ],
                          }
                      ],
                  }
              ],
          }

    }

    `

}