// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Cards(choice string) templ.Component {
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
		switch choice {
		case "image":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"cards\" class=\"flex flex-col md:flex-row justify-center items-center gap-16 mt-8 select-none\"><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;pngFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = imageSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">PNG</p><input type=\"file\" accept=\"image/png\" id=\"pngFile\" name=\"file\" onchange=\"submitForm(&#39;pngFile&#39;)\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;jpgFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = imageSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">JPG</p><input type=\"file\" accept=\"image/jpeg\" id=\"jpgFile\" name=\"file\" onchange=\"submitForm(&#39;jpgFile&#39;)\" hidden></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "document":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"cards\" class=\"flex flex-col md:flex-row justify-center items-center gap-16 mt-8 select-none\"><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;pdfFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = documentSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">PDF</p><input type=\"file\" accept=\"application/pdf\" id=\"pdfFile\" name=\"file\" onchange=\"submitForm(&#39;pdfFile&#39;)\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;docxFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = documentSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">DOCX</p><input type=\"file\" accept=\"application/vnd.openxmlformats-officedocument.wordprocessingml.document\" id=\"docxFile\" name=\"file\" onchange=\"submitForm(&#39;docxFile&#39;)\" hidden></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "audio":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"cards\" class=\"flex flex-col md:flex-row justify-center items-center gap-16 mt-8 select-none\"><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;mp3File&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = audioSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">MP3</p><input type=\"file\" accept=\"audio/mpeg\" id=\"mp3File\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;wavFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = audioSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">WAV</p><input type=\"file\" accept=\"audio/wav\" id=\"wavFile\" hidden></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		case "video":
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"cards\" class=\"flex flex-col md:flex-row justify-center items-center gap-16 mt-8 select-none\"><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;mp4File&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = videoSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">MP4</p><input type=\"file\" accept=\"video/mp4\" id=\"mp4File\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;mkvFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = videoSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">MKV</p><input type=\"file\" accept=\"video/x-matroska\" id=\"mkvFile\" hidden></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		default:
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul id=\"cards\" class=\"flex flex-col md:flex-row justify-center items-center gap-16 mt-8 select-none\"><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;pngFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = imageSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">PNG</p><input type=\"file\" accept=\"image/png\" id=\"pngFile\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;jpgFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = imageSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">JPG</p><input type=\"file\" accept=\"image/jpg\" id=\"jpgFile\" hidden></li><li class=\"flex flex-col items-center p-8 rounded-lg hover:scale-110 bg-[#7AA2E3] cursor-pointer\" onclick=\"document.getElementById(&#39;webpFile&#39;).click();\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = imageSVG().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"text-[#f8f6e3] font-bold text-4xl\">WebP</p><input type=\"file\" accept=\"image/webp\" id=\"webpFile\" hidden></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n\tdocument.body.addEventListener('htmx:afterOnLoad', function (event) {\n\t\tvar fileInputId = event.detail.elt.id;\n\t\tvar response = event.detail.xhr.responseText;\n\t\tvar link = document.createElement('a');\n\t\tlink.href = response;\n\n\t\tif (fileInputId === 'pngFile') {\n\t\t\tlink.download = 'shrinked.png';\n\t\t} else if (fileInputId === 'jpgFile') {\n\t\t\tlink.download = 'shrinked.jpg';\n\t\t}\n\n\t\tdocument.body.appendChild(link);\n\t\tlink.click();\n\t\tdocument.body.removeChild(link);\n\t});\n\n\tfunction submitForm(fileInputId) {\n\t\tvar form = document.createElement('form');\n\t\tform.method = 'POST';\n\t\tform.enctype = 'multipart/form-data';\n\t\tform.style.display = 'none';\n\n\t\tif (fileInputId === 'pngFile') {\n\t\t\tform.action = '/shrink/png';\n\t\t} else if (fileInputId === 'jpgFile') {\n\t\t\tform.action = '/shrink/jpg';\n\t\t} else if (fileInputId === 'pdfFile') {\n\t\t\tform.action = '/shrink/pdf'\n\t\t} else if (fileInputId === 'docxFile') {\n\t\t\tform.action = '/shrink/docx';\n\t\t}\n\n\t\tvar fileInput = document.getElementById(fileInputId);\n\t\tform.appendChild(fileInput.cloneNode(true));\n\n\t\tdocument.body.appendChild(form);\n\t\tform.submit();\n\t\tdocument.body.removeChild(form);\n\t}\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
