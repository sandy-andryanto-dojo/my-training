package main

import ("net/http";"io")


func hello(res http.ResponseWriter, req *http.Request){
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(res, 
	`<doctype html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
			<body>
				<h2>Hello World</h2>
			</body>
		</html>
	`,)
}

func main(){
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}