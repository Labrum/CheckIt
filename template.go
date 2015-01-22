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

package CheckIt

import "text/template"

var headTemp = template.Must(template.New("headTemp").Parse(headStyle)) // HTML template
var openBodyTemp = template.Must(template.New("openBodyTemp").Parse(body))
var pageStartTemp = template.Must(template.New("pageStartTemp").Parse(pageStartText))
var boxTemp = template.Must(template.New("boxTemp").Parse(boxText))
var aboutTemp = template.Must(template.New("aboutTemp").Parse(aboutText))
var pageCloseTemp = template.Must(template.New("pageCloseTemp").Parse(pageCloseText))
var htmlCloseTemp = template.Must(template.New("htmlCloseTemp").Parse(htmlCloseText))

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
    <link href="//maxcdn.bootstrapcdn.com/bootstrap/3.3.2/css/bootstrap.min.css" rel="stylesheet">
    <style>
        /*!
         * Start Bootstrap - 1 Col Portfolio HTML Template (http://startbootstrap.com)
         * Code licensed under the Apache License v2.0.
         * For details, see http://www.apache.org/licenses/LICENSE-2.0.
         */

        body {
            padding-top: 70px; /* Required padding for .navbar-fixed-top. Remove if using .navbar-static-top. Change if height of navigation changes. */
        }

        footer {
            margin: 50px 0;
        }
    </style>
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
                <a class="navbar-brand" href="/">Main</a>
            </div>
            <!-- Collect the nav links, forms, and other content for toggling -->
            <div class="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
                <ul class="nav navbar-nav">
                    <li>
                        <a href="/about">About</a>
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

var pageCloseText = `
        
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
    <script src="//maxcdn.bootstrapcdn.com/bootstrap/3.3.2/js/bootstrap.min.js"></script>
    <script>
        $("textarea").keydown(function(e) {
        if(e.keyCode === 9) { // tab was pressed
            // get caret position/selection
            var start = this.selectionStart;
            var end = this.selectionEnd;

            var $this = $(this);
            var value = $this.val();

            // set textarea value to: text before caret + tab + text after caret
            $this.val(value.substring(0, start)
                        + "\t"
                        + value.substring(end));

            // put caret at right position again (add one for the tab)
            this.selectionStart = this.selectionEnd = start + 1;

            // prevent the focus lose
            e.preventDefault();
        }
    });

    function insertTabs(n) {
        // find the selection start and end
        var cont  = document.getElementById("edit");
        var start = cont.selectionStart;
        var end   = cont.selectionEnd;
        // split the textarea content into two, and insert n tabs
        var v = cont.value;
        var u = v.substr(0, start);
        for (var i=0; i<n; i++) {
            u += "\t";
        }
        u += v.substr(end);
        // set revised content
        cont.value = u;
        // reset caret position after inserted tabs
        cont.selectionStart = start+n;
        cont.selectionEnd = start+n;
    }

    function autoindent(el) {
        var curpos = el.selectionStart;
        var tabs = 0;
        while (curpos > 0) {
            curpos--;
            if (el.value[curpos] == "\t") {
                tabs++;
            } else if (tabs > 0 || el.value[curpos] == "\n") {
                break;
            }
        }
        setTimeout(function() {
            insertTabs(tabs);
        }, 1);
    }

    function preventDefault(e) {
        if (e.preventDefault) {
            e.preventDefault();
        } else {
            e.cancelBubble = true;
        }
    }

    function keyHandler(event) {
        var e = window.event || event;
        if (e.keyCode == 9) { // tab
            insertTabs(1);
            preventDefault(e);
            return false;
        }
        if (e.keyCode == 13) { // enter
            if (e.shiftKey) { // +shift
                compile(e.target);
                preventDefault(e);
                return false;
            } else {
                autoindent(e.target);
            }
        }
        return true;
    }
    var xmlreq

    function save() {
        console.log("share button clicked");
        
        var test = document.getElementsByClassName("codeBody");

        var texts = []

        for (var key in test) {
            texts.push(test[key].value);
        }


        var req1 = new XMLHttpRequest();
        xmlreq = req1;
        req1.onreadystatechange = function(){share();}
        req1.open("POST", window.location.origin+"/share/", true);
        req1.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
        var myJsonString = JSON.stringify(texts);
        req1.send(myJsonString);
    }

    function share(){
        var req1 = xmlreq;

        var s = document.getElementById("sharelink");
        s.className = "sharelink";
        s.innerHTML = req1.responseText;
        s.select();
        
    }

    function autocompile() {
        if(!document.getElementById("autocompile").checked) {
            return;
        }
        compile();
    }

    function compile(name , position) {
        var search = name.concat("edit");
        var prog = document.getElementById(search).value;
        var test = document.getElementsByClassName("codeBody");

        var texts = []

        for (var key in test) {
            texts.push(test[key].value);
        }

        var req = new XMLHttpRequest();
        xmlreq = req;
        
        req.onreadystatechange = function(){compileUpdate(name);}
        position = "/".concat(position)
        nameCat = name.concat(position)
        req.open("POST", window.location.origin+ "/compile/".concat(nameCat), true);
        req.setRequestHeader("Content-Type", "text/plain; charset=utf-8");


        var myJsonString = JSON.stringify(texts);
        req.send(myJsonString);
     
    }

    function compileUpdate(boxId) {
        var req = xmlreq;
        if(!req || req.readyState != 4) {
            return;
        }

        console.log()
        var out = document.getElementById(boxId.concat("output"))
        var err = document.getElementById(boxId.concat("errors"))
        if(req.status == 200) {
           
            out.innerHTML = req.responseText;
            out.className = "alert alert-success"
            err.className = "alert hide alert-danger"
        } else {
            
            err.innerHTML = req.responseText;
            out.className = "alert hide alert-success"
            err.className = "alert alert-danger"
        }
    }
    </script>   
</body>

</html>
`
