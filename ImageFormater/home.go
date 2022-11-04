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
              padding: 20px,
              children: [
                  {
                      type: ui.Column,
                      width: 80%,
                      height: 50%,
                      margin: {left: 1wt, right: 1wt },
                      align: {horizontal: center},
                      background: {
                          type: ui.ShapeDrawable,
                          shape: {
                               name: rect,
                               stroke: #d3d3d3,
                               strokeWidth: 2px,
                               cornerRadius: 20px,
                               dash: [5, 5],
                          },
                      },
                      children: [
                          {
                              type: ui.Image,
                              width: 50px,
                              height: 1:1,
                              src: "assets/icon-add-new.png",
                              margin: {top: 4wt, bottom: 3wt},
                          },{
                              type: ui.Text,
                              text: "Choose an image to convert",
                              font: {size: 15},
                              textColor: #8b8b8b,
                              margin: {bottom: 1wt},
                          }
                      ],
                  },{
                      type: ui.Column,
                      width: 100%,
                      height: 50%,
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
                          },{
                               type: ui.Row,
                               width: 100%,
                               margin: {top: 20px},
                               align: {horizontal: center},
                               children: [
                                   {
                                       type: ui.Text,
                                       text: "Save To:",
                                   },{
                                       type: ui.Text,
                                       width: 57%,
                                       text: "/home/users",
                                       margin: {left: 20px},
                                   },{
                                       type: ui.Button,
                                       text: Change,
                                       margin: {left: 20px},
                                   }
                               ],
                          },{
                               type: ui.Button,
                               text: Convert,
                               margin: {top: 20px, left: 1wt, right: 1wt},
                               theme: [btn, btn_primary],
                          }
                      ],
                  }
              ],
          }

    }

    `

}