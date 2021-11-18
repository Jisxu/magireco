package htmltemplate

var (
	Header = `<html>

<head>
    <base href="https://android.magi-reco.com">
</head>

<body>`
	Footer = `</body>

</html>`
	ContentFormat = `<div>
        <p>%v</p>
        <hr/>
        <p>%v</p>
    </div>
    <hr style="height:1px;border:none;border-top:1px dashed #0066CC;"/>`
)
