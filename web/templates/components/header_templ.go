// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

/*
Implement cards li elements using for loop by iterating over the choices sent by server
*/

func Header() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"fixed top-0 flex md:justify-between rounded-lg mt-8 px-4 w-11/12 md:w-3/5 z-10 bg-[#97E7E1] dark:bg-[#3282B8] bg-opacity-40 dark:bg-opacity-80 shadow-md shadow-[#7AA2E3] dark:shadow-[#BBE1FA]\"><ul class=\"flex items-center gap-4 md:gap-12 font-semibold text-[#7AA2E3] dark:text-[#F8F6E3] select-none cursor-pointer md:text-lg md:font-semibold w-full\"><li id=\"image\" class=\"hover:scale-110\" hx-get=\"/cards\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(`{"choice": "image" }`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/components/header.templ`, Line: 18, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cards\" hx-swap=\"innerHTML\">Image</li><li id=\"document\" class=\"hover:scale-110\" hx-get=\"/cards\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(`{"choice": "document" }`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/components/header.templ`, Line: 28, Col: 39}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cards\" hx-swap=\"innerHTML\">Document</li><li id=\"audio\" class=\"hover:scale-110\" hx-get=\"/cards\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(`{"choice": "audio" }`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/components/header.templ`, Line: 36, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cards\" hx-swap=\"innerHTML\">Audio</li><li id=\"video\" class=\"hover:scale-110\" hx-get=\"/cards\" hx-vals=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(`{"choice": "video" }`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/components/header.templ`, Line: 44, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#cards\" hx-swap=\"innerHTML\">Video</li></ul><button id=\"toggle-theme\" class=\"py-4 ml-6 dark:text-white\"><svg viewBox=\"0 0 24 24\" width=\"24\" height=\"24\" stroke=\"#7AA2E3\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"css-4ym8mv\"><path d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path></svg></button></div><script>\n\tdocument.addEventListener('DOMContentLoaded', function () {\n\t\tconst themeButton = document.getElementById('toggle-theme');\n\t\tlet currentTheme = localStorage.getItem('theme') || 'light';\n\n\t\tfunction updateTheme(newTheme) {\n\t\t\tdocument.documentElement.classList.toggle('dark', newTheme === 'dark');\n\t\t\tlocalStorage.setItem('theme', newTheme);\n\t\t\tcurrentTheme = newTheme;\n\n\t\t\t// Change the theme SVG\n\t\t\tthemeButton.innerHTML = currentTheme === 'dark' ? '<svg viewBox=\"0 0 24 24\" width=\"24\" height=\"24\" stroke=\"currentColor\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"css-4ym8mv\"><circle cx=\"12\" cy=\"12\" r=\"5\"></circle><line x1=\"12\" y1=\"1\" x2=\"12\" y2=\"3\"></line><line x1=\"12\" y1=\"21\" x2=\"12\" y2=\"23\"></line><line x1=\"4.22\" y1=\"4.22\" x2=\"5.64\" y2=\"5.64\"></line><line x1=\"18.36\" y1=\"18.36\" x2=\"19.78\" y2=\"19.78\"></line><line x1=\"1\" y1=\"12\" x2=\"3\" y2=\"12\"></line><line x1=\"21\" y1=\"12\" x2=\"23\" y2=\"12\"></line><line x1=\"4.22\" y1=\"19.78\" x2=\"5.64\" y2=\"18.36\"></line><line x1=\"18.36\" y1=\"5.64\" x2=\"19.78\" y2=\"4.22\"></line></svg>' : '<svg viewBox=\"0 0 24 24\" width=\"24\" height=\"24\" stroke=\"#7AA2E3\" stroke-width=\"2\" fill=\"none\" stroke-linecap=\"round\" stroke-linejoin=\"round\" class=\"css-4ym8mv\"><path d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path></svg>';\n\t\t}\n\n\t\tthemeButton.addEventListener('click', function () {\n\t\t\tconst newTheme = currentTheme === 'dark' ? 'light' : 'dark';\n\t\t\tupdateTheme(newTheme);\n\t\t});\n\n\t\tupdateTheme(currentTheme);\n\n\t\tfunction storeChoice(choice) {\n\t\t\tlocalStorage.setItem('choice', choice);\n\t\t}\n\n\t\tconst imageLi = document.getElementById('image');\n\t\tconst documentLi = document.getElementById('document');\n\t\tconst audioLi = document.getElementById('audio');\n\t\tconst videoLi = document.getElementById('video');\n\n\t\timageLi.addEventListener('click', function () {storeChoice('image');});\n\t\tdocumentLi.addEventListener('click', function () {storeChoice('document');});\n\t\taudioLi.addEventListener('click', function () {storeChoice('audio');});\n\t\tvideoLi.addEventListener('click', function () {storeChoice('video');});\n\t});\n\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
