window.onload = function () {

    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.className = "error";
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var t = JSON.parse(evt.data);
            //console.log(t);
            var item = document.createElement("div");
            item.className = "item";
            //TODO: Clean this up
            item.innerHTML = "<img src='" + t.user.profile_image_url + "'/><b>" + t.user.screen_name + ":</b><br /><i>" + t.text + "</i>";
            //TODO: Prepend vs append
            //item.insertBefore(log,log.firstChild);
            appendLog(item);
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};