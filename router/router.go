package router

import (
	"blog/config"
	"blog/models"
	"log"
	"net/http"
	"text/template"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+0]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	path := config.Cfg.SystemConfig.CurrentDir
	index := path + "/template/index.html"
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(index, home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	var categorys = []models.Category{
		{
			Cid: 0,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid: 0,
			Title: "go博客",
			Content: "不会前端的后端不是一个合格的机器学习工程师",
			UserName: "sorcerer",
			ViewCount: 122,
			CreateAt: "2024-08-24",
			CategoryId:0,
			CategoryName: "go",
			Type:-1,
		},
	}
	var homeresponse = &models.HomeResponse{
		Viewer: config.Cfg.Viewer,
		Categorys: categorys,
		Posts: posts,
		Total: 0,
		Page: 0,
		Pages: []int{0},
		PageEnd: true,
	}
	t.Execute(w, homeresponse)

}
func Router() {
	// 1. 页面views  2.数据  3.静态资源
	http.HandleFunc("/", indexHandler)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}