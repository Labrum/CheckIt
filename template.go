/*
Copyright 2015 Steven Labrum

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import "text/template"

var head = template.Must(template.New("head").Parse(headStyle)) // HTML template
var openBody = template.Must(template.New("openBody").Parse(body))
var pageStart = template.Must(template.New("pageStart").Parse(pageStartText))
var box = template.Must(template.New("box").Parse(boxText))
var about = template.Must(template.New("about").Parse(aboutText))
var contact = template.Must(template.New("contact").Parse(contactText))
var pageClose = template.Must(template.New("pageClose").Parse(pageCloseText))
var htmlClose =template.Must(template.New("htmlClose").Parse(htmlCloseText))

var headStyle = `<!DOCTYPE html>
<html lang="en">

<head>

<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>

    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">

    <title>CheckIt</title>

    <!-- Bootstrap Core CSS -->
    <link href="http://localhost:8088/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom CSS -->
    <link href="http://localhost:8088/css/1-col-portfolio.css" rel="stylesheet">

    <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
        <script src="https://oss.maxcdn.com/libs/html5shiv/3.7.0/html5shiv.js"></script>
        <script src="https://oss.maxcdn.com/libs/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <style>
    	.sharelink { height: auto; }
    	.codeBody { width: 650px; height: 300px}
    </style>

</head>`

var body = `<body>

    <!-- Navigation -->
    <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
        <div class="container">
            <!-- Brand and toggle get grouped for better mobile display -->
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand" href="http://localhost:8088">Main</a>
            </div>
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                    <li>
                        <a href="/about">About</a>
                    </li>
                    <li>
                        <a href="/contact">Contact</a>
                    </li>
                    <li>
                        <a  href="#" onclick="save();return false;">Share</a>
                    </li>
                    <li>
                        <a><textarea id="sharelink" class="sharelink hide" cols="30" rows="1" spellcheck="false"> default text</textarea></a>
                    </li>
                </ul>
            </div>
            <!-- /.navbar-collapse -->
        </div>
        <!-- /.container -->
    </nav>`
var pageStartText = `  <!-- Page Content -->
    <div class="container">

        <!-- Page Heading -->
        <div class="row">
            <div class="col-lg-12">
                <h1 class="page-header">{{.Heading}}
                    <small>{{.SubHeading}}</small>
                </h1>
            </div>
        </div>
        <!-- /.row -->
`
var boxText = `        <!-- Project One -->
        <div id={{.Id}} class="row">
            <div class="col-md-7">
                <textarea class="codeBody" autofocus="true" id={{ print .Id |html}}edit spellcheck="false" onkeydown="keyHandler(event);" onkeyup="autocompile();">{{printf "%s" .Body |html}}</textarea>
            </div>
            <div class="col-md-5">
                <h3>{{.Head}}</h3>
                <h4>{{.SubHead}}</h4>
                <p>{{.Text}}</p>
                <button class="btn btn-primary" id= "compile" onclick="compile('{{ print .Id |html}}','{{ print .Position |html}}');" > Run <span class ='glyphicon glyphicon-chevron-right'></span></button>
                <input type="checkbox" id={{ print .Id |html}}pipe value="output">Pipe Output<br>
                </div>
        </div>
        <div id = {{ print .Id |html}}errors class="alert hide alert-danger"><p></p></div>
        <div id={{ print .Id |html}}output class="alert hide alert-success"><p></p></div>
        <!-- /.row -->
        <hr>`

var aboutText = `        <!-- Project One -->
        <div id=About class="row">
            <div class="col-md-7">
                <h4>{{.Text}}</h4>              
            </div>
            <div class="col-md-5">
                <p>{{.SecondaryText}}</p>                
            </div>
        </div>
        <!-- /.row -->`

var contactText = `        <!-- Project One -->
        <div id=Contact class="row">
            <div class="col-md-7">
                <h4>{{.Text}}</h4>
            </div>
            <div class="col-md-5">
                <ul>
                    <li><h3>{{.Author}}</h3></li>
                    <li><h3>{{.TelNum}}</h3></li>
                    <li><h3>{{.Email}}</h3></li>
                </ul>
            </div>
        </div>
        <!-- /.row -->`

var pageCloseText =`
        
        <hr>

        <!-- Footer -->
        <footer>
            <div class="row">
                <div class="col-lg-12">
                    <p>Copyright &copy; CheckIt 2015</p>
                </div>
            </div>
            <!-- /.row -->
        </footer>

    </div>
`
var htmlCloseText = `    <!-- jQuery -->
    <script src="js/jquery.js"></script>

    <!-- Bootstrap Core JavaScript -->
    <script src="js/bootstrap.min.js"></script>
    <script src="js/index.js"></script>

</body>

</html>
`
