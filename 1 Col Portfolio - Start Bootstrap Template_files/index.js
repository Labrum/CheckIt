

$(document).ready(function(){
  $("#addBox").click(function(){
    $("body").append($(""));
  });
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

function share(){

}

var xmlreq;

function autocompile() {
    if(!document.getElementById("autocompile").checked) {
        return;
    }
    compile();
}

function compile(name) {
    console.log(name)
    var search =name.concat("edit");
    var prog = document.getElementById(search).value;

    var req = new XMLHttpRequest();
    xmlreq = req;
    req.onreadystatechange = function(){compileUpdate(name);}
    req.open("POST", "/compile/".concat(name), true);
    req.setRequestHeader("Content-Type", "text/plain; charset=utf-8");
    req.send(prog); 
}

function compileUpdate(boxId) {
    var req = xmlreq;
    console.log(boxId)
    console.log(req.responseText);
    if(!req || req.readyState != 4) {
        return;
    }
    var out = document.getElementById(boxId.concat("output"))
    var err = document.getElementById(boxId.concat("errors"))
    if(req.status == 200) {
        console.log(boxId.concat("output"));
        
        out.innerHTML = req.responseText;
        out.className = "alert alert-success"
        err.className = "alert hide alert-danger"
    } else {
        console.log(boxId.concat("output"));
        
        err.innerHTML = req.responseText;
        out.className = "alert hide alert-success"
        err.className = "alert alert-danger"
    }
}