package main

import (
	//"net/url"
	"ptp"
	"html/template"
	"net/http"
	"fmt"
	"time"
)

func main() {
	http.HandleFunc("/ajax", ajax)
	http.HandleFunc("/send/", send)
	http.HandleFunc("/test", test)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func test(w http.ResponseWriter, r *http.Request) {
	username:="MemelsMama"
	user:=ptp.NewUser(username, "lol")
	t,_:=template.ParseFiles("client.html")
	t.Execute(w,user)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := `<html>
		<body>
			<input type="text" id="result" />
			<input type="button" id="button" value="Click" onclick="renderButtonClicked();" />
			<script>
				function renderButtonClicked() {
				if (window.XMLHttpRequest) {
					xmlhttp=new XMLHttpRequest();
				} 
				xmlhttp.onreadystatechange = function() {
					if (xmlhttp.readyState==4 && xmlhttp.status==200) {
						document.getElementById("result").value = xmlhttp.responseText;
					}
				};
				var txt = document.getElementById("result").value
				xmlhttp.open("GET", "ajax?" + txt, true);
				console.log("ajax?" + txt)
				xmlhttp.send();	
			}
			</script>
		</body>
	</html>`
	fmt.Fprint(w, tmpl)
}

func ajax(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(r.Host)
	fmt.Fprint(w, time.Now().String())
}