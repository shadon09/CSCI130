package mick

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init(){
	http.HandleFunc("/", index)
}

func index(res http.ResponseWriter, req *http.Request){
	html := ""
	//Build Cookie
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name: "my-cookie",
		Value: id.String(),
		HttpOnly: true,
	}
	http.SetCookie(res,cookie)

	//Store memcache
	ctx := appengine.NewContext(req)
	item := memcache.Item{
		Key: id.String(),
		Value: []byte("Matthew"),
	}
	memcache.Set(ctx,&item)

	//Get uuid from cookie
	cookie, _ = req.Cookie("my-cookie")
	if cookie != nil {
		 html += "UUID from cookie: " + cookie.Value + "<br>"
	}

	//Get uuid and value from memcache
	ctx = appengine.NewContext(req)
	item0, _ := memcache.Get(ctx, id.String())
	if item0 != nil {
		html += "Value from Memcache using uuid: " + string(item0.Value) + "<br>"
	}
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(res, html)

}
