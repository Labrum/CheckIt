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

$(document).ready(function(){
  $("#addBox").click(function(){
    $("body").append($(""));
  });
});

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
    
    var req1 = new XMLHttpRequest();
    xmlreq = req1;
    req1.onreadystatechange = function(){share();}
    req1.open("POST", "http://localhost:8088/share/", true);
    req1.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
    req1.send("");
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

function format(name, language){
    
}

function compile(name , position) {
    var search = name.concat("edit");
    var prog = document.getElementById(search).value;
    var req = new XMLHttpRequest();
    xmlreq = req;
    
    req.onreadystatechange = function(){compileUpdate(name);}
    position = "/".concat(position)
    nameCat = name.concat(position)
    req.open("POST", "http://localhost:8088/compile/".concat(nameCat), true);
    req.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
    req.send(prog);
 
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