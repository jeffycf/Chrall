package main

import (
	"fmt"
	"io"
)

type Handler struct {
	nbRequests int
	server     *GogoServer
}

func (h *Handler) hit() {
	h.server.nbRequests++
	h.nbRequests++
}

func (h *Handler) head(w io.Writer, title string) {
	if (title=="") {
		title = "canop.org:gOgO"
	}
	fmt.Fprintf(w, "<html><head><title>%s</title>", title)
	w.Write([]byte(`
		<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.5.1/jquery.min.js"></script>
		<style type="text/css">
			body {
				background-color: white;
				font-family: Lucida Sans Unicode, Verdana;
				font-size: 14px;
				padding: 0px;
				margin: 0px;
				color: white;
			}
			div.main {
				width: 700px;
				margin:auto;
				margin-top: 20px;
				-webkit-border-radius: 15px;
				-moz-border-radius: 15px;
				border-radius: 15px;
				box-shadow: 0px 4px 10px #666;
				-moz-box-shadow: 0px 4px 10px #666;
				-webkit-box-shadow: 0px 4px 10px #666;
				font-size: 14px;
				margin-bottom: 20px;
				padding: 10px;
				-o-background-size: 100% 100%;
				-moz-background-size: 100% 100%;
				-webkit-background-size: 100% 100%;
				background-size: 100% 100%;
				background: -moz-linear-gradient(
					top,
					dodgerblue,
					indigo
				);
				background: -webkit-gradient(
					linear,
					left top, left bottom,
					from(dodgerblue),
					to(indigo)
				);
			}
			.emphase {
				color: orange;
				font-weight:bold;
				text-shadow: 0 0 0.2em black; 
			}
			a:focus {
				oultine: none;
			}
			a {
				background-color: mediumblue;
				border-top: 1px solid mediumblue;
				border-bottom: 1px solid darkblue;
				margin : 4px;
				padding-left: 8px;
				padding-right: 8px;
				-moz-border-radius: 4px;
				-webkit-border-radius: 4px;
				border-radius: 10px;
				text-decoration: none;
				color: white;
			}
			a:hover {
				border-top: 1px solid darkblue;
				border-bottom: 1px solid mediumblue;
			}
			textarea {
				width: 100%;
			}
			.invisible {
				display: none;
			}
			p#resultContent {
				margin-left: 5px;
			}
		</style>
	</head>
	<body><div class=main>`))
}

func (h *Handler) foot(w io.Writer) {
	fmt.Fprint(w, "</div></body></html>")
}