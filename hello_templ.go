// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Base() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#21295C;`)
	templ_7745c5c3_CSSBuilder.WriteString(`color:#9EB3C2;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-family:Arial;`)
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`flex-direction:column;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSID := templ.CSSID(`Base`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func GetLocation(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_GetLocation_80be`,
		Function: `function __templ_GetLocation_80be(id){if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition((position) => {
        const latitude = position.coords.latitude;
        const longitude = position.coords.longitude;
        console.log(latitude, longitude);
        const element = document.getElementById(id)
        element.setAttribute("data", ` + "`" + `${latitude},${longitude}` + "`" + `)
        return
        },
    function(error) {
        console.log(error)
    }
  );
} else {
    console.log("No Geo location")
}
}`,
		Call:       templ.SafeScript(`__templ_GetLocation_80be`, id),
		CallInline: templ.SafeScriptInline(`__templ_GetLocation_80be`, id),
	}
}

func FetchWeather(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_FetchWeather_38dd`,
		Function: `function __templ_FetchWeather_38dd(id){const element = document.getElementById(id)
    const dataValue = element.getAttribute("data");
    const [latitude, longitude] = dataValue.split(",");

    const apikey = "a235dd6e346a4a2ba64153222242908"
    const apiUrl = ` + "`" + `http://api.weatherapi.com/v1/current.json?key=${apikey}&q=${latitude},${longitude}&aqi=no` + "`" + `;

    fetch(apiUrl)
    .then(response => response.json())
    .then(data => {
        console.log('Weather Data:', data);
        textContent = ` + "`" + `${data.location.country} <br></br>
        ${data.current.condition.text} <br></br>
        ${data.current.feelslike_c}°C<br></br>
        <img src="${data.current.condition.icon}" />` + "`" + `
        element.innerHTML = textContent
        })
    .catch(error => {
        console.error('Error fetching weather:', error);
    });
    
}`,
		Call:       templ.SafeScript(`__templ_FetchWeather_38dd`, id),
		CallInline: templ.SafeScriptInline(`__templ_FetchWeather_38dd`, id),
	}
}

func Hello(name string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><title>Hello Test</title></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 = []any{Base()}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, GetLocation("data-div"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 templ.ComponentScript = GetLocation("data-div")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var3.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var2).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `hello.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, FetchWeather("data-div"))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<h1 onclick=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 templ.ComponentScript = FetchWeather("data-div")
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Hello ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `hello.templ`, Line: 62, Col: 63}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><div id=\"data-div\"></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
