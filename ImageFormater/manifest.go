package main

// Do not change this name, nux need manifest to generate AndroidManifest.xml
const manifest = `
{
    import: {
        ui: "nuxui.org/nuxui/ui",
    },

    application: {
        // display name at luancher
		label: hello,

        // application identifier name
        name: "org.nuxui.samples.ImageFormater",
    },

    permissions: [
        // wifi,
        storage,
        viewPhoto,
        savePhoto,
    ],

    mainWindow: {
        width: 500px,
        height: 1:1,
        title: "Image Foramter",
        content: {
            type: main.Home,
        },
        background: #ffffff,
    },
}
`